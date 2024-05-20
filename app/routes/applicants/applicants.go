package happlication

import (
	applicants "hrms-api/app/service/applicants"
	applicantsStatus "hrms-api/app/service/applicants/status"
	"github.com/gofiber/fiber/v2"
	//jwtvalidate "hrms-api/app/service/jwt/validate"
	//validaterole "hrms-api/app/service/users/validate"
)

func SetupApplication(app *fiber.App) {
	app.Post("/applicants/add", applicants.PostApplicantsData)
	app.Patch("/application/status_update/:id", applicantsStatus.UpdateApplicationStatus)
	app.Get("/applicants",applicants.GetApplicantsData)
	app.Get("/application/status", applicantsStatus.GetApplicationStatus)
	app.Get("/application/manager", applicantsStatus.GetApplicationStatus)
}
//, jwtvalidate.ValidateRefreshToken, validaterole.ValidateAdmin,
//, jwtvalidate.ValidateRefreshToken, validaterole.ValidateHR 