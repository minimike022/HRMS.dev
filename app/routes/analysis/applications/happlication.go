package happlication

import (
	capplication "hrms-api/app/controller/analysis/applications"
	"github.com/gofiber/fiber/v2"
)

func SetupApplicationProgress(app *fiber.App) {
	app.Get("/analysis/progress", capplication.FetchApplicationsProgress)
	app.Get("/analysis/date", capplication.GetApplicantsDate)
} 