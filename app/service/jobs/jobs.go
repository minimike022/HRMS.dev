package sjobs

import (
	Database "hrms-api/app/database"
	mjobs "hrms-api/app/model/jobs"
)

var db = Database.Connect()


func FetchJobs() ([]mjobs.Jobs_List, error) {
	job_position := mjobs.Jobs_List{}
	job_position_array := make([]mjobs.Jobs_List, 0)

	db_query := "CALL fetch_job_positionS"

	db_response, err := db.Query(db_query)

	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
			&job_position.Position_ID,
			&job_position.Position_Name,
			&job_position.Department_Name,
			&job_position.Available_Slot,
			&job_position.Position_Status,
		)
		job_position_array = append(job_position_array, job_position)
	}

	return job_position_array, err
}

func AddJobs(job_position mjobs.JobPosition) error {
	

	db_query := `CALL add_job_slot(?,?,?)`

	db_response, err := db.Query(db_query, 
	job_position.Position_Name,
	job_position.Department_ID,
	job_position.Available_Slot,
	)

	if err != nil {
		panic(err.Error())
	}

	defer db_response.Close()

	return nil 
}

func UpdateJobs(job_position_id string, job_position mjobs.JobPosition) error {

	db_query := `CALL update_job_position(?,?,?,?,?)`

	db_response, err := db.Query(db_query,
	job_position_id, 
	job_position.Position_Name, 
	job_position.Department_ID, 
	job_position.Position_Status,
	job_position.Available_Slot,

	)

	if err != nil {
		panic(err.Error())
	}
	defer db_response.Close()

	return nil
}