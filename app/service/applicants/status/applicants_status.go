package sapplication_status

import (
	Database "hrms-api/app/database"
	mapplication_status "hrms-api/app/model/application_status"
)

var db = Database.Connect()

func FetchStatus() ([]mapplication_status.Application_Status, error) {
	application_status_array := make([]mapplication_status.Application_Status, 0)
	query := `CALL fetch_application_status`

	db_response, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		application_status_model := mapplication_status.Application_Status{}
		db_response.Scan(
			&application_status_model.Status_ID,
			&application_status_model.Applicant_ID,
			&application_status_model.First_Name,
			&application_status_model.Middle_Name,
			&application_status_model.Last_Name,
			&application_status_model.Extension_Name,
			&application_status_model.Position_Name,
			&application_status_model.Application_Status,
			&application_status_model.Interviewee_Name,
			&application_status_model.Interview_Date,
			&application_status_model.Interview_Time,
		)
		application_status_array = append(application_status_array, application_status_model)
	}

	defer db_response.Close()


	return application_status_array, nil
}

func UpdateStatus(application_id string, application_status mapplication_status.ApplicantStatus) error {
	
	db_query := "CALL update_application_status(?,?,?,?,?)"

	_, err := db.Query(db_query, application_id, application_status.Application_Status_ID, application_status.User_Interviewee_ID, application_status.Interview_Date, application_status.Interview_Time)

	if err != nil {
		panic(err.Error())
	}

	return nil
}

func SearchStatus(search_query string) ([]mapplication_status.Application_Status, error) {
	application_status_array := make([]mapplication_status.Application_Status, 0)
	query := `CALL search_app_status(?)`

	db_response, err := db.Query(query, search_query)

	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		application_status_model := mapplication_status.Application_Status{}
		db_response.Scan(
			&application_status_model.Status_ID,
			&application_status_model.Applicant_ID,
			&application_status_model.First_Name,
			&application_status_model.Middle_Name,
			&application_status_model.Last_Name,
			&application_status_model.Extension_Name,
			&application_status_model.Position_Name,
			&application_status_model.Application_Status,
			&application_status_model.Interviewee_Name,
			&application_status_model.Interview_Date,
			&application_status_model.Interview_Time,
		)
		application_status_array = append(application_status_array, application_status_model)
	}

	defer db_response.Close()


	return application_status_array, nil
}