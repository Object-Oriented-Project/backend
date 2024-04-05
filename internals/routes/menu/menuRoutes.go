package menuRoutes

import (
	menuHandler "github.com/Object-Oriented-Project/backend/internals/handler/menu"
	"github.com/gofiber/fiber/v2"
)

func SetupMenuRoutes(route fiber.Router) {
	route.Get("/", menuHandler.GetAllMenus)
	route.Get("/:id", menuHandler.GetMenuByID)

	route.Put("/:id", menuHandler.UpdateMenuByID)

	route.Post("/", menuHandler.CreateMenus)

	route.Delete("/:id", menuHandler.DeleteMenuByID)
}
