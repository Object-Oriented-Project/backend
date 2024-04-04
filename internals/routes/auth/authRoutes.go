package authRoutes

import (
	authHandler "github.com/Object-Oriented-Project/backend/internals/handler/auth"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(route fiber.Router) {
	route.Post("/", authHandler.LoginToken)
}
