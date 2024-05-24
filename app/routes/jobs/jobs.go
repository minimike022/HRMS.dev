package hjobs


import (
	cjobs "hrms-api/app/controller/jobs"
	"github.com/gofiber/fiber/v2"
	jwtvalidate "hrms-api/app/service/jwt/validate"
)

func SetupJobs(app *fiber.App) {
	app.Get("/jobs" ,cjobs.GetJobPosition)
	app.Patch("/jobs/position/update/:id", jwtvalidate.ValidateAccessToken ,cjobs.UpdateJobPosition)
	app.Post("/jobs/add", cjobs.AddJobPosition)
}

// , jwtvalidate.ValidateAccessToken