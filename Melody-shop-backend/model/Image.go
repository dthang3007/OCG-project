package model

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ProductId uint   `json:"product_id"`
	Src       string `json:"src"`
}
