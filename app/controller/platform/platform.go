package cplatformslist

import (
	splatformslist "hrms-api/app/service/platforms"

	"github.com/gofiber/fiber/v2"
)

func FetchPlatformListData(ctx *fiber.Ctx) error {
	platform_list_array, err := splatformslist.FetchPlatformsList()

	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map {
		"platform_list": platform_list_array,
	})
}