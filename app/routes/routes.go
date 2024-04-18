package route

import (
	routesData "hrms-api/app/service"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", routesData.GetData)
	app.Post("/", routesData.PostData)
}





