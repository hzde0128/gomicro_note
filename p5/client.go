package main

import (
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"
)

// consul 服务发现 selector随机选择
// 前提是要启动前面的ProdSrv服务

func main() {
	// consul连接句柄
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"))

	// 获取服务
	getService, err := consulReg.GetService("ProdSrv")
	if err != nil {
		log.Fatalf("get service failed, err:%v\n", err)
		return
	}

	next := selector.Random(getService)

	node, err := next()
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 打印服务节点信息
	log.Println(node.Id, node.Address, node.Metadata)

}
