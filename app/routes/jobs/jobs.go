package hjobs


import (
	jobs "hrms-api/app/service/jobs"
	"github.com/gofiber/fiber/v2"
)

func SetupJobs(app *fiber.App) {
	app.Get("/jobs", jobs.GetJobPosition)
	app.Patch("/jobs/position/update/:id", jobs.UpdateJobPosition)
	app.Post("/jobs/position/add", jobs.AddJobPosition)
}