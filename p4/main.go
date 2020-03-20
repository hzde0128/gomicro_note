package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"net/http"
)

// 调用函数返回json数据
func main(){

	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"))

	r := gin.Default()
	// 路由分组
	v1Group := r.Group("/v1")
	{
		v1Group.Handle("GET", "/prods", func(c *gin.Context) {
			c.JSON(http.StatusOK, NewProdList(5))
		})
	}

	service := web.NewService(
		web.Name("ProdSrv"),
		web.Address("127.0.0.1:8000"),
		web.Handler(r),
		web.Registry(consulReg),
	)
	service.Run()
}
