package models

import "time"

// ECDDCaseManagementFolder represents a case management folder
type ECDDCaseManagementFolder struct {
	ECDDCaseManagementFolderPK string    `json:"ecddCaseManagementFolderPk"`
	FolderName                 string    `json:"folderName"`
	Region                     string    `json:"region"`                       // e.g. MALTA, GIBRALTAR, USA, AUSTRALIA, UK
	CountryID                  *int64    `json:"countryId,omitempty"`         // nullable country identifier
	StateID                    *int64    `json:"stateId,omitempty"`           // nullable state identifier
	LoggedAt                   time.Time `json:"loggedAt"`
	UpdatedBy                  string    `json:"updatedBy"`
}
