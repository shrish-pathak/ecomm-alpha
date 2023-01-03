package handler

import (
	"ecomm-alpha/config"
	"ecomm-alpha/database"
	"ecomm-alpha/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

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
		if err != nil {
			log.Println(err)
		}
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	if ok, errorFields := validateSellerSignUpInput(sellerSUD); ok != true {
		return c.Status(400).JSON(ResponseHTTP{Success: false, Message: "validation error", Data: errorFields})
	}

	var email string

	db.Exec("SELECT email FROM sellers where email = ?", sellerSUD.Email).Scan(&email)

	if email != "" {
		return c.Status(400).JSON(ResponseHTTP{Success: false, Message: "user already exists", Data: nil})
	}
	hash, err := hashPassword(sellerSUD.Password)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}
	sellerSUD.Password = hash

	seller := &sellerSUD.Seller
	if err := db.Create(seller).Error; err != nil {
		log.Println(err)
		return c.Status(500).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = seller.Email
	claims["seller_id"] = seller.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(201).JSON(ResponseHTTP{Success: true, Message: "", Data: t})

}
