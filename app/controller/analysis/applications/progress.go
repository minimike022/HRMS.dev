package cprogress

import (
	sapplication "hrms-api/app/service/analysis/applications"
	"github.com/gofiber/fiber/v2"
)

func FetchApplicationsProgress(ctx *fiber.Ctx) error {
	progress_status_array, err := sapplication.FetchProgress()
	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map {
		"progress_status": progress_status_array,
	})
}

func GetApplicantsDate(ctx *fiber.Ctx) error {
	applicants_data_array, err := sapplication.FetchDate()

	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"All_Applicants": applicants_data_array,
	})
}



