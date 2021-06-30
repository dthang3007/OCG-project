package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ProductId int      `json:"product_id"`
	Product   *Product `json:"product"`
	Quantity  int      `json:"quantity"`
}
