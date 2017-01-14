#!/bin/bash
#
# Copyright (c) 2016 by Cisco Systems, Inc.
# All rights reserved.
#
printf "Generating Python bindings..."
cd ../protos
protoc -I ./ --python_out=../python/src/genpy/ --grpc_out=../python/src/genpy/ --plugin=protoc-gen-grpc=`which grpc_python_plugin` *.proto
echo "Done"
