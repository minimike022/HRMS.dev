package route

import (
	hanalysis "hrms-api/app/routes/analysis"
	happlication "hrms-api/app/routes/applicants"
	hjobs "hrms-api/app/routes/jobs"
	husers "hrms-api/app/routes/users"
	hdepartments "hrms-api/app/routes/department"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	hjobs.SetupJobs(app)
	happlication.SetupApplication(app)
	hanalysis.SetupAnalysis(app)
	husers.SetupUserAccounts(app)
	hdepartments.SetupDepartment(app)
}





