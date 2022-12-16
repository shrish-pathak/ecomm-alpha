package main

import (
	"log"

	"ecomm-alpha/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":5000"))
}
