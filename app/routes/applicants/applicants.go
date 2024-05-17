package happlication

import (
	applicants "hrms-api/app/service/applicants"
	applicantsStatus "hrms-api/app/service/applicants/status"
	"github.com/gofiber/fiber/v2"
	jwtvalidate "hrms-api/app/service/jwt/validate"
	validaterole "hrms-api/app/service/users/validate"
)

func SetupApplication(app *fiber.App) {
	app.Post("/applicants/add", applicants.PostApplicantsData)
	app.Get("/applicants",applicants.GetApplicantsData)
	app.Get("/application/status", jwtvalidate.ValidateRefreshToken, validaterole.ValidateAdmin, applicantsStatus.GetApplicationStatus)
	app.Get("/application/manager", jwtvalidate.ValidateAccessToken, applicantsStatus.GetApplicationStatus)
}

//, jwtvalidate.ValidateRefreshToken, validaterole.ValidateHR 