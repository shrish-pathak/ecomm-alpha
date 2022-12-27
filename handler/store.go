package handler

import (
	"ecomm-alpha/database"
	"ecomm-alpha/models"

	"github.com/gofiber/fiber/v2"
)

func CreateStore(c *fiber.Ctx) error {
	db := database.DB

	store := new(models.Store)

	if err := c.BodyParser(store); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	if err := db.Create(store).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": store})

}
