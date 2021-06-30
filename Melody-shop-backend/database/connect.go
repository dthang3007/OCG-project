package database

import (
	"github.com/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("thang:Thangg@30@/database?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("Could not connect db")
	}
	DB = database
	database.AutoMigrate(&model.User{}, &model.Product{}, &model.Image{}, &model.Category{}, &model.Orders{}, &model.OrderItem{})
}
