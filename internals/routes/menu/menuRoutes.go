package menuRoutes

import (
	"github.com/gofiber/fiber/v2"
	menuHandler "github.com/Object-Oriented-Project/backend/internals/handler/menu"
)

func SetupMenuRoutes(route fiber.Router) {
	route.Get("/", menuHandler.GetAllMenus)
	route.Post("/", menuHandler.CreateMenus)
}