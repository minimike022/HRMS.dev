package main

import (
	route "hrms-api/app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	route.SetupRoutes(app)
	app.Listen(":3000")
}


