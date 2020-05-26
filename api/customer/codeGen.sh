#!/bin/bash
protoc customer.proto  --go_out=plugins=grpc:.  --proto_path=$HOME/Desktop/SELF/customer/api/proto/customer