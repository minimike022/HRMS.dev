package route

import (
	service "hrms-api/app/service"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", service.ReadApplicantsData)
	app.Post("/applicant", service.PostApplicantsData)
	app.Get("/applicant/:id", service.GetApplicationStatus)
}





