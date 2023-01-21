package handler

import (
	"ecomm-alpha/database"
	"ecomm-alpha/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// CreateAddress creates a address
//
//	@Summary		Register a new address data
//	@Description	Register Address
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			address body models.Address	true "Register address"
//	@Success		201
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/address/ [post]
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

// UpdateAddress updates a address
//
//	@Summary		Updates address data
//	@Description	Update Address
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			address body models.Address	true "Update address"
//	@Success		200 {object}	ResponseHTTP{data=string}
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/address/ [put]
func UpdateAddress(c *fiber.Ctx) error {
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

	var addressId string
	err := db.Raw("update address set mobile_no=?,city=?,state=?,zip=?,country=?,address=? returning id;",
		address.MobileNo, address.City, address.State, address.Zip, address.Country, address.Address).Scan(&addressId).Error

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{Success: true, Message: "", Data: addressId})
}
