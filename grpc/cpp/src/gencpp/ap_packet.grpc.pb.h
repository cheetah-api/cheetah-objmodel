// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: ap_packet.proto
// Original file comments:
// @file
// @brief Packet I/O proto file
//
// ----------------------------------------------------------------
//  Copyright (c) 2017 by Cisco Systems, Inc.
//  All rights reserved.
// -----------------------------------------------------------------
//
//
//
#ifndef GRPC_ap_5fpacket_2eproto__INCLUDED
#define GRPC_ap_5fpacket_2eproto__INCLUDED

#include "ap_packet.pb.h"

#include <grpc++/impl/codegen/async_stream.h>
#include <grpc++/impl/codegen/async_unary_call.h>
#include <grpc++/impl/codegen/method_handler_impl.h>
#include <grpc++/impl/codegen/proto_utils.h>
#include <grpc++/impl/codegen/rpc_method.h>
#include <grpc++/impl/codegen/service_type.h>
#include <grpc++/impl/codegen/status.h>
#include <grpc++/impl/codegen/stub_options.h>
#include <grpc++/impl/codegen/sync_stream.h>

namespace grpc {
class CompletionQueue;
class Channel;
class RpcService;
class ServerCompletionQueue;
class ServerContext;
}  // namespace grpc

namespace cheetah {

// @defgroup APPackets
// @ingroup System
// The following RPCs are used to register for packets that the
// client is interested in receiving
// @{
class APPackets final {
 public:
  static constexpr char const* service_full_name() {
    return "cheetah.APPackets";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    // APPacketsRegMsg.Oper = AP_REGOP_REGISTER
    //     Packet registration: Sends a list of Packet registration messages
    //     and expects a list of registration responses.
    //
    // APPacketsRegMsg.Oper = AP_REGOP_UNREGISTER
    //     Packet unregistration: Sends a list of Packet unregistration messages
    //     and expects a list of unregistration responses.
    //
    virtual ::grpc::Status APPacketsRegOp(::grpc::ClientContext* context, const ::cheetah::APPacketsRegMsg& request, ::cheetah::APPacketsRegMsgRsp* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::cheetah::APPacketsRegMsgRsp>> AsyncAPPacketsRegOp(::grpc::ClientContext* context, const ::cheetah::APPacketsRegMsg& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::cheetah::APPacketsRegMsgRsp>>(AsyncAPPacketsRegOpRaw(context, request, cq));
    }
    //
    // Packet notifications
    //
    //
    // This call is used to get a stream of packet notifications matching the
    // set of registrations performed with APPacketsRegOp().
    // The caller must maintain the GRPC channel as long as
    // there is interest in packet notifications. Only sessions that were
    // created through this API will be notified to caller.
    std::unique_ptr< ::grpc::ClientReaderInterface< ::cheetah::APPacketsMsgRsp>> APPacketsInitNotif(::grpc::ClientContext* context, const ::cheetah::APPacketsGetNotifMsg& request) {
      return std::unique_ptr< ::grpc::ClientReaderInterface< ::cheetah::APPacketsMsgRsp>>(APPacketsInitNotifRaw(context, request));
    }
    std::unique_ptr< ::grpc::ClientAsyncReaderInterface< ::cheetah::APPacketsMsgRsp>> AsyncAPPacketsInitNotif(::grpc::ClientContext* context, const ::cheetah::APPacketsGetNotifMsg& request, ::grpc::CompletionQueue* cq, void* tag) {
      return std::unique_ptr< ::grpc::ClientAsyncReaderInterface< ::cheetah::APPacketsMsgRsp>>(AsyncAPPacketsInitNotifRaw(context, request, cq, tag));
    }
  private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::cheetah::APPacketsRegMsgRsp>* AsyncAPPacketsRegOpRaw(::grpc::ClientContext* context, const ::cheetah::APPacketsRegMsg& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientReaderInterface< ::cheetah::APPacketsMsgRsp>* APPacketsInitNotifRaw(::grpc::ClientContext* context, const ::cheetah::APPacketsGetNotifMsg& request) = 0;
    virtual ::grpc::ClientAsyncReaderInterface< ::cheetah::APPacketsMsgRsp>* AsyncAPPacketsInitNotifRaw(::grpc::ClientContext* context, const ::cheetah::APPacketsGetNotifMsg& request, ::grpc::CompletionQueue* cq, void* tag) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel);
    ::grpc::Status APPacketsRegOp(::grpc::ClientContext* context, const ::cheetah::APPacketsRegMsg& request, ::cheetah::APPacketsRegMsgRsp* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::cheetah::APPacketsRegMsgRsp>> AsyncAPPacketsRegOp(::grpc::ClientContext* context, const ::cheetah::APPacketsRegMsg& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::cheetah::APPacketsRegMsgRsp>>(AsyncAPPacketsRegOpRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientReader< ::cheetah::APPacketsMsgRsp>> APPacketsInitNotif(::grpc::ClientContext* context, const ::cheetah::APPacketsGetNotifMsg& request) {
      return std::unique_ptr< ::grpc::ClientReader< ::cheetah::APPacketsMsgRsp>>(APPacketsInitNotifRaw(context, request));
    }
    std::unique_ptr< ::grpc::ClientAsyncReader< ::cheetah::APPacketsMsgRsp>> AsyncAPPacketsInitNotif(::grpc::ClientContext* context, const ::cheetah::APPacketsGetNotifMsg& request, ::grpc::CompletionQueue* cq, void* tag) {
      return std::unique_ptr< ::grpc::ClientAsyncReader< ::cheetah::APPacketsMsgRsp>>(AsyncAPPacketsInitNotifRaw(context, request, cq, tag));
    }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    ::grpc::ClientAsyncResponseReader< ::cheetah::APPacketsRegMsgRsp>* AsyncAPPacketsRegOpRaw(::grpc::ClientContext* context, const ::cheetah::APPacketsRegMsg& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientReader< ::cheetah::APPacketsMsgRsp>* APPacketsInitNotifRaw(::grpc::ClientContext* context, const ::cheetah::APPacketsGetNotifMsg& request) override;
    ::grpc::ClientAsyncReader< ::cheetah::APPacketsMsgRsp>* AsyncAPPacketsInitNotifRaw(::grpc::ClientContext* context, const ::cheetah::APPacketsGetNotifMsg& request, ::grpc::CompletionQueue* cq, void* tag) override;
    const ::grpc::RpcMethod rpcmethod_APPacketsRegOp_;
    const ::grpc::RpcMethod rpcmethod_APPacketsInitNotif_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    // APPacketsRegMsg.Oper = AP_REGOP_REGISTER
    //     Packet registration: Sends a list of Packet registration messages
    //     and expects a list of registration responses.
    //
    // APPacketsRegMsg.Oper = AP_REGOP_UNREGISTER
    //     Packet unregistration: Sends a list of Packet unregistration messages
    //     and expects a list of unregistration responses.
    //
    virtual ::grpc::Status APPacketsRegOp(::grpc::ServerContext* context, const ::cheetah::APPacketsRegMsg* request, ::cheetah::APPacketsRegMsgRsp* response);
    //
    // Packet notifications
    //
    //
    // This call is used to get a stream of packet notifications matching the
    // set of registrations performed with APPacketsRegOp().
    // The caller must maintain the GRPC channel as long as
    // there is interest in packet notifications. Only sessions that were
    // created through this API will be notified to caller.
    virtual ::grpc::Status APPacketsInitNotif(::grpc::ServerContext* context, const ::cheetah::APPacketsGetNotifMsg* request, ::grpc::ServerWriter< ::cheetah::APPacketsMsgRsp>* writer);
  };
  template <class BaseClass>
  class WithAsyncMethod_APPacketsRegOp : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_APPacketsRegOp() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_APPacketsRegOp() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status APPacketsRegOp(::grpc::ServerContext* context, const ::cheetah::APPacketsRegMsg* request, ::cheetah::APPacketsRegMsgRsp* response) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestAPPacketsRegOp(::grpc::ServerContext* context, ::cheetah::APPacketsRegMsg* request, ::grpc::ServerAsyncResponseWriter< ::cheetah::APPacketsRegMsgRsp>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_APPacketsInitNotif : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_APPacketsInitNotif() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_APPacketsInitNotif() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status APPacketsInitNotif(::grpc::ServerContext* context, const ::cheetah::APPacketsGetNotifMsg* request, ::grpc::ServerWriter< ::cheetah::APPacketsMsgRsp>* writer) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestAPPacketsInitNotif(::grpc::ServerContext* context, ::cheetah::APPacketsGetNotifMsg* request, ::grpc::ServerAsyncWriter< ::cheetah::APPacketsMsgRsp>* writer, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncServerStreaming(1, context, request, writer, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_APPacketsRegOp<WithAsyncMethod_APPacketsInitNotif<Service > > AsyncService;
  template <class BaseClass>
  class WithGenericMethod_APPacketsRegOp : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_APPacketsRegOp() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_APPacketsRegOp() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status APPacketsRegOp(::grpc::ServerContext* context, const ::cheetah::APPacketsRegMsg* request, ::cheetah::APPacketsRegMsgRsp* response) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_APPacketsInitNotif : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_APPacketsInitNotif() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_APPacketsInitNotif() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status APPacketsInitNotif(::grpc::ServerContext* context, const ::cheetah::APPacketsGetNotifMsg* request, ::grpc::ServerWriter< ::cheetah::APPacketsMsgRsp>* writer) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_APPacketsRegOp : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithStreamedUnaryMethod_APPacketsRegOp() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::StreamedUnaryHandler< ::cheetah::APPacketsRegMsg, ::cheetah::APPacketsRegMsgRsp>(std::bind(&WithStreamedUnaryMethod_APPacketsRegOp<BaseClass>::StreamedAPPacketsRegOp, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_APPacketsRegOp() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status APPacketsRegOp(::grpc::ServerContext* context, const ::cheetah::APPacketsRegMsg* request, ::cheetah::APPacketsRegMsgRsp* response) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedAPPacketsRegOp(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::cheetah::APPacketsRegMsg,::cheetah::APPacketsRegMsgRsp>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_APPacketsRegOp<Service > StreamedUnaryService;
  template <class BaseClass>
  class WithSplitStreamingMethod_APPacketsInitNotif : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithSplitStreamingMethod_APPacketsInitNotif() {
      ::grpc::Service::MarkMethodStreamed(1,
        new ::grpc::SplitServerStreamingHandler< ::cheetah::APPacketsGetNotifMsg, ::cheetah::APPacketsMsgRsp>(std::bind(&WithSplitStreamingMethod_APPacketsInitNotif<BaseClass>::StreamedAPPacketsInitNotif, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithSplitStreamingMethod_APPacketsInitNotif() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status APPacketsInitNotif(::grpc::ServerContext* context, const ::cheetah::APPacketsGetNotifMsg* request, ::grpc::ServerWriter< ::cheetah::APPacketsMsgRsp>* writer) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with split streamed
    virtual ::grpc::Status StreamedAPPacketsInitNotif(::grpc::ServerContext* context, ::grpc::ServerSplitStreamer< ::cheetah::APPacketsGetNotifMsg,::cheetah::APPacketsMsgRsp>* server_split_streamer) = 0;
  };
  typedef WithSplitStreamingMethod_APPacketsInitNotif<Service > SplitStreamedService;
  typedef WithStreamedUnaryMethod_APPacketsRegOp<WithSplitStreamingMethod_APPacketsInitNotif<Service > > StreamedService;
};
//
// Packet registration operations
//

}  // namespace cheetah


#endif  // GRPC_ap_5fpacket_2eproto__INCLUDED
