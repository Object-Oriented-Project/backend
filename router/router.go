package router

import (
	menuRoutes "github.com/Object-Oriented-Project/backend/internals/routes/menu"
	userRoutes "github.com/Object-Oriented-Project/backend/internals/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("_hc", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api	:= app.Group("/api", logger.New())

	menuRoutes.SetupMenuRoutes(api.Group("/menu"))
	userRoutes.SetupUserRoutes(api.Group("/user"))
}