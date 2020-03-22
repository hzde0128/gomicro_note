package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
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
		web.Address(":8000"),
		web.Handler(r),
	)
	service.Run()
}
