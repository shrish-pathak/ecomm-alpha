package handler

import (
	"ecomm-alpha/database"
	"ecomm-alpha/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

type BuyerSignUpDetails struct {
	models.Buyer
	ConfirmPassword string `json:"confirmPassword" example:"har!@#ryp#$otter123!@#"`
}

// CreateBuyerAccount creates a Buyer account
//
//	@Summary		Register a new Buyer data
//	@Description	Register Buyer
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			buyer body BuyerSignUpDetails	true "Register Buyer"
//	@Success		200	{object}	ResponseHTTP{data=string}
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/Buyer/signup [post]
func CreateBuyerAccount(c *fiber.Ctx) error {
	db := database.DB

	BuyerSUD := new(BuyerSignUpDetails)
	var statusCode int
	if err := c.BodyParser(BuyerSUD); err != nil {
		statusCode = GetStatusCodeFromError(err)
		log.Println(err)
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	if ok, errorFields := validateBuyerSignUpInput(BuyerSUD); ok != true {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{Success: false, Message: "validation error", Data: errorFields})
	}

	var email string

	err := db.Raw("SELECT email FROM Buyers where email = ?", BuyerSUD.Email).Scan(&email).Error

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}
	if email != "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{Success: false, Message: "user already exists", Data: nil})
	}
	hash, err := hashPassword(BuyerSUD.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}
	BuyerSUD.Password = hash

	Buyer := &BuyerSUD.Buyer
	if err := db.Create(Buyer).Error; err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	t, err := createBuyerToken(Buyer)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusCreated).JSON(ResponseHTTP{Success: true, Message: "", Data: t})

}
