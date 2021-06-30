package controller

import (
	"fmt"

	"github.com/app/database"
	"github.com/app/model"
	"github.com/gofiber/fiber/v2"
)

// import (
// 	"fmt"

// 	"github.com/app/model"
// 	"github.com/app/repository"
// 	"github.com/gofiber/fiber/v2"
// )

// func GetAllCategories(c *fiber.Ctx) error {
// 	return c.JSON(repository.Category.GetAllCategories())
// }
func CreateNewCategory(c *fiber.Ctx) error {
	category := new(model.Category)
	error := c.BodyParser(category)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   error,
		})
	}
	result := database.DB.Create(&category)
	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}
	return c.SendString(fmt.Sprintf("New Category is created successfully with id = %d", category.ID))
}
func GetCategoryById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error()) //400: loi may khach va gui loi~
	}
	var category *model.Category

	result := database.DB.Preload("Product").First(&category, id)
	if result.Error != nil {
		fmt.Println(result.Error)
		return c.SendString(result.Error.Error()) //Khong tim thay
	}
	return c.JSON(category)
}

// }
// func DeleteCategoryById(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")
// 	if err != nil {
// 		return c.Status(400).SendString(err.Error()) //400: loi may khach
// 	}
// 	err = repository.Category.DeleteCategoryById(int64(id))
// 	if err != nil {
// 		return c.SendString(err.Error())
// 	}
// 	return c.SendString("Delete successfuly")
// }
// func UpdateCategory(c *fiber.Ctx) error {
// 	updateCategory := new(model.Category)
// 	err := c.BodyParser(&updateCategory)
// 	if err != nil {
// 		return c.Status(400).SendString(err.Error()) //400: loi may khach
// 	}
// 	err = repository.Category.UpdateCategory(updateCategory)
// 	if err != nil {
// 		return c.Status(404).SendString(err.Error())
// 	}
// 	return c.SendString(fmt.Sprintf("Update category with id = %d is successfully updated", updateCategory.CategoryId))
// }
// func UpsertCategory(c *fiber.Ctx) error {
// 	upsertCategory := new(model.Category)
// 	err := c.BodyParser(&upsertCategory)
// 	if err != nil {
// 		return c.Status(400).SendString(err.Error())
// 	}
// 	id := repository.Category.UpsertCategory(upsertCategory)
// 	return c.SendString(fmt.Sprintf("Upsert success with id %d ", id))
// }
