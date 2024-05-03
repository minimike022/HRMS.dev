package happlication

import (
	applicants "hrms-api/app/service/applicants"
	applicantsStatus "hrms-api/app/service/applicants/status"
	"github.com/gofiber/fiber/v2"
	jwtvalidate "hrms-api/app/service/jwt/validate"
)

func SetupApplication(app *fiber.App) {
	app.Post("/applicants/add", applicants.PostApplicantsData)
	app.Get("/applicants", jwtvalidate.ValidateRefreshToken,applicants.GetApplicantsData)
	app.Get("/application/status", jwtvalidate.ValidateAccessToken, applicantsStatus.GetApplicationStatus)
}
