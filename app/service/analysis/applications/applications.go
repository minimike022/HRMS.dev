package sprogress

import (
	"fmt"
	Database "hrms-api/app/database"
	manalysis "hrms-api/app/model/analysis"
)
var db = Database.Connect()

func FetchProgress() ([]manalysis.ProgresStatus, error) {
	progress_status_model := manalysis.ProgresStatus{}
	progress_status_array := make([]manalysis.ProgresStatus,0)

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
