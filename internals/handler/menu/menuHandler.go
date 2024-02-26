package menuHandler

import (
	"log"

	"github.com/Object-Oriented-Project/backend/database"
	"github.com/Object-Oriented-Project/backend/internals/model"
	"github.com/gofiber/fiber/v2"
)

func GetAllMenus(c *fiber.Ctx) error {
	db := database.DB
	var menus []model.Menu
	db.Find(&menus)
	return c.JSON(menus)
}

func CreateMenus(c *fiber.Ctx) error {
	db := database.DB
	menu := new(model.Menu)
	if err := c.BodyParser(menu); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	db.Create(&menu)
	return c.JSON(menu)
}