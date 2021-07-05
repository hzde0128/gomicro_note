#!/bin/bash

protoc --micro_out=./ --go_out=./ models/protos/UserService.proto
protoc-go-inject-tag --input=models/UserService.pb.go
