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

	database.Connect()

	app.Get("/swagger/*", swagger.HandlerDefault)
	router.SetupRoutes(app)

	app.Listen(":" + config.Config("APP_PORT"))
}