package service

import (
	Database "hrms-api/app/database"
	DataModels "hrms-api/app/model"

	"github.com/golang-jwt/jwt/v5"

	//"gopkg.in/gomail.v2"
	"github.com/gofiber/fiber/v2"
)
var jwtSecret = "143"
var db = Database.Connect()

//Guest
func PostApplicantsData(ctx *fiber.Ctx) error {
	applicants_data_model := DataModels.ApplicantsData{}
	err := ctx.BodyParser(&applicants_data_model)

	if err != nil {
		panic(err.Error())
	}
	dbQuery := `INSERT INTO applicants_data (
				position_id, 
				first_name, 
				middle_name, 
				last_name, 
				extension_name, 
				birthdate, 
				age, 
				present_address, 
				highest_education, 
				email_address, 
				facebook_link, 
				bpo_exp, 
				shift_sched, 
				work_report, 
				work_site_location, 
				platform_id, 
				ref_full_name, 
				ref_company, 
				ref_position, 
				ref_contact_num, 
				ref_email, 
				applicant_cv, 
				applicant_portfolio_link
			) 
			VALUES 
			(
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 
				?, ?, ?, ?, ?, ?, ?
			)`

	dbData, err := db.Query(dbQuery,
		applicants_data_model.Position_ID,
		applicants_data_model.First_Name,
		applicants_data_model.Middle_Name,
		applicants_data_model.Last_Name,
		applicants_data_model.Extension_Name,
		applicants_data_model.Birthdate,
		applicants_data_model.Age,
		applicants_data_model.Present_Address,
		applicants_data_model.Highest_Education,
		applicants_data_model.Email_Address,
		applicants_data_model.Facebook_Link,
		applicants_data_model.BPO_Exp,
		applicants_data_model.Shift_Sched,
		applicants_data_model.Work_Report,
		applicants_data_model.Work_Site_Location,
		applicants_data_model.Platform_ID,
		applicants_data_model.Ref_Full_Name,
		applicants_data_model.Ref_Company,
		applicants_data_model.Ref_Position,
		applicants_data_model.Ref_Contact_Num,
		applicants_data_model.Ref_Email,
		applicants_data_model.Applicant_CV,
		applicants_data_model.Applicant_Portfolio_Link)
	if err != nil {
		panic(err.Error())
	}
	defer dbData.Close()

	return ctx.Status(fiber.StatusOK).SendString("Added to database!")
}

func Login(ctx *fiber.Ctx) error {
	login_account_model := DataModels.UserAccount{}
	err := ctx.BodyParser(&login_account_model) 
	if err != nil {
		panic(err.Error())
	}

	db_query := `SELECT account_id, username, password FROM user_accounts WHERE username = ?`
	db_response, err := db.Query(db_query, login_account_model.Username)
	if err != nil {
		panic(err.Error())
	}
	var store_password string
	for db_response.Next() {
		db_response.Scan(&login_account_model.Account_ID, &login_account_model.Username, &store_password)
	}

	if login_account_model.Password != store_password {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Username or Password",
		})
	}
	
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = login_account_model.Account_ID
	
	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err !=nil {
		ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token": tokenString})
}



