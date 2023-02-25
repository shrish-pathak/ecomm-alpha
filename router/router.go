package router

import (
	"ecomm-alpha/handler"
	"ecomm-alpha/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", handler.Hello)
	v1 := api.Group("/v1")
	seller := v1.Group("/seller")
	seller.Post("/signup", handler.CreateSellerAccount)
	seller.Post("/login", handler.SellerLogin)

	buyer := v1.Group("/buyer")
	buyer.Post("/signup", handler.CreateBuyerAccount)
	buyer.Post("/login", handler.BuyerLogin)

	store := v1.Group("/store")

	store.Post("/", middleware.Protected(), handler.CreateStore)
	store.Put("/", middleware.Protected(), handler.UpdateStore)
	store.Patch("/name", middleware.Protected(), handler.PatchStoreName)
	store.Patch("/description", middleware.Protected(), handler.PatchStoreDescription)

	address := v1.Group("/address")

	address.Post("/", middleware.Protected(), handler.CreateAddress)
	address.Put("/", middleware.Protected(), handler.UpdateAddress)

	product := v1.Group("/product")
	product.Post("/", middleware.Protected(), handler.CreateProduct)
	product.Put("/", middleware.Protected(), handler.UpdateProduct)

	cart := v1.Group("/cart")

	cart.Get("/", middleware.Protected(), handler.GetCartItems)
	cart.Post("/", middleware.Protected(), handler.AddToCart)
	cart.Post("/", middleware.Protected(), handler.UpdateCartItem)
	cart.Delete("/", middleware.Protected(), handler.RemoveFromCart)

	order := v1.Group("/order")
	order.Post("/", middleware.Protected(), handler.PlaceOrder)
	order.Post("/cancel", middleware.Protected(), handler.CancelOrder)
	order.Post("/status", middleware.Protected(), handler.GetOrderStatus)

}
