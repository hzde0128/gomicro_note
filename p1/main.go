package main

import (
	"fmt"
	"net/http"

	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
)

// go-micro 体验

func main() {
	service := httpServer.NewServer(server.Address(":8000"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	hd := service.NewHandler(mux)
	service.Handle(hd)
	srv := micro.NewService(
		micro.Server(service),
	)
	srv.Init()
	err := srv.Run()
	if err != nil {
		fmt.Println(err)
	}
}
