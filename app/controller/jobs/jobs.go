package cjobs

import (
	"fmt"
	mjobs "hrms-api/app/model/jobs"
	sjobs "hrms-api/app/service/jobs"

	"github.com/gofiber/fiber/v2"
)

func GetJobPosition(ctx *fiber.Ctx) error {
	count := sjobs.CountJobs()
	fmt.Print(count)
	jobs_list, err := sjobs.FetchJobs()

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

	err = sjobs.UpdateJobs(job_position_id, job_position)

	if err != nil {
		panic(err.Error())
	}


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Update Succesfully!",
	})

}

func SearchJobs(ctx *fiber.Ctx) error {
	search_query := ctx.Query("q")
	
	if len(search_query) > 0 {
		jobs_list, err := sjobs.SearchJobs(search_query)

		if err != nil {
			panic(err.Error())
		}
	
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"job_positions": jobs_list,
		})
	}

	jobs_list, err := sjobs.FetchJobs()

	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"job_positions": jobs_list,
	})
}