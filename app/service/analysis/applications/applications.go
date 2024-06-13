package sapplication

import (
	"fmt"
	Database "hrms-api/app/database"
	mapplication "hrms-api/app/model/analysis/applications"
)
var db = Database.Connect()

func FetchProgress() ([]mapplication.ProgresStatus, error) {
	progress_status_model := mapplication.ProgresStatus{}
	progress_status_array := make([]mapplication.ProgresStatus,0)

	query := `CALL analysis_progress`

	db_response, err := db.Query(query)

	if err != nil {
		panic(err.Error)
	}

	for db_response.Next() {
		db_response.Scan(
			&progress_status_model.Application_Status_ID,
			&progress_status_model.Application_Status_Name,
			&progress_status_model.Application_Status_Count,
		) 
		progress_status_array = append(progress_status_array, progress_status_model)

		fmt.Println(progress_status_array)
	}


	return progress_status_array, nil
}


func FetchDate() ([]mapplication.ApplicantsDate, error) {
	applicants_data_model := new(mapplication.ApplicantsDate)
	applicants_data_array := make([]mapplication.ApplicantsDate, 0)

	db_query := `CALL fetch_all_applicants_data`

	db_response, err := db.Query(db_query)
	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(	
			&applicants_data_model.Position_Name,
			&applicants_data_model.Application_CreatedAt,
		)
		applicants_data_array = append(applicants_data_array, *applicants_data_model)
	}
	
	defer db_response.Close()


	return applicants_data_array, nil
}