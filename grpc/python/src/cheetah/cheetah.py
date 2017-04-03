#
# Copyright (c) 2016 by Cisco Systems, Inc.
# All rights reserved.
#
import abc

import serializers
from genpy import ap_common_types_pb2
from genpy import ap_global_pb2
from genpy import ap_stats_pb2

from grpc.beta import implementations

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
    def system_stats_get(self, *args, **kwargs):
        pass

    @abc.abstractmethod
    def memory_stats_get(self, *args, **kwargs):
        pass

    @abc.abstractmethod
    def dns_stats_get(self, *args, **kwargs):
        pass

    @abc.abstractmethod
    def routes_stats_get(self, *args, **kwargs):
        pass

    @abc.abstractmethod
    def interface_stats_get(self, *args, **kwargs):
        pass

    @abc.abstractmethod
    def wlan_stats_get(self, *args, **kwargs):
        pass

    @abc.abstractmethod
    def client_stats_get(self, *args, **kwargs):
        pass

    @abc.abstractmethod
    def radio_stats_get(self, *args, **kwargs):
        pass

class GrpcClient(AbstractClient):
    TIMEOUT_SECONDS = 20

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
        response = self._stubs[0].APGlobalsGet(serializer,
            self.TIMEOUT_SECONDS)
        return response

    def system_stats_get(self):
        """System Stats Get"""
        serializer = serializers.get_stats_serializer()
        response = self._stubs[1].APSystemStatsGet(serializer,
            self.TIMEOUT_SECONDS)
        return response

    def memory_stats_get(self):
        """Memory Stats Get"""
        serializer = serializers.get_stats_serializer()
        response = self._stubs[1].APMemoryStatsGet(serializer,
            self.TIMEOUT_SECONDS)
        return response

    def dns_stats_get(self):
        """DNS Stats Get"""
        serializer = serializers.get_stats_serializer()
        response = self._stubs[1].APDNSStatsGet(serializer,
            self.TIMEOUT_SECONDS)
        return response

    def routes_stats_get(self):
        """Route Stats Get"""
        serializer = serializers.get_stats_serializer()
        response = self._stubs[1].APRoutesStatsGet(serializer,
            self.TIMEOUT_SECONDS)
        return response

    def interface_stats_get(self):
        """Interface Stats Get"""
        serializer = serializers.get_stats_serializer()
        response = self._stubs[1].APInterfaceStatsGet(serializer,
            self.TIMEOUT_SECONDS)
        return response

    def client_stats_get(self):
        """Client Stats Get"""
        serializer = serializers.get_stats_serializer()
        response = self._stubs[1].APClientStatsGet(serializer,
            self.TIMEOUT_SECONDS)
        return response


    def radio_stats_get(self):
        """Radio Stats Get"""
        serializer = serializers.get_stats_serializer()
        response = self._stubs[1].APRadioStatsGet(serializer,
            self.TIMEOUT_SECONDS)
        return response

    def wlan_stats_get(self):
        """WLAN Stats Get"""
        serializer = serializers.get_stats_serializer()
        response = self._stubs[1].APWLANStatsGet(serializer,
            self.TIMEOUT_SECONDS)
        return response
