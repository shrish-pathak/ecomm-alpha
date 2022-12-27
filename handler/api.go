package handler

import "github.com/gofiber/fiber/v2"

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "hello it works.", "data": nil})
}
