#!/bin/bash

cd models/protos && \
protoc --micro_out=../ --go_out=../ models.proto && \
protoc --micro_out=../ --go_out=../ prodService.proto && \
protoc-go-inject-tag --input=../models.pb.go && \
cd -
