#
# Copyright (c) 2017 by Cisco Systems, Inc.
# All rights reserved.
#
import json
import os
import sys
import threading
import time
import unittest

# Add the generated python bindings to the path
sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.realpath(__file__))))


from cheetah import GrpcClient
from genpy import ap_common_types_pb2
from genpy import ap_global_pb2
from genpy import ap_stats_pb2
from util import util

# gRPC libs
from grpc.beta import implementations

TIMEOUT_SECONDS = 20

stats_types = [
               "Reserved",          # AP_RESERVED = 0
               "SystemStats",       # AP_SYSTEM_STATS = 1
               "MemoryStats",       # AP_MEMORY_STATS = 2
               "InterfaceStats",    # AP_INTERFACE_STATS = 3
               "RoutingStats",      # AP_ROUTING_STATS = 4
               "DNSStats",          # AP_DNS_STATS = 5
               "RadioStats",        # AP_RADIO_STATS = 6
               "WLANStats",         # AP_WLAN_STATS = 7
               "ClientStats"        # AP_CLIENT_STATS = 8
              ]
# Print Received Globals
def print_globals(response):
    if (response.ErrStatus.Status ==
        ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
        util.print_globals(response)
        return True
    else:
        print "Globals response error 0x%x" %(response.ErrStatus.Status)
        return False

# Global notification Callback
# This function is called from the global_init thread context
# To break the stream recv(), return False
def global_init_cback(response, event):
    if response.EventType == ap_global_pb2.AP_GLOBAL_EVENT_TYPE_VERSION:
        if (ap_common_types_pb2.APErrorStatus.AP_SUCCESS ==
                response.ErrStatus.Status) or \
            (ap_common_types_pb2.APErrorStatus.AP_INIT_STATE_CLEAR ==
                response.ErrStatus.Status) or \
            (ap_common_types_pb2.APErrorStatus.AP_INIT_STATE_READY ==
                response.ErrStatus.Status):
            print "Server Returned 0x%x, Server's Version %d.%d.%d" %(
                response.ErrStatus.Status,
                response.InitRspMsg.MajorVer,
                response.InitRspMsg.MinorVer,
                response.InitRspMsg.SubVer)
            # Successfully Initialized
            # This would notify the main thread to proceed
            event.set()
        else:
            return False
    elif response.EventType == ap_global_pb2.AP_GLOBAL_EVENT_TYPE_HEARTBEAT:
        print "Received Event: Heartbeat"
        return True
    elif response.EventType == ap_global_pb2.AP_GLOBAL_EVENT_TYPE_ERROR:
        print "Received Global Error event:", response
        return False
    else:
        print "Received unknown event:", response
        return False

    # Continue looping on events
    return True

# Wait on Global notification events
def global_init(event):
    g_params = ClientTestCase.json_params['global_init']
    json_dump = json.dumps(g_params)
    try:
        response = TestSuite_000_Global.global_notif.global_init(g_params, 
            global_init_cback, event)
        # Should return on errors
        if response.EventType == ap_global_pb2.AP_GLOBAL_EVENT_TYPE_ERROR:
            if (response.ErrStatus.Status ==
                ap_common_types_pb2.APErrorStatus.AP_NOTIF_TERM):
                print "Access Point Session was taken over by another client"
        else:
            # If this session is lost, then most likely the server restarted
            print "global_init: exiting unexpectedly, Server Restart?"
            print "last response from server:", response
    except Exception as e:
        print "Received exception:", e
        print "Server died?"
    os._exit(0)


#
#
class ClientTestCase(unittest.TestCase):
    # Class variables
    test_init = False
    # .json input variables to the test
    json_params = None
    # GRPC channel used for GRPC requests
    client = None

    def setUp(self):
        if not ClientTestCase.test_init:
            # Read the .json template
            filepath = os.path.join(os.path.dirname(__file__), 'template.json')
            with open(filepath) as fp:
                ClientTestCase.json_params = json.loads(fp.read())

            # Setup GRPC channel for RPC tests
            host, port = util.get_server_ip_port()
            ClientTestCase.client = GrpcClient(host, port)

            # Initialize only once
            ClientTestCase.test_init = True

#
# Alphabetical order makes this test run first
#
class TestSuite_000_Global(ClientTestCase):
    # GRPC channel used for Global notifications
    global_notif = None
    # threading.Event() used to sync threads
    global_event = None

    def test_000_global_init(self):
        # Setup a channel for the Global notification thread
        host, port = util.get_server_ip_port()
        TestSuite_000_Global.global_notif = GrpcClient(host, port)

        # Create a synchronization event
        TestSuite_000_Global.global_event = threading.Event()
        # Spawn a thread to wait on notifications
        t = threading.Thread(target = global_init,
                args=(TestSuite_000_Global.global_event,))
        t.start()
        #
        # Wait to hear from the server - Thread is blocked
        print "Waiting to hear from Global event..."
        TestSuite_000_Global.global_event.wait()
        print "Global Event Notification Received! Waiting for events..."

    def test_001_get_globals(self):
        # Get Global info
        response = ClientTestCase.client.global_get()
        err = print_globals(response)
        self.assertTrue(err)


def stats_cback(response, stats_type, event):
    if (response.ErrStatus.Status ==
        ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
        return (response.HasField(stats_types[stats_type]))
    else:
        return (response.ErrStatus.Status ==
                    ap_common_types_pb2.APErrorStatus.AP_NOT_AVAILABLE)

def stats_operation(self, stats_type, time_interval, count=1, event=None):

    response, counter = ClientTestCase.client.stats_get(stats_type, time_interval,
                        stats_cback, count, event)
    self.assertTrue(count==counter)

#
# Test unary statistics
#
class TestSuite_001_Statistics(ClientTestCase):

    time_interval = ap_stats_pb2.AP_STATS_UNARY_OPERATION

    def test_001_get_system_stats(self):
        # Get system stats
        stats_operation(self, ap_stats_pb2.AP_SYSTEM_STATS, self.time_interval)

    def test_002_get_memory_stats(self):
        # Get memory stats
        stats_operation(self, ap_stats_pb2.AP_MEMORY_STATS, self.time_interval)

    def test_003_get_dns_stats(self):
        # Get DNS stats
        stats_operation(self, ap_stats_pb2.AP_DNS_STATS, self.time_interval)

    def test_004_get_route_stats(self):
        # Get Route stats
        stats_operation(self, ap_stats_pb2.AP_ROUTING_STATS, self.time_interval)

    def test_005_get_interface_stats(self):
        # Get Interface stats
        stats_operation(self, ap_stats_pb2.AP_INTERFACE_STATS, self.time_interval)

    def test_006_get_wlan_stats(self):
        # Get WLAN stats
        stats_operation(self, ap_stats_pb2.AP_WLAN_STATS, self.time_interval)

    def test_007_get_client_stats(self):
        # Get Client stats
        stats_operation(self, ap_stats_pb2.AP_CLIENT_STATS, self.time_interval)

    def test_008_get_radio_stats(self):
        # Get Radio stats
        stats_operation(self, ap_stats_pb2.AP_RADIO_STATS, self.time_interval)


def stats_thread(self, stats_type, time_interval, count, event):
    t = threading.Thread(target = stats_operation,
                         args=(self, stats_type, time_interval, count, event))
    t.start()
    #event.wait()
    return t

#
# Test stream statistics
#
class TestSuite_002_Statistics(ClientTestCase):

    # threading.Event() used to sync threads
    stats_event = None

    time_interval = ap_stats_pb2.AP_STATS_MIN_TIME_INTERVAL
    count = 2

    def test_001_stream_get_system_stats(self):
        # Get system stats

        TestSuite_002_Statistics.stats_event = threading.Event()
        t = stats_thread(self, ap_stats_pb2.AP_SYSTEM_STATS, self.time_interval,
                         self.count, TestSuite_002_Statistics.stats_event)
        t.join()
 
    def test_002_stream_get_memory_stats(self):
        # Get memory stats

        # Create a synchronization event
        TestSuite_002_Statistics.stats_event = threading.Event()
        t = stats_thread(self, ap_stats_pb2.AP_MEMORY_STATS, self.time_interval,
                         self.count, TestSuite_002_Statistics.stats_event)
        t.join()

if __name__ == '__main__':
    unittest.main()