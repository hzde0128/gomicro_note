package main

import (
	"log"

	etcd "github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
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
