package main

import (
	"ecomm-alpha/database"
	"log"

	"ecomm-alpha/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":5000"))
}
