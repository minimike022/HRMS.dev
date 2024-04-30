package hrpkg

import (
	Database "hrms-api/app/database"
	DataModels "hrms-api/app/model"
	"github.com/gofiber/fiber/v2"
)

var db = Database.Connect()

func GetApplicantsData(ctx *fiber.Ctx) error {
	applicants_data_model := new(DataModels.ApplicantsData)
	applicants_data_array := make([]DataModels.ApplicantsData, 0)

	db_query := "CALL fetch_applicants_data"

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
			&applicants_data_model.Application_CreatedAt,
		)
		applicants_data_array = append(applicants_data_array, *applicants_data_model)
	}
	
	defer db_response.Close()


	return ctx.Status(fiber.StatusOK).JSON(applicants_data_array)
}

func GetApplicationStatus(ctx *fiber.Ctx) error {
	app_status_data := new(DataModels.ApplicantStatus)
	app_status_model := make([]DataModels.ApplicantStatus, 0)

	query := `CALL fetch_application_status`

	db_response, err := Database.Connect().Query(query)
	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
			&app_status_data.Applicant_ID,
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


func AddJobPosition(ctx *fiber.Ctx) error {
	job_position := DataModels.JobPosition{}
	err := ctx.BodyParser(&job_position)
	if err != nil {
		panic(err.Error())
	}

	db_query := `CALL add_job_slot(?,?,?,?)`
	db_response, err := db.Query(db_query, 
	job_position.Position_Name,
	job_position.Department_ID,
	job_position.Position_Status,
	job_position.Available_Slot,
	)

	if err != nil {
		panic(err.Error())
	}
	defer db_response.Close()

	return ctx.Status(fiber.StatusOK).JSON(job_position)
}

func UpdateJobPosition(ctx *fiber.Ctx) error {
	job_position_id := ctx.Params("id")
	job_position := DataModels.JobPosition{}
	err := ctx.BodyParser(&job_position)
	if err != nil {
		panic(err.Error())
	}

	db_query := `CALL update_job_position(?,?,?,?,?)`

	db_response, err := db.Query(db_query,
	job_position_id, 
	job_position.Position_Name, 
	job_position.Department_ID, 
	job_position.Position_Status,
	job_position.Available_Slot,
	)

	if err != nil {
		panic(err.Error())
	}
	defer db_response.Close()


	return ctx.Status(fiber.StatusOK).SendString("Position Updated")
}
