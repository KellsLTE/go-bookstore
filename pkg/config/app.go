package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect()  {
	d, err := gorm.Open("mysql", "lte:k8FoNi6A4orET17o1unO1oN4Q3Bopu/go_bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDM() *gorm.DB {
	return db
}