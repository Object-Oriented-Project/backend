package main

import (
	"github.com/Object-Oriented-Project/backend/config"
	"github.com/Object-Oriented-Project/backend/database"
	"github.com/Object-Oriented-Project/backend/router"
	_"github.com/Object-Oriented-Project/backend/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)
// @title Cafka API
// @version 1.0
// @description This is a simple API for Cafka cafe.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3001
// @BasePath /

func main() {
	app := fiber.New()
	//CORS
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		return c.Next()
	})

	// Add a route handler for OPTIONS requests
app.Options("/*", func(c *fiber.Ctx) error {
    c.Set("Access-Control-Allow-Origin", "*")
    c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
    c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
    return c.SendStatus(fiber.StatusOK)
})



	database.Connect()

	app.Get("/swagger/*", swagger.HandlerDefault)
	router.SetupRoutes(app)

	app.Listen(":" + config.Config("APP_PORT"))
}