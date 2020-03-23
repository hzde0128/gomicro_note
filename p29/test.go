package main

import (
	"gopkg.in/go-playground/validator.v9"
	"log"
)

type Users struct {
	Username string `validate:"required,min=6,max=20"`
	Password string `validate:"required,min=6,max=20"`
}

func main() {
	user := Users{Username: "jerry", Password: "12345"}
	valid := validator.New()
	err := valid.Struct(user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("验证成功")
}
