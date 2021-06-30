package main

import (
	"github.com/app/database"
	"github.com/app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})
	app.Use(cors.New())

	app.Static("/public", "./public", fiber.Static{
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})
	routes.Setup(app)

	app.Listen(":3000")
}
