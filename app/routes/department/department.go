package hdeparments

import(
	cdepartment "hrms-api/app/controller/jobs/department"
	"github.com/gofiber/fiber/v2"
)

func SetupDepartment(app *fiber.App) {
	app.Get("/departments", cdepartment.GetDepartments)
}