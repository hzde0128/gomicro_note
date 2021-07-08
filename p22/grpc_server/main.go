package main

import (
	"gomicro_note/p22/grpc_server/prods"
	"gomicro_note/p22/models"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {

	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	app := micro.NewService(
		micro.Name("ProdService"),
		micro.Address(":8000"),
		micro.Registry(consulReg),
	)

	app.Init()
	err := models.RegisterProdServiceHandler(app.Server(), new(prods.ProdService))
	if err != nil {
		panic(err)
	}
	app.Run()
}
