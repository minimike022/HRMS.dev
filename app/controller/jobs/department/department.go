package cdepartment

import (
	sdepartment "hrms-api/app/service/jobs/department"

	"github.com/gofiber/fiber/v2"
)

func GetDepartments(ctx *fiber.Ctx) error {
	departments, err := sdepartment.FetchDepartments()

	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"departments": departments,
	})
}