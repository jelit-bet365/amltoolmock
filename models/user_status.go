package models

import "time"

// ECDDUserStatus represents a user's ECDD status
type ECDDUserStatus struct {
	ECDDUserStatusPK            string     `json:"ecddUserStatusPk"`
	UserID                      int64      `json:"userId"`
	UserName                    string     `json:"userName"`
	CountryID                   int64      `json:"countryId"`
	Language                    int64      `json:"languageId"` // Language ID from gopkgs/enums/languages (1=English, 5=German, etc.)
	StateID                     *int64     `json:"stateId,omitempty"`
	ECDDStatus                  int64      `json:"ecddStatus"` // 1=Not Required, 2=In progress, 3=Complete, 4=Suspended-Manual, 5=Suspended-Auto, 6=Closed, 7=Block Process
	ECDDThreshold               float64    `json:"ecddThreshold"`
	ECDDReviewTrigger           int64      `json:"ecddReviewTrigger"`
	ECDDSuspensionDueDate       *time.Time `json:"ecddSuspensionDueDate,omitempty"`
	ECDDMultiplier              float64    `json:"ecddMultiplier"`
	ECDDMultiplierRGFlag        bool       `json:"ecddMultiplierRgFlag"`
	UserLtNetDepositThresholdGBP   float64    `json:"userLtNetDepositThresholdGbp"`
	UserLtDepositThresholdGBP     float64    `json:"userLtDepositThresholdGbp"`
	User12MonthNetDepositThresholdGBP float64 `json:"user12monthNetDepositThresholdGbp"`
	InfoSource                  int64      `json:"infoSource"`
	SignOffStatus               int64      `json:"signOffStatus"`
	DateLastECDDSignOff         *time.Time `json:"dateLastEcddSignOff,omitempty"`
	ECDDRGReviewStatus          int64      `json:"ecddRgReviewStatus"`
	DateLastECDDRGSignOff       *time.Time `json:"dateLastEcddRgSignOff,omitempty"`
	ECDDReportStatus            int64      `json:"ecddReportStatus"`
	ECDDReviewStatus            int64      `json:"ecddReviewStatus"`
	ECDDDocumentStatus          int64      `json:"ecddDocumentStatus"`
	ECDDEscalationStatus        int64      `json:"ecddEscalationStatus"`
	UARStatus                   int64      `json:"uarStatus"`
	LoggedAt                    time.Time  `json:"loggedAt"`
	UpdatedBy                   string     `json:"updatedBy"`
}
