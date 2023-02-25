package handler

import (
	"ecomm-alpha/database"
	"ecomm-alpha/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

type SellerSignUpDetails struct {
	models.Seller
	ConfirmPassword string `json:"confirmPassword" example:"har!@#ryp#$otter123!@#"`
}

// CreateSellerAccount creates a seller account
//
//	@Summary		Register a new seller data
//	@Description	Register seller
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			seller body SellerSignUpDetails	true "Register seller"
//	@Success		200	{object}	ResponseHTTP{data=string}
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/seller/signup [post]
func CreateSellerAccount(c *fiber.Ctx) error {
	db := database.DB

	sellerSUD := new(SellerSignUpDetails)
	var statusCode int
	if err := c.BodyParser(sellerSUD); err != nil {
		statusCode = GetStatusCodeFromError(err)
		log.Println(err)
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	if ok, errorFields := validateSellerSignUpInput(sellerSUD); ok != true {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{Success: false, Message: "validation error", Data: errorFields})
	}

	var email string

	err := db.Raw("SELECT email FROM sellers where email = ?", sellerSUD.Email).Scan(&email).Error

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}
	if email != "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{Success: false, Message: "user already exists", Data: nil})
	}
	hash, err := hashPassword(sellerSUD.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}
	sellerSUD.Password = hash

	seller := &sellerSUD.Seller
	if err := db.Create(seller).Error; err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	t, err := createSellerToken(seller)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusCreated).JSON(ResponseHTTP{Success: true, Message: "", Data: t})

}
