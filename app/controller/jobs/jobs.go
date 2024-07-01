package cjobs

import (
	mjobs "hrms-api/app/model/jobs"
	sjobs "hrms-api/app/service/jobs"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetJobPosition(ctx *fiber.Ctx) error {

	count := sjobs.CountJobs()

	sort_col := `JP.available_slot`
	sort_order := `DESC`

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "0"))
	
	offset := (page - 1) * limit

	if offset == 0 && limit == 0 {
		jobs_list := sjobs.FetchJobs()

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"count": count,
			"job_positions": jobs_list,
		})
	}

	jobs_list, err := sjobs.SortJobs(offset, limit, sort_col, sort_order)

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
	
	search_query := ctx.Query("q")

	sort_col := ctx.Query("sort_col")
	sort_order := ctx.Query("sort_order")

	page, _ := strconv.Atoi(ctx.Query("page","1"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	offset := (page - 1) * limit

	count := sjobs.SearchCount(search_query)
	
	if len(search_query) > 0 {
		jobs_list:= sjobs.SearchJobs(search_query, offset, limit, sort_col, sort_order)
	
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"count": count,
			"job_positions": jobs_list,
		})
	}

	jobs_list, err := sjobs.SortJobs(offset, limit, sort_col, sort_order)

	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"count": count,
		"job_positions": jobs_list,
	})
}