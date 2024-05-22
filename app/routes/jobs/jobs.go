package hjobs


import (
	jobs "hrms-api/app/service/jobs"
	"github.com/gofiber/fiber/v2"
	jwtvalidate "hrms-api/app/service/jwt/validate"
)

func SetupJobs(app *fiber.App) {
	app.Get("/jobs" ,jobs.GetJobPosition)
	app.Patch("/jobs/position/update/:id", jwtvalidate.ValidateAccessToken ,jobs.UpdateJobPosition)
	app.Post("/jobs/add", jobs.AddJobPosition)
}

// , jwtvalidate.ValidateAccessToken