#
# Copyright (c) 2017 by Cisco Systems, Inc.
# All rights reserved.
#

# Standard python libs
import os
import sys

# Add the generated python bindings directory to the path
sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.realpath(__file__))))

# gRPC generated python bindings
from genpy import ap_common_types_pb2
from genpy import ap_stats_pb2

# Utilities
from tutorial import client_init
from tutorial import stats

# gRPC libs
from grpc.beta import implementations

#
# Setup the GRPC channel with the server, and issue RPCs
#
if __name__ == '__main__':
    from util import util
    server_ip, server_port = util.get_server_ip_port()

    print "Using GRPC Server: IP(%s) Port(%s)" %(server_ip, server_port)

    # Create the channel for gRPC.
    channel = implementations.insecure_channel(server_ip, server_port)

    # Spawn a thread to Initialize the client and listen on notifications
    # The thread will run in the background
    client_init.global_init(channel)

    # Create another channel for gRPC requests.
    channel = implementations.insecure_channel(server_ip, server_port)

    # Stats operations

    # System Stats every 6 seconds
    t1=stats.stats_operations(channel, ap_stats_pb2.AP_SYSTEM_STATS, 6)

    # Memory Stats every 11 seconds
    t2=stats.stats_operations(channel, ap_stats_pb2.AP_MEMORY_STATS, 11)

    # Wait till the threads terminate
    t1.join()
    t2.join()

    # Exit and Kill any running GRPC threads.
    os._exit(0)
