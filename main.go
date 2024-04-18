package main

import (
	route "hrms-api/app/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	route.SetupRoutes(app)
	app.Listen(":3000")


}


