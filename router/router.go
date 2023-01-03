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

	store := v1.Group("/store")

	store.Post("/", middleware.Protected(), handler.CreateStore)
	store.Put("/", handler.UpdateStore)
	store.Patch("/", handler.UpdateStore)
}
