#!/bin/bash

export MICRO_REGISTRY=etcd
export MICRO_REGISTRY_ADDRESS=127.0.0.1:2379

export MICRO_API_HANDLER=rpc
export MICRO_API_NAMESPACE=api.hzde.com

micro api
