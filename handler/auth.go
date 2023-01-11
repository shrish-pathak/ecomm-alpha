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

type SellerCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SellerLogin(c *fiber.Ctx) error {

	sellerCredentials := new(SellerCredentials)
	var statusCode int
	if err := c.BodyParser(sellerCredentials); err != nil {
		statusCode = GetStatusCodeFromError(err)
		if err != nil {
			log.Println(err)
		}
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	if ok, errorFields := validateSellerLoginInput(sellerCredentials); ok != true {
		return c.Status(400).JSON(ResponseHTTP{Success: false, Message: "validation error", Data: errorFields})
	}

	seller, err := getSellerByEmail(sellerCredentials.Email)

	if !CheckPasswordHash(sellerCredentials.Password, seller.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseHTTP{Success: false, Message: "Invalid password", Data: nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = seller.Email
	claims["seller_id"] = seller.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}
	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{Success: true, Message: "", Data: t})
}

func getSellerByEmail(e string) (*models.Seller, error) {
	db := database.DB
	seller := new(models.Seller)
	err := db.Raw("select * from sellers where email=?", e).Scan(&seller).Error

	if seller != nil {
		return seller, nil
	}
	return nil, err
}
