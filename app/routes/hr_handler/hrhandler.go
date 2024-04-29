package hrhandler

import (
	hrpkg "hrms-api/app/service/hr_pkg"
	"github.com/gofiber/fiber/v2"
)

func SetupHRhandler(app *fiber.App) {
	app.Get("/applicants", hrpkg.GetApplicantsData)
	app.Get("/application/status", hrpkg.GetApplicationStatus)
	app.Post("/jobs/position/add", hrpkg.AddJobPosition)
	app.Patch("/jobs/position/update/:id", hrpkg.UpdateJobPosition)
}