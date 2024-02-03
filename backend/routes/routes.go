package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/urlshort/controllers"
)

func InitRouter() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/:id?", controllers.GetUser)

	app.Post("/", controllers.CreateUser)

	app.Listen("localhost:3001")
}
