// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: ap_common_types.proto

#ifndef PROTOBUF_ap_5fcommon_5ftypes_2eproto__INCLUDED
#define PROTOBUF_ap_5fcommon_5ftypes_2eproto__INCLUDED

#include <string>

#include <google/protobuf/stubs/common.h>

#if GOOGLE_PROTOBUF_VERSION < 3001000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please update
#error your headers.
#endif
#if 3001000 < GOOGLE_PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/generated_enum_reflection.h>
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)

namespace cheetah {

// Internal implementation detail -- do not call these.
void protobuf_AddDesc_ap_5fcommon_5ftypes_2eproto();
void protobuf_InitDefaults_ap_5fcommon_5ftypes_2eproto();
void protobuf_AssignDesc_ap_5fcommon_5ftypes_2eproto();
void protobuf_ShutdownFile_ap_5fcommon_5ftypes_2eproto();

class APErrorStatus;
class APRadio;
class APSsid;

enum APErrorStatus_APErrno {
  APErrorStatus_APErrno_AP_SUCCESS = 0,
  APErrorStatus_APErrno_AP_NOT_CONNECTED = 1,
  APErrorStatus_APErrno_AP_EAGAIN = 2,
  APErrorStatus_APErrno_AP_ENOMEM = 3,
  APErrorStatus_APErrno_AP_EBUSY = 4,
  APErrorStatus_APErrno_AP_EINVAL = 5,
  APErrorStatus_APErrno_AP_UNSUPPORTED_VER = 6,
  APErrorStatus_APErrno_AP_NOT_AVAILABLE = 7,
  APErrorStatus_APErrno_AP_STREAM_NOT_SUPPORTED = 8,
  APErrorStatus_APErrno_AP_ENOTSUP = 9,
  APErrorStatus_APErrno_AP_SOME_ERR = 10,
  APErrorStatus_APErrno_AP_TIMEOUT = 11,
  APErrorStatus_APErrno_AP_NOTIF_TERM = 12,
  APErrorStatus_APErrno_AP_INIT_START_OFFSET = 1280,
  APErrorStatus_APErrno_AP_INIT_STATE_CLEAR = 1281,
  APErrorStatus_APErrno_AP_INIT_STATE_READY = 1282,
  APErrorStatus_APErrno_AP_INIT_UNSUPPORTED_VER = 1283,
  APErrorStatus_APErrno_AP_INIT_SERVER_NOT_INITIALIZED = 1284,
  APErrorStatus_APErrno_AP_INIT_SERVER_MODE_CHANGE_FAILED = 1285,
  APErrorStatus_APErrno_APErrorStatus_APErrno_INT_MIN_SENTINEL_DO_NOT_USE_ = ::google::protobuf::kint32min,
  APErrorStatus_APErrno_APErrorStatus_APErrno_INT_MAX_SENTINEL_DO_NOT_USE_ = ::google::protobuf::kint32max
};
bool APErrorStatus_APErrno_IsValid(int value);
const APErrorStatus_APErrno APErrorStatus_APErrno_APErrno_MIN = APErrorStatus_APErrno_AP_SUCCESS;
const APErrorStatus_APErrno APErrorStatus_APErrno_APErrno_MAX = APErrorStatus_APErrno_AP_INIT_SERVER_MODE_CHANGE_FAILED;
const int APErrorStatus_APErrno_APErrno_ARRAYSIZE = APErrorStatus_APErrno_APErrno_MAX + 1;

const ::google::protobuf::EnumDescriptor* APErrorStatus_APErrno_descriptor();
inline const ::std::string& APErrorStatus_APErrno_Name(APErrorStatus_APErrno value) {
  return ::google::protobuf::internal::NameOfEnum(
    APErrorStatus_APErrno_descriptor(), value);
}
inline bool APErrorStatus_APErrno_Parse(
    const ::std::string& name, APErrorStatus_APErrno* value) {
  return ::google::protobuf::internal::ParseNamedEnum<APErrorStatus_APErrno>(
    APErrorStatus_APErrno_descriptor(), name, value);
}
enum APRegOp {
  AP_REGOP_RESERVED = 0,
  AP_REGOP_REGISTER = 1,
  AP_REGOP_UNREGISTER = 2,
  AP_REGOP_EOF = 3,
  APRegOp_INT_MIN_SENTINEL_DO_NOT_USE_ = ::google::protobuf::kint32min,
  APRegOp_INT_MAX_SENTINEL_DO_NOT_USE_ = ::google::protobuf::kint32max
};
bool APRegOp_IsValid(int value);
const APRegOp APRegOp_MIN = AP_REGOP_RESERVED;
const APRegOp APRegOp_MAX = AP_REGOP_EOF;
const int APRegOp_ARRAYSIZE = APRegOp_MAX + 1;

const ::google::protobuf::EnumDescriptor* APRegOp_descriptor();
inline const ::std::string& APRegOp_Name(APRegOp value) {
  return ::google::protobuf::internal::NameOfEnum(
    APRegOp_descriptor(), value);
}
inline bool APRegOp_Parse(
    const ::std::string& name, APRegOp* value) {
  return ::google::protobuf::internal::ParseNamedEnum<APRegOp>(
    APRegOp_descriptor(), name, value);
}
enum APObjectOp {
  AP_OBJOP_RESERVED = 0,
  AP_OBJOP_ADD = 1,
  AP_OBJOP_UPDATE = 2,
  AP_OBJOP_DELETE = 3,
  APObjectOp_INT_MIN_SENTINEL_DO_NOT_USE_ = ::google::protobuf::kint32min,
  APObjectOp_INT_MAX_SENTINEL_DO_NOT_USE_ = ::google::protobuf::kint32max
};
bool APObjectOp_IsValid(int value);
const APObjectOp APObjectOp_MIN = AP_OBJOP_RESERVED;
const APObjectOp APObjectOp_MAX = AP_OBJOP_DELETE;
const int APObjectOp_ARRAYSIZE = APObjectOp_MAX + 1;

const ::google::protobuf::EnumDescriptor* APObjectOp_descriptor();
inline const ::std::string& APObjectOp_Name(APObjectOp value) {
  return ::google::protobuf::internal::NameOfEnum(
    APObjectOp_descriptor(), value);
}
inline bool APObjectOp_Parse(
    const ::std::string& name, APObjectOp* value) {
  return ::google::protobuf::internal::ParseNamedEnum<APObjectOp>(
    APObjectOp_descriptor(), name, value);
}
enum APNotifOp {
  AP_NOTIFOP_RESERVED = 0,
  AP_NOTIFOP_ENABLE = 1,
  AP_NOTIFOP_DISABLE = 2,
  APNotifOp_INT_MIN_SENTINEL_DO_NOT_USE_ = ::google::protobuf::kint32min,
  APNotifOp_INT_MAX_SENTINEL_DO_NOT_USE_ = ::google::protobuf::kint32max
};
bool APNotifOp_IsValid(int value);
const APNotifOp APNotifOp_MIN = AP_NOTIFOP_RESERVED;
const APNotifOp APNotifOp_MAX = AP_NOTIFOP_DISABLE;
const int APNotifOp_ARRAYSIZE = APNotifOp_MAX + 1;

const ::google::protobuf::EnumDescriptor* APNotifOp_descriptor();
inline const ::std::string& APNotifOp_Name(APNotifOp value) {
  return ::google::protobuf::internal::NameOfEnum(
    APNotifOp_descriptor(), value);
}
inline bool APNotifOp_Parse(
    const ::std::string& name, APNotifOp* value) {
  return ::google::protobuf::internal::ParseNamedEnum<APNotifOp>(
    APNotifOp_descriptor(), name, value);
}
// ===================================================================

class APErrorStatus : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:cheetah.APErrorStatus) */ {
 public:
  APErrorStatus();
  virtual ~APErrorStatus();

  APErrorStatus(const APErrorStatus& from);

  inline APErrorStatus& operator=(const APErrorStatus& from) {
    CopyFrom(from);
    return *this;
  }

  static const ::google::protobuf::Descriptor* descriptor();
  static const APErrorStatus& default_instance();

  static const APErrorStatus* internal_default_instance();

  void Swap(APErrorStatus* other);

  // implements Message ----------------------------------------------

  inline APErrorStatus* New() const { return New(NULL); }

  APErrorStatus* New(::google::protobuf::Arena* arena) const;
  void CopyFrom(const ::google::protobuf::Message& from);
  void MergeFrom(const ::google::protobuf::Message& from);
  void CopyFrom(const APErrorStatus& from);
  void MergeFrom(const APErrorStatus& from);
  void Clear();
  bool IsInitialized() const;

  size_t ByteSizeLong() const;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input);
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* output) const;
  ::google::protobuf::uint8* SerializeWithCachedSizesToArray(::google::protobuf::uint8* output) const {
    return InternalSerializeWithCachedSizesToArray(false, output);
  }
  int GetCachedSize() const { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const;
  void InternalSwap(APErrorStatus* other);
  void UnsafeMergeFrom(const APErrorStatus& from);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return _internal_metadata_.arena();
  }
  inline void* MaybeArenaPtr() const {
    return _internal_metadata_.raw_arena_ptr();
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const;

  // nested types ----------------------------------------------------

  typedef APErrorStatus_APErrno APErrno;
  static const APErrno AP_SUCCESS =
    APErrorStatus_APErrno_AP_SUCCESS;
  static const APErrno AP_NOT_CONNECTED =
    APErrorStatus_APErrno_AP_NOT_CONNECTED;
  static const APErrno AP_EAGAIN =
    APErrorStatus_APErrno_AP_EAGAIN;
  static const APErrno AP_ENOMEM =
    APErrorStatus_APErrno_AP_ENOMEM;
  static const APErrno AP_EBUSY =
    APErrorStatus_APErrno_AP_EBUSY;
  static const APErrno AP_EINVAL =
    APErrorStatus_APErrno_AP_EINVAL;
  static const APErrno AP_UNSUPPORTED_VER =
    APErrorStatus_APErrno_AP_UNSUPPORTED_VER;
  static const APErrno AP_NOT_AVAILABLE =
    APErrorStatus_APErrno_AP_NOT_AVAILABLE;
  static const APErrno AP_STREAM_NOT_SUPPORTED =
    APErrorStatus_APErrno_AP_STREAM_NOT_SUPPORTED;
  static const APErrno AP_ENOTSUP =
    APErrorStatus_APErrno_AP_ENOTSUP;
  static const APErrno AP_SOME_ERR =
    APErrorStatus_APErrno_AP_SOME_ERR;
  static const APErrno AP_TIMEOUT =
    APErrorStatus_APErrno_AP_TIMEOUT;
  static const APErrno AP_NOTIF_TERM =
    APErrorStatus_APErrno_AP_NOTIF_TERM;
  static const APErrno AP_INIT_START_OFFSET =
    APErrorStatus_APErrno_AP_INIT_START_OFFSET;
  static const APErrno AP_INIT_STATE_CLEAR =
    APErrorStatus_APErrno_AP_INIT_STATE_CLEAR;
  static const APErrno AP_INIT_STATE_READY =
    APErrorStatus_APErrno_AP_INIT_STATE_READY;
  static const APErrno AP_INIT_UNSUPPORTED_VER =
    APErrorStatus_APErrno_AP_INIT_UNSUPPORTED_VER;
  static const APErrno AP_INIT_SERVER_NOT_INITIALIZED =
    APErrorStatus_APErrno_AP_INIT_SERVER_NOT_INITIALIZED;
  static const APErrno AP_INIT_SERVER_MODE_CHANGE_FAILED =
    APErrorStatus_APErrno_AP_INIT_SERVER_MODE_CHANGE_FAILED;
  static inline bool APErrno_IsValid(int value) {
    return APErrorStatus_APErrno_IsValid(value);
  }
  static const APErrno APErrno_MIN =
    APErrorStatus_APErrno_APErrno_MIN;
  static const APErrno APErrno_MAX =
    APErrorStatus_APErrno_APErrno_MAX;
  static const int APErrno_ARRAYSIZE =
    APErrorStatus_APErrno_APErrno_ARRAYSIZE;
  static inline const ::google::protobuf::EnumDescriptor*
  APErrno_descriptor() {
    return APErrorStatus_APErrno_descriptor();
  }
  static inline const ::std::string& APErrno_Name(APErrno value) {
    return APErrorStatus_APErrno_Name(value);
  }
  static inline bool APErrno_Parse(const ::std::string& name,
      APErrno* value) {
    return APErrorStatus_APErrno_Parse(name, value);
  }

  // accessors -------------------------------------------------------

  // optional .cheetah.APErrorStatus.APErrno Status = 1;
  void clear_status();
  static const int kStatusFieldNumber = 1;
  ::cheetah::APErrorStatus_APErrno status() const;
  void set_status(::cheetah::APErrorStatus_APErrno value);

  // @@protoc_insertion_point(class_scope:cheetah.APErrorStatus)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  int status_;
  mutable int _cached_size_;
  friend void  protobuf_InitDefaults_ap_5fcommon_5ftypes_2eproto_impl();
  friend void  protobuf_AddDesc_ap_5fcommon_5ftypes_2eproto_impl();
  friend void protobuf_AssignDesc_ap_5fcommon_5ftypes_2eproto();
  friend void protobuf_ShutdownFile_ap_5fcommon_5ftypes_2eproto();

  void InitAsDefaultInstance();
};
extern ::google::protobuf::internal::ExplicitlyConstructed<APErrorStatus> APErrorStatus_default_instance_;

// -------------------------------------------------------------------

class APRadio : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:cheetah.APRadio) */ {
 public:
  APRadio();
  virtual ~APRadio();

  APRadio(const APRadio& from);

  inline APRadio& operator=(const APRadio& from) {
    CopyFrom(from);
    return *this;
  }

  static const ::google::protobuf::Descriptor* descriptor();
  static const APRadio& default_instance();

  enum RadioCase {
    kName = 1,
    kHandle = 2,
    RADIO_NOT_SET = 0,
  };

  static const APRadio* internal_default_instance();

  void Swap(APRadio* other);

  // implements Message ----------------------------------------------

  inline APRadio* New() const { return New(NULL); }

  APRadio* New(::google::protobuf::Arena* arena) const;
  void CopyFrom(const ::google::protobuf::Message& from);
  void MergeFrom(const ::google::protobuf::Message& from);
  void CopyFrom(const APRadio& from);
  void MergeFrom(const APRadio& from);
  void Clear();
  bool IsInitialized() const;

  size_t ByteSizeLong() const;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input);
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* output) const;
  ::google::protobuf::uint8* SerializeWithCachedSizesToArray(::google::protobuf::uint8* output) const {
    return InternalSerializeWithCachedSizesToArray(false, output);
  }
  int GetCachedSize() const { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const;
  void InternalSwap(APRadio* other);
  void UnsafeMergeFrom(const APRadio& from);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return _internal_metadata_.arena();
  }
  inline void* MaybeArenaPtr() const {
    return _internal_metadata_.raw_arena_ptr();
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // optional string Name = 1;
  private:
  bool has_name() const;
  public:
  void clear_name();
  static const int kNameFieldNumber = 1;
  const ::std::string& name() const;
  void set_name(const ::std::string& value);
  void set_name(const char* value);
  void set_name(const char* value, size_t size);
  ::std::string* mutable_name();
  ::std::string* release_name();
  void set_allocated_name(::std::string* name);

  // optional uint32 Handle = 2;
  private:
  bool has_handle() const;
  public:
  void clear_handle();
  static const int kHandleFieldNumber = 2;
  ::google::protobuf::uint32 handle() const;
  void set_handle(::google::protobuf::uint32 value);

  RadioCase Radio_case() const;
  // @@protoc_insertion_point(class_scope:cheetah.APRadio)
 private:
  inline void set_has_name();
  inline void set_has_handle();

  inline bool has_Radio() const;
  void clear_Radio();
  inline void clear_has_Radio();

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  union RadioUnion {
    RadioUnion() {}
    ::google::protobuf::internal::ArenaStringPtr name_;
    ::google::protobuf::uint32 handle_;
  } Radio_;
  mutable int _cached_size_;
  ::google::protobuf::uint32 _oneof_case_[1];

  friend void  protobuf_InitDefaults_ap_5fcommon_5ftypes_2eproto_impl();
  friend void  protobuf_AddDesc_ap_5fcommon_5ftypes_2eproto_impl();
  friend void protobuf_AssignDesc_ap_5fcommon_5ftypes_2eproto();
  friend void protobuf_ShutdownFile_ap_5fcommon_5ftypes_2eproto();

  void InitAsDefaultInstance();
};
extern ::google::protobuf::internal::ExplicitlyConstructed<APRadio> APRadio_default_instance_;

// -------------------------------------------------------------------

class APSsid : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:cheetah.APSsid) */ {
 public:
  APSsid();
  virtual ~APSsid();

  APSsid(const APSsid& from);

  inline APSsid& operator=(const APSsid& from) {
    CopyFrom(from);
    return *this;
  }

  static const ::google::protobuf::Descriptor* descriptor();
  static const APSsid& default_instance();

  enum SsidCase {
    kName = 1,
    kHandle = 2,
    SSID_NOT_SET = 0,
  };

  static const APSsid* internal_default_instance();

  void Swap(APSsid* other);

  // implements Message ----------------------------------------------

  inline APSsid* New() const { return New(NULL); }

  APSsid* New(::google::protobuf::Arena* arena) const;
  void CopyFrom(const ::google::protobuf::Message& from);
  void MergeFrom(const ::google::protobuf::Message& from);
  void CopyFrom(const APSsid& from);
  void MergeFrom(const APSsid& from);
  void Clear();
  bool IsInitialized() const;

  size_t ByteSizeLong() const;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input);
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* output) const;
  ::google::protobuf::uint8* SerializeWithCachedSizesToArray(::google::protobuf::uint8* output) const {
    return InternalSerializeWithCachedSizesToArray(false, output);
  }
  int GetCachedSize() const { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const;
  void InternalSwap(APSsid* other);
  void UnsafeMergeFrom(const APSsid& from);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return _internal_metadata_.arena();
  }
  inline void* MaybeArenaPtr() const {
    return _internal_metadata_.raw_arena_ptr();
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // optional string Name = 1;
  private:
  bool has_name() const;
  public:
  void clear_name();
  static const int kNameFieldNumber = 1;
  const ::std::string& name() const;
  void set_name(const ::std::string& value);
  void set_name(const char* value);
  void set_name(const char* value, size_t size);
  ::std::string* mutable_name();
  ::std::string* release_name();
  void set_allocated_name(::std::string* name);

  // optional uint32 Handle = 2;
  private:
  bool has_handle() const;
  public:
  void clear_handle();
  static const int kHandleFieldNumber = 2;
  ::google::protobuf::uint32 handle() const;
  void set_handle(::google::protobuf::uint32 value);

  SsidCase Ssid_case() const;
  // @@protoc_insertion_point(class_scope:cheetah.APSsid)
 private:
  inline void set_has_name();
  inline void set_has_handle();

  inline bool has_Ssid() const;
  void clear_Ssid();
  inline void clear_has_Ssid();

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  union SsidUnion {
    SsidUnion() {}
    ::google::protobuf::internal::ArenaStringPtr name_;
    ::google::protobuf::uint32 handle_;
  } Ssid_;
  mutable int _cached_size_;
  ::google::protobuf::uint32 _oneof_case_[1];

  friend void  protobuf_InitDefaults_ap_5fcommon_5ftypes_2eproto_impl();
  friend void  protobuf_AddDesc_ap_5fcommon_5ftypes_2eproto_impl();
  friend void protobuf_AssignDesc_ap_5fcommon_5ftypes_2eproto();
  friend void protobuf_ShutdownFile_ap_5fcommon_5ftypes_2eproto();

  void InitAsDefaultInstance();
};
extern ::google::protobuf::internal::ExplicitlyConstructed<APSsid> APSsid_default_instance_;

// ===================================================================


// ===================================================================

#if !PROTOBUF_INLINE_NOT_IN_HEADERS
// APErrorStatus

// optional .cheetah.APErrorStatus.APErrno Status = 1;
inline void APErrorStatus::clear_status() {
  status_ = 0;
}
inline ::cheetah::APErrorStatus_APErrno APErrorStatus::status() const {
  // @@protoc_insertion_point(field_get:cheetah.APErrorStatus.Status)
  return static_cast< ::cheetah::APErrorStatus_APErrno >(status_);
}
inline void APErrorStatus::set_status(::cheetah::APErrorStatus_APErrno value) {
  
  status_ = value;
  // @@protoc_insertion_point(field_set:cheetah.APErrorStatus.Status)
}

inline const APErrorStatus* APErrorStatus::internal_default_instance() {
  return &APErrorStatus_default_instance_.get();
}
// -------------------------------------------------------------------

// APRadio

// optional string Name = 1;
inline bool APRadio::has_name() const {
  return Radio_case() == kName;
}
inline void APRadio::set_has_name() {
  _oneof_case_[0] = kName;
}
inline void APRadio::clear_name() {
  if (has_name()) {
    Radio_.name_.DestroyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
    clear_has_Radio();
  }
}
inline const ::std::string& APRadio::name() const {
  // @@protoc_insertion_point(field_get:cheetah.APRadio.Name)
  if (has_name()) {
    return Radio_.name_.GetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  return *&::google::protobuf::internal::GetEmptyStringAlreadyInited();
}
inline void APRadio::set_name(const ::std::string& value) {
  // @@protoc_insertion_point(field_set:cheetah.APRadio.Name)
  if (!has_name()) {
    clear_Radio();
    set_has_name();
    Radio_.name_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  Radio_.name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:cheetah.APRadio.Name)
}
inline void APRadio::set_name(const char* value) {
  if (!has_name()) {
    clear_Radio();
    set_has_name();
    Radio_.name_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  Radio_.name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(value));
  // @@protoc_insertion_point(field_set_char:cheetah.APRadio.Name)
}
inline void APRadio::set_name(const char* value, size_t size) {
  if (!has_name()) {
    clear_Radio();
    set_has_name();
    Radio_.name_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  Radio_.name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(
      reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:cheetah.APRadio.Name)
}
inline ::std::string* APRadio::mutable_name() {
  if (!has_name()) {
    clear_Radio();
    set_has_name();
    Radio_.name_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  // @@protoc_insertion_point(field_mutable:cheetah.APRadio.Name)
  return Radio_.name_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* APRadio::release_name() {
  // @@protoc_insertion_point(field_release:cheetah.APRadio.Name)
  if (has_name()) {
    clear_has_Radio();
    return Radio_.name_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  } else {
    return NULL;
  }
}
inline void APRadio::set_allocated_name(::std::string* name) {
  if (!has_name()) {
    Radio_.name_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  clear_Radio();
  if (name != NULL) {
    set_has_name();
    Radio_.name_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
        name);
  }
  // @@protoc_insertion_point(field_set_allocated:cheetah.APRadio.Name)
}

// optional uint32 Handle = 2;
inline bool APRadio::has_handle() const {
  return Radio_case() == kHandle;
}
inline void APRadio::set_has_handle() {
  _oneof_case_[0] = kHandle;
}
inline void APRadio::clear_handle() {
  if (has_handle()) {
    Radio_.handle_ = 0u;
    clear_has_Radio();
  }
}
inline ::google::protobuf::uint32 APRadio::handle() const {
  // @@protoc_insertion_point(field_get:cheetah.APRadio.Handle)
  if (has_handle()) {
    return Radio_.handle_;
  }
  return 0u;
}
inline void APRadio::set_handle(::google::protobuf::uint32 value) {
  if (!has_handle()) {
    clear_Radio();
    set_has_handle();
  }
  Radio_.handle_ = value;
  // @@protoc_insertion_point(field_set:cheetah.APRadio.Handle)
}

inline bool APRadio::has_Radio() const {
  return Radio_case() != RADIO_NOT_SET;
}
inline void APRadio::clear_has_Radio() {
  _oneof_case_[0] = RADIO_NOT_SET;
}
inline APRadio::RadioCase APRadio::Radio_case() const {
  return APRadio::RadioCase(_oneof_case_[0]);
}
inline const APRadio* APRadio::internal_default_instance() {
  return &APRadio_default_instance_.get();
}
// -------------------------------------------------------------------

// APSsid

// optional string Name = 1;
inline bool APSsid::has_name() const {
  return Ssid_case() == kName;
}
inline void APSsid::set_has_name() {
  _oneof_case_[0] = kName;
}
inline void APSsid::clear_name() {
  if (has_name()) {
    Ssid_.name_.DestroyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
    clear_has_Ssid();
  }
}
inline const ::std::string& APSsid::name() const {
  // @@protoc_insertion_point(field_get:cheetah.APSsid.Name)
  if (has_name()) {
    return Ssid_.name_.GetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  return *&::google::protobuf::internal::GetEmptyStringAlreadyInited();
}
inline void APSsid::set_name(const ::std::string& value) {
  // @@protoc_insertion_point(field_set:cheetah.APSsid.Name)
  if (!has_name()) {
    clear_Ssid();
    set_has_name();
    Ssid_.name_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  Ssid_.name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:cheetah.APSsid.Name)
}
inline void APSsid::set_name(const char* value) {
  if (!has_name()) {
    clear_Ssid();
    set_has_name();
    Ssid_.name_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  Ssid_.name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(value));
  // @@protoc_insertion_point(field_set_char:cheetah.APSsid.Name)
}
inline void APSsid::set_name(const char* value, size_t size) {
  if (!has_name()) {
    clear_Ssid();
    set_has_name();
    Ssid_.name_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  Ssid_.name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(
      reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:cheetah.APSsid.Name)
}
inline ::std::string* APSsid::mutable_name() {
  if (!has_name()) {
    clear_Ssid();
    set_has_name();
    Ssid_.name_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  // @@protoc_insertion_point(field_mutable:cheetah.APSsid.Name)
  return Ssid_.name_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* APSsid::release_name() {
  // @@protoc_insertion_point(field_release:cheetah.APSsid.Name)
  if (has_name()) {
    clear_has_Ssid();
    return Ssid_.name_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  } else {
    return NULL;
  }
}
inline void APSsid::set_allocated_name(::std::string* name) {
  if (!has_name()) {
    Ssid_.name_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  clear_Ssid();
  if (name != NULL) {
    set_has_name();
    Ssid_.name_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
        name);
  }
  // @@protoc_insertion_point(field_set_allocated:cheetah.APSsid.Name)
}

// optional uint32 Handle = 2;
inline bool APSsid::has_handle() const {
  return Ssid_case() == kHandle;
}
inline void APSsid::set_has_handle() {
  _oneof_case_[0] = kHandle;
}
inline void APSsid::clear_handle() {
  if (has_handle()) {
    Ssid_.handle_ = 0u;
    clear_has_Ssid();
  }
}
inline ::google::protobuf::uint32 APSsid::handle() const {
  // @@protoc_insertion_point(field_get:cheetah.APSsid.Handle)
  if (has_handle()) {
    return Ssid_.handle_;
  }
  return 0u;
}
inline void APSsid::set_handle(::google::protobuf::uint32 value) {
  if (!has_handle()) {
    clear_Ssid();
    set_has_handle();
  }
  Ssid_.handle_ = value;
  // @@protoc_insertion_point(field_set:cheetah.APSsid.Handle)
}

inline bool APSsid::has_Ssid() const {
  return Ssid_case() != SSID_NOT_SET;
}
inline void APSsid::clear_has_Ssid() {
  _oneof_case_[0] = SSID_NOT_SET;
}
inline APSsid::SsidCase APSsid::Ssid_case() const {
  return APSsid::SsidCase(_oneof_case_[0]);
}
inline const APSsid* APSsid::internal_default_instance() {
  return &APSsid_default_instance_.get();
}
#endif  // !PROTOBUF_INLINE_NOT_IN_HEADERS
// -------------------------------------------------------------------

// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace cheetah

#ifndef SWIG
namespace google {
namespace protobuf {

template <> struct is_proto_enum< ::cheetah::APErrorStatus_APErrno> : ::google::protobuf::internal::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::cheetah::APErrorStatus_APErrno>() {
  return ::cheetah::APErrorStatus_APErrno_descriptor();
}
template <> struct is_proto_enum< ::cheetah::APRegOp> : ::google::protobuf::internal::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::cheetah::APRegOp>() {
  return ::cheetah::APRegOp_descriptor();
}
template <> struct is_proto_enum< ::cheetah::APObjectOp> : ::google::protobuf::internal::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::cheetah::APObjectOp>() {
  return ::cheetah::APObjectOp_descriptor();
}
template <> struct is_proto_enum< ::cheetah::APNotifOp> : ::google::protobuf::internal::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::cheetah::APNotifOp>() {
  return ::cheetah::APNotifOp_descriptor();
}

}  // namespace protobuf
}  // namespace google
#endif  // SWIG

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_ap_5fcommon_5ftypes_2eproto__INCLUDED