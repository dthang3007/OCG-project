package model

import "gorm.io/gorm"

type Entity interface {
	Count(db *gorm.DB, where interface{}, search string) int64
	Take(db *gorm.DB, limit int, offset int, order string, sort string, where interface{}, search string) interface{}
}
