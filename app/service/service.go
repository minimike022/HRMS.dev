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
	appData := new(DataModels.ApplicantsData)
	applicantsData := make([]DataModels.ApplicantsData, 0)
	fmt.Println(appData)

	dbRows, err := db.Query("SELECT * FROM applicants_data")
	if err != nil {
		panic(err.Error())
	}

	for dbRows.Next() {
		dbRows.Scan(
			&appData.Applicant_ID,
			&appData.Position_ID,
			&appData.First_Name,
			&appData.Middle_Name,
			&appData.Last_Name,
			&appData.Extension_Name,
			&appData.Birthdate,
			&appData.Age,
			&appData.Present_Address,
			&appData.Highest_Education,
			&appData.Email_Address,
			&appData.Facebook_Link,
			&appData.BPO_Exp,
			&appData.Shift_Sched,
			&appData.Work_Report,
			&appData.Work_Site_Location,
			&appData.Platform_ID,
			&appData.Ref_Full_Name,
			&appData.Ref_Company,
			&appData.Ref_Position,
			&appData.Ref_Contact_Num,
			&appData.Ref_Email,
			&appData.Applicant_CV,
			&appData.Applicant_Portfolio_Link,
			&appData.Applicant_Status_ID)
		applicantsData = append(applicantsData, *appData)

	}
	fmt.Println(applicantsData)
	return ctx.Status(fiber.StatusOK).JSON(applicantsData)
}

//Guest
func PostApplicantsData(ctx *fiber.Ctx) error {
	var appData DataModels.ApplicantsData
	//db := Database.Connect()
	applicationData := new(DataModels.ApplicantsData)
	err := ctx.BodyParser(applicationData)

	if err != nil {
		panic(err.Error())
	}

	appData = DataModels.ApplicantsData{
		Position_ID:              applicationData.Position_ID,
		First_Name:               applicationData.First_Name,
		Middle_Name:              applicationData.Middle_Name,
		Last_Name:                applicationData.Last_Name,
		Extension_Name:           applicationData.Extension_Name,
		Birthdate:                applicationData.Birthdate,
		Age:                      applicationData.Age,
		Present_Address:          applicationData.Present_Address,
		Highest_Education:        applicationData.Highest_Education,
		Email_Address:            applicationData.Email_Address,
		Facebook_Link:            applicationData.Facebook_Link,
		BPO_Exp:                  applicationData.BPO_Exp,
		Shift_Sched:              applicationData.Shift_Sched,
		Work_Report:              applicationData.Work_Report,
		Work_Site_Location:       applicationData.Work_Site_Location,
		Platform_ID:              applicationData.Platform_ID,
		Ref_Full_Name:            applicationData.Ref_Full_Name,
		Ref_Company:              applicationData.Ref_Company,
		Ref_Position:             applicationData.Ref_Position,
		Ref_Contact_Num:          applicationData.Ref_Contact_Num,
		Ref_Email:                applicationData.Ref_Email,
		Applicant_CV:             applicationData.Applicant_CV,
		Applicant_Portfolio_Link: applicationData.Applicant_Portfolio_Link,
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
		appData.Position_ID,
		appData.First_Name,
		appData.Middle_Name,
		appData.Last_Name,
		appData.Extension_Name,
		appData.Birthdate,
		appData.Age,
		appData.Present_Address,
		appData.Highest_Education,
		appData.Email_Address,
		appData.Facebook_Link,
		appData.BPO_Exp,
		appData.Shift_Sched,
		appData.Work_Report,
		appData.Work_Site_Location,
		appData.Platform_ID,
		appData.Ref_Full_Name,
		appData.Ref_Company,
		appData.Ref_Position,
		appData.Ref_Contact_Num,
		appData.Ref_Email,
		appData.Applicant_CV,
		appData.Applicant_Portfolio_Link)

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

	dbData, err := Database.Connect().Query(query)
	if err != nil {
		panic(err.Error())
	}

	for dbData.Next() {
		dbData.Scan(
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
	dbData.Close()
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

	query := `INSERT INTO user_accounts (username, password, user_role, user_name, user_position, department_id)
	VALUES (?,?,?,?,?,?)`
	
	db_query, err := db.Query(query, 
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
	defer db_query.Close()

	return ctx.Status(fiber.StatusOK).SendString("User Added!")
}

//Admin
func ReadUserAccounts(ctx *fiber.Ctx) error {
	user_accounts_data := new(DataModels.UserAccount)
	user_accounts_array := make([]DataModels.UserAccount,0)

	query := "SELECT * FROM user_accounts"

	db_query, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	for db_query.Next() {
		db_query.Scan(
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
	db_query.Close()
	
	return ctx.Status(fiber.StatusOK).JSON(user_accounts_array)
}

func ChangeAccountStatus(ctx *fiber.Ctx) error {
	user_account_id := ctx.Params("account_id")
	account_status := "Active"

	query := `UPDATE TABLE user_accounts
	SET account_status = ?
	WHERE account_id = ?
	`
	db_query, err := db.Query(query, account_status, user_account_id)
	if err != nil{
		panic(err.Error())
	}
	defer db_query.Close()

	return ctx.Status(fiber.StatusOK).JSON(db_query)
}
