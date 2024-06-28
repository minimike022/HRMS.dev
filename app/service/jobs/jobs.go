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
	db_response, _ := db.Query(query, search_query)
	for db_response.Next() {
		db_response.Scan(
			&count,
		)
	}
	defer db_response.Close()

	return count
}

func FetchJobs(offset int, limit int, sort_col string, sort_order string) ([]mjobs.Jobs_List, error) {
	fmt.Println("Hello")
	job_position := mjobs.Jobs_List{}
	job_position_array := make([]mjobs.Jobs_List, 0)

	query := `SELECT JP.position_id , JP.position_name,JP.department_id, DP.department_name,  JP.available_slot, JP.position_status FROM job_position as JP
INNER JOIN department as DP ON JP.department_id = DP.department_id `

	if sort_col != "" && sort_order != "" {
		query += ` ORDER BY ` + sort_col + ` ` + sort_order + ` `
	}

	query += `LIMIT ?, ?;`
	fmt.Println(query)

	stmt, _ := db.Prepare(query)
	db_response, err := stmt.Query(offset, limit)
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

func SearchJobs(search_query string, offset int, limit int, sort_col string, sort_order string) ([]mjobs.Jobs_List) {
	search_result := make([]mjobs.Jobs_List, 0)
	search_model := mjobs.Jobs_List{}

	query := `SELECT JP.position_id , JP.position_name,JP.department_id, DP.department_name,  JP.available_slot, JP.position_status FROM job_position as JP
INNER JOIN department as DP ON JP.department_id = DP.department_id `

	query += `WHERE JP.position_name LIKE ? `

	if sort_col != "" && sort_order != "" {
		query += ` ORDER BY ` + sort_col + ` ` + sort_order + ` `
	}

	query += `LIMIT ?, ?`

	stmt, _ := db.Prepare(query)

	db_response, err := stmt.Query("%"+search_query+"%", offset, limit)

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

	return search_result

}
