package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"net/http"
)

// consul 服务注册
func main(){

	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"))

	r := gin.Default()
	r.Handle("GET", "/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})
	})
	service := web.NewService(
		web.Address("127.0.0.1:8000"),
		web.Handler(r),
		web.Registry(consulReg),
	)
	service.Run()
}
