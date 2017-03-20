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

# Print System Stats
def print_system_stats(response):
    if (response.ErrStatus.Status ==
        ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
        util.print_system_stats(response)
    else:
        print "System stats response error 0x%x" %(response.ErrStatus.Status)
        return False
    return True

# Print Memory Stats
def print_memory_stats(response):
    if (response.ErrStatus.Status ==
        ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
        util.print_memory_stats(response)
    else:
        print "Memory stats response error 0x%x" %(response.ErrStatus.Status)
        return False
    return True

# Print DNS Stats
def print_dns_stats(response):
    if (response.ErrStatus.Status ==
        ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
        util.print_dns_stats(response)
    else:
        print "DNS stats response error 0x%x" %(response.ErrStatus.Status)
        return False
    return True

# Print Route Stats
def print_route_stats(response):
    if (response.ErrStatus.Status ==
        ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
        util.print_route_stats(response)
    else:
        print "Route stats response error 0x%x" %(response.ErrStatus.Status)
        return False
    return True

# Print Interface Stats
def print_interface_stats(response):
    if (response.ErrStatus.Status ==
        ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
        util.print_interface_stats(response)
    else:
        print "Interface stats response error 0x%x" %(response.ErrStatus.Status)
        return False
    return True

# Print Client Stats
def print_client_stats(response):
    if (response.ErrStatus.Status ==
        ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
        util.print_client_stats(response)
    else:
        print "Client stats response error 0x%x" %(response.ErrStatus.Status)
        return False
    return True

# Print Radio Stats
def print_radio_stats(response):
    if (response.ErrStatus.Status ==
        ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
        util.print_radio_stats(response)
    else:
        print "Radio stats response error 0x%x" %(response.ErrStatus.Status)
        return False
    return True

# Print WLAN Stats
def print_wlan_stats(response):
    if (response.ErrStatus.Status ==
        ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
        util.print_wlan_stats(response)
    else:
        print "WLAN stats response error 0x%x" %(response.ErrStatus.Status)
        return False
    return True

# Print Generic Stats
def print_generic_stats(response):
    if (response.ErrStatus.Status ==
        ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
        util.print_generic_stats(response)
    else:
        print "Generic stats response error 0x%x" %(response.ErrStatus.Status)
        return False
    return True

# Print Received Globals
def print_globals(response):
    if (response.ErrStatus.Status ==
        ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
        util.print_globals(response)
    else:
        print "Globals response error 0x%x" %(response.ErrStatus.Status)
        return False
    return True

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

class TestSuite_001_Statistics(ClientTestCase):

    def test_001_get_system_stats(self):
        # Get system stats
        response = ClientTestCase.client.system_stats_get()
        err = print_system_stats(response)
        self.assertTrue(err)

    def test_002_get_memory_stats(self):
        # Get memory stats
        response = ClientTestCase.client.memory_stats_get()
        err = print_memory_stats(response)
        self.assertTrue(err)

    def test_003_get_dns_stats(self):
        # Get DNS stats
        response = ClientTestCase.client.dns_stats_get()
        err = print_dns_stats(response)
        self.assertTrue(err)

    def test_004_get_route_stats(self):
        # Get Route stats
        response = ClientTestCase.client.routes_stats_get()
        err = print_route_stats(response)
        self.assertTrue(err)

    def test_005_get_interface_stats(self):
        # Get Interface stats
        response = ClientTestCase.client.interface_stats_get()
        err = print_interface_stats(response)
        self.assertTrue(err)

    def test_006_get_wlan_stats(self):
        # Get WLAN stats
        response = ClientTestCase.client.wlan_stats_get()
        err = print_wlan_stats(response)
        self.assertTrue(err)

    def test_007_get_client_stats(self):
        # Get Client stats
        response = ClientTestCase.client.client_stats_get()
        err = print_client_stats(response)
        self.assertTrue(err)

    def test_008_get_radio_stats(self):
        # Get Radio stats
        response = ClientTestCase.client.radio_stats_get()
        err = print_radio_stats(response)
        self.assertTrue(err)

    def test_009_get_generic_stats(self):
        # Get Generic stats
        response = ClientTestCase.client.generic_stats_get()
        err = print_generic_stats(response)
        self.assertTrue(err)


if __name__ == '__main__':
    unittest.main()
