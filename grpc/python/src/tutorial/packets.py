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
from genpy import ap_packet_pb2

# gRPC libs
from grpc.beta import implementations

# Utilities
from util import util
from tutorial import client_init

#
# Packets thread
#    channel: GRPC channel
#
def packets_reg_thread(channel, oper, packets_type, packets_subtype):
    # Create the gRPC stub.
    stub = ap_packet_pb2.beta_create_APPackets_stub(channel)

    # Register the packets. Create APPacketsRegMsg
    packets_reg_msg = ap_packet_pb2.APPacketsRegMsg()

    # Add Registrtaion operation
    packets_reg_msg.Oper = oper
    #
    # Add packets to the list
    #
    packets = packets_reg_msg.PacketHdr.add()
    packets.MsgType = packets_type
    if (packets.MsgType == ap_packet_pb2.AP_MSG_TYPE_MGMT):
        packets.mgmt = packets_subtype
    elif (packets.MsgType == ap_packet_pb2.AP_MSG_TYPE_CTRL):
        packets.ctrl = packets_subtype
    elif (packets.MsgType == ap_packet_pb2.AP_MSG_TYPE_DATA):
        packets.data = packets_subtype
    elif (packets.MsgType == ap_packet_pb2.AP_MSG_TYPE_CISCO):
        packets.cisco = packets_subtype
    else:
        print "Invalid packet type"
        os._exit(0)

    Timeout = 365*24*60*60 # Seconds
    response = stub.APPacketsRegOp(packets_reg_msg, Timeout)
    if (response.ErrStatus.Status ==
        ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
        print "Registration Success"
    else:
        print "Registration Failed"
        print response.Results

def packets_notif_thread(channel):
    # Create the gRPC stub.
    stub = ap_packet_pb2.beta_create_APPackets_stub(channel)

    # Get the packets. Create APPacketsGetNotifMsg
    packets_notif_msg = ap_packet_pb2.APPacketsGetNotifMsg()

    Timeout = 365*24*60*60 # Seconds
    for response in stub.APPacketsInitNotif(packets_notif_msg, Timeout):
        if (response.ErrStatus.Status ==
            ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
            if (response.PacketHdr.MsgType == ap_packet_pb2.AP_MSG_TYPE_MGMT):
                print "Got Management packet %d" %(response.PacketHdr.mgmt)
            elif (response.PacketHdr.MsgType == ap_packet_pb2.AP_MSG_TYPE_CTRL):
                print "Got Control packet %d" %(response.PacketHdr.ctrl)
            elif (response.PacketHdr.MsgType == ap_packet_pb2.AP_MSG_TYPE_DATA):
                print "Got Data packet %d" %(response.PacketHdr.data)
            elif (response.PacketHdr.MsgType == ap_packet_pb2.AP_MSG_TYPE_CISCO):
                print "Got Cisco packet %d" %(response.PacketHdr.cisco)
            elif ((response.PacketHdr.MsgType == ap_packet_pb2.AP_MSG_TYPE_RESERVED) and
                  (response.PacketLen == 0)):
                print "Packet notification request success"
            else:
                print "Got Unclassified packet"
        else:
            print "Packets config response error 0x%x" %(response.ErrStatus.Status)
            os._exit(0)

def packets_reg_operations(channel, regop, packets_type, packets_subtype):
    t = threading.Thread(target = packets_reg_thread,
                         args=(channel, regop, packets_type, packets_subtype))
    t.start()
    return t

def packets_notif_operations(channel):
    t = threading.Thread(target = packets_notif_thread, args=(channel,))
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

    # Packets reg operations
    t1=packets_reg_operations(channel,
                              ap_common_types_pb2.AP_REGOP_REGISTER,
                              ap_packet_pb2.AP_MSG_TYPE_MGMT,
                              ap_packet_pb2.AP_MGMT_MSG_SUBTYPE_AUTH)

    t1.join()

    # Packets notif operation
    t2=packets_notif_operations(channel)
    t2.join()

    # Exit and Kill any running GRPC threads.
    os._exit(0)
