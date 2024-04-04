package menuHandler

import (
	"fmt"
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

// GetMenuByID godoc
// @Summary Get menu by id
// @Description Get menu by id
// @Tags Menu
// @Accept json
// @Produce json
// @Success 200 {object} model.Menu
// @Router /api/menu/:id [get]
func GetMenuByID(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")

	var menu model.Menu
	result := db.First(&menu, id)

	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Cannot find menu",
		})
	}

	return c.JSON(menu)
}

// DeleteMenuByID godoc
// @Summary Delete menu by id
// @Description Delete menu by id
// @Tags Menu
// @Accept json
// @Produce json
// @Success 200 {null}
// @Router /api/menu/:id [delete]
func DeleteMenuByID(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")

	var menu model.Menu
	result := db.Delete(&menu, id) // Soft Delete
	// result := db.Unscoped().Delete(&menu, id) // Hard Delete

	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot delete menu",
		})
	}

	fmt.Println("Successfully delete menu at id = ", id)
	return nil
}

// UpdateMenuByID godoc
// @Summary Update menu by id
// @Description Update menu by id
// @Tags Menu
// @Accept json
// @Produce json
// @Success 200 {object} model.Menu
// @Router /api/menu/:id [get]
func UpdateMenuByID(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")

	// Find exited menu from DB as Struct
	menu := new(model.Menu)
	db.Table("public.menus").Select("*").Where("id = ?", id).Scan(menu)

	if err := c.BodyParser(menu); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Update in DB
	db.Save(&menu)
	return c.JSON(fiber.Map{
		"Update": "Succesfully",
	})
}
