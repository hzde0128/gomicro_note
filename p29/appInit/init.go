package appInit

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gomicro_note/p29/models"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:000000aa@tcp(127.0.0.1:3306)/user?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	// 自动建表
	db.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return db
}
