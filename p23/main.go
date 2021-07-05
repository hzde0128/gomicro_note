package main

import (
	"gomicro_note/p23/models"
	"gomicro_note/p23/test"

	etcd "github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	myService := micro.NewService(
		micro.Name("test.hzde.com"),
		micro.Address(":8000"),
		micro.Registry(etcdReg),
	)

	models.RegisterTestServiceHandler(myService.Server(), new(test.TestService))

	myService.Init()

	myService.Run()
}
