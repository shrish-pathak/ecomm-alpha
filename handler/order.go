package handler

import (
	"ecomm-alpha/database"
	"ecomm-alpha/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type OrderStatus struct {
	models.Order
	models.CancelledOrder
}

// PlaceOrder creates a new order
//
//	@Summary		Register a new order data
//	@Description	Register order
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			order body models.Order true "Register order"
//	@Success		201
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/order/ [post]
func PlaceOrder(c *fiber.Ctx) error {
	/*
		-get all cart items,
		-prepare all order field values from cart items,
		-create order
		-create orderItem entries,
		-delete all cart items for user
	*/
	buyerId := uint(c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["buyer_id"].(float64))

	cartItems, err := getBuyerCartItems(buyerId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	log.Println(cartItems)

	order := new(models.Order)
	var statusCode int
	if err := c.BodyParser(order); err != nil {
		statusCode = GetStatusCodeFromError(err)
		log.Println(err)
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	for _, itm := range *cartItems {
		order.TotalAmount += itm.Product.Price * float64(itm.Quantity)
	}

	order.BuyerID = buyerId
	// order.Tax = TODO:

	db := database.DB

	if err := db.Create(order).Error; err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	orderItems := make([]models.OrderItem, len(*cartItems))

	for i, cItms := range *cartItems {
		orderItems[i].OrderId = order.ID
		orderItems[i].ProductID = cItms.ProductID
		orderItems[i].Quantity = cItms.Quantity
		orderItems[i].Price = cItms.Product.Price
		orderItems[i].Discount = cItms.Product.Discount
	}

	if err := db.Create(orderItems).Error; err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	err = emptyBuyerCart(buyerId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}
	return c.SendStatus(fiber.StatusCreated)
}

// CancelOrder creates a cancel order entry
//
//	@Summary		Register cancel order data
//	@Description	Register cancel order
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			cancelOrder body models.CancelledOrder true "Register order"
//	@Success		201
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/order/cancel [post]
func CancelOrder(c *fiber.Ctx) error {
	db := database.DB

	cancelledOrder := new(models.CancelledOrder)

	var statusCode int
	if err := c.BodyParser(cancelledOrder); err != nil {
		statusCode = GetStatusCodeFromError(err)
		log.Println(err)
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	if err := db.Create(cancelledOrder).Error; err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// GetOrderStatus shows order status
//
//	@Summary		Shows order status data
//	@Description	Shows order status
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			order body models.Order true "Get order status"
//	@Success		201
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/order/status [post]
func GetOrderStatus(c *fiber.Ctx) error {
	db := database.DB

	order := new(models.Order)

	var statusCode int
	if err := c.BodyParser(order); err != nil {
		statusCode = GetStatusCodeFromError(err)
		log.Println(err)
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	t1 := new(OrderStatus)
	err := db.Raw("select * from orders as o left join cancel_order as co on o.id=co.order_id where id=?", order.ID).Scan(t1).Error

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{Success: true, Message: "", Data: t1})
}

// UpdateOrder updates order
//
//	@Summary		Update order data
//	@Description	Update order
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			order body models.Order true "Update order"
//	@Success		200
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/order/status [patch]
func UpdateOrderStatus(c *fiber.Ctx) error {
	db := database.DB

	order := new(models.Order)

	var statusCode int
	if err := c.BodyParser(order); err != nil {
		statusCode = GetStatusCodeFromError(err)
		log.Println(err)
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	buyerId := uint(c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["buyer_id"].(float64))

	var orderId string
	err := db.Raw("update orders set status=? where order_id=? and buyer_id=? returning id", order.Status, order.ID, buyerId).Scan(&orderId).Error

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	return c.SendStatus(fiber.StatusOK)
}
