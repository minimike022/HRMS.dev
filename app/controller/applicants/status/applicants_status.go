package capplication_status

import (
	sapplication_status "hrms-api/app/service/applicants/status"
	mapplication_status "hrms-api/app/model/application_status"
	"github.com/gofiber/fiber/v2"
)

func GetApplicationStatus(ctx *fiber.Ctx) error {
	search_query := ctx.Query("q")

	if len(search_query) > 0 {
		application_status, err := sapplication_status.SearchStatus(search_query)

		if err != nil {
			panic(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"application_status": application_status, 
		})
	}

	application_status, err := sapplication_status.FetchStatus()

	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"application_status": application_status, 
	})
}

func UpdateApplicationStatus(ctx *fiber.Ctx) error {
	application_id := ctx.Params("id")
	application_status_model := mapplication_status.ApplicantStatus{}

	err := ctx.BodyParser(&application_status_model)

	if err != nil {
		panic(err.Error())
	}

	err = sapplication_status.UpdateStatus(application_id, application_status_model)

	if err != nil {
		panic(err. Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Status Updated",
	})
}