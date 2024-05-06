package login

import (
	Database "hrms-api/app/database"
	model_users "hrms-api/app/model/users"
	generatejwt "hrms-api/app/service/jwt/generate"
	util "hrms-api/app/util"
	"time"
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
		db_response.Scan(
			&login_account_model.Account_ID, 
			&login_account_model.Username, 
			&store_password, 
			&login_account_model.User_Role, 
			&login_account_model.User_Name)
	}

	err = util.CompareHash(store_password, login_account_model.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error" : "Incorrect Username or Password",
		})
	}

	refresh_token, err := generatejwt.GenerateRefreshToken(login_account_model.User_Name, login_account_model.Account_ID, login_account_model.User_Role)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
	}

	access_token, err := generatejwt.GenerateAccessToken(login_account_model.User_Name, login_account_model.Account_ID, login_account_model.User_Role)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
	}

	refresh_cookie := fiber.Cookie {
		Name: "refresh_token",
		Value: refresh_token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	access_cookie := fiber.Cookie {
		Name: "access_token",
		Value: access_token,
		Expires: time.Now().Add(time.Minute * 15),
		HTTPOnly: true,
	}

	ctx.Cookie(&refresh_cookie)
	ctx.Cookie(&access_cookie)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Logged In",
	})
}