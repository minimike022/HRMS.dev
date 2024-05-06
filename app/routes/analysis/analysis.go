package hanalysis

import (
	platforms "hrms-api/app/service/analysis/platforms"
	jwtvalidate "hrms-api/app/service/jwt/validate"

	"github.com/gofiber/fiber/v2"
)

func SetupAnalysis(app *fiber.App) {
	SetupPlatformAnalysis(app)
}

//Platform Analysis
func SetupPlatformAnalysis(app *fiber.App) {
	app.Get("/analysis/platform", jwtvalidate.ValidateAccessToken, platforms.GetPlatformData)
}
