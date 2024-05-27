package happlication

import (
	capplicants "hrms-api/app/controller/applicants"
	capplicants_status "hrms-api/app/controller/applicants/status"
	"github.com/gofiber/fiber/v2"
	//jwtvalidate "hrms-api/app/service/jwt/validate"
	//validaterole "hrms-api/app/service/users/validate"
)

func SetupApplication(app *fiber.App) {
	app.Post("/applicants/add", capplicants.AddApplicantsData)
	app.Patch("/application/status_update/:id", capplicants_status.UpdateApplicationStatus)
	app.Get("/applicants/:app_id", capplicants.GetApplicantsData)
	app.Get("/application/status", capplicants_status.GetApplicationStatus)
}
//, jwtvalidate.ValidateRefreshToken, validaterole.ValidateAdmin,
//, jwtvalidate.ValidateRefreshToken, validaterole.ValidateHR