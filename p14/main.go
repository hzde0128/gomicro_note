package main

// 作为客户端，调用p13的rpc服务

import (
	"context"
	"gomicro_note/p13/models"

	etcd "github.com/asim/go-micro/plugins/registry/etcd/v3"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	service := httpServer.NewServer(
		server.Name("testService.client"),
		server.Address(":9000"),
		server.Registry(etcdReg),
	)

	myService := micro.NewService(micro.Name("ProdService.client"))
	prodService := models.NewProdService("ProdService", myService.Client())
	v1Group := r.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", func(c *gin.Context) {
			var prodReq models.ProdRequest
			err := c.Bind(&prodReq)
			if err != nil {
				c.JSON(500, gin.H{
					"status": err.Error()})
				return
			}
			prodRes, err := prodService.GetProdList(context.Background(), &prodReq)
			if err != nil {
				c.JSON(500, gin.H{
					"status": err.Error()})
				return
			}
			c.JSON(200, gin.H{
				"data": prodRes.Data,
			})
		})
	}

	hd := service.NewHandler(r)
	service.Handle(hd)

	srv := micro.NewService(
		micro.Server(service),
	)

	srv.Init()
	srv.Run()
}
