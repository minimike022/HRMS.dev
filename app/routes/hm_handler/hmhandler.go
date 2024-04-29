package hmhandler

import (
	hmpkg "hrms-api/app/service/hm_pkg"
	"github.com/gofiber/fiber/v2"

)

func SetupHMhandler(app *fiber.App) {
	app.Get("/user_account/applicants_data/:id", hmpkg.ManagerApplicantsData)
}