package models

import "time"

// ECDDCaseManagementFolder represents a case management folder
type ECDDCaseManagementFolder struct {
	ECDDCaseManagementFolderPK string    `json:"ecdd_case_management_folder_pk"`
	FolderName                 string    `json:"folder_name"`
	Region                     string    `json:"region"`                       // e.g. MALTA, GIBRALTAR, USA, AUSTRALIA, UK
	CountryID                  *int64    `json:"country_id,omitempty"`         // nullable country identifier
	StateID                    *int64    `json:"state_id,omitempty"`           // nullable state identifier
	LoggedAt                   time.Time `json:"logged_at"`
	UpdatedBy                  string    `json:"updated_by"`
}