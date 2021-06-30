package controller

import (
	"fmt"
	"strconv"

	"github.com/app/database"
	"github.com/app/model"
	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	order := c.Query("order", "ID")
	sort := c.Query("sort", "asc")
	category, _ := strconv.Atoi(c.Query("category", "0"))
	search := c.Query("q", " ")
	return c.JSON(model.Paginate(database.DB, &model.Product{}, page, order, sort, category, search))
}
func CreateNewProduct(c *fiber.Ctx) error {
	product := new(model.Product)
	error := c.BodyParser(&product)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   error,
		})
	}
	result := database.DB.Create(&product)
	if result.Error != nil {
		c.SendString(result.Error.Error())
	}

	return c.JSON(product)
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error()) //400: loi may khach va gui loi~
	}
	var product *model.Product

	result := database.DB.Preload("Images").First(&product, id)
	if result.Error != nil {
		fmt.Println(result.Error)
		return c.SendString(result.Error.Error()) //Khong tim thay
	}
	return c.JSON(product)
}

func DeleteProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error()) //400: loi may khach
	}
	result := database.DB.Delete(&model.Product{}, id)
	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}
	return c.SendString("Delete successfuly")
}

func UpdateProduct(c *fiber.Ctx) error {
	updateProduct := new(model.Product)
	err := c.BodyParser(&updateProduct)
	if err != nil {
		return c.Status(400).SendString(err.Error()) //400: loi may khach
	}
	result := database.DB.Model(&model.Product{}).Updates(&updateProduct)
	if result.Error != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.SendString(fmt.Sprintf("Update product with id = %d is successfully updated", updateProduct.ID))
}

// func GetProductByCategory(c *fiber.Ctx) error {
// 	category_id, err := strconv.Atoi(c.Query("category_id", "1"))
// 	if err != nil {
// 		return c.Status(400).SendString(err.Error()) //400: loi may khach va gui loi~
// 	}
// 	var product *[]model.Product
// 	database.DB.Where("product_id=?", category_id).Find(&product)
// 	page, _ := strconv.Atoi(c.Query("page", "1"))
// 	order := c.Query("order", "ID")
// 	sort := c.Query("sort", "asc")
// 	return c.JSON(model.Paginate(database.DB, product, page, order, sort))

// }

// func UpsertProduct(c *fiber.Ctx) error {
// 	upsertProduct := new(model.Product)
// 	err := c.BodyParser(&upsertProduct)
// 	if err != nil {
// 		return c.Status(400).SendString(err.Error())
// 	}
// 	id := repository.Product.UpsertProduct(upsertProduct)
// 	return c.SendString(fmt.Sprintf("Upsert success with id %d ", id))
// }
