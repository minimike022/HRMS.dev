package platforms

import (
	Database "hrms-api/app/database"
	model_platform "hrms-api/app/model/platform"

	"github.com/gofiber/fiber/v2"
)

var db = Database.Connect()

func GetPlatformData(ctx *fiber.Ctx) error {
	posting_platform_model := new(model_platform.PostingPlatform)
	posting_platform_array := make([]model_platform.PostingPlatform,0)

	db_query := `CALL fetch_posting_data()`
	db_response, err := db.Query(db_query)
	if err != nil {
		panic(err.Error())
	}
	defer db_response.Close()
	for db_response.Next() {
		db_response.Scan(
			&posting_platform_model.Platform_ID,
			&posting_platform_model.Platform_Name,
			&posting_platform_model.Platform_Count)
		posting_platform_array = append(posting_platform_array, *posting_platform_model)
	}

	return ctx.Status(fiber.StatusOK).JSON(posting_platform_array)
}