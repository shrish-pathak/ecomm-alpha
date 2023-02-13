package handler

import (
	"ecomm-alpha/database"
	"ecomm-alpha/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type OrderStatus struct {
	models.Order
	CancelledOrderID *uuid.UUID `json:"cancelledOrderID"`
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

	buyerId := uuid.MustParse(c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["buyer_id"].(string))

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

	buyerId := uuid.MustParse(c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["buyer_id"].(string))

	var statusCode int
	if err := c.BodyParser(cancelledOrder); err != nil {
		statusCode = GetStatusCodeFromError(err)
		log.Println(err)
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	order := new(models.Order)
	err := db.Raw("select * from orders where id=? and buyer_id=?", cancelledOrder.OrderId, buyerId).Scan(order).Error
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}
	if order != nil {
		if err := db.Create(cancelledOrder).Error; err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
		}
		return c.SendStatus(fiber.StatusCreated)
	}

	return c.SendStatus(fiber.StatusBadRequest)
}

// GetOrderStatus shows order status
//
//	@Summary		Shows order status data
//	@Description	Shows order status
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			order body models.Order true "Get order status"
//	@Success		200 {object}	ResponseHTTP{data=OrderStatus}
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

	buyerId := uuid.MustParse(c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["buyer_id"].(string))

	t1 := new(OrderStatus)

	err := db.Table("orders").Select("orders.*, cancelled_orders.id as cancelled_order_id").
		Joins("left join cancelled_orders on orders.id = cancelled_orders.order_id").Where("orders.id=? and order.buyer_id=?", order.ID, buyerId).
		Scan(t1).Error

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{Success: true, Message: "", Data: t1})
}

// func UpdateOrderStatus()  {
//Note: System will update order status when integrated with delivery system.
// }
