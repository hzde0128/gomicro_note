package appInit

import (
	"gomicro_note/p28/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/user?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	db.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return db
}
