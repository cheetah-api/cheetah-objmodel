#include <iostream>
#include <memory>
#include <string>

#include <grpc++/grpc++.h>

#include "../gencpp/ap_global.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;

int
main(int argc, char** argv)
{

  /* Create the channel for gRPC */
  auto channel = grpc::CreateChannel(("localhost:57777"),
                        grpc::InsecureChannelCredentials());

  return (0);
}
