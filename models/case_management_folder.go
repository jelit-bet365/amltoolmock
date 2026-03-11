package models

import "time"

// ECDDCaseManagementFolder represents a case management folder
type ECDDCaseManagementFolder struct {
	ECDDCaseManagementFolderPK string    `json:"ecdd_case_management_folder_pk"`
	FolderName                 string    `json:"folder_name"`
	LoggedAt                   time.Time `json:"logged_at"`
	UpdatedBy                  string    `json:"updated_by"`
}