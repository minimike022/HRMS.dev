package happlication

import (
	applicants "hrms-api/app/service/applicants"
	"github.com/gofiber/fiber/v2"
	jwt "hrms-api/app/service/jwt"
)

func SetupApplication(app *fiber.App) {
	app.Post("/applicants/add", applicants.PostApplicantsData)
	app.Get("/applicants", jwt.JwtAuth(), applicants.GetApplicantsData)
}
