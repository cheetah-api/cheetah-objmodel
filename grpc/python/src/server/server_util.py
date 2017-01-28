#
# Copyright (c) 2016 by Cisco Systems, Inc.
# All rights reserved.
#

# Standard python libs
import os

# Name of interface on the access point
WIRED_INTERFACE='wired0'

# 
# Get the GRPC Server IP address and port number
#
def get_server_ip_port():

    # Set GRPC Server's IP in the environment
    try:
        f = os.popen('ifconfig ' +  WIRED_INTERFACE  + ' | ' +
                     'grep "inet\ addr" | cut -d: -f2 | cut -d" " -f1')
    except Exception as e:
        print str(e)
	os.exit()
        
    os.environ['SERVER_IP'] = f.read()
    os.environ['SERVER_PORT'] = '57777'

    return (os.environ['SERVER_IP'],
            int(os.environ['SERVER_PORT']))
