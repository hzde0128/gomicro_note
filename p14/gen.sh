#!/bin/bash

protoc --micro_out=./ --go_out=./ models/protos/prodService.proto && \
protoc-go-inject-tag --input=models/prodService.pb.go
