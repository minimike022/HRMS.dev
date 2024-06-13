package hplatforms

import(
	cplatforms "hrms-api/app/controller/analysis/platforms"
	"github.com/gofiber/fiber/v2"
)

func SetupPlatformAnalysis(app *fiber.App) {
	app.Get("/analysis/platform", cplatforms.GetPlatformData)
}