package cplatforms

import (
	splatforms "hrms-api/app/service/analysis/platforms"
	"github.com/gofiber/fiber/v2"
)

func GetPlatformData(ctx *fiber.Ctx) error {
	posting_platform_array, err := splatforms.GetPlatform()
	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"platform_analysis": posting_platform_array,
	})
}
