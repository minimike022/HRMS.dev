package login

import (
	Database "hrms-api/app/database"
	model_users "hrms-api/app/model/users"
	jwt "hrms-api/app/service/jwt"
	util "hrms-api/app/util"

	"github.com/gofiber/fiber/v2"
)

var db = Database.Connect()


func Login(ctx *fiber.Ctx) error {
	var store_password string
	login_account_model := model_users.UserAccount{}
	err := ctx.BodyParser(&login_account_model) 
	if err != nil {
		panic(err.Error())
	}

	db_query := `CALL user_login(?)`
	db_response, err := db.Query(db_query, login_account_model.Username)
	if err != nil {
		panic(err.Error())
	}
	
	for db_response.Next() {
		db_response.Scan(&login_account_model.Account_ID, &login_account_model.Username, &store_password)
	}

	err = util.CompareHash(store_password, login_account_model.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Username or Password",
		})
	}

	tokenString, err := jwt.GenerateToken(login_account_model.User_Name, login_account_model.Account_ID)
	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": tokenString,
	})
}