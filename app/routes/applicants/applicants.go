package happlicants

import (
	capplicants "hrms-api/app/controller/applicants"
	"github.com/gofiber/fiber/v2"
	//jwtvalidate "hrms-api/app/service/jwt/validate"
	//validaterole "hrms-api/app/service/users/validate"
)

func SetupApplication(app *fiber.App) {
	app.Post("/applicants/add", capplicants.AddApplicantsData)
	app.Get("/applicants/:app_id", capplicants.GetApplicantsData)
	app.Get("/new_applicants", capplicants.FetchNewApplicantsData)
}
//, jwtvalidate.ValidateRefreshToken, validaterole.ValidateAdmin,
//, jwtvalidate.ValidateRefreshToken, validaterole.ValidateHR