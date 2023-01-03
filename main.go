package main

import (
	"ecomm-alpha/database"
	"log"

	"ecomm-alpha/router"

	_ "ecomm-alpha/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title Ecomm-Alpha
// @version 1.0
// @description This is the documentation for Ecomm-Alpha
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5000
// @BasePath /swagger/
func main() {
	app := fiber.New()
	app.Use(recover.New())
	database.ConnectDB()
	app.Get("/swagger/*", swagger.HandlerDefault)
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":5000"))
}
