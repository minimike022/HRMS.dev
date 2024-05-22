package hdeparments

import(
	department "hrms-api/app/service/jobs/department"
	"github.com/gofiber/fiber/v2"
)

func SetupDepartment(app *fiber.App) {
	app.Get("/departments", department.GetDepartments)
}