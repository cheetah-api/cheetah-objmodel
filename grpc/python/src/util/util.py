#
# Copyright (c) 2016 by Cisco Systems, Inc.
# All rights reserved.
#

# Standard python libs
import os
import ipaddress

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
    print "AP ID     : %s" %(response.ID)
    print "Uptime    : %d" %(response.Uptime)
    print "When      : %s" %(response.When)
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

if __name__ == '__main__':
    server_ip, server_port = get_server_ip_port()
    print "Using GRPC Server IP(%s) Port(%s)" %(server_ip, server_port)
    
