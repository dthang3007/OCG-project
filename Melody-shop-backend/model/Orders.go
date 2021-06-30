package model

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	UserId     int          `json:"user_id"`
	User       *User        `json:"user" `
	OrderItems []*OrderItem `json:"order_item" gorm:"foreignKey:OrderId"`
	
}
type OrderItem struct {
	OrderId   int
	ProductId int
	Product   *Product
	Quantity  int64
}
