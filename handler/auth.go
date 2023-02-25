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
type BuyerCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// SellerLogin does the login
//
//	@Summary		Login seller
//	@Description	Login seller
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			store body SellerCredentials true "Login Seller"
//	@Success		200	{object}	ResponseHTTP{data=string}
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/seller/login [post]
func SellerLogin(c *fiber.Ctx) error {

	sellerCredentials := new(SellerCredentials)
	var statusCode int
	if err := c.BodyParser(sellerCredentials); err != nil {
		statusCode = GetStatusCodeFromError(err)
		log.Println(err)
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	if ok, errorFields := validateSellerLoginInput(sellerCredentials); ok != true {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{Success: false, Message: "validation error", Data: errorFields})
	}

	seller, err := getSellerByEmail(sellerCredentials.Email)

	if !CheckPasswordHash(sellerCredentials.Password, seller.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseHTTP{Success: false, Message: "Invalid password", Data: nil})
	}

	t, err := createSellerToken(seller)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}
	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{Success: true, Message: "", Data: t})
}

// BuyerLogin does the login
//
//	@Summary		Login Buyer
//	@Description	Login Buyer
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			store body BuyerCredentials true "Login Buyer"
//	@Success		200	{object}	ResponseHTTP{data=string}
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/buyer/login [post]
func BuyerLogin(c *fiber.Ctx) error {

	BuyerCredentials := new(BuyerCredentials)
	var statusCode int
	if err := c.BodyParser(BuyerCredentials); err != nil {
		statusCode = GetStatusCodeFromError(err)
		log.Println(err)
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	if ok, errorFields := validateBuyerLoginInput(BuyerCredentials); ok != true {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{Success: false, Message: "validation error", Data: errorFields})
	}

	Buyer, err := getBuyerByEmail(BuyerCredentials.Email)

	if !CheckPasswordHash(BuyerCredentials.Password, Buyer.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseHTTP{Success: false, Message: "Invalid password", Data: nil})
	}

	t, err := createBuyerToken(Buyer)
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
func getBuyerByEmail(e string) (*models.Buyer, error) {
	db := database.DB
	buyer := new(models.Buyer)
	err := db.Raw("select * from buyers where email=?", e).Scan(&buyer).Error

	if buyer != nil {
		return buyer, nil
	}
	return nil, err
}

func createSellerToken(seller *models.Seller) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = seller.Email
	claims["seller_id"] = seller.ID
	claims["user_type"] = "seller"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(config.Config("SECRET")))
	return t, err
}
func createBuyerToken(buyer *models.Buyer) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = buyer.Email
	claims["buyer_id"] = buyer.ID
	claims["user_type"] = "buyer"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(config.Config("SECRET")))
	return t, err
}
