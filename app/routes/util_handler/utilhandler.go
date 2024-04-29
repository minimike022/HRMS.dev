package utilhandler

import (
	util "hrms-api/app/util"
	"github.com/gofiber/fiber/v2"
)

func SetupUtilHandler(app *fiber.App) {
	app.Patch("/application/change_status/:id", util.UpdateApplicationStatus)
	app.Post("/application/add", util.PostApplicantsData)
	app.Get("/user/login", util.Login)
	app.Get("/analysis/platform", util.GetPlatformData)
}