package mapplicants

import (
	"encoding/json"
)

type ApplicantsData struct {
	Applicant_ID   int `json:"applicant_id"`
	Position_ID    int `json:"position_id"`
	Position_Name    string `json:"position_name"`
	First_Name     string `json:"first_name"`
	Middle_Name    string `json:"middle_name"`
	Last_Name      string `json:"last_name"`
	Extension_Name string `json:"extension_name"`
	Birthdate      string `json:"birthdate"`
	Age int `json:"age"` 
	Present_Address string `json:"present_address"`
	Highest_Education string `json:"highest_education"`
	Degree_Course string `json:"degree_course"`
	School_Name string `json:"school_name"`
	Mobile_Number string `json:"mobile_number"`
	Email_Address string `json:"email_address"`
	Facebook_Link string `json:"facebook_link"`
	BPO_Exp string `json:"bpo_exp"`
	Shift_Sched string `json:"shift_sched"`
	Work_Report string `json:"work_report"`
	Work_Site_Location string `json:"work_site_location"`
	Platforms json.RawMessage `json:"platforms"`
	Ref_Full_Name string `json:"ref_full_name"`
	Ref_Company string `json:"ref_company"`
	Ref_Position string `json:"ref_position"`
	Ref_Contact_Num string `json:"ref_contact_num"`
	Ref_Email  string `json:"ref_email"`
	Applicant_CV string `json:"applicant_cv"`
	Applicant_Portfolio_Link string `json:"applicant_portfolio"`
	Application_CreatedAt string
}

type NewApplicants struct {
	Applicant_ID int `json:"app_id"`
	First_Name string `json:"first_name"`
	Last_Name string  `json:"last_name"`
	Extension_Name string `json:"ext_name"`
	Job_Position string `json:"job_position"`
	Application_Date string  `json:"app_date"`
}




