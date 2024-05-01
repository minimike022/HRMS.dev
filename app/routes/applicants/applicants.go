package happlication

import (
	applicants "hrms-api/app/service/applicants"
	"github.com/gofiber/fiber/v2"
)

func SetupApplication(app *fiber.App) {
	app.Post("/applicants/add", applicants.PostApplicantsData)
	app.Get("/applicants", applicants.GetApplicantsData)
}
