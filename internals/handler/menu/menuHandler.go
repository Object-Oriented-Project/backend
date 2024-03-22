package menuHandler

import (
	"log"

	"github.com/Object-Oriented-Project/backend/database"
	"github.com/Object-Oriented-Project/backend/internals/model"
	"github.com/gofiber/fiber/v2"
)

// GetAllMenus godoc
// @Summary Get all menus
// @Description Get all menus
// @Tags Menu
// @Accept json
// @Produce json
// @Success 200 {array} model.Menu
// @Router /api/menu [get]
func GetAllMenus(c *fiber.Ctx) error {
	db := database.DB
	var menus []model.Menu
	db.Find(&menus)
	return c.JSON(menus)
}

// CreateMenus godoc
// @Summary Create a menu
// @Description Create a menu
// @Tags Menu
// @Accept json
// @Produce json
// @Param menu body model.Menu true "Menu"
// @Success 200 {object} model.Menu
// @Router /api/menu [post]
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