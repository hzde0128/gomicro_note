package main

import (
	"fmt"
	"net/http"

	"github.com/micro/go-micro/v2/web"
)

// go-micro 体验

func main() {
	service := web.NewService(web.Address(":8000"))

	service.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	err := service.Run()
	if err != nil {
		fmt.Println(err)
	}
}
