#!/bin/bash

protoc --micro_out=./ --go_out=./ models/protos/test.proto