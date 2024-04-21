package route

import (
	service "hrms-api/app/service"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", service.GetData)
	app.Post("/postData", service.PostData)
	//app.Get("/putData",routesData.PutData)
}





