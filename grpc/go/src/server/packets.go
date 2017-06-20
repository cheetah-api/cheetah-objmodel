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
	"google.golang.org/grpc/peer"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"syscall"
	"time"
)

/* cheetah packages */
import (
	pb "gengo"
	"server_util"
)

const (
	INTERFACE_NAME  = "aptrace0"
	MAX_PACKET_SIZE = 2048
	PRIV_PROTOCOL   = 4482 //htons(0x8211)
)

type PacketConfig struct {
	srv     pb.APPackets_APPacketsGetServer
	msgtype pb.APMsgType
	mgmt    pb.APMgmtMsgSubtype
	ctrl    pb.APCtrlMsgSubtype
	data    pb.APDataMsgSubtype
	cisco   pb.APCiscoMsgSubtype
}

// Clients packet configuration
type ClientInfo struct {
	pconfig     map[string]PacketConfig
	pmux        sync.Mutex
	pconfigured bool
}

var ci ClientInfo

/*
 * Add, update or delete the config for client
 */
func updatePacketConfig(cfg PacketConfig) bool {
	dbg.Println("update configuration for srv:", cfg.srv)
	var args []byte

	pr, ok := peer.FromContext(cfg.srv.Context())
	if !ok {
		fmt.Println("failed to get peer from ctx")
		return false
	}

	if pr.Addr == net.Addr(nil) {
		fmt.Println("failed to get peer address")
		return false
	}
	addr := pr.Addr.String()

	dbg.Println("peer address:", addr)

	ci.pmux.Lock()
	if val, ok := ci.pconfig[addr]; ok {
		val.srv = cfg.srv
		val.msgtype = cfg.msgtype
		val.mgmt |= cfg.mgmt
		val.ctrl |= cfg.ctrl
		val.data |= cfg.data
		val.cisco |= cfg.cisco
		ci.pconfig[addr] = val
	} else {
		ci.pconfig[addr] = cfg
	}

	enable := (len(ci.pconfig) != 0)
	ci.pmux.Unlock()

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
		dbg.Println("srv: %x type:%x mgmt: %x ctrl: %x data: %x cisco:%x",
			v.srv, v.msgtype, v.mgmt, v.ctrl, v.data, v.cisco)
	}
	ci.pmux.Unlock()
}

func PacketSubscribed(v PacketConfig, response *pb.APPacketsMsgRsp) bool {
	matched := false

	if v.msgtype&response.PacketHdr.MsgType != 0 {
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

// Get Packets
func (s *PacketsServer) APPacketsGet(msg *pb.APPacketsMsg, srv pb.APPackets_APPacketsGetServer) error {
	var config PacketConfig

	dbg.Println("Received APPacketGet call", srv)

	/* Create response message */
	response := &pb.APPacketsMsgRsp{}
	response.ErrStatus = &pb.APErrorStatus{}
	response.ErrStatus.Status = pb.APErrorStatus_AP_SUCCESS

	ok := true
	rsvd_subtype := false
	for _, request := range msg.PacketHdr {
		switch request.MsgType {
		case pb.APMsgType_AP_MSG_TYPE_MGMT:
			config.mgmt |= request.GetMgmt()
			_, ok = pb.APMgmtMsgSubtype_name[int32(request.GetMgmt())]
			rsvd_subtype = (request.GetMgmt() == pb.APMgmtMsgSubtype_AP_MGMT_MSG_SUBTYPE_RESERVED)

		case pb.APMsgType_AP_MSG_TYPE_CTRL:
			config.ctrl |= request.GetCtrl()
			_, ok = pb.APCtrlMsgSubtype_name[int32(request.GetCtrl())]
			rsvd_subtype = (request.GetCtrl() == pb.APCtrlMsgSubtype_AP_CTRL_MSG_SUBTYPE_RESERVED)

		case pb.APMsgType_AP_MSG_TYPE_DATA:
			config.data |= request.GetData()
			_, ok = pb.APDataMsgSubtype_name[int32(request.GetData())]
			rsvd_subtype = (request.GetData() == pb.APDataMsgSubtype_AP_DATA_MSG_SUBTYPE_RESERVED)

		case pb.APMsgType_AP_MSG_TYPE_CISCO:
			config.cisco |= request.GetCisco()
			_, ok = pb.APCiscoMsgSubtype_name[int32(request.GetCisco())]
			rsvd_subtype = (request.GetCisco() == pb.APCiscoMsgSubtype_AP_CISCO_MSG_SUBTYPE_RESERVED)

		case pb.APMsgType_AP_MSG_TYPE_RESERVED:
			fallthrough

		default:
			ok = false
		}

		if !ok || rsvd_subtype {
			dbg.Println("Invalid request type:", request.MsgType, "subtype:", request.GetSubtype())
			response.ErrStatus.Status = pb.APErrorStatus_AP_EINVAL
			break
		}

		if response.ErrStatus.Status == pb.APErrorStatus_AP_SUCCESS {
			config.srv = srv
			config.msgtype = request.MsgType
			ok = updatePacketConfig(config)
			if !ok {
				response.ErrStatus.Status = pb.APErrorStatus_AP_EINVAL
				break
			} else {
				listPacketConfig()
			}
		}
	}

	/* send response to config request */
	response.PacketHdr = &pb.APPacketHdr{}
	response.PacketHdr.MsgType = pb.APMsgType_AP_MSG_TYPE_RESERVED
	response.PacketLen = 0

	srv.Send(response)

	/* start infinite loop for this client */
	for {
		time.Sleep(time.Second)
		/*
		   open a channel  here
		   if err := srv.Send(hbt); err != nil {
		       return err
		   }
		*/
	}
}

/* open socket to receive packets from kernel */
func OpenCaptureSocket() {
	/* Initialize Global packet config */
	ci = ClientInfo{pconfig: make(map[string]PacketConfig)}

	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, PRIV_PROTOCOL)
	if err != nil {
		fmt.Println("socket listen error", err)
	}

	if_info, err := net.InterfaceByName(INTERFACE_NAME)
	if err != nil {
		fmt.Println("interface not found", INTERFACE_NAME)
	}

	addr := syscall.SockaddrLinklayer{
		Protocol: PRIV_PROTOCOL,
		Ifindex:  if_info.Index,
	}

	err = syscall.Bind(fd, &addr)
	if err != nil {
		fmt.Println("socket bind error", err)
	}

	err = syscall.SetLsfPromisc(INTERFACE_NAME, true)
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
