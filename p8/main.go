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

// 商品服务
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
		server.Registry(etcdReg),
	)

	hd := service.NewHandler(r)
	service.Handle(hd)

	srv := micro.NewService(
		micro.Server(service),
	)
	// 通过命令行参数启动
	// --server_address 指定地址端口，或者环境变量$MICRO_SERVER_ADDRESS]
	// 运行2个服务
	// go run main.go prodModels.go --server_address  127.0.0.1:8000
	// go run main.go prodModels.go --server_address  127.0.0.1:8001
	srv.Init()
	srv.Run()
}
