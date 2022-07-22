package config

import (
	"github.com/jinzhu/gorm"
	"log"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:@(127.0.0.1:3306)/go_book?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db = d

}

func GetDB() *gorm.DB {
	return db
}
