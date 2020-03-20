#!/bin/bash

cd models/protos && protoc --micro_out=../ --go_out=../ *.proto && cd -
