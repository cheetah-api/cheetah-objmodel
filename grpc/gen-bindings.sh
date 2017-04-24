#!/bin/bash
#
# Copyright (c) 2017 by Cisco Systems, Inc.
# All rights reserved.
#
echo "Generating bindings..."
cd python && ./gen-bindings.sh 
cd ../cpp && ./gen-bindings.sh
cd ../go && ./gen-bindings.sh
