// Generated by the gRPC protobuf plugin.
// If you make any local change, they will be lost.
// source: ap_global.proto

#include "ap_global.pb.h"
#include "ap_global.grpc.pb.h"

#include <grpc++/impl/codegen/async_stream.h>
#include <grpc++/impl/codegen/async_unary_call.h>
#include <grpc++/impl/codegen/channel_interface.h>
#include <grpc++/impl/codegen/client_unary_call.h>
#include <grpc++/impl/codegen/method_handler_impl.h>
#include <grpc++/impl/codegen/rpc_service_method.h>
#include <grpc++/impl/codegen/service_type.h>
#include <grpc++/impl/codegen/sync_stream.h>
namespace cheetah {

static const char* APGlobal_method_names[] = {
  "/cheetah.APGlobal/APGlobalInitNotif",
  "/cheetah.APGlobal/APGlobalsGet",
};

std::unique_ptr< APGlobal::Stub> APGlobal::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  std::unique_ptr< APGlobal::Stub> stub(new APGlobal::Stub(channel));
  return stub;
}

APGlobal::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_APGlobalInitNotif_(APGlobal_method_names[0], ::grpc::RpcMethod::SERVER_STREAMING, channel)
  , rpcmethod_APGlobalsGet_(APGlobal_method_names[1], ::grpc::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::ClientReader< ::cheetah::APGlobalNotif>* APGlobal::Stub::APGlobalInitNotifRaw(::grpc::ClientContext* context, const ::cheetah::APInitMsg& request) {
  return new ::grpc::ClientReader< ::cheetah::APGlobalNotif>(channel_.get(), rpcmethod_APGlobalInitNotif_, context, request);
}

::grpc::ClientAsyncReader< ::cheetah::APGlobalNotif>* APGlobal::Stub::AsyncAPGlobalInitNotifRaw(::grpc::ClientContext* context, const ::cheetah::APInitMsg& request, ::grpc::CompletionQueue* cq, void* tag) {
  return new ::grpc::ClientAsyncReader< ::cheetah::APGlobalNotif>(channel_.get(), cq, rpcmethod_APGlobalInitNotif_, context, request, tag);
}

::grpc::Status APGlobal::Stub::APGlobalsGet(::grpc::ClientContext* context, const ::cheetah::APGlobalsGetMsg& request, ::cheetah::APGlobalsGetMsgRsp* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_APGlobalsGet_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::cheetah::APGlobalsGetMsgRsp>* APGlobal::Stub::AsyncAPGlobalsGetRaw(::grpc::ClientContext* context, const ::cheetah::APGlobalsGetMsg& request, ::grpc::CompletionQueue* cq) {
  return new ::grpc::ClientAsyncResponseReader< ::cheetah::APGlobalsGetMsgRsp>(channel_.get(), cq, rpcmethod_APGlobalsGet_, context, request);
}

APGlobal::Service::Service() {
  (void)APGlobal_method_names;
  AddMethod(new ::grpc::RpcServiceMethod(
      APGlobal_method_names[0],
      ::grpc::RpcMethod::SERVER_STREAMING,
      new ::grpc::ServerStreamingHandler< APGlobal::Service, ::cheetah::APInitMsg, ::cheetah::APGlobalNotif>(
          std::mem_fn(&APGlobal::Service::APGlobalInitNotif), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      APGlobal_method_names[1],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< APGlobal::Service, ::cheetah::APGlobalsGetMsg, ::cheetah::APGlobalsGetMsgRsp>(
          std::mem_fn(&APGlobal::Service::APGlobalsGet), this)));
}

APGlobal::Service::~Service() {
}

::grpc::Status APGlobal::Service::APGlobalInitNotif(::grpc::ServerContext* context, const ::cheetah::APInitMsg* request, ::grpc::ServerWriter< ::cheetah::APGlobalNotif>* writer) {
  (void) context;
  (void) request;
  (void) writer;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status APGlobal::Service::APGlobalsGet(::grpc::ServerContext* context, const ::cheetah::APGlobalsGetMsg* request, ::cheetah::APGlobalsGetMsgRsp* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace cheetah

