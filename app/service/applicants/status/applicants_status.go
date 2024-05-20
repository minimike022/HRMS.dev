package applicantsStatus

import (
	Database "hrms-api/app/database"
	model_applicants "hrms-api/app/model/applicants"
	application_status "hrms-api/app/model/application_status"
	status_list "hrms-api/app/model/status_list"
	model_jobs "hrms-api/app/model/jobs"
	model_users "hrms-api/app/model/users"
	"github.com/gofiber/fiber/v2"
)

var db = Database.Connect()

func GetApplicationStatus(ctx *fiber.Ctx) error {
	applicants_data := new(model_applicants.ApplicantsData)
	application_status := new(application_status.ApplicantStatus)
	status_list := new(status_list.ApplicationStatusList)
	job_position := new(model_jobs.JobPosition)
	user_accounts := new(model_users.UserAccount)

	app_data := make([]application_status.Handler_Application_Status, 0)

	query := `CALL fetch_application_status`

	db_response, err := Database.Connect().Query(query)
	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
			&applicants_data.First_Name,
			&applicants_data.Middle_Name,
			&applicants_data.Last_Name,
			&applicants_data.Extension_Name,
			&job_position.Position_Name,
			&user_accounts.User_Name,
			&status_list.Application_Status_Name,
			&application_status.Interview_Date,
			&application_status.Interview_Time,
		)
		app_status_model = append(app_status_model, *applicants_data, )

	}
	defer db_response.Close()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": app_status_model,
	})
}

func UpdateApplicationStatus(ctx *fiber.Ctx) error {
	application_id := ctx.Params("id")
	application_status_model := model_applicants.ApplicantStatus{}
	application_status_list := model_applicants.Application_Status_List{}

	err := ctx.BodyParser(&application_status_model)

	if err != nil {
		panic(err.Error())
	}

	db_query := "CALL update_application_status(?,?,?,?,?)"

	_, err = db.Query(db_query, application_id, application_status_list.Application_Status_ID, application_status_model.User_Interviewee_ID, application_status_model.Interview_Date, application_status_model.Interview_Time)
	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON("Application Status Updated")
}