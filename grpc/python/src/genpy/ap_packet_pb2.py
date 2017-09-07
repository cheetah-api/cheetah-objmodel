# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ap_packet.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import ap_common_types_pb2 as ap__common__types__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='ap_packet.proto',
  package='cheetah',
  syntax='proto3',
  serialized_pb=_b('\n\x0f\x61p_packet.proto\x12\x07\x63heetah\x1a\x15\x61p_common_types.proto\"\x16\n\x14\x41PPacketsGetNotifMsg\"Z\n\x0f\x41PPacketsRegMsg\x12\x1e\n\x04Oper\x18\x01 \x01(\x0e\x32\x10.cheetah.APRegOp\x12\'\n\tPacketHdr\x18\x02 \x03(\x0b\x32\x14.cheetah.APPacketHdr\"j\n\x12\x41PPacketsRegMsgRsp\x12)\n\tErrStatus\x18\x01 \x01(\x0b\x32\x16.cheetah.APErrorStatus\x12)\n\x07Results\x18\x02 \x01(\x0b\x32\x18.cheetah.APPacketsRegMsg\"\xeb\x01\n\x0b\x41PPacketHdr\x12#\n\x07MsgType\x18\x01 \x01(\x0e\x32\x12.cheetah.APMsgType\x12)\n\x04mgmt\x18\x02 \x01(\x0e\x32\x19.cheetah.APMgmtMsgSubtypeH\x00\x12)\n\x04\x63trl\x18\x03 \x01(\x0e\x32\x19.cheetah.APCtrlMsgSubtypeH\x00\x12)\n\x04\x64\x61ta\x18\x04 \x01(\x0e\x32\x19.cheetah.APDataMsgSubtypeH\x00\x12+\n\x05\x63isco\x18\x05 \x01(\x0e\x32\x1a.cheetah.APCiscoMsgSubtypeH\x00\x42\t\n\x07Subtype\"\x8b\x01\n\x0f\x41PPacketsMsgRsp\x12)\n\tErrStatus\x18\x01 \x01(\x0b\x32\x16.cheetah.APErrorStatus\x12\'\n\tPacketHdr\x18\x02 \x01(\x0b\x32\x14.cheetah.APPacketHdr\x12\x11\n\tPacketLen\x18\x03 \x01(\r\x12\x11\n\tPacketBuf\x18\x04 \x01(\x0c*~\n\tAPMsgType\x12\x18\n\x14\x41P_MSG_TYPE_RESERVED\x10\x00\x12\x14\n\x10\x41P_MSG_TYPE_MGMT\x10\x01\x12\x14\n\x10\x41P_MSG_TYPE_CTRL\x10\x02\x12\x14\n\x10\x41P_MSG_TYPE_DATA\x10\x03\x12\x15\n\x11\x41P_MSG_TYPE_CISCO\x10\x04*\xaf\x01\n\x10\x41PMgmtMsgSubtype\x12 \n\x1c\x41P_MGMT_MSG_SUBTYPE_RESERVED\x10\x00\x12\x1d\n\x19\x41P_MGMT_MSG_SUBTYPE_ASSOC\x10\x01\x12\x1c\n\x18\x41P_MGMT_MSG_SUBTYPE_AUTH\x10\x02\x12\x1d\n\x19\x41P_MGMT_MSG_SUBTYPE_PROBE\x10\x04\x12\x1d\n\x17\x41P_MGMT_MSG_SUBTYPE_ALL\x10\xff\xff\x03*S\n\x10\x41PCtrlMsgSubtype\x12 \n\x1c\x41P_CTRL_MSG_SUBTYPE_RESERVED\x10\x00\x12\x1d\n\x17\x41P_CTRL_MSG_SUBTYPE_ALL\x10\xff\xff\x03*\xc9\x01\n\x10\x41PDataMsgSubtype\x12 \n\x1c\x41P_DATA_MSG_SUBTYPE_RESERVED\x10\x00\x12\x1b\n\x17\x41P_DATA_MSG_SUBTYPE_ARP\x10\x01\x12\x1c\n\x18\x41P_DATA_MSG_SUBTYPE_DHCP\x10\x02\x12\x1b\n\x17\x41P_DATA_MSG_SUBTYPE_EAP\x10\x04\x12\x1c\n\x18\x41P_DATA_MSG_SUBTYPE_ICMP\x10\x08\x12\x1d\n\x17\x41P_DATA_MSG_SUBTYPE_ALL\x10\xff\xff\x03*t\n\x11\x41PCiscoMsgSubtype\x12!\n\x1d\x41P_CISCO_MSG_SUBTYPE_RESERVED\x10\x00\x12\x1c\n\x18\x41P_CISCO_MSG_SUBTYPE_NDP\x10\x01\x12\x1e\n\x18\x41P_CISCO_MSG_SUBTYPE_ALL\x10\xff\xff\x03\x32\xa5\x01\n\tAPPackets\x12G\n\x0e\x41PPacketsRegOp\x12\x18.cheetah.APPacketsRegMsg\x1a\x1b.cheetah.APPacketsRegMsgRsp\x12O\n\x12\x41PPacketsInitNotif\x12\x1d.cheetah.APPacketsGetNotifMsg\x1a\x18.cheetah.APPacketsMsgRsp0\x01\x62\x06proto3')
  ,
  dependencies=[ap__common__types__pb2.DESCRIPTOR,])
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

_APMSGTYPE = _descriptor.EnumDescriptor(
  name='APMsgType',
  full_name='cheetah.APMsgType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='AP_MSG_TYPE_RESERVED', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_MSG_TYPE_MGMT', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_MSG_TYPE_CTRL', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_MSG_TYPE_DATA', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_MSG_TYPE_CISCO', index=4, number=4,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=655,
  serialized_end=781,
)
_sym_db.RegisterEnumDescriptor(_APMSGTYPE)

APMsgType = enum_type_wrapper.EnumTypeWrapper(_APMSGTYPE)
_APMGMTMSGSUBTYPE = _descriptor.EnumDescriptor(
  name='APMgmtMsgSubtype',
  full_name='cheetah.APMgmtMsgSubtype',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='AP_MGMT_MSG_SUBTYPE_RESERVED', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_MGMT_MSG_SUBTYPE_ASSOC', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_MGMT_MSG_SUBTYPE_AUTH', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_MGMT_MSG_SUBTYPE_PROBE', index=3, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_MGMT_MSG_SUBTYPE_ALL', index=4, number=65535,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=784,
  serialized_end=959,
)
_sym_db.RegisterEnumDescriptor(_APMGMTMSGSUBTYPE)

APMgmtMsgSubtype = enum_type_wrapper.EnumTypeWrapper(_APMGMTMSGSUBTYPE)
_APCTRLMSGSUBTYPE = _descriptor.EnumDescriptor(
  name='APCtrlMsgSubtype',
  full_name='cheetah.APCtrlMsgSubtype',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='AP_CTRL_MSG_SUBTYPE_RESERVED', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_CTRL_MSG_SUBTYPE_ALL', index=1, number=65535,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=961,
  serialized_end=1044,
)
_sym_db.RegisterEnumDescriptor(_APCTRLMSGSUBTYPE)

APCtrlMsgSubtype = enum_type_wrapper.EnumTypeWrapper(_APCTRLMSGSUBTYPE)
_APDATAMSGSUBTYPE = _descriptor.EnumDescriptor(
  name='APDataMsgSubtype',
  full_name='cheetah.APDataMsgSubtype',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='AP_DATA_MSG_SUBTYPE_RESERVED', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_DATA_MSG_SUBTYPE_ARP', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_DATA_MSG_SUBTYPE_DHCP', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_DATA_MSG_SUBTYPE_EAP', index=3, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_DATA_MSG_SUBTYPE_ICMP', index=4, number=8,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_DATA_MSG_SUBTYPE_ALL', index=5, number=65535,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1047,
  serialized_end=1248,
)
_sym_db.RegisterEnumDescriptor(_APDATAMSGSUBTYPE)

APDataMsgSubtype = enum_type_wrapper.EnumTypeWrapper(_APDATAMSGSUBTYPE)
_APCISCOMSGSUBTYPE = _descriptor.EnumDescriptor(
  name='APCiscoMsgSubtype',
  full_name='cheetah.APCiscoMsgSubtype',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='AP_CISCO_MSG_SUBTYPE_RESERVED', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_CISCO_MSG_SUBTYPE_NDP', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AP_CISCO_MSG_SUBTYPE_ALL', index=2, number=65535,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1250,
  serialized_end=1366,
)
_sym_db.RegisterEnumDescriptor(_APCISCOMSGSUBTYPE)

APCiscoMsgSubtype = enum_type_wrapper.EnumTypeWrapper(_APCISCOMSGSUBTYPE)
AP_MSG_TYPE_RESERVED = 0
AP_MSG_TYPE_MGMT = 1
AP_MSG_TYPE_CTRL = 2
AP_MSG_TYPE_DATA = 3
AP_MSG_TYPE_CISCO = 4
AP_MGMT_MSG_SUBTYPE_RESERVED = 0
AP_MGMT_MSG_SUBTYPE_ASSOC = 1
AP_MGMT_MSG_SUBTYPE_AUTH = 2
AP_MGMT_MSG_SUBTYPE_PROBE = 4
AP_MGMT_MSG_SUBTYPE_ALL = 65535
AP_CTRL_MSG_SUBTYPE_RESERVED = 0
AP_CTRL_MSG_SUBTYPE_ALL = 65535
AP_DATA_MSG_SUBTYPE_RESERVED = 0
AP_DATA_MSG_SUBTYPE_ARP = 1
AP_DATA_MSG_SUBTYPE_DHCP = 2
AP_DATA_MSG_SUBTYPE_EAP = 4
AP_DATA_MSG_SUBTYPE_ICMP = 8
AP_DATA_MSG_SUBTYPE_ALL = 65535
AP_CISCO_MSG_SUBTYPE_RESERVED = 0
AP_CISCO_MSG_SUBTYPE_NDP = 1
AP_CISCO_MSG_SUBTYPE_ALL = 65535



_APPACKETSGETNOTIFMSG = _descriptor.Descriptor(
  name='APPacketsGetNotifMsg',
  full_name='cheetah.APPacketsGetNotifMsg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=51,
  serialized_end=73,
)


_APPACKETSREGMSG = _descriptor.Descriptor(
  name='APPacketsRegMsg',
  full_name='cheetah.APPacketsRegMsg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Oper', full_name='cheetah.APPacketsRegMsg.Oper', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='PacketHdr', full_name='cheetah.APPacketsRegMsg.PacketHdr', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=75,
  serialized_end=165,
)


_APPACKETSREGMSGRSP = _descriptor.Descriptor(
  name='APPacketsRegMsgRsp',
  full_name='cheetah.APPacketsRegMsgRsp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ErrStatus', full_name='cheetah.APPacketsRegMsgRsp.ErrStatus', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Results', full_name='cheetah.APPacketsRegMsgRsp.Results', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=167,
  serialized_end=273,
)


_APPACKETHDR = _descriptor.Descriptor(
  name='APPacketHdr',
  full_name='cheetah.APPacketHdr',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='MsgType', full_name='cheetah.APPacketHdr.MsgType', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='mgmt', full_name='cheetah.APPacketHdr.mgmt', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ctrl', full_name='cheetah.APPacketHdr.ctrl', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='data', full_name='cheetah.APPacketHdr.data', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='cisco', full_name='cheetah.APPacketHdr.cisco', index=4,
      number=5, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='Subtype', full_name='cheetah.APPacketHdr.Subtype',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=276,
  serialized_end=511,
)


_APPACKETSMSGRSP = _descriptor.Descriptor(
  name='APPacketsMsgRsp',
  full_name='cheetah.APPacketsMsgRsp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ErrStatus', full_name='cheetah.APPacketsMsgRsp.ErrStatus', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='PacketHdr', full_name='cheetah.APPacketsMsgRsp.PacketHdr', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='PacketLen', full_name='cheetah.APPacketsMsgRsp.PacketLen', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='PacketBuf', full_name='cheetah.APPacketsMsgRsp.PacketBuf', index=3,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=514,
  serialized_end=653,
)

_APPACKETSREGMSG.fields_by_name['Oper'].enum_type = ap__common__types__pb2._APREGOP
_APPACKETSREGMSG.fields_by_name['PacketHdr'].message_type = _APPACKETHDR
_APPACKETSREGMSGRSP.fields_by_name['ErrStatus'].message_type = ap__common__types__pb2._APERRORSTATUS
_APPACKETSREGMSGRSP.fields_by_name['Results'].message_type = _APPACKETSREGMSG
_APPACKETHDR.fields_by_name['MsgType'].enum_type = _APMSGTYPE
_APPACKETHDR.fields_by_name['mgmt'].enum_type = _APMGMTMSGSUBTYPE
_APPACKETHDR.fields_by_name['ctrl'].enum_type = _APCTRLMSGSUBTYPE
_APPACKETHDR.fields_by_name['data'].enum_type = _APDATAMSGSUBTYPE
_APPACKETHDR.fields_by_name['cisco'].enum_type = _APCISCOMSGSUBTYPE
_APPACKETHDR.oneofs_by_name['Subtype'].fields.append(
  _APPACKETHDR.fields_by_name['mgmt'])
_APPACKETHDR.fields_by_name['mgmt'].containing_oneof = _APPACKETHDR.oneofs_by_name['Subtype']
_APPACKETHDR.oneofs_by_name['Subtype'].fields.append(
  _APPACKETHDR.fields_by_name['ctrl'])
_APPACKETHDR.fields_by_name['ctrl'].containing_oneof = _APPACKETHDR.oneofs_by_name['Subtype']
_APPACKETHDR.oneofs_by_name['Subtype'].fields.append(
  _APPACKETHDR.fields_by_name['data'])
_APPACKETHDR.fields_by_name['data'].containing_oneof = _APPACKETHDR.oneofs_by_name['Subtype']
_APPACKETHDR.oneofs_by_name['Subtype'].fields.append(
  _APPACKETHDR.fields_by_name['cisco'])
_APPACKETHDR.fields_by_name['cisco'].containing_oneof = _APPACKETHDR.oneofs_by_name['Subtype']
_APPACKETSMSGRSP.fields_by_name['ErrStatus'].message_type = ap__common__types__pb2._APERRORSTATUS
_APPACKETSMSGRSP.fields_by_name['PacketHdr'].message_type = _APPACKETHDR
DESCRIPTOR.message_types_by_name['APPacketsGetNotifMsg'] = _APPACKETSGETNOTIFMSG
DESCRIPTOR.message_types_by_name['APPacketsRegMsg'] = _APPACKETSREGMSG
DESCRIPTOR.message_types_by_name['APPacketsRegMsgRsp'] = _APPACKETSREGMSGRSP
DESCRIPTOR.message_types_by_name['APPacketHdr'] = _APPACKETHDR
DESCRIPTOR.message_types_by_name['APPacketsMsgRsp'] = _APPACKETSMSGRSP
DESCRIPTOR.enum_types_by_name['APMsgType'] = _APMSGTYPE
DESCRIPTOR.enum_types_by_name['APMgmtMsgSubtype'] = _APMGMTMSGSUBTYPE
DESCRIPTOR.enum_types_by_name['APCtrlMsgSubtype'] = _APCTRLMSGSUBTYPE
DESCRIPTOR.enum_types_by_name['APDataMsgSubtype'] = _APDATAMSGSUBTYPE
DESCRIPTOR.enum_types_by_name['APCiscoMsgSubtype'] = _APCISCOMSGSUBTYPE

APPacketsGetNotifMsg = _reflection.GeneratedProtocolMessageType('APPacketsGetNotifMsg', (_message.Message,), dict(
  DESCRIPTOR = _APPACKETSGETNOTIFMSG,
  __module__ = 'ap_packet_pb2'
  # @@protoc_insertion_point(class_scope:cheetah.APPacketsGetNotifMsg)
  ))
_sym_db.RegisterMessage(APPacketsGetNotifMsg)

APPacketsRegMsg = _reflection.GeneratedProtocolMessageType('APPacketsRegMsg', (_message.Message,), dict(
  DESCRIPTOR = _APPACKETSREGMSG,
  __module__ = 'ap_packet_pb2'
  # @@protoc_insertion_point(class_scope:cheetah.APPacketsRegMsg)
  ))
_sym_db.RegisterMessage(APPacketsRegMsg)

APPacketsRegMsgRsp = _reflection.GeneratedProtocolMessageType('APPacketsRegMsgRsp', (_message.Message,), dict(
  DESCRIPTOR = _APPACKETSREGMSGRSP,
  __module__ = 'ap_packet_pb2'
  # @@protoc_insertion_point(class_scope:cheetah.APPacketsRegMsgRsp)
  ))
_sym_db.RegisterMessage(APPacketsRegMsgRsp)

APPacketHdr = _reflection.GeneratedProtocolMessageType('APPacketHdr', (_message.Message,), dict(
  DESCRIPTOR = _APPACKETHDR,
  __module__ = 'ap_packet_pb2'
  # @@protoc_insertion_point(class_scope:cheetah.APPacketHdr)
  ))
_sym_db.RegisterMessage(APPacketHdr)

APPacketsMsgRsp = _reflection.GeneratedProtocolMessageType('APPacketsMsgRsp', (_message.Message,), dict(
  DESCRIPTOR = _APPACKETSMSGRSP,
  __module__ = 'ap_packet_pb2'
  # @@protoc_insertion_point(class_scope:cheetah.APPacketsMsgRsp)
  ))
_sym_db.RegisterMessage(APPacketsMsgRsp)


try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities


  class APPacketsStub(object):
    """@defgroup APPackets
    @ingroup System
    The following RPCs are used to register for packets that the
    client is interested in receiving
    @{

    Packet registration operations

    """

    def __init__(self, channel):
      """Constructor.

      Args:
        channel: A grpc.Channel.
      """
      self.APPacketsRegOp = channel.unary_unary(
          '/cheetah.APPackets/APPacketsRegOp',
          request_serializer=APPacketsRegMsg.SerializeToString,
          response_deserializer=APPacketsRegMsgRsp.FromString,
          )
      self.APPacketsInitNotif = channel.unary_stream(
          '/cheetah.APPackets/APPacketsInitNotif',
          request_serializer=APPacketsGetNotifMsg.SerializeToString,
          response_deserializer=APPacketsMsgRsp.FromString,
          )


  class APPacketsServicer(object):
    """@defgroup APPackets
    @ingroup System
    The following RPCs are used to register for packets that the
    client is interested in receiving
    @{

    Packet registration operations

    """

    def APPacketsRegOp(self, request, context):
      """APPacketsRegMsg.Oper = AP_REGOP_REGISTER
      Packet registration: Sends a list of Packet registration messages
      and expects a list of registration responses.

      APPacketsRegMsg.Oper = AP_REGOP_UNREGISTER
      Packet unregistration: Sends a list of Packet unregistration messages
      and expects a list of unregistration responses.

      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def APPacketsInitNotif(self, request, context):
      """
      Packet notifications


      This call is used to get a stream of packet notifications matching the
      set of registrations performed with APPacketsRegOp().
      The caller must maintain the GRPC channel as long as
      there is interest in packet notifications. Only sessions that were
      created through this API will be notified to caller.
      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')


  def add_APPacketsServicer_to_server(servicer, server):
    rpc_method_handlers = {
        'APPacketsRegOp': grpc.unary_unary_rpc_method_handler(
            servicer.APPacketsRegOp,
            request_deserializer=APPacketsRegMsg.FromString,
            response_serializer=APPacketsRegMsgRsp.SerializeToString,
        ),
        'APPacketsInitNotif': grpc.unary_stream_rpc_method_handler(
            servicer.APPacketsInitNotif,
            request_deserializer=APPacketsGetNotifMsg.FromString,
            response_serializer=APPacketsMsgRsp.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        'cheetah.APPackets', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


  class BetaAPPacketsServicer(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    """@defgroup APPackets
    @ingroup System
    The following RPCs are used to register for packets that the
    client is interested in receiving
    @{

    Packet registration operations

    """
    def APPacketsRegOp(self, request, context):
      """APPacketsRegMsg.Oper = AP_REGOP_REGISTER
      Packet registration: Sends a list of Packet registration messages
      and expects a list of registration responses.

      APPacketsRegMsg.Oper = AP_REGOP_UNREGISTER
      Packet unregistration: Sends a list of Packet unregistration messages
      and expects a list of unregistration responses.

      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def APPacketsInitNotif(self, request, context):
      """
      Packet notifications


      This call is used to get a stream of packet notifications matching the
      set of registrations performed with APPacketsRegOp().
      The caller must maintain the GRPC channel as long as
      there is interest in packet notifications. Only sessions that were
      created through this API will be notified to caller.
      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)


  class BetaAPPacketsStub(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    """@defgroup APPackets
    @ingroup System
    The following RPCs are used to register for packets that the
    client is interested in receiving
    @{

    Packet registration operations

    """
    def APPacketsRegOp(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """APPacketsRegMsg.Oper = AP_REGOP_REGISTER
      Packet registration: Sends a list of Packet registration messages
      and expects a list of registration responses.

      APPacketsRegMsg.Oper = AP_REGOP_UNREGISTER
      Packet unregistration: Sends a list of Packet unregistration messages
      and expects a list of unregistration responses.

      """
      raise NotImplementedError()
    APPacketsRegOp.future = None
    def APPacketsInitNotif(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """
      Packet notifications


      This call is used to get a stream of packet notifications matching the
      set of registrations performed with APPacketsRegOp().
      The caller must maintain the GRPC channel as long as
      there is interest in packet notifications. Only sessions that were
      created through this API will be notified to caller.
      """
      raise NotImplementedError()


  def beta_create_APPackets_server(servicer, pool=None, pool_size=None, default_timeout=None, maximum_timeout=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_deserializers = {
      ('cheetah.APPackets', 'APPacketsInitNotif'): APPacketsGetNotifMsg.FromString,
      ('cheetah.APPackets', 'APPacketsRegOp'): APPacketsRegMsg.FromString,
    }
    response_serializers = {
      ('cheetah.APPackets', 'APPacketsInitNotif'): APPacketsMsgRsp.SerializeToString,
      ('cheetah.APPackets', 'APPacketsRegOp'): APPacketsRegMsgRsp.SerializeToString,
    }
    method_implementations = {
      ('cheetah.APPackets', 'APPacketsInitNotif'): face_utilities.unary_stream_inline(servicer.APPacketsInitNotif),
      ('cheetah.APPackets', 'APPacketsRegOp'): face_utilities.unary_unary_inline(servicer.APPacketsRegOp),
    }
    server_options = beta_implementations.server_options(request_deserializers=request_deserializers, response_serializers=response_serializers, thread_pool=pool, thread_pool_size=pool_size, default_timeout=default_timeout, maximum_timeout=maximum_timeout)
    return beta_implementations.server(method_implementations, options=server_options)


  def beta_create_APPackets_stub(channel, host=None, metadata_transformer=None, pool=None, pool_size=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_serializers = {
      ('cheetah.APPackets', 'APPacketsInitNotif'): APPacketsGetNotifMsg.SerializeToString,
      ('cheetah.APPackets', 'APPacketsRegOp'): APPacketsRegMsg.SerializeToString,
    }
    response_deserializers = {
      ('cheetah.APPackets', 'APPacketsInitNotif'): APPacketsMsgRsp.FromString,
      ('cheetah.APPackets', 'APPacketsRegOp'): APPacketsRegMsgRsp.FromString,
    }
    cardinalities = {
      'APPacketsInitNotif': cardinality.Cardinality.UNARY_STREAM,
      'APPacketsRegOp': cardinality.Cardinality.UNARY_UNARY,
    }
    stub_options = beta_implementations.stub_options(host=host, metadata_transformer=metadata_transformer, request_serializers=request_serializers, response_deserializers=response_deserializers, thread_pool=pool, thread_pool_size=pool_size)
    return beta_implementations.dynamic_stub(channel, 'cheetah.APPackets', cardinalities, options=stub_options)
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)
