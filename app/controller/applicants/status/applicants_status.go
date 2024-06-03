package capplication_status

import (
	"fmt"
	mapplication_status "hrms-api/app/model/application_status"
	sapplication_status "hrms-api/app/service/applicants/status"

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
	fmt.Println(application_id)
	application_status_model := mapplication_status.ApplicantStatus{}

	err := ctx.BodyParser(&application_status_model)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(application_status_model)
	err = sapplication_status.UpdateStatus(application_id, application_status_model)

	if err != nil {
		panic(err. Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Status Updated",
	})
}

func FetchStatusList(ctx *fiber.Ctx) error {
	status_list, err := sapplication_status.FetchList()

	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status_list": status_list,
	})

} 