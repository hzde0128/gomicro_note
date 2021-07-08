package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"

	"gomicro_note/p27/controllers"
	"gomicro_note/p27/models"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	app := micro.NewService(
		micro.Name("api.hzde.com.user"),
		micro.Address(":8000"),
		micro.Registry(consulReg),
	)

	models.RegisterUserServiceHandler(app.Server(), new(controllers.UserService))
	app.Init()

	app.Run()
}
