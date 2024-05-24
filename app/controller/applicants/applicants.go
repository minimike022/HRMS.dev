package capplicants

import (
	"github.com/gofiber/fiber/v2"
	mapplicants "hrms-api/app/model/applicants"
	sapplicants "hrms-api/app/service/applicants"
	"time"
)

func AddApplicantsData(ctx *fiber.Ctx) error {
	created_at := time.Now().Format("2006-01-02 15:04:05") 
	applicants_data_model := mapplicants.ApplicantsData{}
	err := ctx.BodyParser(&applicants_data_model)

	if err != nil {
		panic(err.Error())
	}

	err = sapplicants.AddApplicants(created_at, applicants_data_model)

	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Applicants Added!",
	})
}


func GetApplicantsData(ctx *fiber.Ctx) error {
	applicants_data_array, err := sapplicants.GetApplicantsData()

	if err != nil {
		panic(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Applicants": applicants_data_array,
	})
}
