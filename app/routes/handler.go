package route

import (
	//service "hrms-api/app/service"
	"github.com/gofiber/fiber/v2"
	hrhandler "hrms-api/app/routes/hr_handler"
	hmhandler "hrms-api/app/routes/hm_handler"
	adminhandler "hrms-api/app/routes/admin_handler"
)

func SetupRoutes(app *fiber.App) {
	hrhandler.SetupHRhandler(app)
	hmhandler.SetupHMhandler(app)
	adminhandler.SetupAdminhandler(app)
}





