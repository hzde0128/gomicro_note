package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"gomicro_note/p23/models"
	"gomicro_note/p23/test"
)

func main(){
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
