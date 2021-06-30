package routes

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	productRouter := app.Group("/api/product")
	categoryRouter := app.Group("/api/category")
	// reviewRouter := app.Group("/api/review")
	imageRouter := app.Group("/api/image")
	authenRouter := app.Group("/api/authen")
	// orderRouter := app.Group("/api/order")
	// ReviewApi(&reviewRouter)     //Review
	ProductApi(&productRouter)   //Product
	CategoryApi(&categoryRouter) //Category
	ImagesApi(&imageRouter)      //Images
	AuthenApi(&authenRouter)
	// OrderApi(&orderRouter)
}
