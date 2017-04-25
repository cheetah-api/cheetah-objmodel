#
# Copyright (c) 2017 by Cisco Systems, Inc.
# All rights reserved.
#

# Standard python libs
import os
import sys
import socket
import threading
import grpc
import time
import subprocess, re

# Add the generated python bindings to the path
sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.realpath(__file__))))

# gRPC generated python bindings
#import pdb
#pdb.set_trace()

from genpy import ap_global_pb2
from genpy import ap_stats_pb2
from genpy import ap_common_types_pb2
from genpy import ap_version_pb2

from concurrent import futures
from datetime import datetime

# gRPC libs
from grpc.beta import implementations

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

ext_table = ["none", "above", "auto", "below"]

# Global data
MAX_RADIO = 2
PROC_MEMINFO = '/proc/meminfo'
PROC_SLABINFO = '/proc/slabinfo'
ETC_RESOLV_CONF = '/etc/resolv.conf'

PROC_SYSTEM_INFO = '/proc/aptrace/sysinfo/system'
PROC_CLIENT_INFO = '/proc/aptrace/sysinfo/clients'
PROC_WLAN_INFO = '/proc/aptrace/sysinfo/wlans'
PROC_RADIO_INFO = '/proc/aptrace/sysinfo/radios'
PROC_WIRED_INFO = '/proc/aptrace/sysinfo/wired'
CLIENT_IP_TABLE = '/click/client_ip_table/list'

#
# Access Point Global functions
#
class APGlobal ():

  def APGlobalInitNotif(self, request, context):

    print "Received Global init notification, major ver:", request.MajorVer

    init_resp=ap_global_pb2.APGlobalNotif()
    init_resp.EventType=ap_global_pb2.AP_GLOBAL_EVENT_TYPE_VERSION
    init_resp.ErrStatus.Status=ap_common_types_pb2.APErrorStatus.AP_SUCCESS
    init_resp.InitRspMsg.MajorVer = ap_version_pb2.AP_MAJOR_VERSION
    init_resp.InitRspMsg.MinorVer = ap_version_pb2.AP_MINOR_VERSION
    init_resp.InitRspMsg.SubVer   = ap_version_pb2.AP_SUB_VERSION

    for i in range(100):
        yield(init_resp)
        init_resp.EventType=ap_global_pb2.AP_GLOBAL_EVENT_TYPE_HEARTBEAT
        time.sleep(10)


  def APGlobalsGet(self, request, context):
    print "Received GlobalsGet response"

    get_resp=ap_global_pb2.APGlobalsGetMsgRsp()
    get_resp.ErrStatus.Status=ap_common_types_pb2.APErrorStatus.AP_SUCCESS
    get_resp.MaxRadioNameLength=16
    get_resp.MaxSsidNameLength=16

    return (get_resp)


#
# Access Point Statistics
#
def get_meminfo(meminfo):
    try:
        f = open(PROC_MEMINFO, 'r')
        for line in f:
            if line.startswith('MemTotal'):
                meminfo.Total_kB = int(line.split()[1])
            elif line.startswith('MemFree'):
                meminfo.Available_kB = int(line.split()[1])
            else:
                continue
        f.close()
    except Exception as e:
        print str(e)

def get_slabinfo(slabinfo):
    try:
        f = open(PROC_SLABINFO, 'r')

        mostActiveObjs = 0
        topList = []
        for line in f:
            # skip line version
            if line.startswith('slabinfo'):
                continue

            if line.startswith('#'):
                keys = line.split()[1:]
                continue

            slab = line.split()

            if int(slab[1]) > mostActiveObjs:
                mostActiveObjs = int(slab[1])
                topList = slab

        slabinfo.Name = topList[keys.index('name')]
        slabinfo.ObjSize = int(topList[keys.index('<objsize>')])
        slabinfo.ActiveObjs = int(topList[keys.index('<active_objs>')])
        slabinfo.NumObjs = int(topList[keys.index('<num_objs>')])

        f.close()

    except Exception as e:
        print str(e)


def get_wired_info(resp):
    try:
        f = open(PROC_WIRED_INFO, 'r')
        for line in f:
            if ("Link status") in line and ("up") in line:
                resp.Link = True
            if ("duplex") in line and ("full") in line:
                resp.FullDuplex = True
            if ("speed") in line:
                resp.Speed = 0 #line.split(":")[1].strip()
        f.close()

    except Exception as e:
        print str(e)

def get_interface_stats(interface, module):
    head, sep, tail = module.partition("RX bytes:")
    interface.RxBytes = int(tail.split()[0])

    head, sep, tail = module.partition("RX packets:")
    interface.RxPkts = int(tail.split()[0])

    interface.RxDiscards = int(tail.split("dropped:")[1].split()[0])

    head, sep, tail = module.partition("TX bytes:")
    interface.TxBytes = int(tail.split()[0])

    head, sep, tail = module.partition("TX packets:")

def get_wlan_stats(wlan_info, fields):
    wlan_info.Wlan.ID = fields['ID']
    wlan_info.Wlan.SSID = fields['SSID']
    wlan_info.RadioIndex = int(fields['RadioIndex'])
    wlan_info.BSSID = fields['BSSID']
    wlan_info.Dev = fields['Dev']
    wlan_info.NumClients = int(fields['NumClients'])
    wlan_info.Counter.TxMcastPkts = int(fields['TxMcastPkts'])
    wlan_info.Counter.TxMcastBytes = int(fields['TxMcastBytes'])

def get_client_stats(client_info, fields):
    client_info.MAC = fields['MAC']
    client_info.RadioIndex = int(fields['RadioIndex'])
    client_info.Band = fields['Band']
    client_info.Wlan.ID = fields['SSID']
    client_info.Wlan.SSID = fields['BSSID']
    client_info.ConnectedTimeSec = int(fields['ConnectedTimeSec'])
    client_info.InactiveTimeMilliSec = int(fields['InactiveTimeMilliSec'])
    client_info.RSSI = int(fields['RSSI'])       # change to int
    client_info.NF = int(fields['NF'])       # change to int
    strval = fields['PerAntennaRSSI']
    rlist = strval[1:-1].split(',')
    for val in rlist:
        client_info.AntennaRSSI.append(int(val.strip()))
    client_info.TxBitRate = int(fields['TxBitRate'])
    client_info.TxUnicastBytes = int(fields['TxUnicastBytes'])
    client_info.TxUnicastPkts = int(fields['TxUnicastPkts'])
    client_info.RxBytes = int(fields['RxBytes'])
    client_info.RxPkts = int(fields['RxPkts'])

#
#==============================================
# APStatistics service implementation
#==============================================
#
class APStatistics ():

#
# APSystemsStatsGet
#
  def APSystemStatsGet(self, request, context):
    print "Received system stats get request"

    resp=ap_stats_pb2.APSystemStatsMsgRsp()

    try:
        f = open(PROC_SYSTEM_INFO, 'r')
        for line in f:
            # skip empty lines
            if line.strip() == '':
                continue
            if line.startswith('System Information'):
                continue
            values = line.split(':', 1)
            if values[0].strip() == "ID":
                resp.ID = values[1].strip()
            elif values[0].strip() == "Uptime":
                resp.Uptime = int(values[1].strip())
            else:
                continue
        resp.ErrStatus.Status=ap_common_types_pb2.APErrorStatus.AP_SUCCESS
        f.close()
    except Exception as e:
        resp.ErrStatus.Status=ap_common_types_pb2.APErrorStatus.AP_NOT_AVAILABLE
        print str(e)

    resp.When = str(datetime.now())
    resp.SerialNumber = open('/MERAKI_SERIAL', 'r').read()
    resp.ProductId = open('/AP_PLATFORM_NAME', 'r').read()

    return (resp)

#
# APMemoryStatsGet
#
  def APMemoryStatsGet(self, request, context):
    print "Received memory stats get request"

    resp=ap_stats_pb2.APMemoryStatsMsgRsp()
    resp.ErrStatus.Status=ap_common_types_pb2.APErrorStatus.AP_SUCCESS

    # MemInfo
    get_meminfo(resp.ProcMemInfo)

    # SlabInfo
    get_slabinfo(resp.TopProcSlabInfo)

    return (resp)

#
# APDNSStatsGet
#
  def APDNSStatsGet(self, request, context):
    print "Received DNS stats get request"

    resp=ap_stats_pb2.APDNSServersMsgRsp()

    try:
        f = open(ETC_RESOLV_CONF, 'r')
        for line in f:
            if line.startswith('nameserver'):
                resp.IP.append(line.split()[1])
        f.close()
        resp.ErrStatus.Status=ap_common_types_pb2.APErrorStatus.AP_SUCCESS

    except Exception as e:
        resp.ErrStatus.Status=ap_common_types_pb2.APErrorStatus.AP_NOT_AVAILABLE
        print str(e)

    return (resp)

#
# APRoutesStatsGet
#
  def APRoutesStatsGet(self, request, context):
    print "Received Route stats get request"

    resp=ap_stats_pb2.APRoutesMsgRsp()

    table = subprocess.Popen('route', shell=True,
                             stdout=subprocess.PIPE).stdout.read()
    lines = table.split('\n')
    for line in lines:
        if line.strip() == '':
            continue
        if line.startswith('Kernel'):
            continue
        if line.startswith('Destination'):
            keys = line.split()
            continue
        values = line.split()

        ipv4_route=resp.IPv4Routes.add()
        ipv4_route.Destination = values[keys.index('Destination')]
        ipv4_route.Gateway = values[keys.index('Gateway')]
        ipv4_route.Genmask = values[keys.index('Genmask')]
        ipv4_route.Flags = values[keys.index('Flags')]
        ipv4_route.Metric = int(values[keys.index('Metric')])
        ipv4_route.Ref = int(values[keys.index('Ref')])
        ipv4_route.Use = int(values[keys.index('Use')])
        ipv4_route.Iface = values[keys.index('Iface')]

    resp.ErrStatus.Status=ap_common_types_pb2.APErrorStatus.AP_SUCCESS

    return (resp)


#
# APRadioStatsGet
#
  def APRadioStatsGet(self, request, context):
    print "Received Radio stats get request"

    import server_util

    resp = ap_stats_pb2.APRadioStatsMsgRsp()
    record_count = 0

    try:
        f = open(PROC_RADIO_INFO, 'r')
        for line in f:
            # skip empty lines
            if line.strip() == '':
                continue

            # skip first line of each record
            if line.startswith('radio num'):
                if record_count > 0:
                    radioDFS = radio_info.DFS
                    radioDFS.CacState  = int(dfs['CacState'])
                    radioDFS.RadarDetected  = (dfs['RadarDetected'] == 'TRUE')
                radio_info = resp.Radios.add()
                record_count += 1
                util_flag = False
                counter_flag = False
                dfs_flag = False
                continue

            # skips descriptions lines before the first record
            if record_count == 0:
                continue

            values = line.split(':', 1)
            for i in range(len(values)):
                values[i] = values[i].strip()

            if values[0] == "Dev":
                radio_info.Dev = values[1]
            elif values[0] == "Band":
                radio_info.Band = values[1]
            elif values[0] == "Channel":
                radio_info.Channel = int(values[1])
            elif values[0] == "SecondaryChannel":
                radio_info.SecondaryChannel = ext_table.index(values[1])
            elif values[0] == "Bandwidth":
                radio_info.Bandwidth = int(values[1])
            elif values[0] == "NoiseFloor":
                radio_info.NoiseFloor = int(values[1])
            elif values[0] == "MaxTxPower":
                radio_info.MaxTxPower = int(values[1])
            elif values[0] == "Utilization":
                util = {}
                util_flag = True
            elif values[0] == "PerAntennaRSSI":
                util_flag = False
                radioUtil = radio_info.Utilization
                radioUtil.All = float(util['All'])
                radioUtil.Tx = float(util['Tx'])
                radioUtil.RxInBSS = float(util['RxInBSS'])
                radioUtil.RxOtherBSS = float(util['RxOtherBSS'])
                radioUtil.NonWifi = float(util['NonWifi'])
                rlist = values[1][1:-1].split(',')
                for val in rlist:
                    radio_info.AntennaRSSI.append(int(val.strip()))
            elif values[0] == "Counter":
                counter = {}
                counter_flag = True
            elif values[0] == "DFS":
                dfs = {}
                dfs_flag = True
                counter_flag = False
                radioCounter = radio_info.Counter
                radioCounter.TxBytes  = int(counter['TxBytes'])
                radioCounter.TxPkts   = int(counter['TxPkts'])
                radioCounter.TxMgmt   = int(counter['TxMgmt'])
                radioCounter.TxErrors = int(counter['TxErrors'])
                radioCounter.RxBytes  = int(counter['RxBytes'])
                radioCounter.RxPkts   = int(counter['RxPkts'])
                radioCounter.RxMgmt   = int(counter['RxMgmt'])
                radioCounter.RxErrors = int(counter['RxErrors'])
            elif util_flag == True:
                util[values[0]] = values[1]
            elif counter_flag == True:
                counter[values[0]] = values[1]
            elif dfs_flag == True:
                dfs[values[0]] = values[1]

        # update the last record
        if record_count > 0:
            radioDFS = radio_info.DFS
            radioDFS.CacState  = int(dfs['CacState'])
            radioDFS.RadarDetected  = (dfs['RadarDetected'] == 'TRUE')

        f.close()
    except Exception as e:
        resp.ErrStatus.Status = ap_common_types_pb2.APErrorStatus.AP_NOT_AVAILABLE
        print str(e)

    resp.ErrStatus.Status=ap_common_types_pb2.APErrorStatus.AP_SUCCESS
    return (resp)



#
# APWLANStatsGet
#
  def APWLANStatsGet(self, request, context):
    print "Received WLAN stats get request"

    import server_util

    resp = ap_stats_pb2.APWLANStatsMsgRsp()
    record_count = 0
    fields = {}

    try:
        f = open(PROC_WLAN_INFO, 'r')
        for line in f:
            # skip first line of each record
            if line.startswith('wlan num:'):
                # Fill last record with values
                if record_count > 0:
                    get_wlan_stats(wlan_info, fields)
                wlan_info = resp.WLANEntries.add()
                record_count += 1
                fields = {}
                continue

            # skip empty lines
            if line.strip() == '':
                continue

            if record_count == 0:
                continue

            values = line.split(':', 1)
            for i in range(len(values)):
                values[i] = values[i].strip()

            fields[values[0]] = values[1]

        f.close()

        if len(fields) != 0:
            get_wlan_stats(wlan_info, fields)

    except Exception as e:
        resp.ErrStatus.Status = ap_common_types_pb2.APErrorStatus.AP_NOT_AVAILABLE
        print str(e)

    resp.ErrStatus.Status=ap_common_types_pb2.APErrorStatus.AP_SUCCESS
    return (resp)


#
# APClientStatsGet
#
  def APClientStatsGet(self, request, context):
    print "Received Client stats get request"

    import server_util

    resp=ap_stats_pb2.APClientStatsMsgRsp()
    record_count = 0
    fields = {}

    try:
        f = open(PROC_CLIENT_INFO, 'r')
        for line in f:
            # skip first line of each record
            if line.startswith('client num:'):
                # Fill last record with values
                if record_count > 0:
                    get_client_stats(client_info, fields)
                client_info = resp.Clients.add()
                record_count += 1
                fields = {}
                continue

            # skip empty lines
            if line.strip() == '':
                continue

            if record_count == 0:
                continue

            values = line.split(':', 1)
            for i in range(len(values)):
                values[i] = values[i].strip()

            fields[values[0]] = values[1]

        f.close()

        if len(fields) != 0:
            get_client_stats(client_info, fields)

    except Exception as e:
        resp.ErrStatus.Status = ap_common_types_pb2.APErrorStatus.AP_NOT_AVAILABLE
        print str(e)

    resp.ErrStatus.Status=ap_common_types_pb2.APErrorStatus.AP_SUCCESS
    return (resp)


#
# APInterfaceStatsGet
#
  def APInterfaceStatsGet(self, request, context):
    print "Received Interface stats get request"

    resp = ap_stats_pb2.APInterfaceStatsMsgRsp()

    interface=resp.Interfaces.add()
    ifname = "wired0"
    module = subprocess.Popen('ifconfig ' + ifname, shell=True,
                              stdout=subprocess.PIPE).stdout.read()
    if module.strip() != '':
        interface.Name = module.split()[0]

        get_wired_info(interface)
        get_interface_stats(interface, module)


    return (resp)

## End class

#
#==============================================
# main
#==============================================
#
if __name__ == '__main__':
  import server_util
  server_ip, server_port = server_util.get_server_ip_port()

  print "Starting GRPC Server IP(%s) Port(%s)" %(server_ip, server_port)

  # Create the server
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

  # Add APGlobal servicer
  ap_global_pb2.add_APGlobalServicer_to_server(APGlobal(), server)

  # Add APStatistics servicer
  ap_stats_pb2.add_APStatisticsServicer_to_server(APStatistics(), server)

  server.add_insecure_port('[::]:' + str(server_port))
  server.start()

  try:
    while True:
      time.sleep(_ONE_DAY_IN_SECONDS)
  except KeyboardInterrupt:
    server.stop(0)
