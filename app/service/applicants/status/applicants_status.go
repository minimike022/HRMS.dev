package sapplication_status

import (
	"fmt"
	Database "hrms-api/app/database"
	mapplication_status "hrms-api/app/model/application_status"
)

var db = Database.Connect()

func CountStatus() (count int) {
	var status_count int
	query := `CALL count_application_status`
	db_response, _ := db.Query(query)

	for db_response.Next() {
		db_response.Scan(
			&status_count,
		)
	}
	defer db_response.Close()

	return status_count
}
func SearchCount(search_query string) (search_count int) {
	var count int
	query := `CALL count_search_applicants(?)`
	db_response, _ := db.Query(query, search_query)
	for db_response.Next() {
		db_response.Scan(
			&count,
		)
	}
	defer db_response.Close()

	return count
}

func FetchStatus(page_limit int, offset int, sort_col string, sort_order string) ([]mapplication_status.Application_Status, error) {
	application_status_array := make([]mapplication_status.Application_Status, 0)
	query := `SELECT APS.status_id,
AD.applicant_id,
AD.first_name, 
AD.middle_name, 
AD.last_name, 
AD.extension_name,
JP.position_name, 
ASL.application_status_name,
UA.account_id,
UA.user_name, 
APS.interview_date, APS.interview_time
FROM application_status as APS
LEFT JOIN applicants_data as AD ON 
APS.applicant_id = AD.applicant_id
LEFT JOIN user_accounts as UA ON APS.user_interviewee_id = UA.account_id
LEFT JOIN job_position as JP ON 
APS.position_id = JP.position_id
LEFT JOIN application_status_list as ASL ON APS.application_status_id = ASL.application_status_id
` + ` `

	if sort_col != "" && sort_order != "" {
		query += `ORDER BY` + ` ` + sort_col + ` ` + sort_order + ` `
	}

	query += `LIMIT ?, ?;`
	fmt.Println()
	stmt, _ := db.Prepare(query)

	db_response, err := stmt.Query(offset, page_limit)

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
			&application_status_model.Interviewee_ID,
			&application_status_model.Interviewee_Name,
			&application_status_model.Interview_Date,
			&application_status_model.Interview_Time,
		)
		application_status_array = append(application_status_array, application_status_model)
		fmt.Println(application_status_array)
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

func SearchStatus(search_query string, page_limit int, offset int, sort_col string, sort_order string) ([]mapplication_status.Application_Status, error) {
	application_status_array := make([]mapplication_status.Application_Status, 0)
	query := `SELECT APS.status_id,
APS.applicant_id,
AD.first_name, 
AD.middle_name, 
AD.last_name, 
AD.extension_name,
JP.position_name, 
ASL.application_status_name,
UA.account_id,
UA.user_name, 
APS.interview_date, APS.interview_time
FROM application_status as APS
LEFT JOIN applicants_data as AD ON 
APS.applicant_id = AD.applicant_id
LEFT JOIN user_accounts as UA ON APS.user_interviewee_id = UA.account_id
LEFT JOIN job_position as JP ON 
APS.position_id = JP.position_id
LEFT JOIN application_status_list as ASL ON APS.application_status_id = ASL.application_status_id
WHERE AD.first_name LIKE ? OR AD.last_name LIKE ? OR AD.extension_name LIKE ? OR JP.position_name LIKE ?
` + ` `

	query += `ORDER BY` + ` ` + sort_col + ` ` + sort_order + ` `

	search_query = "%" + search_query + "%"

	query += `LIMIT ?, ?;`
	stmt, _ := db.Prepare(query)
	db_response, err := stmt.Query(search_query, search_query, search_query, search_query, offset, page_limit)

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
			&application_status_model.Interviewee_ID,
			&application_status_model.Interviewee_Name,
			&application_status_model.Interview_Date,
			&application_status_model.Interview_Time,
		)
		application_status_array = append(application_status_array, application_status_model)
		fmt.Println(application_status_array)
	}

	defer db_response.Close()

	return application_status_array, nil
}

func FetchList() ([]mapplication_status.Status_List, error) {
	application_status_array := make([]mapplication_status.Status_List, 0)
	application_status_model := mapplication_status.Status_List{}

	query := `CALL fetch_status_list`

	db_response, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
			&application_status_model.Application_Status_ID,
			&application_status_model.Application_Status_Name,
		)
		application_status_array = append(application_status_array, application_status_model)
	}

	defer db_response.Close()

	return application_status_array, nil

}
