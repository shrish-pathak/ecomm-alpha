package handler

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "hello it works.", "data": nil})
}