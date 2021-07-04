package main

import (
	"log"
	"net/http"

	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
)

// 使用gin框架

func main() {
	r := gin.Default()
	r.Handle("GET", "/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})
	})
	service := httpServer.NewServer(
		server.Name("demo_service"),
		server.Address(":8000"),
	)
	hd := service.NewHandler(r)
	service.Handle(hd)
	srv := micro.NewService(
		micro.Server(service),
	)
	srv.Init()
	if err := srv.Run(); err != nil {
		log.Print(err.Error())
	}
}
