package adminpkg
import (
	Database "hrms-api/app/database"
	DataModels "hrms-api/app/model"

	//"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
)

var db = Database.Connect()

func AddUserAccount(ctx *fiber.Ctx) error {
	user_account_data := DataModels.UserAccount{}
	err := ctx.BodyParser(&user_account_data)
	if err != nil  {
		panic(err.Error())
	}

	db_query := `INSERT INTO user_accounts (username, password, user_role, user_name, user_position, department_id)
	VALUES (?,?,?,?,?,?)`
	
	db_response, err := db.Query(db_query, 
		user_account_data.Username,
		user_account_data.Password,
		user_account_data.User_Role,
		user_account_data.User_Name,
		user_account_data.User_Position,
		user_account_data.Department_ID,
	)
	if err != nil {
		panic(err.Error())
	}
	defer db_response.Close()

	return ctx.Status(fiber.StatusOK).SendString("User Added!")
}

func ReadUserAccounts(ctx *fiber.Ctx) error {
	user_accounts_data := new(DataModels.UserAccount)
	user_accounts_array := make([]DataModels.UserAccount,0)

	db_query := "SELECT * FROM user_accounts"

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
		)
		user_accounts_array = append(user_accounts_array, *user_accounts_data)

	}
	defer db_response.Close()
	
	return ctx.Status(fiber.StatusOK).JSON(user_accounts_array)
}

func ChangeAccountStatus(ctx *fiber.Ctx) error {
	user_account_id := ctx.Params("id")
	var account_status string
	err := ctx.BodyParser(&account_status)
	if err !=nil {
		panic(err.Error())
	}
	query := `UPDATE user_accounts
	SET account_status = ?
	WHERE account_id = ?
	`
	_, err = db.Query(query, account_status, user_account_id)
	if err != nil{
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).SendString("Updated")
}

func GetSourceData(ctx *fiber.Ctx) error {
	posting_platform_model := new(DataModels.PostingPlatform)
	posting_platform_array := make([]DataModels.PostingPlatform,0)

	db_query := `SELECT job_posting_platform.platform_id, job_posting_platform.platform_name, 
	COUNT(*) FROM applicants_data
	INNER JOIN Job_posting_platform ON job_posting_platform.platform_id = applicants_data.platform_id
	GROUP BY platform_name
	`
	db_response, err := db.Query(db_query)
	if err != nil {
		panic(err.Error())
	}
	defer db_response.Close()
	for db_response.Next() {
		db_response.Scan(
			&posting_platform_model.Platform_ID,
			&posting_platform_model.Platform_Name,
			&posting_platform_model.Platform_Count)
		posting_platform_array = append(posting_platform_array, *posting_platform_model)
	}

	return ctx.Status(fiber.StatusOK).JSON(posting_platform_array)
} 