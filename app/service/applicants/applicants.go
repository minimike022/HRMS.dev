package sapplicants

import (
	Database "hrms-api/app/database"
	mapplicants "hrms-api/app/model/applicants"
	//jwt "hrms-api/app/service/jwt"
)
var db = Database.Connect()

func AddApplicants(createdAt string, applicants_data_model mapplicants.ApplicantsData) error {
	
	dbQuery := `CALL add_applicants(?, ?, ?, ?, ?, ?, ?, ?, ?, 
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,?)`

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
		applicants_data_model.Degree_Course,
		applicants_data_model.School_Name,
		applicants_data_model.Mobile_Number,
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
		createdAt)

	if err != nil {
		panic(err.Error())
	}

	defer dbData.Close()

	return nil
}


func GetApplicantsData(applicants_id string) ([]mapplicants.ApplicantsData, error) {
	applicants_data_model := new(mapplicants.ApplicantsData)
	applicants_data_array := make([]mapplicants.ApplicantsData, 0)

	db_query := `CALL fetch_applicants_data(?)`

	db_response, err := db.Query(db_query, applicants_id)
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


	return applicants_data_array, nil
}

func FetchNewApplicants() ([]mapplicants.NewApplicants, error) {
	new_app := mapplicants.NewApplicants{}
	new_app_array := make([]mapplicants.NewApplicants, 0)
	query := "CALL fetch_new_applicants"
	db_response, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
			&new_app.Applicant_ID,
			&new_app.First_Name,
			&new_app.Last_Name,
			&new_app.Extension_Name,
			&new_app.Job_Position,
			&new_app.Application_Date,
		)
		new_app_array = append(new_app_array, new_app)
	}
	defer db_response.Close()

	return new_app_array, nil
}

