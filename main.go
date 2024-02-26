package main

import (

	"github.com/Object-Oriented-Project/backend/config"
	"github.com/Object-Oriented-Project/backend/database"
	"github.com/Object-Oriented-Project/backend/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()

	router.SetupRoutes(app)

	app.Listen(":" + config.Config("APP_PORT"))
}