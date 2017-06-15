### GO quick Tutorial

The reader is assumed familiar with the equivalent Python quick start tutorial. This tutorial shows how the same example can be writen in Go.

Before you run it, you need to have installed Go on your system:
   https://golang.org/doc/install

For example, for Go 1.8 on amd64:
  wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz


Set up $GOPATH appropriately:
   https://golang.org/doc/code.html
   https://github.com/golang/go/wiki/SettingGOPATH

# Use the following command to install gRPC
go get google.golang.org/grpc

# Next, install the protoc plugin for Go (optional if you have the Go bindings already)
go get -a github.com/golang/protobuf/protoc-gen-go

# Install the cheetah-objmodel
go get github.com/cheetah-api/cheetah-objmodel

# Test/Build the tutorial
export SERVER_IP=192.168.122.192
export SERVER_PORT=57777

# This would require setting $GOARCH to the appropriate architecture (amd64, arm, etc)
# For applications deployed on the Cisco AP3700/AP3800 Access Points, sent GOARCH to arm
cd cheetah-objmodel/grpc/go/src/tutorial
go run quickstart.go
