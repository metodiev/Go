package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/dialect/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "dbUserandPassword@simplerest?charset=utf8parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}