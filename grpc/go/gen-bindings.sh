#!/bin/bash
#
# Copyright (c) 2017 by cisco Systems, Inc.
# All rights reserved.
#
cd ../protos
printf "Generating Go bindings..."
protoc -I ./ *.proto --go_out=plugins=grpc:../go/src/gengo/
echo "Done"
