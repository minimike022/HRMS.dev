package applicants

import (
	Database "hrms-api/app/database"
	model_applicants "hrms-api/app/model/applicants"

	//jwt "hrms-api/app/service/jwt"
	"time"

	"github.com/gofiber/fiber/v2"
)
var db = Database.Connect()

func PostApplicantsData(ctx *fiber.Ctx) error {
	created_at := time.Now().Format("2006-01-02 15:04:05") 
	applicants_data_model := model_applicants.ApplicantsData{}
	err := ctx.BodyParser(&applicants_data_model)

	if err != nil {
		panic(err.Error())
	}
	dbQuery := `CALL add_applicants(?, ?, ?, ?, ?, ?, ?, ?, ?, 
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	dbData, err := db.Query(dbQuery,
		applicants_data_model.Position_ID,
		applicants_data_model.First_Name,
		applicants_data_model.Middle_Name,
		applicants_data_model.Last_Name,
		applicants_data_model.Extension_Name,
		applicants_data_model.Birthdate,
		applicants_data_model.Age,
		applicants_data_model.Present_Address,
		applicants_data_model.Highest_Education,
		applicants_data_model.Email_Address,
		applicants_data_model.Facebook_Link,
		applicants_data_model.BPO_Exp,
		applicants_data_model.Shift_Sched,
		applicants_data_model.Work_Report,
		applicants_data_model.Work_Site_Location,
		applicants_data_model.Platforms,
		applicants_data_model.Ref_Full_Name,
		applicants_data_model.Ref_Company,
		applicants_data_model.Ref_Position,
		applicants_data_model.Ref_Contact_Num,
		applicants_data_model.Ref_Email,
		applicants_data_model.Applicant_CV,
		applicants_data_model.Applicant_Portfolio_Link,
		created_at)
	if err != nil {
		panic(err.Error())
	}

	defer dbData.Close()

	return ctx.Status(fiber.StatusOK).SendString("Added to database!")
}

func GetApplicantsData(ctx *fiber.Ctx) error {

	applicants_data_model := new(model_applicants.ApplicantsData)
	applicants_data_array := make([]model_applicants.ApplicantsData, 0)

	db_query := "CALL fetch_applicants_data"

	db_response, err := db.Query(db_query)
	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(	
			&applicants_data_model.Applicant_ID,
			&applicants_data_model.First_Name,
			&applicants_data_model.Middle_Name,
			&applicants_data_model.Last_Name,
			&applicants_data_model.Extension_Name,
			&applicants_data_model.Birthdate,
			&applicants_data_model.Age,
			&applicants_data_model.Present_Address,
			&applicants_data_model.Mobile_Number,
			&applicants_data_model.Email_Address,
			&applicants_data_model.Facebook_Link,
			&applicants_data_model.Position_Name,
			&applicants_data_model.BPO_Exp,
			&applicants_data_model.Shift_Sched,
			&applicants_data_model.Work_Report,
			&applicants_data_model.Work_Site_Location,
			&applicants_data_model.Highest_Education,
			&applicants_data_model.Degree_Course,
			&applicants_data_model.School_Name,
			&applicants_data_model.Platforms,
			&applicants_data_model.Ref_Full_Name,
			&applicants_data_model.Ref_Company,
			&applicants_data_model.Ref_Position,
			&applicants_data_model.Ref_Contact_Num,
			&applicants_data_model.Ref_Email,
			&applicants_data_model.Applicant_CV,
			&applicants_data_model.Applicant_Portfolio_Link,
			&applicants_data_model.Application_CreatedAt,
		)
		applicants_data_array = append(applicants_data_array, *applicants_data_model)
	}
	
	defer db_response.Close()


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"applicants_data": applicants_data_array,
	})
}