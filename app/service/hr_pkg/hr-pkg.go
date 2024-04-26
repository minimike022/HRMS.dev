package hrpkg

import (
	Database "hrms-api/app/database"
	DataModels "hrms-api/app/model"

	//"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
)

var db = Database.Connect()

func GetApplicantsData(ctx *fiber.Ctx) error {
	applicants_data_model := new(DataModels.ApplicantsData)
	applicants_data_array := make([]DataModels.ApplicantsData, 0)

	db_query := `SELECT * FROM applicants_data`

	db_response, err := db.Query(db_query)
	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
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
	defer db_response.Close()


	return ctx.Status(fiber.StatusOK).JSON(applicants_data_array)
}

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

