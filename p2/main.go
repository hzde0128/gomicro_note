package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"net/http"
)

// 使用gin框架

func main() {
	r := gin.Default()
	r.Handle("GET", "/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})
	})
	service := web.NewService(
		web.Name("demo_service"),
		web.Address("127.0.0.1:8000"),
		web.Handler(r),
		)
	service.Run()
}
