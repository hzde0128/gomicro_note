#!/bin/bash

# go get -u github.com/favadi/protoc-go-inject-tag
protoc --micro_out=./ --go_out=./ models/protos/prods.proto && \
protoc-go-inject-tag --input=./models/prods.pb.go
