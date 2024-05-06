package applicantsStatus

import (
	Database "hrms-api/app/database"
	model_applicants "hrms-api/app/model/applicants"
	"time"
	"github.com/gofiber/fiber/v2"
)

var db = Database.Connect()

func GetApplicationStatus(ctx *fiber.Ctx) error {
	app_status_data := new(model_applicants.ApplicantStatus)
	app_status_model := make([]model_applicants.ApplicantStatus, 0)

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
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": app_status_model,
	})
}

func UpdateApplicationStatus(ctx *fiber.Ctx) error {
	updated_at := time.Now().Format("2006-01-02 15:04:05") 
	application_id := ctx.Params("id")
	application_status_model := model_applicants.ApplicantStatus{}

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