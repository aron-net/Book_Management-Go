package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/dislects/mysql"
)

var (
	db * gorm.DB
)

func Connect()  {
	d, err := gorm.Open("mysql", "akhil:Axlesharm@120/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}