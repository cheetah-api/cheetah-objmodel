### GO quick Tutorial

The reader is assumed familiar with the equivalent Python quick start tutorial. This tutorial shows how the same example can be writen in Go.
Please refer to <http://golang.org> for more information about GO.

Before you run it, you need to have installed Go on your system, typically the following steps are sufficient:

# Assuming you've installed the cheetah-objmode here
$HOME/cheetah-objmodel

# Get Golang

cd $HOME
wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
tar xzvf go1.8.linux-amd64.tar.gz

# Set ENV variables
export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
export GOARCH=arm
export GOPATH=$HOME/grpc/go

# Use the following command to install gRPC
go get google.golang.org/grpc

# Next, install the protoc plugin for Go (optional if you have the Go bindings already)
go get -a github.com/golang/protobuf/protoc-gen-go

# Test/Build the tutorial

export SERVER_IP=192.168.122.192
export SERVER_PORT=57777

# This would require setting $GOARCH to the appropriate architecture (amd64, etc)
cd $GOPATH/src/tutorial
go run quickstart.go
