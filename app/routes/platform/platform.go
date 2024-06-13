package hplatformslist

import (
	cplatformslist "hrms-api/app/controller/platform"
	"github.com/gofiber/fiber/v2"
)

func SetupPlatformList(app *fiber.App) {
	app.Get("/platforms", cplatformslist.FetchPlatformListData)
}