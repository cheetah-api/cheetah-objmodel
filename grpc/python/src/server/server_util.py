#
# Copyright (c) 2016 by Cisco Systems, Inc.
# All rights reserved.
#

# Standard python libs
import os
import socket
import subprocess
import re

CLIENT_IP_TABLE = '/click/client_ip_table/list'

# Name of interface on the access point
WIRED_INTERFACE='wired0'

#
# GRPC IP address of interface
#
def get_ip_address(ifname):
    try:
        f = os.popen('ifconfig ' +  ifname  + ' | ' +
                     'grep "inet\ addr" | cut -d: -f2 | cut -d" " -f1')
        return f.read()
    except Exception as e:
        print str(e)
        return None

#
# Get the GRPC Server IP address and port number
#
def get_server_ip_port():

    # Set GRPC Server's IP in the environment
    os.environ['SERVER_IP'] = get_ip_address(WIRED_INTERFACE)
    os.environ['SERVER_PORT'] = '57777'

    return (os.environ['SERVER_IP'],
            int(os.environ['SERVER_PORT']))

#
# Generic function to call wcp handlers
#
def get_wcp_data(handler):
    command = "wcp_read " + handler

    output_proc = subprocess.Popen(command, shell=True, stdout=subprocess.PIPE)
    output = output_proc.stdout.read()
    return (output)


def get_client_count_per_ssid(ssid):
    count = 0

    with open(CLIENT_IP_TABLE) as f:
       for line in f:
           columns = line.split()
           if (columns[2] == ssid):
               count += 1

    return count


def get_mcast_pkts(i, ssid, mcast_counters):
   handler = "wcp/RadDrv" + str(i) + ".vap_stats"
   wcp_data = get_wcp_data(handler)
   mylist = re.split("\n", wcp_data)

   mcast_counters.TxMcastPkts = mcast_counters.TxMcastBytes = 0

   for line in mylist[1:]:
       if (line == ''):
           continue
       values = line.split()
       if (ssid == values[0]):
           mcast_counters.TxMcastPkts = int(values[3])
           mcast_counters.TxMcastBytes = int(values[5])
           return mcast_counters
