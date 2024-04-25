package service

import (
	//Database "hrms-api/app/database"
	"fmt"
	Database "hrms-api/app/database"
	DataModels "hrms-api/app/model"

	"github.com/gofiber/fiber/v2"
)

// func GetData(ctx *fiber.Ctx) error {

// }
var db = Database.Connect()

//All
func ReadApplicantsData(ctx *fiber.Ctx) error {
	applicants_data_model := new(DataModels.ApplicantsData)
	applicants_data_array := make([]DataModels.ApplicantsData, 0)

	dbRows, err := db.Query("SELECT * FROM applicants_data")
	if err != nil {
		panic(err.Error())
	}

	for dbRows.Next() {
		dbRows.Scan(
			&applicants_data_model.Applicant_ID,
			&applicants_data_model.Position_ID,
			&applicants_data_model.First_Name,
			&applicants_data_model.Middle_Name,
			&applicants_data_model.Last_Name,
			&applicants_data_model.Extension_Name,
			&applicants_data_model.Birthdate,
			&applicants_data_model.Age,
			&applicants_data_model.Present_Address,
			&applicants_data_model.Highest_Education,
			&applicants_data_model.Email_Address,
			&applicants_data_model.Facebook_Link,
			&applicants_data_model.BPO_Exp,
			&applicants_data_model.Shift_Sched,
			&applicants_data_model.Work_Report,
			&applicants_data_model.Work_Site_Location,
			&applicants_data_model.Platform_ID,
			&applicants_data_model.Ref_Full_Name,
			&applicants_data_model.Ref_Company,
			&applicants_data_model.Ref_Position,
			&applicants_data_model.Ref_Contact_Num,
			&applicants_data_model.Ref_Email,
			&applicants_data_model.Applicant_CV,
			&applicants_data_model.Applicant_Portfolio_Link,
			&applicants_data_model.Applicant_Status_ID)
		applicants_data_array = append(applicants_data_array, *applicants_data_model)

	}
	fmt.Println(applicants_data_array)
	return ctx.Status(fiber.StatusOK).JSON(applicants_data_array)
}

//Guest
func PostApplicantsData(ctx *fiber.Ctx) error {
	applicants_data_model := DataModels.ApplicantsData{}
	fmt.Println(applicants_data_model)
	//db := Database.Connect()
	applicants_data := new(DataModels.ApplicantsData)
	fmt.Println(applicants_data)
	err := ctx.BodyParser(applicants_data)

	if err != nil {
		panic(err.Error())
	}

	applicants_data_model = DataModels.ApplicantsData{
		Position_ID:              applicants_data.Position_ID,
		First_Name:               applicants_data.First_Name,
		Middle_Name:              applicants_data.Middle_Name,
		Last_Name:                applicants_data.Last_Name,
		Extension_Name:           applicants_data.Extension_Name,
		Birthdate:                applicants_data.Birthdate,
		Age:                      applicants_data.Age,
		Present_Address:          applicants_data.Present_Address,
		Highest_Education:        applicants_data.Highest_Education,
		Email_Address:            applicants_data.Email_Address,
		Facebook_Link:            applicants_data.Facebook_Link,
		BPO_Exp:                  applicants_data.BPO_Exp,
		Shift_Sched:              applicants_data.Shift_Sched,
		Work_Report:              applicants_data.Work_Report,
		Work_Site_Location:       applicants_data.Work_Site_Location,
		Platform_ID:              applicants_data.Platform_ID,
		Ref_Full_Name:            applicants_data.Ref_Full_Name,
		Ref_Company:              applicants_data.Ref_Company,
		Ref_Position:             applicants_data.Ref_Position,
		Ref_Contact_Num:          applicants_data.Ref_Contact_Num,
		Ref_Email:                applicants_data.Ref_Email,
		Applicant_CV:             applicants_data.Applicant_CV,
		Applicant_Portfolio_Link: applicants_data.Applicant_Portfolio_Link,
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

	dbData, err := Database.Connect().Query(dbQuery,
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

//All
func GetApplicationStatus(ctx *fiber.Ctx) error {
	app_status_data := new(DataModels.ApplicantStatus)
	app_status_model := make([]DataModels.ApplicantStatus, 0)

	query := `SELECT applicants_data.first_name, applicants_data.middle_name, applicants_data.last_name,
	job_position.position_name,
	department.department_name,
	application_status_list.application_status_name,
	user_accounts.user_name
	FROM application_status
	INNER JOIN applicants_data ON application_status.applicant_id = applicants_data.applicant_id
	INNER JOIN application_status_list ON application_status.application_status_id = application_status_list.application_status_id
	INNER JOIN job_position ON applicants_data.position_id = job_position.position_id
	INNER JOIN department ON job_position.department_id = department.department_id
	INNER JOIN user_accounts ON job_position.department_id = user_accounts.department_id`

	db_response, err := Database.Connect().Query(query)
	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
			&app_status_data.Applicant_First_Name,
			&app_status_data.Applicant_Middle_Name,
			&app_status_data.Applicant_Last_Name,
			&app_status_data.Job_Position_Name,
			&app_status_data.Department_Name,
			&app_status_data.Application_Status,
			&app_status_data.User_Interviewee_Name,
		)
		app_status_model = append(app_status_model, *app_status_data)

	}
	defer db_response.Close()
	return ctx.Status(fiber.StatusOK).JSON(app_status_model)
}

// func GetApplicationStatus(ctx *fiber.Ctx) error {
// 	app_status_data := new(DataModels.ApplicantStatus)
// 	app_status_model := make([]DataModels.ApplicantStatus, 0)
// 	fmt.Println(app_status_data)
// 	query := `SELECT first_name, middle_name, last_name,
// 	application_status_list.application_status_name,
// 	job_position.position_name,
// 	department.department_name,
// 	user_accounts.user_name
// 	FROM application_status
// 	INNER JOIN application_status_list ON applicants_data.application_status_id = application_status_list.application_status_id
// 	INNER JOIN job_position ON applicants_data.position_id = job_position.position_id
// 	INNER JOIN department ON job_position.department_id = department.department_id
// 	INNER JOIN user_accounts ON job_position.department_id = user_accounts.department_id`

// 	dbData, err := Database.Connect().Query(query)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	for dbData.Next() {
// 		dbData.Scan(
// 			&app_status_data.Applicant_First_Name,
// 			&app_status_data.Applicant_Middle_Name,
// 			&app_status_data.Applicant_Last_Name,
// 			&app_status_data.Job_Position_Name,
// 			&app_status_data.Department_Name,
// 			&app_status_data.Application_Status,
// 			&app_status_data.User_Interviewee_Name,
// 		)
// 			app_status_model = append(app_status_model,*app_status_data)

// 	}
// 	fmt.Println(app_status_model)
// 	return ctx.Status(fiber.StatusOK).JSON(app_status_model)
// }

//Admin
func AddUserAccount(ctx *fiber.Ctx) error {
	var user_account_data DataModels.UserAccount
	user_account := new(DataModels.UserAccount)
	err := ctx.BodyParser(user_account)
	if err != nil  {
		panic(err.Error())
	}

	user_account_data = DataModels.UserAccount{
		Username: user_account.Username,
		Password: user_account.Password,
		User_Role: user_account.User_Role,
		User_Position: user_account.User_Position,
		Department_ID: user_account.Department_ID,
		Account_Status: user_account.Account_Status,
	}

	db_query := `INSERT INTO user_accounts (username, password, user_role, user_name, user_position, department_id)
	VALUES (?,?,?,?,?,?)`
	
	db_response, err := db.Query(db_query, 
		user_account_data.Username,
		user_account.Password,
		user_account.User_Role,
		user_account.User_Name,
		user_account.User_Position,
		user_account.Department_ID,
	)
	if err != nil {
		panic(err.Error())
	}
	defer db_response.Close()

	return ctx.Status(fiber.StatusOK).SendString("User Added!")
}

//Admin
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
	account_status := "Inactive"

	query := `UPDATE user_accounts
	SET account_status = ?
	WHERE account_id = ?
	`
	_, err := db.Query(query, account_status, user_account_id)
	if err != nil{
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).SendString("Updated")
}
