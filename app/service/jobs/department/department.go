package departments

import (
	Database "hrms-api/app/database"
	model_department "hrms-api/app/model/department"

	"github.com/gofiber/fiber/v2"
)

var db = Database.Connect()

func GetDepartments(ctx *fiber.Ctx) error {
	departments_model := new(model_department.Department)
	department_array := make([]model_department.Department, 0)

	query := `CALL fetch_departments`
	
	db_response, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
			&departments_model.Department_ID,
			&departments_model.Department_Name,
			&departments_model.Department_Status,
		)

		department_array = append(department_array, *departments_model)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"departments": department_array,
	})
}
