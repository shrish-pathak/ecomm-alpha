package middleware

import (
	"ecomm-alpha/config"
	"ecomm-alpha/handler"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected protect routes
func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.Config("SECRET")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(handler.ResponseHTTP{Success: false, Message: "Missing or malformed JWT", Data: nil})

	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(handler.ResponseHTTP{Success: false, Message: "Invalid or expired JWT", Data: nil})
	}
}
