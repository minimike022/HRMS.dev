package model_applicants

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
	Applicant_Status_ID int
	Application_CreatedAt string
}

type ApplicantStatus struct {
	Applicant_ID int
	Applicant_First_Name string
	Applicant_Middle_Name string
	Applicant_Last_Name string
	Application_Status string
	Job_Position_Name string
	Department_Name string
	User_Interviewee_Name string	
	Application_Status_ID int
	//Interview_Date string
	//Interview_Time string
	
}
