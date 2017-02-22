#!/bin/bash
#
# Copyright (c) 2017 by Cisco Systems, Inc.
# All rights reserved.
#
printf "Generating C++ bindings..."
cd ../protos
protoc -I ./ --cpp_out=../cpp/src/gencpp/ --grpc_out=../cpp/src/gencpp/ --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` *.proto
echo "Done"
