package sjobs

import (
	"fmt"
	Database "hrms-api/app/database"
	mjobs "hrms-api/app/model/jobs"
)

var db = Database.Connect()

func CountJobs() (jobs_count int) {
	var count int
	query := `CALL count_jobs`
	db_response, _ := db.Query(query)

	for db_response.Next() {
		db_response.Scan(
			&count,
		)
	}
	defer db_response.Close()
	
	return count
}

func SearchCount(search_query string) (search_count int) {
	var count int
	query := `CALL count_search_jobs(?)`
	db_response, _ := db.Query(query,search_query)
	for db_response.Next() {
		db_response.Scan(
			&count,
		)
	}
	defer db_response.Close()

	return count
}

func FetchJobs(page int, offset int) ([]mjobs.Jobs_List, error) {
	job_position := mjobs.Jobs_List{}
	job_position_array := make([]mjobs.Jobs_List, 0)

	db_query := "CALL fetch_job_positions(?, ?)"

	db_response, err := db.Query(db_query, page, offset)

	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
			&job_position.Position_ID,
			&job_position.Position_Name,
			&job_position.Department_ID,
			&job_position.Department_Name,
			&job_position.Available_Slot,
			&job_position.Position_Status,
		)
		job_position_array = append(job_position_array, job_position)
	}

	defer db_response.Close()

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

	fmt.Println(job_position)

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

func SearchJobs (search_query string, page int, offset int) ([]mjobs.Jobs_List, error) {
	search_result := make([]mjobs.Jobs_List, 0)
	search_model := mjobs.Jobs_List{}

	query := `CALL search_jobs(?,?,?)`

	db_response, err := db.Query(query, search_query, page, offset)

	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
			&search_model.Position_ID,
			&search_model.Position_Name,
			&search_model.Department_ID,
			&search_model.Department_Name,
			&search_model.Available_Slot,
			&search_model.Position_Status,
		)
		search_result = append(search_result, search_model)
	}

	defer db_response.Close()

	return search_result, nil

}