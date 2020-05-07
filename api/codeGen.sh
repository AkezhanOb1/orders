#!/bin/bash
protoc order.proto  --go_out=plugins=grpc:.