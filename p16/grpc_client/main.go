package main

// 装饰器的使用

import (
	"context"
	"fmt"
	"gomicro_note/p16/grpc_client/models"
	"gomicro_note/p16/grpc_client/routers"

	etcd "github.com/asim/go-micro/plugins/registry/etcd/v3"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/metadata"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
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

	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	myService := micro.NewService(
		micro.Name("ProdService.client"),
		micro.WrapClient(newLogWrapper),
	)
	prodService := models.NewProdService("ProdService", myService.Client())

	service := httpServer.NewServer(
		server.Name("ProdService.client"),
		server.Address(":9000"),
		server.Registry(etcdReg),
	)

	hd := service.NewHandler(routers.InitRouter(prodService))
	service.Handle(hd)

	srv := micro.NewService(
		micro.Server(service),
	)

	srv.Init()
	srv.Run()
}
