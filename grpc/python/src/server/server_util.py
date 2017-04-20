#
# Copyright (c) 2016 by Cisco Systems, Inc.
# All rights reserved.
#

# Standard python libs
import os
import socket
import subprocess
import re

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
