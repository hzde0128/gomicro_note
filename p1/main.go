package main

import (
	"fmt"
	"github.com/micro/go-micro/web"
	"net/http"
)

// go-micro 体验

func main(){
	service := web.NewService(web.Address("192.168.10.117:8000"))

	service.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	err := service.Run()
	if err != nil {
		fmt.Println(err)
	}
}
