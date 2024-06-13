package mapplication

type ProgresStatus struct {
	Application_Status_ID int `json:"app_status_id"`
	Application_Status_Name string `json:"app_status_name"`
	Application_Status_Count int `json:"app_status_count"`
}

type ApplicantsDate struct {
	Position_Name    string `json:"position_name"`
	Application_CreatedAt string `json:"created_At"`
}