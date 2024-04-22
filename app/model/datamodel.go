package DataModels

type ApplicantsData struct {
	Applicant_ID   int
	Position_ID    int
	First_Name     string
	Middle_Name    string
	Last_Name      string
	Extension_Name string
	Birthdate      string
	Age int
	Present_Address string
	Highest_Education string
	Email_Address string
	Facebook_Link string
	BPO_Exp string
	Shift_Sched string
	Work_Report string
	Work_Site_Location string
	Platform_ID int
	Ref_Full_Name string
	Ref_Company string
	Ref_Position string
	Ref_Contact_Num string
	Ref_Email  string
	Applicant_CV string
	Applicant_Portfolio_Link string

}

type ApplicantStatus struct {
	Application_ID int
	Applicant_ID int
	Application_Status_ID int
	User_Interviewee_ID int
	Interview_Date string
	Interview_Time string
	Application_Status string
}

type DepartmentList struct {
	Department_ID int
	Department_Name string
	Department_Status string
}

type JobPosition struct {
	Position_ID int
	Position_Name string
	Department_ID int
	Position_Status string
}

type PostingPlatform struct {
	Platform_ID int
	Platform_Name string
	Platform_Status string
}

type ApplicationStatusList struct {
	Application_Status_ID int
	Application_Status_Name string
}

type UserAccount struct {
	Account_ID int
	Username string
	Password string
	User_Role string
	User_Name string
	User_Position string
	Department_ID int
	Account_Satus string
}
