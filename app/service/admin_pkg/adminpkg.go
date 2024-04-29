package adminpkg

import (
	Database "hrms-api/app/database"
	DataModels "hrms-api/app/model"
	"time"
	"golang.org/x/crypto/bcrypt"
	//"github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/fiber/v2"
)

var db = Database.Connect()

//Admin Create User Account
func AddUserAccount(ctx *fiber.Ctx) error {
	created_at := time.Now().Format("2006-01-02 15:04:05") 
	user_account_data := DataModels.UserAccount{}
	err := ctx.BodyParser(&user_account_data)
	if err != nil  {
		panic(err.Error())
	}

	//Password Hashing
	user_hashed_password, err := bcrypt.GenerateFromPassword([]byte(user_account_data.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}

	//Calling Procedured Query
	db_query := `CALL add_user_accounts(?,?,?,?,?,?,?)`
	
	db_response, err := db.Query(db_query, 
		user_account_data.Username,
		user_hashed_password,
		user_account_data.User_Role,
		user_account_data.User_Name,
		user_account_data.User_Position,
		user_account_data.Department_ID,
		created_at,
	)
	if err != nil {
		panic(err.Error())
	}
	defer db_response.Close()


	return ctx.Status(fiber.StatusOK).SendString("User Added!")
}

//Fetching User Accounts
func GetUserAccounts(ctx *fiber.Ctx) error {
	user_accounts_data := new(DataModels.UserAccount)
	user_accounts_array := make([]DataModels.UserAccount,0)

	//Calling Procedured Query
	db_query := "CALL fetch_user_accounts"

	db_response, err := db.Query(db_query)
	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
			&user_accounts_data.Account_ID,
			&user_accounts_data.Username,
			&user_accounts_data.Password,
			&user_accounts_data.User_Role,
			&user_accounts_data.User_Name,
			&user_accounts_data.User_Position,
			&user_accounts_data.Department_ID,
			&user_accounts_data.Account_Status,
			&user_accounts_data.CreatedAt,
		)
		user_accounts_array = append(user_accounts_array, *user_accounts_data)

	}
	defer db_response.Close()
	
	return ctx.Status(fiber.StatusOK).JSON(user_accounts_array)
}

//Updating User Account Status
func UpdateAccountStatus(ctx *fiber.Ctx) error {
	user_accounts_model := DataModels.UserAccount{}
	user_account_id := ctx.Params("id")
	
	err := ctx.BodyParser(&user_accounts_model)
	if err !=nil {
		panic(err.Error())
	}
	query := `CALL update_account_status(?, ?)`
	_, err = db.Query(query, user_account_id, user_accounts_model.Account_Status)
	if err != nil{
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).SendString("Updated")
}

