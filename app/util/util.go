package util

import (
	Database "hrms-api/app/database"
	DataModels "hrms-api/app/model"
	"time"
	jwtware "github.com/gofiber/contrib/jwt"
	jwt "hrms-api/app/service/jwt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

)
var jwtSecret = "143"
var db = Database.Connect()

func JwtAuth() func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
	})
} 

func UpdateApplicationStatus(ctx *fiber.Ctx) error {
	updated_at := time.Now().Format("2006-01-02 15:04:05") 
	application_id := ctx.Params("id")
	application_status_model := DataModels.ApplicantStatus{}

	err := ctx.BodyParser(&application_status_model)

	if err != nil {
		panic(err.Error())
	}

	db_query := "CALL update_application_status(?,?, ?)"

	_, err = db.Query(db_query, application_id, application_status_model.Application_Status_ID, updated_at)
	if err != nil {
		panic(err.Error())
	}
	return ctx.Status(fiber.StatusOK).SendString("Application Status Updated!")
}

func PostApplicantsData(ctx *fiber.Ctx) error {
	created_at := time.Now().Format("2006-01-02 15:04:05") 
	applicants_data_model := DataModels.ApplicantsData{}
	err := ctx.BodyParser(&applicants_data_model)

	if err != nil {
		panic(err.Error())
	}
	dbQuery := `CALL add_applicants(?, ?, ?, ?, ?, ?, ?, ?, ?, 
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

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
		applicants_data_model.Applicant_Portfolio_Link,
		created_at)
	if err != nil {
		panic(err.Error())
	}
	defer dbData.Close()

	return ctx.Status(fiber.StatusOK).SendString("Added to database!")
}

func Login(ctx *fiber.Ctx) error {
	var store_password string
	login_account_model := DataModels.UserAccount{}
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

	err = bcrypt.CompareHashAndPassword([]byte(store_password), []byte(login_account_model.Password))

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

func GetPlatformData(ctx *fiber.Ctx) error {
	posting_platform_model := new(DataModels.PostingPlatform)
	posting_platform_array := make([]DataModels.PostingPlatform,0)

	db_query := `CALL fetch_posting_data()`
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