package handler

import (
	"ecomm-alpha/database"
	"ecomm-alpha/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// CreateProduct creates a Product for seller
//
//	@Summary		Register a new Product data
//	@Description	Register Product
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			product body models.Product true "Register Product"
//	@Success		201
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/product/ [post]
func CreateProduct(c *fiber.Ctx) error {
	db := database.DB

	//TODO: add check for seller id exist in given store id and then create product for that
	product := new(models.Product)
	var statusCode int
	if err := c.BodyParser(product); err != nil {
		statusCode = GetStatusCodeFromError(err)
		log.Println(err)
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	if ok, errorFields := validateProductInput(product); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{Success: false, Message: "validation error", Data: errorFields})
	}

	if err := db.Create(product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// UpdateProduct updates the Product of seller
//
//	@Summary		Updates the Product data
//	@Description	Update Product
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			product body models.Product true "Update Product"
//	@Success		200
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/product/ [put]

func UpdateProduct(c *fiber.Ctx) error {
	db := database.DB

	product := new(models.Product)
	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)

	var statusCode int
	if err := c.BodyParser(product); err != nil {
		statusCode = GetStatusCodeFromError(err)
		log.Println(err)
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	if ok, errorFields := validateProductInput(product); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{Success: false, Message: "validation error", Data: errorFields})
	}

	var userID uint
	if user["user_type"] == "seller" {
		userID = uint(user["seller_id"].(float64))
	}
	if user["user_type"] == "buyer" {
		userID = uint(user["buyer_id"].(float64))
	}

	/*
		if store exists by userID and storeID matches with product.StoreID then update else not
	*/
	err := db.Raw("update product set title=?, description=?,price=?,discount=? where exists (select 1 from store where store.id=product.store_id and seller_id=?)", product.Title, product.Description, product.Price, product.Discount, userID)

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	return c.SendStatus(fiber.StatusOK)
}
