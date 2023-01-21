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

	store := v1.Group("/store")

	store.Post("/", middleware.Protected(), handler.CreateStore)
	store.Put("/", middleware.Protected(), handler.UpdateStore)
	store.Patch("/name", middleware.Protected(), handler.PatchStoreName)
	store.Patch("/description", middleware.Protected(), handler.PatchStoreDescription)

	address := v1.Group("/address")

	address.Post("/", middleware.Protected(), handler.CreateAddress)
	address.Put("/", middleware.Protected(), handler.UpdateAddress)
}
