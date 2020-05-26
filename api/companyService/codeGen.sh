#!/bin/bash
protoc companyServices.proto  --go_out=plugins=grpc:.  --proto_path=$HOME/Desktop/SELF/diplomaProject/api/proto/business/companyServices/