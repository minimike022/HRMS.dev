package mjobs

type JobPosition struct {
	Position_ID int `json:"position_id"`
	Position_Name string `json:"position_name"`
	Department_ID int `json:"department_id"`
	Position_Status string `json:"position_status"`
	Available_Slot int	`json:"available_slot"`
}

type Jobs_List struct {
	Position_ID int `json:"position_id"`
	Position_Name string `json:"position_name"`
	Department_Name string `json:"department_name"`
	Available_Slot int `json:"available_slot"`
	Position_Status string `json:"position_status"`
}