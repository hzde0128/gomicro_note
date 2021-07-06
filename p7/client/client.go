package main

import (
	"log"
	"time"

	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

// consul 通过轮询获取服务
// 使用前面的方法同时运行3个相同的服务
// 并测试其它节点关闭时不影响业务

func main() {

	// etcd 连接句柄
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"))
	for {
		// 获取服务
		getService, err := etcdReg.GetService("ProdSrv")
		if err != nil {
			log.Fatalf("get service failed, err:%v\n", err)
			return
		}

		next := selector.RoundRobin(getService)

		node, err := next()
		if err != nil {
			log.Fatalln(err)
			return
		}

		log.Println(node.Address)
		time.Sleep(time.Second)
	}

}
