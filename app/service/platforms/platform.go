package splatformslist

import (
	"fmt"
	Database "hrms-api/app/database"
	mplatformslist "hrms-api/app/model/platforms"
)

var db = Database.Connect()

func FetchPlatformsList() ([]mplatformslist.Platform_List, error) {
	platform_list_model := mplatformslist.Platform_List{}
	platform_list_array := make([]mplatformslist.Platform_List, 0)

	query := `CALL fetch_posting_data()`

	db_response, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	for db_response.Next() {
		db_response.Scan(
			&platform_list_model.Platform_ID,
			&platform_list_model.Platform_Name,
		)
		platform_list_array = append(platform_list_array, platform_list_model)
		fmt.Println(platform_list_array)
	}

	return platform_list_array, nil
	
}