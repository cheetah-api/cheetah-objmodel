/*
 *
 * Copyright (c) 2016, Cisco Systems Inc.
 * All rights reserved.
 */

package main

/* standard packages */
import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"syscall"
	"time"

	context "golang.org/x/net/context"
)

/* cheetah packages */
import (
	pb "gengo"
	"server_util"
)

const (
	MAX_PACKET_SIZE = 2048
	PRIV_PROTOCOL   = 4482 //htons(0x8211)
)

type PacketConfig struct {
	mgmt  pb.APMgmtMsgSubtype
	ctrl  pb.APCtrlMsgSubtype
	data  pb.APDataMsgSubtype
	cisco pb.APCiscoMsgSubtype
	srv   pb.APPackets_APPacketsInitNotifServer
}

// Clients packet configuration
type ClientInfo struct {
	pconfig     map[string]PacketConfig
	pmux        sync.Mutex
	pconfigured bool
}

var ci ClientInfo

func validateRegRequest(request *pb.APPacketHdr) bool {
	rsvd_subtype := false

	msgtype := request.GetMsgType()
	dbg.Println("Received msgtype", msgtype)

	if msgtype == pb.APMsgType_AP_MSG_TYPE_RESERVED {
		dbg.Println("Reserved message type", msgtype)
		return false
	}

	_, ok := pb.APMsgType_name[int32(msgtype)]
	if !ok {
		dbg.Println("Invalid message type")
		return false
	}

	switch request.MsgType {
	case pb.APMsgType_AP_MSG_TYPE_MGMT:
		_, ok = pb.APMgmtMsgSubtype_name[int32(request.GetMgmt())]
		rsvd_subtype = (request.GetMgmt() == pb.APMgmtMsgSubtype_AP_MGMT_MSG_SUBTYPE_RESERVED)

	case pb.APMsgType_AP_MSG_TYPE_CTRL:
		_, ok = pb.APCtrlMsgSubtype_name[int32(request.GetCtrl())]
		rsvd_subtype = (request.GetCtrl() == pb.APCtrlMsgSubtype_AP_CTRL_MSG_SUBTYPE_RESERVED)

	case pb.APMsgType_AP_MSG_TYPE_DATA:
		_, ok = pb.APDataMsgSubtype_name[int32(request.GetData())]
		rsvd_subtype = (request.GetData() == pb.APDataMsgSubtype_AP_DATA_MSG_SUBTYPE_RESERVED)

	case pb.APMsgType_AP_MSG_TYPE_CISCO:
		_, ok = pb.APCiscoMsgSubtype_name[int32(request.GetCisco())]
		rsvd_subtype = (request.GetCisco() == pb.APCiscoMsgSubtype_AP_CISCO_MSG_SUBTYPE_RESERVED)

	case pb.APMsgType_AP_MSG_TYPE_RESERVED:
		fallthrough

	default:
		ok = false
	}

	if rsvd_subtype {
		dbg.Println("Reserved message subtype")
		ok = false
	}

	return ok
}

func validateRegOper(addr string, msg *pb.APPacketsRegMsg) bool {
	ret := true

	if (msg.Oper != pb.APRegOp_AP_REGOP_REGISTER) &&
		(msg.Oper != pb.APRegOp_AP_REGOP_UNREGISTER) {
		dbg.Println("Invalid Oper", msg.Oper)
		return false
	}

	ci.pmux.Lock()
	if _, ok := ci.pconfig[addr]; ok {
		if msg.Oper == pb.APRegOp_AP_REGOP_REGISTER {
			dbg.Println("Duplicate Registration")
			ret = false /* duplicate registration */
		}
	} else {
		if msg.Oper == pb.APRegOp_AP_REGOP_UNREGISTER {
			dbg.Println("Non-existent Unregistration")
			ret = false /* duplicate unregistration */
		}
	}
	ci.pmux.Unlock()

	return ret
}

/*
 * Add, update or delete the config for client
 */
func updatePacketConfig(addr string,
	msg *pb.APPacketsRegMsg, cfg PacketConfig) {

	dbg.Println("update configuration for client:")

	ci.pmux.Lock()
	val, found := ci.pconfig[addr]
	if msg.Oper == pb.APRegOp_AP_REGOP_REGISTER {
		if found {
			val.mgmt |= cfg.mgmt
			val.ctrl |= cfg.ctrl
			val.data |= cfg.data
			val.cisco |= cfg.cisco
			ci.pconfig[addr] = val
		} else {
			/* add to map */
			cfg.srv = nil /* not initialized */
			ci.pconfig[addr] = cfg
		}
	} else {
		if found {
			val.mgmt = val.mgmt &^ cfg.mgmt
			val.ctrl = val.ctrl &^ cfg.ctrl
			val.data = val.data &^ cfg.data
			val.cisco = val.cisco &^ cfg.cisco
			if (val.mgmt == 0) && (val.ctrl == 0) &&
				(val.data == 0) && (val.cisco == 0) {
				/* remove from map */
				delete(ci.pconfig, addr)
			}
		}
	}

	ci.pmux.Unlock()

	return
}

func getSubscriberCount() int {
	count := 0
	ci.pmux.Lock()
	for _, v := range ci.pconfig {
		if v.srv != nil {
			count++
		}
	}
	ci.pmux.Unlock()

	return count
}

/* Update srv for client */
func updateClientSrv(srv pb.APPackets_APPacketsInitNotifServer) bool {
	var args []byte

	dbg.Println("update client srv:", srv)

	addr, err := server_util.GetAddressFromCtx(srv.Context())
	if err == false {
		fmt.Println("srv update failed: peer address", addr, "not found")
		return false
	}

	ci.pmux.Lock()
	if val, ok := ci.pconfig[addr]; ok {
		val.srv = srv
		ci.pconfig[addr] = val
	} else {
		fmt.Println("srv update failed: subscription not found")
		ci.pmux.Unlock()
		return false
	}
	ci.pmux.Unlock()

	enable := (getSubscriberCount() > 0)
	if ci.pconfigured != enable {
		dbg.Println("Packet Configuration:", ci.pconfigured, "->", enable)
		if enable {
			args = []byte("+tohost")
		} else {
			args = []byte("-tohost")
		}
		err := ioutil.WriteFile("/proc/aptrace/debug", args, 0644)
		if err != nil {
			log.Println("Error in configuring capture", err)
		} else {
			ci.pconfigured = enable
		}
	}

	return true
}

func listPacketConfig() {
	ci.pmux.Lock()
	for k, v := range ci.pconfig {
		dbg.Println("Addr:", k)
		dbg.Println("Config:")
		dbg.Println("type:%x mgmt: %x ctrl: %x data: %x cisco:%x",
			v.mgmt, v.ctrl, v.data, v.cisco)
	}
	ci.pmux.Unlock()
}

func PacketSubscribed(v PacketConfig, response *pb.APPacketsMsgRsp) bool {
	matched := false

	if v.srv == nil {
		return false
	}

	switch response.PacketHdr.MsgType {
	case pb.APMsgType_AP_MSG_TYPE_MGMT:
		matched = (v.mgmt&response.PacketHdr.GetMgmt() != 0)

	case pb.APMsgType_AP_MSG_TYPE_CTRL:
		matched = (v.ctrl&response.PacketHdr.GetCtrl() != 0)

	case pb.APMsgType_AP_MSG_TYPE_DATA:
		matched = (v.data&response.PacketHdr.GetData() != 0)

	case pb.APMsgType_AP_MSG_TYPE_CISCO:
		matched = (v.cisco&response.PacketHdr.GetCisco() != 0)

	case pb.APMsgType_AP_MSG_TYPE_RESERVED:
		fallthrough

	default:
		matched = false

	}

	return matched
}

/*
 * Find type and subtype and send the packet for subscribed clients
 */

func sendPacketResponse(buffer []byte, plen int) {
	var msgtype, subtype, length uint32

	/* Create response message */
	response := &pb.APPacketsMsgRsp{}
	response.ErrStatus = &pb.APErrorStatus{}
	response.ErrStatus.Status = pb.APErrorStatus_AP_SUCCESS

	offset := 0
	msgtype = binary.BigEndian.Uint32(buffer[offset:])
	offset = offset + 4
	subtype = binary.BigEndian.Uint32(buffer[offset:])
	offset = offset + 4
	length = binary.BigEndian.Uint32(buffer[offset:])
	offset = offset + 4

	response.PacketHdr = &pb.APPacketHdr{}
	response.PacketHdr.MsgType = pb.APMsgType(msgtype)
	switch response.PacketHdr.MsgType {
	case pb.APMsgType_AP_MSG_TYPE_MGMT:
		mgmt := &pb.APPacketHdr_Mgmt{}
		mgmt.Mgmt = pb.APMgmtMsgSubtype(subtype)
		response.PacketHdr.Subtype = mgmt

	case pb.APMsgType_AP_MSG_TYPE_CTRL:
		ctrl := &pb.APPacketHdr_Ctrl{}
		ctrl.Ctrl = pb.APCtrlMsgSubtype(subtype)
		response.PacketHdr.Subtype = ctrl

	case pb.APMsgType_AP_MSG_TYPE_DATA:
		data := &pb.APPacketHdr_Data{}
		data.Data = pb.APDataMsgSubtype(subtype)
		response.PacketHdr.Subtype = data

	case pb.APMsgType_AP_MSG_TYPE_CISCO:
		cisco := &pb.APPacketHdr_Cisco{}
		cisco.Cisco = pb.APCiscoMsgSubtype(subtype)
		response.PacketHdr.Subtype = cisco

	case pb.APMsgType_AP_MSG_TYPE_RESERVED:
		fallthrough

	default:
		dbg.Println("Invalid message type", msgtype)
	}

	response.PacketLen = length
	response.PacketBuf = make([]byte, length)
	bytes_copied := copy(response.PacketBuf, buffer[offset:])
	copy(response.PacketBuf, buffer[offset:])
	dbg.Println("\tbytes_copied :", bytes_copied)
	dbg.Println("type: ", response.PacketHdr.MsgType)
	dbg.Println("subtype: ", response.PacketHdr.Subtype)
	dbg.Println("PacketLen: ", response.PacketLen)
	for _, v := range ci.pconfig {
		if PacketSubscribed(v, response) {
			dbg.Println("Sending response to", v.srv)
			v.srv.Send(response)
		}
	}
	return
}

/*
 * APPackets Server implementation
 */
type PacketsServer struct{}

/* Get Packets Registration RPC */
func (s *PacketsServer) APPacketsRegOp(ctx context.Context,
	msg *pb.APPacketsRegMsg) (*pb.APPacketsRegMsgRsp, error) {
	var hdrs []*pb.APPacketHdr
	var config PacketConfig

	dbg.Println("\nReceived APPacketGet call")

	/* Create response message */
	response := &pb.APPacketsRegMsgRsp{}
	response.ErrStatus = &pb.APErrorStatus{}
	response.ErrStatus.Status = pb.APErrorStatus_AP_SUCCESS

	addr, err := server_util.GetAddressFromCtx(ctx)
	if err == false {
		response.ErrStatus.Status = pb.APErrorStatus_AP_EINVAL
		fmt.Println("failed to get peer address")
		return response, nil
	}

	for _, request := range msg.PacketHdr {
		if validateRegRequest(request) == false {
			/* add to response */
			hdrs = append(hdrs, request)
			response.ErrStatus.Status = pb.APErrorStatus_AP_EINVAL
			continue
		}

		/* validate operation */
		if validateRegOper(addr, msg) == false {
			/* add to response */
			hdrs = append(hdrs, request)
			response.ErrStatus.Status = pb.APErrorStatus_AP_EINVAL
			continue
		}

		switch request.MsgType {
		case pb.APMsgType_AP_MSG_TYPE_MGMT:
			config.mgmt |= request.GetMgmt()

		case pb.APMsgType_AP_MSG_TYPE_CTRL:
			config.ctrl |= request.GetCtrl()

		case pb.APMsgType_AP_MSG_TYPE_DATA:
			config.data |= request.GetData()

		case pb.APMsgType_AP_MSG_TYPE_CISCO:
			config.cisco |= request.GetCisco()
		}
	}

	response.Results = &pb.APPacketsRegMsg{}
	response.Results.Oper = msg.Oper
	if len(hdrs) > 0 {
		response.Results.PacketHdr = []*pb.APPacketHdr{}
		response.Results.PacketHdr = hdrs
	} else {
		/* save client config */
		updatePacketConfig(addr, msg, config)
		listPacketConfig()
	}

	return response, nil
}

/* Packet notification init */
func (s *PacketsServer) APPacketsInitNotif(msg *pb.APPacketsGetNotifMsg, srv pb.APPackets_APPacketsInitNotifServer) error {
	/* Create response message */
	response := &pb.APPacketsMsgRsp{}
	response.ErrStatus = &pb.APErrorStatus{}
	response.ErrStatus.Status = pb.APErrorStatus_AP_SUCCESS

	/* update srv for client */
	if updateClientSrv(srv) == false {
		response.ErrStatus.Status = pb.APErrorStatus_AP_EINVAL
		srv.Send(response)
		return nil
	}

	/* response to config request */
	response.PacketHdr = &pb.APPacketHdr{}
	response.PacketHdr.MsgType = pb.APMsgType_AP_MSG_TYPE_RESERVED
	response.PacketLen = 0

	srv.Send(response)
	for {
		/* keep connection alive */
		time.Sleep(time.Second)
	}

	return nil
}

/* open socket to receive packets from kernel */
func OpenCaptureSocket() {
	/* Initialize Global packet config */
	ci = ClientInfo{pconfig: make(map[string]PacketConfig)}

	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, PRIV_PROTOCOL)
	if err != nil {
		fmt.Println("socket listen error", err)
	}

	if_info, err := net.InterfaceByName(server_util.GetCaptureInterfaceName())
	if err != nil {
		fmt.Println("interface not found", server_util.GetCaptureInterfaceName())
	}

	addr := syscall.SockaddrLinklayer{
		Protocol: PRIV_PROTOCOL,
		Ifindex:  if_info.Index,
	}

	err = syscall.Bind(fd, &addr)
	if err != nil {
		fmt.Println("socket bind error", err)
	}

	err = syscall.SetLsfPromisc(server_util.GetCaptureInterfaceName(), true)
	if err != nil {
		fmt.Println("Promisc setting  error", err)
	}

	buffer := make([]byte, MAX_PACKET_SIZE)

	rfds := &syscall.FdSet{}
	for {
		server_util.FD_ZERO(rfds)
		server_util.FD_SET(rfds, fd)
		_, err := syscall.Select(fd+1, rfds, nil, nil, nil)
		if err != nil {
			log.Fatalln(err)
		}

		// Got data on socket
		if server_util.FD_ISSET(rfds, fd) {
			plen, _, err := syscall.Recvfrom(fd, buffer, 0)
			if err != nil {
				log.Fatalln(err)
			}
			sendPacketResponse(buffer, plen)
		}
	}
}
