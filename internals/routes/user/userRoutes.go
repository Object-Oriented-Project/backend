package userRoutes

import (
	userHandler "github.com/Object-Oriented-Project/backend/internals/handler/user"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(route fiber.Router) {
	route.Get("/", userHandler.GetAllUsers)
	// route.Get("/:id", user.GetUser)
	route.Post("/", userHandler.CreateUser)
	// route.Put("/:id", user.UpdateUser)
	// route.Delete("/:id", user.DeleteUser)
	// route.Post("/register", user.Register)
	// route.Post("/login", user.Login)
}
