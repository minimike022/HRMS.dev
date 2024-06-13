package route

import (
	hplatforms "hrms-api/app/routes/analysis/platforms"
	happlicants "hrms-api/app/routes/applicants"
	hjobs "hrms-api/app/routes/jobs"
	husers "hrms-api/app/routes/users"
	hdepartments "hrms-api/app/routes/department"
	hstatus "hrms-api/app/routes/applicants/status"
	happlication "hrms-api/app/routes/analysis/applications"
	hplatformslist "hrms-api/app/routes/platform"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	hjobs.SetupJobs(app)
	happlicants.SetupApplication(app)
	hplatforms.SetupPlatformAnalysis(app)
	husers.SetupUserAccounts(app)
	hdepartments.SetupDepartment(app)
	hstatus.SetupApplicationStatus(app)
	hplatformslist.SetupPlatformList(app)
	happlication.SetupApplicationProgress(app)
}





