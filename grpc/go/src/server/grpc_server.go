/*
 *
 * Copyright (c) 2016, Cisco Systems Inc.
 * All rights reserved.
 */

package main

/* standard packages */
import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

/* cheetah packages */
import (
	pb "gengo"
	"server_util"
	"util"
)

const (
	PROC_MEMINFO         = "/proc/meminfo"
	PROC_SLABINFO        = "/proc/slabinfo"
	PROC_SYSTEM_INFO     = "/proc/aptrace/sysinfo/system"
	PROC_CLIENT_INFO     = "/proc/aptrace/sysinfo/clients"
	PROC_WLAN_INFO       = "/proc/aptrace/sysinfo/wlans"
	PROC_RADIO_INFO      = "/proc/aptrace/sysinfo/radios"
	PROC_WIRED_INFO      = "/proc/aptrace/sysinfo/wired"
	FILE_ETC_RESOLV_CONF = "/etc/resolv.conf"
	FILE_MERAKI_SERIAL   = "/MERAKI_SERIAL"
	FILE_PLATFORM_NAME   = "/AP_PLATFORM_NAME"
	FILE_APPHOST_CFG     = "/tmp/apphostcfg"
)

/* Command line arguments */
var (
	debug = flag.Bool("debug", false, "Enable debugging")
)

/* Debug related */
type Debug bool

var dbg Debug

func (d Debug) Println(a ...interface{}) {
	if d {
		fmt.Println(a...)
	}
}

/*
 * APGlobal Server implementation
 */
type GlobalServer struct{}

func (s *GlobalServer) APGlobalInitNotif(request *pb.APInitMsg, srv pb.APGlobal_APGlobalInitNotifServer) error {
	dbg.Println("Received Global init notification, major ver:", request.MajorVer)

	/* Create a APGlobalsGetMsg */
	init_resp := &pb.APGlobalNotif{
		EventType: pb.APGlobalNotifType_AP_GLOBAL_EVENT_TYPE_VERSION,
		ErrStatus: &pb.APErrorStatus{
			Status: pb.APErrorStatus_AP_SUCCESS,
		},
		Event: &pb.APGlobalNotif_InitRspMsg{
			&pb.APInitMsgRsp{
				MajorVer: uint32(pb.APVersion_AP_MAJOR_VERSION),
				MinorVer: uint32(pb.APVersion_AP_MINOR_VERSION),
				SubVer:   uint32(pb.APVersion_AP_SUB_VERSION),
			},
		},
	}

	srv.Send(init_resp)

	// send apphosting config if any
	sendAppHostingConfig(srv)

	// send continuous heartbeats
	sendHeartbeatNotification(srv)

	return nil
}

func (s *GlobalServer) APGlobalsGet(ctx context.Context, in *pb.APGlobalsGetMsg) (*pb.APGlobalsGetMsgRsp, error) {
	dbg.Println("Received Global Get call")

	/* create APGlobalsGetMsgRsp */
	globalAPResp := &pb.APGlobalsGetMsgRsp{}
	globalAPResp.ErrStatus = &pb.APErrorStatus{}
	globalAPResp.ErrStatus.Status = pb.APErrorStatus_AP_SUCCESS
	globalAPResp.MaxRadioNameLength = 16
	globalAPResp.MaxSsidNameLength = 16

	return globalAPResp, nil
}

func sendHeartbeatNotification(srv pb.APGlobal_APGlobalInitNotifServer) error {
	hbt := &pb.APGlobalNotif{}
	hbt.EventType = pb.APGlobalNotifType_AP_GLOBAL_EVENT_TYPE_HEARTBEAT
	hbt.ErrStatus = &pb.APErrorStatus{}
	hbt.ErrStatus.Status = pb.APErrorStatus_AP_SUCCESS
	for {
		time.Sleep(time.Second)
		if err := srv.Send(hbt); err != nil {
			return err
		}
	}

	return nil
}

func sendAppHostingConfig(srv pb.APGlobal_APGlobalInitNotifServer) {
	var port uint64

	appcfg := &pb.APGlobalNotif{}
	appcfg.EventType = pb.APGlobalNotifType_AP_GLOBAL_EVENT_TYPE_CONFIG
	appcfg.ErrStatus = &pb.APErrorStatus{}
	appcfg.ErrStatus.Status = pb.APErrorStatus_AP_SUCCESS

	keys := []string{"token", "proxyurl", "proxyport"}
	sep := "="
	fmap := server_util.GetFieldsFromFile(FILE_APPHOST_CFG, keys, sep)

	if len(fmap) > 0 {
		resp := &pb.APGlobalNotif_CfgRspMsg{}
		resp.CfgRspMsg = &pb.APCfgMsgRsp{}
		resp.CfgRspMsg.Token = fmap[keys[0]]
		resp.CfgRspMsg.ProxyURL = fmap[keys[1]]
		port, _ = strconv.ParseUint(fmap[keys[2]], 10, 32)
		resp.CfgRspMsg.ProxyPort = uint32(port)
		appcfg.Event = resp
	} else {
		appcfg.ErrStatus.Status = pb.APErrorStatus_AP_NOT_AVAILABLE
		fmt.Println("APP config not found")
	}

	srv.Send(appcfg)

	return
}

/*
 * APStatistics Server implementation
 */

type StatsServer struct{}

type fn func() (response *pb.APStatsMsgRsp, err error)

var statsfn = map[pb.APStatsType]fn{
	pb.APStatsType_AP_SYSTEM_STATS:    APSystemStatsGet,
	pb.APStatsType_AP_MEMORY_STATS:    APMemoryStatsGet,
	pb.APStatsType_AP_INTERFACE_STATS: APInterfaceStatsGet,
	pb.APStatsType_AP_ROUTING_STATS:   APRoutingStatsGet,
	pb.APStatsType_AP_DNS_STATS:       APDNSStatsGet,
	pb.APStatsType_AP_RADIO_STATS:     APRadioStatsGet,
	pb.APStatsType_AP_WLAN_STATS:      APWLANStatsGet,
	pb.APStatsType_AP_CLIENT_STATS:    APClientStatsGet,
}

// Get statistics
func (s *StatsServer) APStatsGet(msg *pb.APStatsMsg, srv pb.APStatistics_APStatsGetServer) error {
	for _, request := range msg.StatsRequest {
		/* Get the valid stats function */
		_, ok := statsfn[request.StatsType]
		if !ok {
			fmt.Println("invalid request ", request.StatsType)
			return nil
		}

		// Check for min allowed interval for push
		if (request.TimeInterval != uint32(pb.StatsTimeInterval_AP_STATS_UNARY_OPERATION)) &&
			(request.TimeInterval < uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL)) {
			request.TimeInterval = uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL)
		}

		processAPStatsGet(request, srv)
	}

	return nil
}

// Process stats request
func processAPStatsGet(request *pb.APStatsRequest, srv pb.APStatistics_APStatsGetServer) {
	getstats, _ := statsfn[request.StatsType]
	for {
		stats_resp, err := getstats()
		if err != nil {
			log.Println("stats processing failed : %v", err)
			return
		}
		srv.Send(stats_resp)

		// If one time request, break out
		if request.TimeInterval == uint32(pb.StatsTimeInterval_AP_STATS_UNARY_OPERATION) {
			return
		}

		// Otherwise, loop for specified interval
		time.Sleep(time.Duration(request.TimeInterval) * time.Second)
	}

	return
}

// Get system level statistics
func APSystemStatsGet() (*pb.APStatsMsgRsp, error) {
	keys := []string{"ID", "Uptime"}
	var uptime uint64

	dbg.Println("Received APSystemStatsGet call")

	/* Create system stats response */
	resp := &pb.APStatsMsgRsp_SystemStats{}
	resp.SystemStats = &pb.APSystemStatsMsgRsp{}

	sep := ":"
	fmap := server_util.GetFieldsFromFile(PROC_SYSTEM_INFO, keys, sep)
	if len(fmap) > 0 {
		resp.SystemStats.ID = fmap[keys[0]]
		uptime, _ = strconv.ParseUint(fmap[keys[1]], 10, 32)
	}
	resp.SystemStats.Uptime = uint32(uptime)
	resp.SystemStats.When = time.Now().String()
	resp.SystemStats.SerialNumber =
		strings.Trim(server_util.ReadFileAsString(FILE_MERAKI_SERIAL), "\n")
	resp.SystemStats.ProductId =
		strings.Trim(server_util.ReadFileAsString(FILE_PLATFORM_NAME), "\n")

	/* Create response message */
	m := &pb.APStatsMsgRsp{}
	m.ErrStatus = &pb.APErrorStatus{}
	m.ErrStatus.Status = pb.APErrorStatus_AP_SUCCESS
	m.MsgRsp = resp

	return m, nil
}

func getMemInfo() *pb.MemInfo {
	keys := []string{"MemTotal", "MemFree"}
	var memval uint64
	meminfo := &pb.MemInfo{}

	sep := ":"
	fmap := server_util.GetFieldsFromFile(PROC_MEMINFO, keys, sep)
	if len(fmap) > 0 {
		memval, _ = strconv.ParseUint(strings.Split(fmap[keys[0]], " ")[0], 10, 32)
		meminfo.TotalKB = uint32(memval)
		memval, _ = strconv.ParseUint(strings.Split(fmap[keys[1]], " ")[0], 10, 32)
		meminfo.AvailableKB = uint32(memval)
	}

	return meminfo
}

func getTopSlabInfo() (*pb.SlabInfo, pb.APErrorStatus_APErrno) {
	slabinfo := &pb.SlabInfo{}
	var mostActiveObjs, temp uint64
	var topList, slab, keys []string
	var sline string
	var index int
	var errno error
	var retval pb.APErrorStatus_APErrno

	if file, err := os.Open(PROC_SLABINFO); err == nil {
		// make sure it gets closed
		defer file.Close()

		mostActiveObjs = 0
		// create a new reader and read the file line by line
		bf := bufio.NewReader(file)
		for {
			line, isPrefix, errno := bf.ReadLine()

			// loop termination - EOF
			if errno == io.EOF {
				break
			}

			// loop termination - error
			if errno != nil {
				log.Println("error in reading file err", file.Name(), errno)
				break
			}

			// loop termination - error
			if isPrefix {
				log.Println("Error: Unexpected long line reading", file.Name())
				break
			}

			sline = string(line)
			// skip line version
			if strings.HasPrefix(sline, "slabinfo") {
				continue
			}

			if strings.HasPrefix(sline, "#") {
				keys = strings.Fields(sline)[1:]
				continue
			}

			slab = strings.Fields(sline)
			temp, _ = strconv.ParseUint(slab[1], 10, 32)
			if temp > mostActiveObjs {
				mostActiveObjs = temp
				topList = slab
			}
		}

		if errno == nil || errno == io.EOF {
			index = server_util.GetFieldIndex(keys, "name")
			if index != -1 {
				slabinfo.Name = topList[index]
			}
			index = server_util.GetFieldIndex(keys, "<objsize>")
			if index != -1 {
				temp, _ = strconv.ParseUint(topList[index], 10, 32)
				slabinfo.ObjSize = int32(temp)
			}
			index = server_util.GetFieldIndex(keys, "<active_objs>")
			if index != -1 {
				temp, _ = strconv.ParseUint(topList[index], 10, 32)
				slabinfo.ActiveObjs = int32(temp)
			}
			index = server_util.GetFieldIndex(keys, "<num_objs>")
			if index != -1 {
				temp, _ = strconv.ParseUint(topList[index], 10, 32)
				slabinfo.NumObjs = int32(temp)
			}
			retval = pb.APErrorStatus_AP_SUCCESS
		} else {
			retval = pb.APErrorStatus_AP_NOT_AVAILABLE
		}
	} else {
		retval = pb.APErrorStatus_AP_EINVAL
	}

	return slabinfo, retval
}

// Get memory statistics
func APMemoryStatsGet() (*pb.APStatsMsgRsp, error) {
	dbg.Println("Received APMemoryStatsGet call")
	var err pb.APErrorStatus_APErrno

	/* Create memory stats response */
	resp := &pb.APStatsMsgRsp_MemoryStats{}
	resp.MemoryStats = &pb.APMemoryStatsMsgRsp{}

	/* fill memory info */
	resp.MemoryStats.ProcMemInfo = &pb.MemInfo{}
	resp.MemoryStats.ProcMemInfo = getMemInfo()

	/* fill slab info */
	resp.MemoryStats.TopProcSlabInfo = &pb.SlabInfo{}
	resp.MemoryStats.TopProcSlabInfo, err = getTopSlabInfo()

	/* Create response message */
	m := &pb.APStatsMsgRsp{}
	m.ErrStatus = &pb.APErrorStatus{}
	m.ErrStatus.Status = err
	if m.ErrStatus.Status == pb.APErrorStatus_AP_SUCCESS {
		m.MsgRsp = resp
	}

	return m, nil
}

func getInterfaceStats(ifnames []string) ([]*pb.APInterfaceEntry, int) {
	var interfaces []*pb.APInterfaceEntry
	var info *pb.APInterfaceEntry
	var record_count int
	var sval uint64

	keys := []string{"Link status", "Port speed", "Port duplex"}
	sep := ":"
	fmap := server_util.GetFieldsFromFile(PROC_WIRED_INFO, keys, sep)

	if len(fmap) > 0 {
		record_count = 0
		for _, iname := range ifnames {
			info = new(pb.APInterfaceEntry)

			info.Link = (strings.Compare(fmap[keys[0]], "up") == 0)
			info.FullDuplex = (strings.Compare(fmap[keys[2]], "full") == 0)
			intval, err := strconv.Atoi(strings.Trim(fmap[keys[1]], "G"))
			if err == nil {
				info.Speed = uint32(intval)
			}

			info.RxBytes, _ = server_util.GetIfStats(iname, "rx_bytes")
			sval, _ = server_util.GetIfStats(iname, "rx_packets")
			info.RxPkts = uint32(sval)
			sval, _ = server_util.GetIfStats(iname, "rx_dropped")
			info.RxDiscards = uint32(sval)
			info.RxBytes, _ = server_util.GetIfStats(iname, "tx_bytes")
			sval, _ = server_util.GetIfStats(iname, "tx_packets")
			info.TxPkts = uint32(sval)

			record_count++
			interfaces = append(interfaces, info)
		}
	}

	return interfaces, record_count
}

// Get Interface statistics
func APInterfaceStatsGet() (*pb.APStatsMsgRsp, error) {
	ifnames := []string{"wired0"}

	dbg.Println("Received APInterfaceStatsGet call")

	interfaces, cnt := getInterfaceStats(ifnames)

	/* Create Interface stats response */
	resp := &pb.APStatsMsgRsp_InterfaceStats{}
	resp.InterfaceStats = &pb.APInterfaceStatsMsgRsp{}

	m := &pb.APStatsMsgRsp{}
	m.ErrStatus = &pb.APErrorStatus{}
	if cnt > 0 {
		m.ErrStatus.Status = pb.APErrorStatus_AP_SUCCESS
		resp.InterfaceStats.Interfaces = []*pb.APInterfaceEntry{}
		resp.InterfaceStats.Interfaces = interfaces
		m.MsgRsp = resp
	} else {
		m.ErrStatus.Status = pb.APErrorStatus_AP_NOT_AVAILABLE
	}

	return m, nil
}

func fillRoute(keys []string, values []string) *pb.IPv4Route {
	var intval int

	route := new(pb.IPv4Route)
	route.Destination = values[server_util.GetFieldIndex(keys, "Destination")]
	route.Gateway = values[server_util.GetFieldIndex(keys, "Gateway")]
	route.Genmask = values[server_util.GetFieldIndex(keys, "Genmask")]
	route.Flags = values[server_util.GetFieldIndex(keys, "Flags")]
	intval, _ = strconv.Atoi(values[server_util.GetFieldIndex(keys, "Metric")])
	route.Metric = uint32(intval)
	intval, _ = strconv.Atoi(values[server_util.GetFieldIndex(keys, "Ref")])
	route.Ref = uint32(intval)
	intval, _ = strconv.Atoi(values[server_util.GetFieldIndex(keys, "Use")])
	route.Use = uint32(intval)
	route.Iface = values[server_util.GetFieldIndex(keys, "Iface")]

	return route
}

func getRoutes() ([]*pb.IPv4Route, int) {
	var routes []*pb.IPv4Route
	var record_count int
	var keys, values []string

	out, err := exec.Command("route").Output()
	if err != nil {
		log.Println("Error in reading routes", err)
		return nil, 0
	}

	record_count = 0
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "Kernel") {
			continue
		}

		if strings.HasPrefix(line, "Destination") {
			keys = strings.Fields(line)
			continue
		}

		values = strings.Fields(line)
		routes = append(routes, fillRoute(keys, values))
		record_count++
	}

	return routes, record_count
}

// Get Routing statistics
func APRoutingStatsGet() (*pb.APStatsMsgRsp, error) {
	dbg.Println("Received APRoutingStatsGet call")

	routes, cnt := getRoutes()

	/* Create Routing stats response */
	resp := &pb.APStatsMsgRsp_RoutingStats{}
	resp.RoutingStats = &pb.APRoutingStatsMsgRsp{}

	/* Create response message */
	m := &pb.APStatsMsgRsp{}
	m.ErrStatus = &pb.APErrorStatus{}
	if cnt > 0 {
		m.ErrStatus.Status = pb.APErrorStatus_AP_SUCCESS
		resp.RoutingStats.IPv4Routes = []*pb.IPv4Route{}
		resp.RoutingStats.IPv4Routes = routes
		m.MsgRsp = resp
	} else {
		m.ErrStatus.Status = pb.APErrorStatus_AP_NOT_AVAILABLE
	}

	return m, nil
}

// Get DNS servers
func APDNSStatsGet() (*pb.APStatsMsgRsp, error) {
	var record_count int
	var sline string
	var iplist []string
	var retval pb.APErrorStatus_APErrno

	dbg.Println("Received APDNSStatsGet call")

	/* Create DNS stats response */
	resp := &pb.APStatsMsgRsp_DNSStats{}
	resp.DNSStats = &pb.APDNSStatsMsgRsp{}

	if file, err := os.Open(FILE_ETC_RESOLV_CONF); err == nil {
		// make sure it gets closed
		defer file.Close()

		record_count = 0
		// create a new reader and read the file line by line
		bf := bufio.NewReader(file)
		for {
			line, isPrefix, errno := bf.ReadLine()

			// loop termination - EOF
			if errno == io.EOF {
				break
			}

			// loop termination - error
			if errno != nil {
				log.Println("error in reading file err", file.Name(), errno)
				break
			}

			// loop termination - error
			if isPrefix {
				log.Println("Error: Unexpected long line reading", file.Name())
				break
			}

			sline = string(line)
			// skip line version
			if strings.HasPrefix(sline, "nameserver") {
				iplist = append(iplist, strings.Fields(sline)[1])
				record_count++
			}
		}

		if record_count > 0 {
			retval = pb.APErrorStatus_AP_SUCCESS
		} else {
			retval = pb.APErrorStatus_AP_NOT_AVAILABLE
		}
	} else {
		retval = pb.APErrorStatus_AP_EINVAL
	}

	/* Create response message */
	m := &pb.APStatsMsgRsp{}
	m.ErrStatus = &pb.APErrorStatus{}
	m.ErrStatus.Status = retval
	if m.ErrStatus.Status == pb.APErrorStatus_AP_SUCCESS {
		resp.DNSStats.IP = iplist
		m.MsgRsp = resp
	}

	return m, nil
}

func fillRadioEntry(radio *pb.RadioEntry, values map[string]string) {
	var val uint64
	var int_val int64
	var float_val float64

	switch option := values["DATATYPE"]; option {
	case "simple":
		radio.Dev = values["Dev"]
		radio.Band = values["Band"]
		val, _ = strconv.ParseUint(values["Channel"], 10, 32)
		radio.Channel = uint32(val)
		val, _ = strconv.ParseUint(values["SecondaryChannel"], 10, 32)
		radio.SecondaryChannel = uint32(val)
		val, _ = strconv.ParseUint(values["Bandwidth"], 10, 32)
		radio.Bandwidth = uint32(val)
		int_val, _ = strconv.ParseInt(values["NoiseFloor"], 10, 32)
		radio.NoiseFloor = int32(int_val)
		val, _ = strconv.ParseUint(values["MaxTxPower"], 10, 32)
		radio.MaxTxPower = uint32(val)

	case "Utilization":
		radio.Utilization = &pb.RadioUtilization{}
		float_val, _ = strconv.ParseFloat(values["All"], 32)
		radio.Utilization.All = float32(float_val)
		float_val, _ = strconv.ParseFloat(values["Tx"], 32)
		radio.Utilization.Tx = float32(float_val)
		float_val, _ = strconv.ParseFloat(values["RxInBSS"], 32)
		radio.Utilization.RxInBSS = float32(float_val)
		float_val, _ = strconv.ParseFloat(values["RxOtherBSS"], 32)
		radio.Utilization.RxOtherBSS = float32(float_val)
		float_val, _ = strconv.ParseFloat(values["NonWifi"], 32)
		radio.Utilization.NonWifi = float32(float_val)

	case "PerAntennaRSSI":
		rssi_values := strings.Split(strings.Trim(values["PerAntennaRSSI"], "[]"), ",")
		for _, v := range rssi_values {
			int_val, _ = strconv.ParseInt(strings.Trim(v, " "), 10, 32)
			radio.AntennaRSSI = append(radio.AntennaRSSI, int32(int_val))
		}

	case "Counter":
		radio.Counter = &pb.RadioCounters{}
		radio.Counter.TxBytes, _ = strconv.ParseUint(values["TxBytes"], 10, 32)
		val, _ = strconv.ParseUint(values["TxPkts"], 10, 32)
		radio.Counter.TxPkts = uint32(val)
		val, _ = strconv.ParseUint(values["TxMgmt"], 10, 32)
		radio.Counter.TxMgmt = uint32(val)
		val, _ = strconv.ParseUint(values["TxErrors"], 10, 32)
		radio.Counter.TxErrors = uint32(val)
		radio.Counter.RxBytes, _ = strconv.ParseUint(values["RxBytes"], 10, 32)
		val, _ = strconv.ParseUint(values["RxPkts"], 10, 32)
		radio.Counter.RxPkts = uint32(val)
		val, _ = strconv.ParseUint(values["RxMgmt"], 10, 32)
		radio.Counter.RxMgmt = uint32(val)
		val, _ = strconv.ParseUint(values["RxErrors"], 10, 32)
		radio.Counter.RxErrors = uint32(val)

	case "DFS":
		radio.DFS = &pb.DfsState{}
		val, _ = strconv.ParseUint(values["CacState"], 10, 32)
		radio.DFS.CacState = uint32(val)
		radio.DFS.RadarDetected, _ = strconv.ParseBool(values["RadarDetected"])

	default:
		fmt.Println("Invalid radio stat type: ", option)
	}

	// reset map
	for key := range values {
		delete(values, key)
	}

	return
}

func getRadioEntries() ([]*pb.RadioEntry, pb.APErrorStatus_APErrno) {
	var entries []*pb.RadioEntry
	var radio *pb.RadioEntry
	var sline string
	var fmap map[string]string
	var record_count, index int
	var retval pb.APErrorStatus_APErrno

	simple_keys := []string{
		"Dev",
		"Band",
		"Channel",
		"SecondaryChannel",
		"Bandwidth",
		"NoiseFloor",
		"MaxTxPower",
	}

	mline_keys := []string{
		"Utilization",
		"PerAntennaRSSI",
		"Counter",
		"DFS",
	}

	record_count = 0
	if file, err := os.Open(PROC_RADIO_INFO); err == nil {
		// make sure it gets closed
		defer file.Close()

		// create a new reader and read the file line by line
		fmap = make(map[string]string)
		bf := bufio.NewReader(file)
		for {
			line, isPrefix, errno := bf.ReadLine()

			// loop termination - EOF
			if errno == io.EOF {
				break
			}

			// loop termination - error
			if errno != nil {
				log.Println("error in reading file err", file.Name(), errno)
				break
			}

			// loop termination - error
			if isPrefix {
				log.Println("Error: Unexpected long line reading", file.Name())
				break
			}

			sline = strings.TrimSpace(string(line))
			if strings.HasPrefix(sline, "Radios Information:") || (sline == "") {
				continue
			}

			if strings.HasPrefix(sline, "radio num:") {
				if record_count > 0 {
					fillRadioEntry(radio, fmap)
					entries = append(entries, radio)
				}
				radio = new(pb.RadioEntry)
				fmap["DATATYPE"] = "simple"
				record_count++
				continue
			}

			values := strings.SplitN(sline, ":", 2)
			// trim leading and trailing whitespaces
			values[0] = strings.TrimSpace(values[0])
			values[1] = strings.TrimSpace(values[1])

			index = server_util.GetFieldIndex(simple_keys, values[0])
			if index != -1 {
				if fmap["DATATYPE"] == mline_keys[len(mline_keys)-1] {
					fillRadioEntry(radio, fmap)
				}
				fmap[values[0]] = values[1]
			} else {
				index = server_util.GetFieldIndex(mline_keys, values[0])
				if index != -1 {
					if len(fmap) > 0 {
						fillRadioEntry(radio, fmap)
					}
					/* start new data type */
					fmap["DATATYPE"] = values[0]
					/* handle special case - value on same line */
					if fmap["DATATYPE"] == "PerAntennaRSSI" {
						fmap[values[0]] = values[1]
					}
				} else {
					fmap[values[0]] = values[1]
				}
			}
		}

		if len(fmap) > 0 {
			fillRadioEntry(radio, fmap)
			entries = append(entries, radio)
			retval = pb.APErrorStatus_AP_SUCCESS
		} else {
			retval = pb.APErrorStatus_AP_NOT_AVAILABLE
		}
	} else {
		retval = pb.APErrorStatus_AP_EINVAL
	}

	return entries, retval
}

// Get radio statistics
func APRadioStatsGet() (*pb.APStatsMsgRsp, error) {
	dbg.Println("Received APRadioStatsGet call")
	var err pb.APErrorStatus_APErrno

	radios, err := getRadioEntries()

	/* Create Radio stats response */
	resp := &pb.APStatsMsgRsp_RadioStats{}
	resp.RadioStats = &pb.APRadioStatsMsgRsp{}

	/* Create response message */
	m := &pb.APStatsMsgRsp{}
	m.ErrStatus = &pb.APErrorStatus{}
	m.ErrStatus.Status = err
	if m.ErrStatus.Status == pb.APErrorStatus_AP_SUCCESS {
		resp.RadioStats.Radios = []*pb.RadioEntry{}
		resp.RadioStats.Radios = radios
		m.MsgRsp = resp
	}

	return m, nil
}

func fillWlanEntry(values map[string]string) *pb.WLANEntry {
	var val uint64

	wlan := new(pb.WLANEntry)
	wlan.Wlan = &pb.WLAN{}
	wlan.Wlan.ID = values["ID"]
	wlan.Wlan.SSID = values["SSID"]
	val, _ = strconv.ParseUint(values["RadioIndex"], 10, 32)
	wlan.RadioIndex = uint32(val)
	wlan.BSSID = values["BSSID"]
	wlan.Dev = values["Dev"]
	val, _ = strconv.ParseUint(values["NumClients"], 10, 32)
	wlan.NumClients = int32(val)
	wlan.Counter = &pb.MulticastCounter{}
	val, _ = strconv.ParseUint(values["TxMcastPkts"], 10, 32)
	wlan.Counter.TxMcastPkts = uint32(val)
	wlan.Counter.TxMcastBytes, _ = strconv.ParseUint(values["TxMcastBytes"], 10, 32)

	return wlan
}

func getWlanEntries() ([]*pb.WLANEntry, pb.APErrorStatus_APErrno) {
	var entries []*pb.WLANEntry
	var wlan *pb.WLANEntry
	var sline string
	var fmap map[string]string
	var record_count int
	var retval pb.APErrorStatus_APErrno

	record_count = 0
	if file, err := os.Open(PROC_WLAN_INFO); err == nil {
		// make sure it gets closed
		defer file.Close()

		// create a new reader and read the file line by line
		fmap = make(map[string]string)
		bf := bufio.NewReader(file)
		for {
			line, isPrefix, errno := bf.ReadLine()

			// loop termination - EOF
			if errno == io.EOF {
				break
			}

			// loop termination - error
			if errno != nil {
				log.Println("error in reading file err", file.Name(), errno)
				break
			}

			// loop termination - error
			if isPrefix {
				log.Println("Error: Unexpected long line reading", file.Name())
				break
			}

			sline = string(line)
			if strings.HasPrefix(sline, "Wlans Information") {
				continue
			}

			if strings.TrimSpace(sline) == "" {
				continue
			}

			if strings.HasPrefix(sline, "wlan num:") {
				if record_count > 0 {
					wlan = fillWlanEntry(fmap)
					entries = append(entries, wlan)
				}
				/* reset map */
				for key := range fmap {
					delete(fmap, key)
				}
				record_count++
				continue
			}

			values := strings.SplitN(sline, ":", 2)
			fmap[strings.TrimSpace(values[0])] = strings.TrimSpace(values[1])
		}

		if len(fmap) > 0 {
			wlan = fillWlanEntry(fmap)
			entries = append(entries, wlan)
			retval = pb.APErrorStatus_AP_SUCCESS
		} else {
			retval = pb.APErrorStatus_AP_NOT_AVAILABLE
		}
	} else {
		retval = pb.APErrorStatus_AP_EINVAL
	}

	return entries, retval
}

// Get WLAN statistics
func APWLANStatsGet() (*pb.APStatsMsgRsp, error) {
	dbg.Println("Received APWLANStatsGet call")
	var err pb.APErrorStatus_APErrno

	wlans, err := getWlanEntries()

	/* Create Wlan stats response */
	resp := &pb.APStatsMsgRsp_WLANStats{}
	resp.WLANStats = &pb.APWLANStatsMsgRsp{}

	/* Create response message */
	m := &pb.APStatsMsgRsp{}
	m.ErrStatus = &pb.APErrorStatus{}
	m.ErrStatus.Status = err
	if m.ErrStatus.Status == pb.APErrorStatus_AP_SUCCESS {
		resp.WLANStats.WLANEntries = []*pb.WLANEntry{}
		resp.WLANStats.WLANEntries = wlans
		m.MsgRsp = resp
	}

	return m, nil
}

func fillClientEntry(values map[string]string) *pb.APClientEntry {
	var val uint64
	var int_val int64

	client := new(pb.APClientEntry)
	client.MAC = values["MAC"]
	val, _ = strconv.ParseUint(values["RadioIndex"], 10, 32)
	client.RadioIndex = uint32(val)
	client.Band = values["Band"]

	client.Wlan = &pb.WLAN{}
	client.Wlan.ID = values["SSID"]
	client.Wlan.SSID = values["BSSID"]

	val, _ = strconv.ParseUint(values["ConnectedTimeSec"], 10, 32)
	client.ConnectedTimeSec = uint32(val)
	val, _ = strconv.ParseUint(values["InactiveTimeMilliSec"], 10, 32)
	client.InactiveTimeMilliSec = uint32(val)
	val, _ = strconv.ParseUint(values["RSSI"], 10, 32)
	client.RSSI = int32(val)
	val, _ = strconv.ParseUint(values["NF"], 10, 32)
	client.NF = int32(val)

	rssi_values := strings.Split(strings.Trim(values["PerAntennaRSSI"], "[]"), ",")
	for _, v := range rssi_values {
		int_val, _ = strconv.ParseInt(strings.Trim(v, " "), 10, 32)
		client.AntennaRSSI = append(client.AntennaRSSI, int32(int_val))
	}

	val, _ = strconv.ParseUint(values["TxBitRate"], 10, 32)
	client.TxBitRate = int32(val)
	client.TxUnicastBytes, _ = strconv.ParseUint(values["TxUnicastBytes"], 10, 32)
	val, _ = strconv.ParseUint(values["TxUnicastPkts"], 10, 32)
	client.TxUnicastPkts = uint32(val)
	client.RxBytes, _ = strconv.ParseUint(values["RxBytes"], 10, 32)
	val, _ = strconv.ParseUint(values["RxPkts"], 10, 32)
	client.RxPkts = uint32(val)

	return client
}

func getClientEntries() ([]*pb.APClientEntry, pb.APErrorStatus_APErrno) {
	var entries []*pb.APClientEntry
	var client *pb.APClientEntry
	var sline string
	var fmap map[string]string
	var record_count int
	var retval pb.APErrorStatus_APErrno

	record_count = 0
	if file, err := os.Open(PROC_CLIENT_INFO); err == nil {
		// make sure it gets closed
		defer file.Close()

		// create a new reader and read the file line by line
		fmap = make(map[string]string)
		bf := bufio.NewReader(file)
		for {
			line, isPrefix, errno := bf.ReadLine()

			// loop termination - EOF
			if errno == io.EOF {
				break
			}

			// loop termination - error
			if errno != nil {
				log.Println("error in reading file err", file.Name(), errno)
				break
			}

			// loop termination - error
			if isPrefix {
				log.Println("Error: Unexpected long line reading", file.Name())
				break
			}

			sline = string(line)
			if strings.HasPrefix(sline, "Clients:") {
				continue
			}

			if strings.TrimSpace(sline) == "" {
				continue
			}

			if strings.HasPrefix(sline, "client num:") {
				if record_count > 0 {
					client = fillClientEntry(fmap)
					entries = append(entries, client)
				}
				/* reset map */
				for key := range fmap {
					delete(fmap, key)
				}
				record_count++
				continue
			}

			values := strings.SplitN(sline, ":", 2)
			fmap[strings.TrimSpace(values[0])] = strings.TrimSpace(values[1])
		}

		if len(fmap) > 0 {
			client = fillClientEntry(fmap)
			entries = append(entries, client)
			retval = pb.APErrorStatus_AP_SUCCESS
		} else {
			retval = pb.APErrorStatus_AP_NOT_AVAILABLE
		}
	} else {
		retval = pb.APErrorStatus_AP_EINVAL
	}

	return entries, retval
}

// Get Client statistics
func APClientStatsGet() (*pb.APStatsMsgRsp, error) {
	dbg.Println("Received APClientStatsGet call")
	var err pb.APErrorStatus_APErrno

	clients, err := getClientEntries()

	/* Create Client stats response */
	resp := &pb.APStatsMsgRsp_ClientStats{}
	resp.ClientStats = &pb.APClientStatsMsgRsp{}

	/* Create response message */
	m := &pb.APStatsMsgRsp{}
	m.ErrStatus = &pb.APErrorStatus{}
	m.ErrStatus.Status = err
	if m.ErrStatus.Status == pb.APErrorStatus_AP_SUCCESS {
		resp.ClientStats.Clients = []*pb.APClientEntry{}
		resp.ClientStats.Clients = clients
		m.MsgRsp = resp
	}

	return m, nil
}

func main() {
	/* parse any command line arguments */
	flag.Parse()

	if *debug {
		dbg = true
	}

	/* get server IP adn Port from Env */
	server, port := util.GetServerIPPort()

	lis, err := net.Listen("tcp", server+":"+port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	/* Create the server */
	grpcServer := grpc.NewServer()

	/* Add APGlobal server */
	pb.RegisterAPGlobalServer(grpcServer, &GlobalServer{})

	/* Add APStatistics server */
	pb.RegisterAPStatisticsServer(grpcServer, &StatsServer{})

	/* Add APPackets server */
	pb.RegisterAPPacketsServer(grpcServer, &PacketsServer{})

	/* open socket to listen packets from kernel */
	go OpenCaptureSocket()

	/* serve requests */
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start global server: %v", err)
	}
}
