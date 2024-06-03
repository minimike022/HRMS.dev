package musers

type UserAccount struct {
	Account_ID int
	Username string 
	Password string 
	User_Role string
	User_Name string
	User_Position string
	Department_ID int
	Account_Status string
	CreatedAt string
}

type Users struct {
	Account_ID int `json:"account_id"`
	User_Role string `json:"user_role"`
	User_Name string `json:"user_name"`
	Department_Name string `json:"department_name"`
}