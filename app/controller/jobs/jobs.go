package cjobs

import (
	mjobs "hrms-api/app/model/jobs"
	cjobs "hrms-api/app/service/jobs"
	sjobs "hrms-api/app/service/jobs"

	"github.com/gofiber/fiber/v2"
)

func GetJobPosition(ctx *fiber.Ctx) error {
	jobs_list, err := cjobs.FetchJobs()

	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"job_positions": jobs_list,
	})
}

func AddJobPosition(ctx *fiber.Ctx) error {
	job_position := mjobs.JobPosition{}
	err := ctx.BodyParser(&job_position)
	if err != nil {
		panic(err.Error())
	}

	err = sjobs.AddJobs(job_position)

	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Added Succesfully!",
	})
}

func UpdateJobPosition(ctx *fiber.Ctx) error {
	job_position_id := ctx.Params("id")

	job_position := mjobs.JobPosition{}

	err := ctx.BodyParser(&job_position)
	
	if err != nil {
		panic(err.Error())
	}

	err = cjobs.UpdateJobs(job_position_id, job_position)

	if err != nil {
		panic(err.Error())
	}


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Update Succesfully!",
	})

}