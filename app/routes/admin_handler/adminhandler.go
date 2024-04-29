package adminhandler

import (
	"github.com/gofiber/fiber/v2"
	adminpkg "hrms-api/app/service/admin_pkg"
)

func SetupAdminhandler(app *fiber.App) {
	app.Patch("/user_accounts/change_status/:id", adminpkg.UpdateAccountStatus)
	app.Get("/user_accounts", adminpkg.GetUserAccounts)
	app.Post("/user_accounts/add", adminpkg.AddUserAccount)
}	