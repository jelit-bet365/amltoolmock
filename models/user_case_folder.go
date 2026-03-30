package models

import "time"

// ECDDUserCaseManagementFolder represents the many-to-many relationship between users and case folders
type ECDDUserCaseManagementFolder struct {
	ECDDUserCaseManagementFolderPK string    `json:"ecddUserCaseManagementFolderPk"`
	FolderPK                       string    `json:"ecddCaseManagementFolderPk"` // FK to ECDDCaseManagementFolder
	UserStatusPK                   string    `json:"ecddUserStatusPk"`  // FK to ECDDUserStatus
	LoggedAt                       time.Time `json:"loggedAt"`
	UpdatedBy                      string    `json:"updatedBy"`
}
