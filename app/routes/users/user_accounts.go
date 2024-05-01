package husers

import (
	"github.com/gofiber/fiber/v2"
	user_accounts "hrms-api/app/service/users/accounts"
	login "hrms-api/app/service/users/login"
)


func SetupUserAccounts(app *fiber.App) {
	app.Get("/user/login", login.Login)
	app.Post("/user/accounts/add", user_accounts.AddUserAccount)
	app.Get("/users/accounts", user_accounts.GetUserAccounts)
	app.Patch("/users/accounts/update/:id", user_accounts.UpdateAccountStatus)
}