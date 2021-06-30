package controller

import (
	"fmt"

	"github.com/app/database"
	"github.com/app/model"
	"github.com/gofiber/fiber/v2"
)

// func GetImageById(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")
// 	if err != nil {
// 		return c.Status(400).SendString(err.Error()) //400: loi may khach va gui loi~
// 	}
// 	image, err := repository.Image.FindImageById(int64(id)) //Khong loi 400 thi se de repo xu ly tiep
// 	if err != nil {
// 		return c.Status(404).SendString(err.Error()) //Khong tim thay
// 	}
// 	return c.JSON(image)
// }
func GetAllImages(c *fiber.Ctx) error {
	var image *[]model.Image

	database.DB.Find(&image)
	return c.JSON(image)

}

func CreateNewImage(c *fiber.Ctx) error {
	image := new(model.Image)
	error := c.BodyParser(image)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   error,
		})
	}
	result := database.DB.Create(&image)
	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}
	return c.SendString(fmt.Sprintf("New image is created successfully with id = %d", image.ID))
}

// func UpdateImage(c *fiber.Ctx) error {
// 	updatedimage := new(model.Image)

// 	err := c.BodyParser(&updatedimage)
// 	// if error
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Cannot parse JSON",
// 			"error":   err,
// 		})
// 	}

// 	err = repository.Image.UpdateImage(updatedimage)
// 	if err != nil {
// 		return c.Status(404).SendString(err.Error())
// 	}

// 	return c.SendString(fmt.Sprintf("image with id = %d is successfully updated", updatedimage.Id))
// }

// func UpsertImage(c *fiber.Ctx) error {
// 	image := new(model.Image)

// 	err := c.BodyParser(&image)
// 	// if error
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Cannot parse JSON",
// 			"error":   err,
// 		})
// 	}

// 	id := repository.Image.UpdateImage(image)
// 	return c.SendString(fmt.Sprintf("image with id = %d is successfully upserted", id))
// }
// func DeleteImageById(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")
// 	if err != nil {
// 		return c.Status(400).SendString(err.Error())
// 	}
// 	err = repository.Image.DeleteImageById(int64(id))

// 	if err != nil {
// 		return c.Status(404).SendString(err.Error())
// 	} else {

// 		return c.SendString("delete review successfully")
// 	}
// }
