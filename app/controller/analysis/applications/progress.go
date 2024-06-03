package cprogress

import (
	sprogress "hrms-api/app/service/analysis/applications"
	"github.com/gofiber/fiber/v2"
)

func FetchApplicationsProgress(ctx *fiber.Ctx) error {
	progress_status_array, err := sprogress.FetchProgress()
	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map {
		"progress_status": progress_status_array,
	})
}


