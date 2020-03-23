package main

import (
	"gomicro_note/p15/grpc_client/models"
	"gomicro_note/p15/grpc_client/routers"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
)

func main() {

	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	myService := micro.NewService(micro.Name("ProdService.client"))
	prodService := models.NewProdService("ProdService", myService.Client())

	service := web.NewService(
		web.Name("ProdService.client"),
		web.Address(":9000"),
		web.Handler(routers.InitRouter(prodService)),
		web.Registry(etcdReg),
	)

	service.Init()
	service.Run()
}
