package sdepartments

import (
	Database "hrms-api/app/database"
	mdepartment "hrms-api/app/model/department"
)

var db = Database.Connect()

func FetchDepartments() ([]mdepartment.Department, error) {
	departments_model := new(mdepartment.Department)
	department_array := make([]mdepartment.Department, 0)

	query := `CALL fetch_departments`
	
	db_response, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
			&departments_model.Department_ID,
			&departments_model.Department_Name,
			&departments_model.Department_Status,
		)

		department_array = append(department_array, *departments_model)
	}

	return department_array, nil
}
