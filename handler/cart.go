package handler

import (
	"ecomm-alpha/database"
	"ecomm-alpha/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// AddToCart creates a new cart item
//
//	@Summary		Register a new cart item data
//	@Description	Register cart item
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			cartItem body models.CartItem true "Register cart item"
//	@Success		201
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/cart/ [post]
func AddToCart(c *fiber.Ctx) error {
	db := database.DB
	cartItemEntry := new(models.CartItem)

	var statusCode int
	if err := c.BodyParser(cartItemEntry); err != nil {
		statusCode = GetStatusCodeFromError(err)

		log.Println(err)
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	if cartItemEntry.Quantity > 0 == false {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{Success: false, Message: "invalid quantity", Data: nil})
	}

	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["buyer_id"].(string)

	cartItemEntry.BuyerID = uuid.MustParse(userId)

	if getAvailableQuantity(cartItemEntry.ProductID) > cartItemEntry.Quantity {
		if err := db.Create(cartItemEntry).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
		}
		err := updateProductAvailableQuantity(cartItemEntry.ProductID, int(-cartItemEntry.Quantity))

		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
		}

		return c.SendStatus(fiber.StatusCreated)
	}
	return c.Status(fiber.StatusExpectationFailed).JSON(ResponseHTTP{Success: false, Message: "item out of stock", Data: nil})
}

// GetCartItems gets cart items
//
//	@Summary		Get cart items data
//	@Description	Get cart items
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Success		200 {object}	ResponseHTTP{data=models.CartItem}
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/cart/ [get]
func GetCartItems(c *fiber.Ctx) error {

	userId := uuid.MustParse(c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["buyer_id"].(string))

	cartItems, err := getBuyerCartItems(userId)

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	type CartItems struct {
	}
	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{Success: true, Message: "", Data: cartItems})
}

// UpdateCartItem updates cart item
//
//	@Summary		Update cart item data
//	@Description	Update cart item
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			cartItem body models.CartItem true "Update cart item"
//	@Success		200
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/cart/ [put]
func UpdateCartItem(c *fiber.Ctx) error {
	cartItemEntry := new(models.CartItem)

	var statusCode int
	if err := c.BodyParser(cartItemEntry); err != nil {
		statusCode = GetStatusCodeFromError(err)

		log.Println(err)
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["buyer_id"].(string)

	cartItemEntry.BuyerID = uuid.MustParse(userId)

	if cartItemEntry.Quantity > 0 == false {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{Success: false, Message: "invalid quantity", Data: nil})
	}

	oldCartItemEntry := new(models.CartItem)
	db := database.DB
	err := db.Raw("select * from cart_item where id=?", cartItemEntry.ID).Scan(oldCartItemEntry).Error
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}
	if cartItemEntry.Quantity > oldCartItemEntry.Quantity {
		//subtract from available
		quantityDifference := cartItemEntry.Quantity - oldCartItemEntry.Quantity

		err = updateProductAvailableQuantity(cartItemEntry.ProductID, int(-quantityDifference))
		if err != nil {
			log.Println(err)

			return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
		}

	} else {
		//add to available
		quantityDifference := oldCartItemEntry.Quantity - cartItemEntry.Quantity

		err = updateProductAvailableQuantity(cartItemEntry.ProductID, int(quantityDifference))
		if err != nil {
			log.Println(err)

			return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
		}

	}
	var cartItemId string
	err = db.Raw("update cart_items set quantity=? where id=? and buyer_id=? returning id", cartItemEntry.Quantity, cartItemEntry.ID, cartItemEntry.BuyerID).Scan(cartItemId).Error
	if err != nil {
		log.Println(err)

		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	return c.SendStatus(fiber.StatusOK)
}

// RemoveFromCart deletes cart item
//
//	@Summary		Delete cart item data
//	@Description	Delete cart item
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			cartItem body models.CartItem true "Delete cart item"
//	@Success		200
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/cart/ [delete]
func RemoveFromCart(c *fiber.Ctx) error {
	db := database.DB
	cartItemEntry := new(models.CartItem)

	var statusCode int
	if err := c.BodyParser(cartItemEntry); err != nil {
		statusCode = GetStatusCodeFromError(err)

		log.Println(err)
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	userId := uuid.MustParse(c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["buyer_id"].(string))

	cartItemEntry.BuyerID = userId

	//Todo: implement new logic based on quantity
	var cartItemId string
	err := db.Raw("delete from cart_items where id=? and buyer_id=?", cartItemEntry.ID, cartItemEntry.BuyerID).Scan(&cartItemId).Error

	log.Println(cartItemId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	err = updateProductAvailableQuantity(cartItemEntry.ProductID, int(cartItemEntry.Quantity))

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	return c.SendStatus(fiber.StatusOK)
}

func getAvailableQuantity(productId uuid.UUID) uint {

	db := database.DB

	var availableQuantity uint

	err := db.Raw("select available_quantity from products where id=?", productId).Scan(availableQuantity).Error
	if err != nil {
		log.Println(err)
	}

	return availableQuantity
}

func getBuyerCartItems(buyerId uuid.UUID) (*[]models.CartItem, error) {
	db := database.DB

	cartItems := new([]models.CartItem)

	err := db.Model(models.CartItem{}).Where("buyer_id=?", buyerId).Preload("Product").Find(&cartItems).Error

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return cartItems, nil
}

func emptyBuyerCart(buyerId uuid.UUID) error {
	db := database.DB

	cartItemIds := new(string)
	err := db.Raw("delete from cart_items where buyer_id=?", buyerId).Scan(cartItemIds).Error

	log.Println(cartItemIds)
	return err
}
func updateProductAvailableQuantity(productId uuid.UUID, quantity int) error {
	db := database.DB

	var pid string
	err := db.Raw("update products set available_quantity=available_quantity+? where id=?", quantity, productId).Scan(&pid).Error

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
