package route

import (
	hanalysis "hrms-api/app/routes/analysis"
	happlication "hrms-api/app/routes/applicants"
	hjobs "hrms-api/app/routes/jobs"
	husers "hrms-api/app/routes/users"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	hjobs.SetupJobs(app)
	happlication.SetupApplication(app)
	hanalysis.SetupAnalysis(app)
	husers.SetupUserAccounts(app)
}





