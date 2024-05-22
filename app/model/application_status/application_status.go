package application_status

type ApplicantStatus struct {
	Status_ID int
	Job_Position_ID int
	User_Interviewee_ID int `json:"interviewee_id"` 
	Application_Status_ID int `json:"app_status_id"`
	Interview_Date string `json:"int_date"`
	Interview_Time string `json:"int_time"`
}

type Application_Status struct {
	Status_ID int `json:"application_status_id"`
	First_Name string `json:"first_name"`
	Middle_Name string `json:"middle_name"`
	Last_Name string `json:"last_name"`
	Extension_Name string `json:"extension_name"`
	Position_Name string `json:"position_name"`
	Interviewee_Name string `json:"interviewee_name"`
	Application_Status string `json:"application_status"`
	Interview_Date string `json:"interview_date"`
	Interview_Time string `json:"interview_time"`
}