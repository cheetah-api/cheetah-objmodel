#
# Copyright (c) 2016 by Cisco Systems, Inc.
# All rights reserved.
#
import abc

import serializers
from genpy import ap_common_types_pb2
from genpy import ap_global_pb2
from genpy import ap_stats_pb2
from genpy import ap_packet_pb2
from util import util

from grpc.beta import implementations

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

class Operation(object):
    ADD = 1
    UPDATE = 2
    DELETE = 3

class AbstractClient(object):
    __metaclass__ = abc.ABCMeta

    @abc.abstractmethod
    def global_init(self, *args, **kwargs):
        pass

    @abc.abstractmethod
    def global_get(self, *args, **kwargs):
        pass

    @abc.abstractmethod
    def stats_get(self, *args, **kwargs):
        pass

    @abc.abstractmethod
    def pkts_get(self, *args, **kwargs):
        pass


class GrpcClient(AbstractClient):
    TIMEOUT_SECONDS = 365*24*60*60 # Seconds

    def __init__(self, host, port, channel_credentials=None):
        if channel_credentials is None:
            # Instantiate insecure channel object.
            channel = implementations.insecure_channel(host, port)
        else:
            # Instantiate secure channel object.
            channel = implementations.secure_channel(host, port,
                                                     channel_credentials)
        self._stubs = (
            # 0
            ap_global_pb2.beta_create_APGlobal_stub(channel),
            # 1
            ap_stats_pb2.beta_create_APStatistics_stub(channel),
            # 2
            ap_packet_pb2.beta_create_APPackets_stub(channel),
        )

    def global_init(self, g_params, cback_func, event):
        """Global Init"""
        serializer = serializers.global_init_serializer(g_params)
        # Expect a stream of APGlobalNotif - XXX Use large timeout for now
        for response in self._stubs[0].APGlobalInitNotif(serializer,
            3600*24*365):
            if not cback_func(response, event):
                break
        # Returns on exit
        return response

    def global_get(self):
        """Global Get"""
        serializer = serializers.global_get_serializer()
        response = self._stubs[0].APGlobalsGet(serializer, self.TIMEOUT_SECONDS)
        return response

    def stats_get(self, stats_type, time_interval, cback_func, count, event):
        """Stats Get"""

        local_counter = 0

        #serializer = serializers.get_stats_serializer(stats_type, time_interval)
        serializer = serializers.get_stats_serializer()

        # Create Stats element
        stats = serializer.StatsRequest.add()
        stats.StatsType = stats_type
        stats.TimeInterval = time_interval

        for response in self._stubs[1].APStatsGet(serializer, self.TIMEOUT_SECONDS):
            local_counter += 1
            if not cback_func(response, stats_type, event):
                break
            if (count == local_counter):
                break

        # Terminated
        # This would notify the main thread to proceed
        #if not event is None:
           #event.set()
        return (response, local_counter)


    def pkts_get(self, serializer, cback_func, negative, event):
        """Packets Get"""

        local_counter = 0
        for response in self._stubs[2].APPacketsGet(serializer, self.TIMEOUT_SECONDS):
            local_counter += 1
            return cback_func(response, negative, event)

        # Terminated
        # This would notify the main thread to proceed
        #if not event is None:
           #event.set()
        #return (response, local_counter)
