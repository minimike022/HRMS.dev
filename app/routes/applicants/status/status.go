package hstatus

import (
	capplicants_status "hrms-api/app/controller/applicants/status"
	"github.com/gofiber/fiber/v2"
)

func SetupApplicationStatus(app *fiber.App) {
	app.Get("/status/list", capplicants_status.FetchStatusList)
	app.Get("/application/status", capplicants_status.GetApplicationStatus)
	app.Patch("/application/status/:id", capplicants_status.UpdateApplicationStatus)
}