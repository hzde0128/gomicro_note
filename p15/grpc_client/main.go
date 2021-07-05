package main

import (
	"gomicro_note/p15/grpc_client/models"
	"gomicro_note/p15/grpc_client/routers"

	etcd "github.com/asim/go-micro/plugins/registry/etcd/v3"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
)

func main() {

	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	service := httpServer.NewServer(
		server.Name("ProdService.client"),
		server.Address(":9000"),
		server.Registry(etcdReg),
	)

	myService := micro.NewService(micro.Name("ProdService.client"))
	prodService := models.NewProdService("ProdService", myService.Client())

	hd := service.NewHandler(routers.InitRouter(prodService))
	service.Handle(hd)

	srv := micro.NewService(
		micro.Server(service),
	)

	srv.Init()
	srv.Run()
}
