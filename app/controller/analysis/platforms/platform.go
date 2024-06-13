package cplatforms

import (
	splatforms "hrms-api/app/service/analysis/platforms"
	splatformslist "hrms-api/app/service/platforms"
	"github.com/gofiber/fiber/v2"
)

func GetPlatformData(ctx *fiber.Ctx) error {
	posting_platform_array, _ := splatforms.GetPlatform()
	platforms_list, _ := splatformslist.FetchPlatformsList()
	

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"platform_analysis": posting_platform_array,
		"platform_list": platforms_list,
	})
}
