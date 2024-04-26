package adminhandler

import (
	//service "hrms-api/app/service"
	"github.com/gofiber/fiber/v2"
	adminpkg "hrms-api/app/service/admin_pkg"

)

func SetupAdminhandler(app *fiber.App) {
	app.Post("/user_accounts/add", adminpkg.AddUserAccount)
}