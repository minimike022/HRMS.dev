package cjobs

import (
	mjobs "hrms-api/app/model/jobs"
	sjobs "hrms-api/app/service/jobs"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetJobPosition(ctx *fiber.Ctx) error {
	count := sjobs.CountJobs()

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "0"))

	offset := (page - 1) * limit

	jobs_list, err := sjobs.FetchJobs(offset, limit)

	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"count": count,
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
	count := sjobs.CountJobs()

	page, _ := strconv.Atoi(ctx.Query("page","1"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	offset := (page - 1) * limit

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

	jobs_list, err := sjobs.FetchJobs(page, offset)

	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"count": count,
		"job_positions": jobs_list,
	})
}