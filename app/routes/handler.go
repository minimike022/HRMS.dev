package route

import (
	service "hrms-api/app/service"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", service.ReadApplicantsData)
	app.Get("/application/status", service.GetApplicationStatus)
	app.Post("/applicant", service.PostApplicantsData)
	app.Post("/user_account/add", service.AddUserAccount)
	app.Get("/user_account", service.ReadUserAccounts)
	app.Patch("/user_account/changeStatus/:account_id", service.ChangeAccountStatus)
}





