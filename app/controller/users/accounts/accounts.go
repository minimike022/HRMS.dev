package caccounts

import (
	"fmt"
	saccounts "hrms-api/app/service/users/accounts"
	musers "hrms-api/app/model/users"
	util "hrms-api/app/util"
	"log"
	//"time"

	//"github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/fiber/v2"
)

func FetchUsers(ctx *fiber.Ctx) error {
	users , err := saccounts.FetchUsers()

	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
	"users": users,
	})
	
}

func FetchUserAccounts(ctx *fiber.Ctx) error {
	user_accounts, err := saccounts.FetchUserAccounts()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Print(user_accounts)


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_accounts": user_accounts,
	})
}

func AddUserAccounts(ctx *fiber.Ctx) error {
	user_account_data := musers.UserAccount{}
	err := ctx.BodyParser(&user_account_data)
	if err != nil  {
		panic(err.Error())
	}

	//Password Hashing
	var user_hashed_password = util.HashedPassword(user_account_data.Password)

	add_response := saccounts.AddUser(user_account_data, user_hashed_password)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map {
		"Response Status": add_response,
		"msg": "User Added!",
	})
}

func UpdateUserAccounts(ctx *fiber.Ctx) error {
	user_accounts_model := musers.UserAccount{}
	user_account_id := ctx.Params("id")
	
	err := ctx.BodyParser(&user_accounts_model)
	if err != nil {
		panic(err.Error())
	}

	err = saccounts.UpdateUser(user_account_id, user_accounts_model)
	if err != nil {
		panic(err.Error())
	}


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Updated Successful",
	})
}