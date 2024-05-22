package applicantsStatus

import (
	Database "hrms-api/app/database"
	application_status "hrms-api/app/model/application_status"

	"github.com/gofiber/fiber/v2"
)

var db = Database.Connect()

func GetApplicationStatus(ctx *fiber.Ctx) error {
	application_status_model := new(application_status.Application_Status)
	application_status_array := make([]application_status.Application_Status, 0)

	query := `CALL fetch_application_status`

	db_response, err := Database.Connect().Query(query)
	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
			&application_status_model.Status_ID,
			&application_status_model.First_Name,
			&application_status_model.Middle_Name,
			&application_status_model.Last_Name,
			&application_status_model.Extension_Name,
			&application_status_model.Position_Name,
			&application_status_model.Interviewee_Name,
			&application_status_model.Application_Status,
			&application_status_model.Interview_Date,
			&application_status_model.Interview_Time,
		)
		application_status_array = append(application_status_array, *application_status_model, )

	}
	defer db_response.Close()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"application_status": application_status_array,
	})
}

func UpdateApplicationStatus(ctx *fiber.Ctx) error {
	application_id := ctx.Params("id")
	application_status_model := application_status.ApplicantStatus{}

	err := ctx.BodyParser(&application_status_model)

	if err != nil {
		panic(err.Error())
	}

	db_query := "CALL update_application_status(?,?,?,?,?)"

	_, err = db.Query(db_query, application_id, application_status_model.Application_Status_ID, application_status_model.User_Interviewee_ID, application_status_model.Interview_Date, application_status_model.Interview_Time)
	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON("Application Status Updated")
}