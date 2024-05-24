package saccounts

import (
	Database "hrms-api/app/database"
	musers "hrms-api/app/model/users"
)

var db = Database.Connect()


func FetchUser() ([]musers.UserAccount, error) {
	user_accounts_data := musers.UserAccount{}
	user_accounts_array := make([]musers.UserAccount,0)
	//Calling Procedured Query
	db_query := "CALL fetch_user_accounts"

	db_response, err := db.Query(db_query)
	if err != nil {

		panic(err.Error())
	}

	for db_response.Next() {
		
		db_response.Scan(
			&user_accounts_data.Account_ID,
			&user_accounts_data.Username,
			&user_accounts_data.Password,
			&user_accounts_data.User_Role,
			&user_accounts_data.User_Name,
			&user_accounts_data.User_Position,
			&user_accounts_data.Department_ID,
			&user_accounts_data.Account_Status,
			&user_accounts_data.CreatedAt,
		)
		user_accounts_array = append(user_accounts_array, user_accounts_data)
	}
	defer db_response.Close()
	
	return user_accounts_array, nil
}

func AddUser(user_accounts musers.UserAccount, user_hashed_password string) error {
	db_query := `CALL add_user_accounts(?,?,?,?,?,?,?)`

	db_response, err := db.Query(db_query, 
		user_accounts.Username,
		user_hashed_password,
		user_accounts.User_Role,
		user_accounts.User_Name,
		user_accounts.User_Position,
		user_accounts.Department_ID,
	)

	if err != nil {
		panic(err.Error())
	}
	
	defer db_response.Close()

	return nil
}

func UpdateUser(account_id string, user_accounts musers.UserAccount) error {
	query := `CALL update_account_status(?, ?)`

	_, err := db.Query(query, account_id, user_accounts.Account_Status)

	if err != nil {
		panic(err.Error())
	}

	return nil
}