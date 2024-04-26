package route

import (
	service "hrms-api/app/service"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", service.ReadApplicantsData)
	app.Get("/application/status", service.GetApplicationStatus)
	app.Post("/application/add", service.PostApplicantsData)
	app.Post("/user_account/add", service.AddUserAccount)
	app.Get("/user_account", service.ReadUserAccounts)
	app.Patch("/user_account/change_status/:id", service.ChangeAccountStatus)
	app.Get("/application/manager/:id", service.ManagerApplicantsData)
	app.Get("/application/source", service.GetSourceData)
}





