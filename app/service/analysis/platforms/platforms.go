package splatforms

import (
	Database "hrms-api/app/database"
	mplatform "hrms-api/app/model/analysis/platforms"
)

var db = Database.Connect()

func GetPlatform() ([]mplatform.PostingPlatform,error) {
	posting_platform_model := new(mplatform.PostingPlatform)
	posting_platform_array := make([]mplatform.PostingPlatform,0)

	db_query := `CALL fetch_platform_data()`
	
	db_response, err := db.Query(db_query)

	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
			&posting_platform_model.Platforms)
		posting_platform_array = append(posting_platform_array, *posting_platform_model)
	}
	defer db_response.Close()
	return posting_platform_array, nil
}