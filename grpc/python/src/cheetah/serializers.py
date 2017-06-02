#
# Copyright (c) 2016 by Cisco Systems, Inc.
# All rights reserved.
#
import collections

from util import util

from genpy import ap_common_types_pb2
from genpy import ap_global_pb2
from genpy import ap_stats_pb2
from genpy import ap_packet_pb2

def global_init_serializer(init):
    """Global Init Message serializer."""
    serializer = ap_global_pb2.APInitMsg()
    if 'major' in init:
        serializer.MajorVer = init['major']
    if 'minor' in init:
        serializer.MinorVer = init['minor']
    if 'sub_ver' in init:
        serializer.SubVer = init['sub_ver']
    return serializer

def global_get_serializer():
    """Global Get Message serializer."""
    serializer = ap_global_pb2.APGlobalsGetMsg()
    return serializer

def get_stats_serializer():
    """Get Stats Message serializer."""
    serializer = ap_stats_pb2.APStatsMsg()
    return serializer

def get_pkts_serializer():
    """Get Pkts Message serializer."""
    serializer = ap_packet_pb2.APPacketsMsg()
    return serializer
