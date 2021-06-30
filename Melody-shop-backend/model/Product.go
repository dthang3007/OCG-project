package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryId int `json:"category_id"`
	Category   *Category
	Name       string   `json:"name"`
	Price      float64  `json:"price"`
	Describe   string   `json:"describe"`
	TradeMark  string   `json:"trade_mark"`
	Origin     string   `json:"origin"`
	Version    string   `json:"version"`
	SKU        int64    `json:"sku"`
	Images     []*Image `json:"images" gorm:"- foreignKey:ProductId"`
}

func (product *Product) Count(db *gorm.DB, category_id interface{}, search string) int64 {
	var total int64

	if category_id == 0 {
		db.Where("name Like ?", "%"+search+"%").Model(&Product{}).Count(&total)
	} else {
		db.Where("name Like ?", "%"+search+"%").Where("category_id=?", category_id).Model(&Product{}).Count(&total)
	}
	return total
}
func (product *Product) Take(db *gorm.DB, limit int, offset int, order string, sort string, category_id interface{}, search string) interface{} {
	var products []Product
	if category_id == 0 {
		db.Where("name Like ?", "%"+search+"%").Order(order + " " + sort).Offset(offset).Limit(limit).Preload("Images").Find(&products)
	} else {
		db.Where("name Like ?", "%"+search+"%").Where("category_id=?", category_id).Order(order + " " + sort).Offset(offset).Limit(limit).Preload("Images").Find(&products)

	}

	return products
}
