package hrhandler

import (
	hrpkg "hrms-api/app/service/hr_pkg"
	"github.com/gofiber/fiber/v2"
)

func SetupHRhandler(app *fiber.App) {
	app.Get("/applicants", hrpkg.GetApplicantsData)
	app.Get("/analysis", hrpkg.GetSourceData)
}