package service

import (
	//Database "hrms-api/app/database"
	Database "hrms-api/app/database"
	DataModels "hrms-api/app/model"

	"github.com/gofiber/fiber/v2"
)

// func GetData(ctx *fiber.Ctx) error {

// }
func ReadApplicantsData(ctx *fiber.Ctx) error  {
	db := Database.Connect()
	appData := new(DataModels.ApplicantsData)
	applicantsData := make([]DataModels.ApplicantsData, 0)

	dbRows, err := db.Query("SELECT * FROM applicants_data")
	if err !=nil {
		panic(err.Error())
	}

	for dbRows.Next() {
		dbRows.Scan(&appData.Applicant_ID,&appData.Position_ID,&appData.First_Name,&appData.Middle_Name,&appData.Last_Name,&appData.Extension_Name, &appData.Birthdate,
			&appData.Age, &appData.Present_Address,&appData.Highest_Education, &appData.Email_Address, &appData.Facebook_Link, &appData.BPO_Exp,
			&appData.Shift_Sched, &appData.Work_Report, &appData.Work_Site_Location, &appData.Platform_ID, &appData.Ref_Full_Name, &appData.Ref_Company,
			&appData.Ref_Position, &appData.Ref_Contact_Num, &appData.Ref_Email, &appData.Applicant_CV, &appData.Applicant_Portfolio_Link)
		applicantsData = append(applicantsData,*appData)
	}

	return ctx.Status(fiber.StatusOK).JSON(applicantsData)
}

func PostApplicantsData(ctx *fiber.Ctx) error {
	var appData DataModels.ApplicantsData
	//db := Database.Connect()
	applicationData := new(DataModels.ApplicantsData)
	err := ctx.BodyParser(applicationData)

	if err != nil {
		panic(err.Error())
	}


	appData = DataModels.ApplicantsData {
		Position_ID: applicationData.Position_ID,
		First_Name: applicationData.First_Name,
		Middle_Name: applicationData.Middle_Name,
		Last_Name: applicationData.Last_Name,
		Extension_Name: applicationData.Extension_Name,
		Birthdate: applicationData.Birthdate,
		Age: applicationData.Age,
		Present_Address: applicationData.Present_Address,
		Highest_Education: applicationData.Highest_Education,
		Email_Address: applicationData.Email_Address,
		Facebook_Link: applicationData.Facebook_Link,
		BPO_Exp: applicationData.BPO_Exp,
		Shift_Sched: applicationData.Shift_Sched,
		Work_Report: applicationData.Work_Report,
		Work_Site_Location: applicationData.Work_Site_Location,
		Platform_ID: applicationData.Platform_ID,
		Ref_Full_Name: applicationData.Ref_Full_Name,
		Ref_Company: applicationData.Ref_Company,
		Ref_Position: applicationData.Ref_Position,
		Ref_Contact_Num: applicationData.Ref_Contact_Num,
		Ref_Email: applicationData.Ref_Email,
		Applicant_CV: applicationData.Applicant_CV,
		Applicant_Portfolio_Link: applicationData.Applicant_Portfolio_Link,
	}
	dbQuery := `INSERT INTO applicants_data (
				position_id, 
				first_name, 
				middle_name, 
				last_name, 
				extension_name, 
				birthdate, 
				age, 
				present_address, 
				highest_education, 
				email_address, 
				facebook_link, 
				bpo_exp, 
				shift_sched, 
				work_report, 
				work_site_location, 
				platform_id, 
				ref_full_name, 
				ref_company, 
				ref_position, 
				ref_contact_num, 
				ref_email, 
				applicant_cv, 
				applicant_portfolio_link
			) 
			VALUES 
			(
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 
				?, ?, ?, ?, ?, ?, ?
			)`
	  

	dbData, err := Database.Connect().Query(dbQuery, 
	appData.Position_ID, 
	appData.First_Name,
	appData.Middle_Name, 
	appData.Last_Name, 
	appData.Extension_Name, 
	appData.Birthdate,
	appData.Age, 
	appData.Present_Address,
	appData.Highest_Education, 
	appData.Email_Address, 
	appData.Facebook_Link, 
	appData.BPO_Exp,
	appData.Shift_Sched, 
	appData.Work_Report,
	appData.Work_Site_Location, 
	appData.Platform_ID, 
	appData.Ref_Full_Name, 
	appData.Ref_Company,
	appData.Ref_Position, 
	appData.Ref_Contact_Num,
	appData.Ref_Email, 
	appData.Applicant_CV, 
	appData.Applicant_Portfolio_Link)

	if err != nil {
		panic(err.Error())
	}
	defer dbData.Close()

	return ctx.Status(fiber.StatusOK).SendString("Added to database!")
}

func GetApplicationStatus(ctx *fiber.Ctx) error {
	appData := new(DataModels.ApplicantsData)
	applicantsData := make([]DataModels.ApplicantsData, 0)
	applicantID := ctx.Params("id")

	dbData, err := Database.Connect().Query("SELECT * FROM applicants_data WHERE applicant_id = ?", applicantID)
	if err != nil {
		panic(err.Error())
	}

	for dbData.Next() {
		dbData.Scan(&appData.Applicant_ID,&appData.Position_ID,&appData.First_Name,&appData.Middle_Name,&appData.Last_Name,&appData.Extension_Name, &appData.Birthdate,
			&appData.Age, &appData.Present_Address,&appData.Highest_Education, &appData.Email_Address, &appData.Facebook_Link, &appData.BPO_Exp,
			&appData.Shift_Sched, &appData.Work_Report, &appData.Work_Site_Location, &appData.Platform_ID, &appData.Ref_Full_Name, &appData.Ref_Company,
			&appData.Ref_Position, &appData.Ref_Contact_Num, &appData.Ref_Email, &appData.Applicant_CV, &appData.Applicant_Portfolio_Link)
		applicantsData = append(applicantsData,*appData)
	}
	dbData.Close()
	return ctx.Status(fiber.StatusOK).JSON(applicantsData)
}

