#
# Copyright (c) 2016 by Cisco Systems, Inc.
# All rights reserved.
#

# Standard python libs
import os

#
# Get the GRPC Server IP address and port number
#
def get_server_ip_port():
    # Get GRPC Server's IP from the environment
    if 'SERVER_IP' not in os.environ.keys():
        print "Need to set the SERVER_IP env variable e.g."
        print "export SERVER_IP='172.16.239.136'"
        os._exit(0)
    
    # Get GRPC Server's Port from the environment
    if 'SERVER_PORT' not in os.environ.keys():
        print "Need to set the SERVER_PORT env variable e.g."
        print "export SERVER_PORT='57777'"
        os._exit(0)
    
    return (os.environ['SERVER_IP'], int(os.environ['SERVER_PORT']))

# Print Globals
def print_globals(response):
    print "----Global statistics----"
    print "Max Radio Name Len  : %d" %(response.MaxRadioNameLength)
    print "Max Ssid Name Len   : %d" %(response.MaxSsidNameLength)
    print

# Print System Stats
def print_system_stats(response):
    print "----System statistics----"
    print "AP ID       : %s" %(response.ID)
    print "Serial No   : %s" %(response.SerialNumber)
    print "Product ID  : %s" %(response.ProductId)
    print "Uptime      : %d" %(response.Uptime)
    print "When        : %s" %(response.When)
    print

# Print Memory Stats
def print_memory_stats(response):
    print "----MemInfo----"
    print "Total KB      : %d" %(response.ProcMemInfo.Total_kB)
    print "Available KB  : %d" %(response.ProcMemInfo.Available_kB)
    print

    print "----TopSlabInfo----"
    print "Name : %s" %(response.TopProcSlabInfo.Name)
    print "Object size: %d" %(response.TopProcSlabInfo.ObjSize)
    print "Active objects : %d" %(response.TopProcSlabInfo.ActiveObjs)
    print "Number of objects : %d" %(response.TopProcSlabInfo.NumObjs)
    print

# Print DNS Stats
def print_dns_stats(response):
    print "----DNS statistics----"
    for ip in response.IP:
        print "DNS Server     : %s" %(ip)
    print

# Print Route Stats
def print_route_stats(response):
    print "----Route statistics----"
    for route in response.IPv4Routes:
        print "Destination : %s" %(route.Destination)
        print "Gateway     : %s" %(route.Gateway)
        print "Mask        : %s" %(route.Genmask)
        print "Flags       : %s" %(route.Flags)
        print "Metric      : %d" %(route.Metric)
        print "Ref         : %d" %(route.Ref)
        print "Use         : %d" %(route.Use)
        print "Interface   : %s" %(route.Iface)
        print

# Print Interface Stats
def print_interface_stats(resp):
    print "----Interface statistics----"
    for interface in resp.Interfaces:
        print "Name        : %s" %(interface.Name)
        print "Link        : %s" %(interface.Link)
        print "Full Duplex : %s" %(interface.FullDuplex)
        print "Speed       : %d" %(interface.Speed)
        print "RxBytes     : %d" %(interface.RxBytes)
        print "RxPkts      : %d" %(interface.RxPkts)
        print "RxDiscards  : %d" %(interface.RxDiscards)
        print "TxBytes     : %d" %(interface.TxBytes)
        print "TxPkts      : %d" %(interface.TxPkts)
        print

# Print WLAN Stats
def print_wlan_stats(resp):
    print "----WLAN statistics----"
    for wlan in resp.WLANEntries:
        print "ESSID           : %s" %(wlan.Wlan.ID)
        print "-----------------------"
        print " SSID           : %s" %(wlan.Wlan.SSID)
        print " Dev            : %s" %(wlan.Dev)
        print " BSSID          : %s" %(wlan.BSSID)
        print " RadioIndex     : %s" %(wlan.RadioIndex)
        print " Num Clients    : %d" %(wlan.NumClients)
        print " Mcast TX Pkts  : %d" %(wlan.Counter.TxMcastPkts)
        print " Mcast TX Bytes : %d" %(wlan.Counter.TxMcastBytes)
        print

# Print Client Stats
def print_client_stats(resp):
    print "----Client statistics----"
    for client in resp.Clients:
        print "MAC                    : %s" %(client.MAC)
        print "RadioIndex             : %d" %(client.RadioIndex)
        print "Band                   : %s" %(client.Band)
        print "ESSID                  : %s" %(client.Wlan.ID)
        print "SSID                   : %s" %(client.Wlan.SSID)
        print "ConnectedTimeSec       : %d" %(client.ConnectedTimeSec)
        print "InactiveTimeMilliSec   : %d" %(client.InactiveTimeMilliSec)
        print "RSSI                   : %d" %(client.RSSI)
        print "NF                     : %d" %(client.NF)
        print "PerAntennaRSSI         : "
        for rssi in client.AntennaRSSI:
            print "\tRSSI            : %d" %(rssi)
        print "TxBitRate              : %d" %(client.TxBitRate)
        print "TxUnicastBytes         : %d" %(client.TxUnicastBytes)
        print "TxUnicastPkts          : %d" %(client.TxUnicastPkts)
        print "RxBytes                : %d" %(client.RxBytes)
        print "RxPkts                 : %d" %(client.RxPkts)

# Print Radio Stats
def print_radio_stats(resp):
    print "----Radio statistics----"
    for radio in resp.Radios:
        print "Device             : %s" %(radio.Dev)
        print "Band               : %s" %(radio.Band)
        print "Channel            : %d" %(radio.Channel)
        print "Secondary Channel  : %d" %(radio.SecondaryChannel)
        print "Bandwidth          : %d" %(radio.Bandwidth)
        print "NoiseFloor         : %d" %(radio.NoiseFloor)
        print "MaxTxPower         : %d" %(radio.MaxTxPower)
        print "Utilization        : "
        print "\tAll              : %f" %(radio.Utilization.All)
        print "\tTx               : %f" %(radio.Utilization.Tx)
        print "\tRxInBSS          : %f" %(radio.Utilization.RxInBSS)
        print "\tRxOtherBSS       : %f" %(radio.Utilization.RxOtherBSS)
        print "\tNonWifi          : %f" %(radio.Utilization.NonWifi)
        print "PerAntennaRSSI     : "
        for rssi in radio.AntennaRSSI:
            print "\tRSSI        : %d" %(rssi)
        print "Counters           : "
        print "\tTxBytes              : %d" %(radio.Counter.TxBytes)
        print "\tTxPkts               : %d" %(radio.Counter.TxPkts)
        print "\tTxMgmt               : %d" %(radio.Counter.TxMgmt)
        print "\tTxErrors             : %d" %(radio.Counter.TxErrors)
        print "\tRxBytes              : %d" %(radio.Counter.RxBytes)
        print "\tRxPkts               : %d" %(radio.Counter.RxPkts)
        print "\tRxMgmt               : %d" %(radio.Counter.RxMgmt)
        print "\tRxErrors             : %d" %(radio.Counter.RxErrors)
        print "DFS State          : "
        print "\tCacState              : %d" %(radio.DFS.CacState)
        print "\tRadarDetected         : %d" %(radio.DFS.RadarDetected)

#
#============================================
# main
#============================================
#
if __name__ == '__main__':
    server_ip, server_port = get_server_ip_port()
    print "Using GRPC Server: IP(%s) Port(%s)" %(server_ip, server_port)
    
