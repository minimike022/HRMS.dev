package hanalysis


import (
	platforms "hrms-api/app/service/analysis/platforms"
	"github.com/gofiber/fiber/v2"
)

func SetupAnalysis(app *fiber.App) {
	SetupPlatformAnalysis(app)
}

//Platform Analysis
func SetupPlatformAnalysis(app *fiber.App) {
	app.Get("/analysis/platform", platforms.GetPlatformData)
}
