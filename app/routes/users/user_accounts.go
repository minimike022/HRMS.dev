package husers

import (
	"github.com/gofiber/fiber/v2"
	user_accounts "hrms-api/app/service/users/accounts"
	login "hrms-api/app/service/users/login"
	jwtvalidate "hrms-api/app/service/jwt/validate"
	validaterole "hrms-api/app/service/users/validate"
)

func SetupUserAccounts(app *fiber.App) {
	app.Post("/user/login", login.Login)
	app.Post("/user/accounts/add", jwtvalidate.ValidateAccessToken, validaterole.ValidateAdmin, user_accounts.AddUserAccount)
	app.Get("/users/accounts", jwtvalidate.ValidateAccessToken, validaterole.ValidateAdmin, user_accounts.GetUserAccounts)
	app.Patch("/users/accounts/update/:id", jwtvalidate.ValidateAccessToken, user_accounts.UpdateAccountStatus)
}