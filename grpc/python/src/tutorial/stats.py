#
# Copyright (c) 2017 by Cisco Systems, Inc. 
# All rights reserved.
#

# Standard python libs
import os
import sys
import threading

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
# Statistics thread
#    channel: GRPC channel
#
def stats_thread(channel, stats_type, time_interval):
    # Create the gRPC stub.
    stub = ap_stats_pb2.beta_create_APStatistics_stub(channel)

    # Get the system level stats. Create APStatsMsg
    stats_msg = ap_stats_pb2.APStatsMsg()

    #
    # Add system stats to the list, once every 5 seconds
    #
    stats = stats_msg.StatsRequest.add()
    #system_stats.StatsType = ap_stats_pb2.AP_SYSTEM_STATS
    stats.StatsType = stats_type
    stats.TimeInterval = time_interval

    Timeout = 365*24*60*60 # Seconds
    for response in stub.APStatsGet(stats_msg, Timeout):
        if (response.ErrStatus.Status ==
            ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
            if response.HasField("SystemStats"):
                util.print_system_stats(response.SystemStats)
            elif response.HasField("MemoryStats"):
                util.print_memory_stats(response.MemoryStats)
        else:
            print "Stats response error 0x%x" %(response.ErrStatus.Status)
            os._exit(0)


def stats_operations(channel, stats_type, time_interval):
    t = threading.Thread(target = stats_thread, args=(channel, stats_type, time_interval))
    t.start()
    return t

#
# Setup the GRPC channel with the server, and issue RPCs
#
if __name__ == '__main__':
    from util import util
    server_ip, server_port = util.get_server_ip_port()

    print "Using GRPC Server: IP(%s) Port(%s)" %(server_ip, server_port)

    # Create the channel for Server notifications.
    channel = implementations.insecure_channel(server_ip, server_port)

    # Spawn a thread to Initialize the client and listen on notifications
    # The thread will run in the background
    client_init.global_init(channel)

    # Create another channel for gRPC requests.
    channel = implementations.insecure_channel(server_ip, server_port)

    # Stats operations 
    t1=stats_operations(channel, 1, 10)
    t2=stats_operations(channel, 2, 5)
    t1.join()
    t2.join()

    # Exit and Kill any running GRPC threads.
    os._exit(0)
