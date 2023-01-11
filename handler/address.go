package handler

import (
	"ecomm-alpha/database"
	"ecomm-alpha/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CreateAddress(c *fiber.Ctx) error {
	db := database.DB
	address := new(models.Address)
	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	var statusCode int
	if err := c.BodyParser(address); err != nil {
		statusCode = GetStatusCodeFromError(err)

		log.Println(err)

		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	if ok, errorFields := validateAddressInput(address); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{Success: false, Message: "validation error", Data: errorFields})
	}

	if user["user_type"] == "seller" {
		address.UserID = uint(user["seller_id"].(float64))
	}
	if user["user_type"] == "buyer" {
		address.UserID = uint(user["buyer_id"].(float64))
	}

	if err := db.Create(address).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	return c.SendStatus(fiber.StatusCreated)
}
