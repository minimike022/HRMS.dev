package happlication

import (
	applicants "hrms-api/app/service/applicants"
	applicantsStatus "hrms-api/app/service/applicants/status"
	"github.com/gofiber/fiber/v2"
	//jwt "hrms-api/app/service/jwt"
)

func SetupApplication(app *fiber.App) {
	app.Post("/applicants/add", applicants.PostApplicantsData)
	app.Get("/applicants", applicants.GetApplicantsData)
	app.Get("/application/status", applicantsStatus.GetApplicationStatus)
}
