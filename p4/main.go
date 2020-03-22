package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
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

	service := web.NewService(
		web.Name("ProdSrv"),
		web.Address(":8000"),
		web.Handler(r),
		web.Registry(etcdReg),
	)
	service.Run()
}
