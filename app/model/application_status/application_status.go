package application_status

import (
	model_applicants "hrms-api/app/model/applicants"
	status_list "hrms-api/app/model/status_list"
	model_jobs "hrms-api/app/model/jobs"
	model_users "hrms-api/app/model/users"

)

type ApplicantStatus struct {
	Status_ID int
	Job_Position_ID int
	User_Interviewee_ID int `json:"interviewee_id"` 
	Application_Status_ID string `json:"app_status_id"`
	Interview_Date string `json:"int_date"`
	Interview_Time string `json:"int_time"`
}

type Handler_Application_Status interface {
	Application_Status (ApplicantStatus, status_list.ApplicationStatusList, model_applicants.ApplicantsData, model_jobs.JobPosition, model_users.UserAccount) error
}