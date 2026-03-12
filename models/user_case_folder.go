package models

import "time"

// ECDDUserCaseManagementFolder represents the many-to-many relationship between users and case folders
type ECDDUserCaseManagementFolder struct {
	ECDDUserCaseManagementFolderPK string    `json:"ecdd_user_case_management_folder_pk"`
	FolderPK                       string    `json:"case_management_folder_pk"` // FK to ECDDCaseManagementFolder
	UserStatusPK                   string    `json:"user_status_pk"`  // FK to ECDDUserStatus
	LoggedAt                       time.Time `json:"logged_at"`
	UpdatedBy                      string    `json:"updated_by"`
}