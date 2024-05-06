package husers

import (
	"github.com/gofiber/fiber/v2"
	user_accounts "hrms-api/app/service/users/accounts"
	login "hrms-api/app/service/users/login"
	jwtvalidate "hrms-api/app/service/jwt/validate"
	validaterole "hrms-api/app/service/users/validate"
)

func SetupUserAccounts(app *fiber.App) {
	app.Get("/user/login", login.Login)
	app.Post("/user/accounts/add", jwtvalidate.ValidateRefreshToken, validaterole.ValidateAdmin, user_accounts.AddUserAccount)
	app.Get("/users/accounts", jwtvalidate.ValidateRefreshToken, validaterole.ValidateAdmin, user_accounts.GetUserAccounts)
	app.Patch("/users/accounts/update/:id", jwtvalidate.ValidateRefreshToken, user_accounts.UpdateAccountStatus)
}