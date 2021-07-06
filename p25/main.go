package main

import (
	"gomicro_note/p25/models"
	"gomicro_note/p25/test"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	myService := micro.NewService(
		micro.Name("api.hzde.com.test"),
		micro.Address(":8000"),
		micro.Registry(etcdReg),
	)

	models.RegisterTestServiceHandler(myService.Server(), new(test.TestService))

	myService.Init()

	myService.Run()
}
