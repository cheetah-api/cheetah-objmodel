/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: ap_version.proto */

#ifndef PROTOBUF_C_ap_5fversion_2eproto__INCLUDED
#define PROTOBUF_C_ap_5fversion_2eproto__INCLUDED

#include <protobuf-c/protobuf-c.h>

PROTOBUF_C__BEGIN_DECLS

#if PROTOBUF_C_VERSION_NUMBER < 1003000
# error This file was generated by a newer version of protoc-c which is incompatible with your libprotobuf-c headers. Please update your headers.
#elif 1003000 < PROTOBUF_C_MIN_COMPILER_VERSION
# error This file was generated by an older version of protoc-c which is incompatible with your libprotobuf-c headers. Please regenerate this file with a newer version of protoc-c.
#endif




/* --- enums --- */

/*
 * Access Point API version.
 * This is used in the Global init message exchange to handshake client/server
 * Version numbers.
 */
typedef enum _Cheetah__APVersion {
  CHEETAH__APVERSION__AP_VERSION_UNUSED = 0,
  CHEETAH__APVERSION__AP_MAJOR_VERSION = 0,
  CHEETAH__APVERSION__AP_MINOR_VERSION = 0,
  CHEETAH__APVERSION__AP_SUB_VERSION = 1
    PROTOBUF_C__FORCE_ENUM_TO_BE_INT_SIZE(CHEETAH__APVERSION)
} Cheetah__APVersion;

/* --- messages --- */

/* --- per-message closures --- */


/* --- services --- */


/* --- descriptors --- */

extern const ProtobufCEnumDescriptor    cheetah__apversion__descriptor;

PROTOBUF_C__END_DECLS


#endif  /* PROTOBUF_C_ap_5fversion_2eproto__INCLUDED */