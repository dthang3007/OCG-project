package model

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, entity Entity, page int, order string, sort string, where interface{}, search string) fiber.Map {
	limit := 12
	offset := (page - 1) * limit

	data := entity.Take(db, limit, offset, order, sort, where, search)
	total := entity.Count(db, where, search)

	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": int(math.Ceil(float64(total) / float64(limit))),
		},
	}
}
