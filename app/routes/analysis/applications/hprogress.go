package hprogress

import (
	cprogress "hrms-api/app/controller/analysis/applications"
	"github.com/gofiber/fiber/v2"
)

func SetupApplicationProgress(app *fiber.App) {
	app.Get("/analysis/progress", cprogress.FetchApplicationsProgress)
	app.Get("/analysis/date", cprogress.GetApplicantsDate)
} 