package jobs

import (
	Database "hrms-api/app/database"
	model_jobs "hrms-api/app/model/jobs"

	"github.com/gofiber/fiber/v2"
)

var db = Database.Connect()


func GetJobPosition(ctx *fiber.Ctx) error {
	job_position := model_jobs.JobPosition{}
	job_position_array := make([]model_jobs.JobPosition, 0)

	db_query := "CALL fetch_jobs_position"

	db_response, err := db.Query(db_query)

	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
			&job_position.Position_ID,
			&job_position.Position_Name,
			&job_position.Department_ID,
			&job_position.Position_Status,
			&job_position.Available_Slot,
		)
		job_position_array = append(job_position_array, job_position)
	}

	return ctx.Status(fiber.StatusOK).JSON(job_position_array)
}

func AddJobPosition(ctx *fiber.Ctx) error {
	job_position := model_jobs.JobPosition{}
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
	job_position := model_jobs.JobPosition{}
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