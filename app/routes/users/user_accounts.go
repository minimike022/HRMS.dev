package husers

import (
	"github.com/gofiber/fiber/v2"
	caccounts "hrms-api/app/controller/users/accounts"
	login "hrms-api/app/service/users/login"
	jwtvalidate "hrms-api/app/service/jwt/validate"
	validaterole "hrms-api/app/service/users/validate"
)

func SetupUserAccounts(app *fiber.App) {
	app.Get("/users", caccounts.FetchUsers)
	app.Post("/user/login", login.Login)
	app.Post("/user/accounts/add", jwtvalidate.ValidateAccessToken, validaterole.ValidateAdmin, caccounts.AddUserAccounts)
	app.Get("/users/accounts", jwtvalidate.ValidateAccessToken, validaterole.ValidateAdmin, caccounts.FetchUserAccounts)
	app.Patch("/users/accounts/update/:id", jwtvalidate.ValidateAccessToken, caccounts.UpdateUserAccounts)
}