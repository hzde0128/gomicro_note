package main

// 熔断的使用

import (
	"context"
	"fmt"
	"gomicro_note/p22/grpc_client/routers"
	"gomicro_note/p22/grpc_client/wrappers"
	"gomicro_note/p22/models"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

func newLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}

func main() {

	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	myService := micro.NewService(
		micro.Name("ProdService.client"),
		micro.WrapClient(newLogWrapper),
		micro.WrapClient(wrappers.NewProdsWrapper),
	)
	prodService := models.NewProdService("ProdService", myService.Client())

	service := web.NewService(
		web.Name("ProdService.client"),
		web.Address(":9000"),
		web.Handler(routers.InitRouter(prodService)),
		web.Registry(consulReg),
	)

	service.Init()
	service.Run()
}
