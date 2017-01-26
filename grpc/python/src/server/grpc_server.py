#
# Copyright (c) 2017 by Cisco Systems, Inc.
# All rights reserved.
#

# Standard python libs
import os
import sys
import threading
import grpc
import time
import subprocess

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

# Global data
MAX_RADIO = 2
PROC_MEMINFO = '/proc/meminfo'
PROC_SLABINFO = '/proc/slabinfo'
ETC_RESOLV_CONF = '/etc/resolv.conf'

PROC_SYSTEM_INFO = '/proc/aptrace/sysinfo/system'
PROC_CLIENT_INFO = '/proc/aptrace/sysinfo/clients'
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
                resp.Link = "true"
            if ("duplex") in line and ("full") in line:
                resp.FullDuplex = "true"
            if ("speed") in line:
                resp.Speed = line.split(":")[1].strip()
        f.close()

    except Exception as e:
        print str(e)

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
                resp.Uptime = values[1].strip()
            else:
                continue
        f.close()
    except Exception as e:
        # using the exception path to test. Change to error eventually
        resp.ID = "52:00:00:ca:fe:01"
        resp.Uptime = 10000

    resp.When = str(datetime.now())
    resp.ErrStatus.Status=ap_common_types_pb2.APErrorStatus.AP_SUCCESS

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
    pass

#
# APWLANStatsGet
#
  def APWLANStatsGet(self, request, context):
    pass

#
# APClientStatsGet
#
  def APClientStatsGet(self, request, context):
    pass


#
# APInterfaceStatsGet
#
  def APInterfaceStatsGet(self, request, context):
    print "Received Interface stats get request"

    resp = ap_stats_pb2.APInterfaceStatsMsgRsp()

    interface=resp.Interfaces.add()
    ifname = "ens33"
    module = subprocess.Popen('ifconfig ' + ifname, shell=True,
                              stdout=subprocess.PIPE).stdout.read()

    if module.strip() != '':
        interface.Name = module.split()[0]

        get_wired_info(interface)

        head, sep, tail = module.partition("RX bytes:")
        interface.RxBytes = int(tail.split()[0])

        head, sep, tail = module.partition("RX packets:")
        interface.RxPkts = int(tail.split()[0])

        interface.RxDiscards = int(tail.split("dropped:")[1].split()[0])

        head, sep, tail = module.partition("TX bytes:")
        interface.TxBytes = int(tail.split()[0])

        head, sep, tail = module.partition("TX packets:")
        interface.TxPkts = int(tail.split()[0])

    return (resp)

    

## End class 

#
#==============================================
# main
#==============================================
#
if __name__ == '__main__':
  from util import util
  server_ip, server_port = util.get_server_ip_port()

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

