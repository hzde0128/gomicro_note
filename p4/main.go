package main

import (
	"net/http"

	etcd "github.com/asim/go-micro/plugins/registry/etcd/v3"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
)

// 调用函数返回json数据
func main() {

	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"))

	r := gin.Default()
	// 路由分组
	v1Group := r.Group("/v1")
	{
		v1Group.Handle("GET", "/prods", func(c *gin.Context) {
			c.JSON(http.StatusOK, NewProdList(5))
		})
	}

	service := httpServer.NewServer(
		server.Name("ProdSrv"),
		server.Address(":8000"),
		server.Registry(etcdReg),
	)

	hd := service.NewHandler(r)
	service.Handle(hd)
	srv := micro.NewService(
		micro.Server(service),
	)

	srv.Run()
}
