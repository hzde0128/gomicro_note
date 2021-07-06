#!/bin/bash

cd models/protos
protoc --go_out=../ Models.proto
protoc --micro_out=../ --go_out=../ UserService.proto
protoc-go-inject-tag --input=../Models.pb.go
protoc-go-inject-tag --input=../UserService.pb.go
cd -