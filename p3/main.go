package main

import (
	"log"
	"net/http"

	etcd "github.com/asim/go-micro/plugins/registry/etcd/v3"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
)

// etcd 服务发现
func main() {

	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"))

	r := gin.Default()
	r.Handle("GET", "/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})
	})
	service := httpServer.NewServer(
		server.Address(":8000"),
		server.Registry(etcdReg),
	)
	hd := service.NewHandler(r)
	service.Handle(hd)
	srv := micro.NewService(
		micro.Server(service),
	)
	if err := srv.Run(); err != nil {
		log.Println(err.Error())
	}
}
