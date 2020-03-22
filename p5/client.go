package main

import (
	"log"

	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

// consul 服务发现 selector随机选择
// 前提是要启动前面的ProdSrv服务

func main() {
	// consul连接句柄
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"))

	// 获取服务
	getService, err := etcdReg.GetService("ProdSrv")
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
	log.Println(node.Address)

}
