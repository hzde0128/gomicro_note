package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
)

// consul 服务发现
func main() {

	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"))

	r := gin.Default()
	r.Handle("GET", "/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})
	})
	service := web.NewService(
		web.Address(":8000"),
		web.Handler(r),
		web.Registry(consulReg),
	)
	service.Run()
}
