package routes

import (
	"github.com/app/controller"
	"github.com/gofiber/fiber/v2"
)

func ProductApi(router *fiber.Router) {
	(*router).Get("/", controller.GetAllProduct)
	(*router).Post("/", controller.CreateNewProduct)
	(*router).Get("/:id", controller.GetProductById)
	(*router).Delete("/:id", controller.DeleteProductById)
	(*router).Patch("", controller.UpdateProduct)
}

func ImagesApi(router *fiber.Router) {
	(*router).Post("/", controller.CreateNewImage)
	(*router).Get("/", controller.GetAllImages)
}

func AuthenApi(router *fiber.Router) {
	(*router).Post("/register", controller.Register)
	(*router).Post("/login", controller.Login)

}

func CategoryApi(router *fiber.Router) {
	// (*router).Get("/", controller.GetAllCategories)
	(*router).Post("/", controller.CreateNewCategory)
	(*router).Get("/:id", controller.GetCategoryById)
	// (*router).Delete("/:id", controller.DeleteCategoryById)
	// (*router).Patch("", controller.UpdateCategory)
	// (*router).Put("", controller.UpsertProduct)
}

// func OrderApi(router *fiber.Router){
// 	(*router).Post("/",controller.CreateNewOrder)
// }

// func ReviewApi(router *fiber.Router) {
// 	(*router).Get("/", controller.GetAllReviews)

// 	(*router).Get("/:id", controller.GetReviewById) //Xem chi tiết một bản ghi

// 	(*router).Delete("/:id", controller.DeleteReviewById) //Xoá một bản ghi

// 	(*router).Post("", controller.CreateReview) //INSERT: Tạo một bản ghi

// 	(*router).Patch("", controller.UpdateReview) //UPDATE: Cập nhật một bản ghi

// 	(*router).Put("", controller.UpsertReview)
// }

// func ImagesApi(router *fiber.Router) {

// 	(*router).Get("/:id", controller.GetImageById) //Xem chi tiết một bản ghi

// 	(*router).Delete("/:id", controller.DeleteImageById) //Xoá một bản ghi

// 	(*router).Post("", controller.CreateNewImage) //INSERT: Tạo một bản ghi

// 	(*router).Patch("", controller.UpdateReview) //UPDATE: Cập nhật một bản ghi

// 	(*router).Put("", controller.UpsertReview)
// }
