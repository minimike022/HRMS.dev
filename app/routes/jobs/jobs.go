package hjobs


import (
	cjobs "hrms-api/app/controller/jobs"
	"github.com/gofiber/fiber/v2"
)

func SetupJobs(app *fiber.App) {
	app.Get("/jobs" ,cjobs.GetJobPosition)
	app.Get("/jobs/search" ,cjobs.SearchJobs)
	app.Patch("/jobs/update/:id", cjobs.UpdateJobPosition)
	app.Post("/jobs/add", cjobs.AddJobPosition)
}

// , jwtvalidate.ValidateAccessToken