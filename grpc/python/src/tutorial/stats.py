#
# Copyright (c) 2017 by Cisco Systems, Inc. 
# All rights reserved.
#

# Standard python libs
import os
import sys

# Add the generated python bindings to the path
sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.realpath(__file__))))

# gRPC generated python bindings
from genpy import ap_global_pb2
from genpy import ap_common_types_pb2
from genpy import ap_version_pb2
from genpy import ap_stats_pb2

# gRPC libs
from grpc.beta import implementations

# Utilities
from util import util
from tutorial import client_init

#
# Statistics operations
#    channel: GRPC channel
#
def system_stats_operation(channel):
    # Create the gRPC stub.
    stub = ap_stats_pb2.beta_create_APStatistics_stub(channel)

    # Get the system level stats. Create a APStasGetMsg
    stats_get = ap_stats_pb2.APStatsGetMsg()

    #
    # Make an RPC call to get the system stats
    #
    Timeout = 10 # Seconds
    response = stub.APSystemStatsGet(stats_get, Timeout)
    if (response.ErrStatus.Status ==
        ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
        util.print_system_stats(response)
    else:
        print "System stats response error 0x%x" %(response.ErrStatus.Status)
        os._exit(0)


def memory_stats_operation(channel):
    # Create the gRPC stub.
    stub = ap_stats_pb2.beta_create_APStatistics_stub(channel)

    # Get the system level stats. Create a APStasGetMsg
    stats_get = ap_stats_pb2.APStatsGetMsg()

    #
    # Make an RPC call to get the system stats
    #
    Timeout = 10 # Seconds
    response = stub.APMemoryStatsGet(stats_get, Timeout)
    if (response.ErrStatus.Status ==
        ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
        util.print_memory_stats(response)
    else:
        print "System stats response error 0x%x" %(response.ErrStatus.Status)
        os._exit(0)

#
# Setup the GRPC channel with the server, and issue RPCs
#
if __name__ == '__main__':
    from util import util
    server_ip, server_port = util.get_server_ip_port()

    print "Using GRPC Server IP(%s) Port(%s)" %(server_ip, server_port)

    # Create the channel for Server notifications.
    channel = implementations.insecure_channel(server_ip, server_port)

    # Spawn a thread to Initialize the client and listen on notifications
    # The thread will run in the background
    client_init.global_init(channel)

    # Create another channel for gRPC requests.
    channel = implementations.insecure_channel(server_ip, server_port)

    # Send RPCs for Stats
    system_stats_operation(channel)
    memory_stats_operation(channel)

    # Exit and Kill any running GRPC threads.
    os._exit(0)
