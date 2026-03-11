package models

import "time"

// ECDDUserStatus represents a user's ECDD status
type ECDDUserStatus struct {
	ECDDUserStatusPK            string     `json:"ecdd_user_status_pk"`
	UserID                      int64      `json:"user_id"`
	UserName                    string     `json:"user_name"`
	CountryID                   int64      `json:"country_id"`
	Language                    string     `json:"language"` // ISO 639-1 language code, e.g. "EN"
	StateID                     *int64     `json:"state_id,omitempty"`
	ECDDStatus                  int64      `json:"ecdd_status"` // 1=Not Required, 2=In progress, 3=Complete, 4=Suspended-Manual, 5=Suspended-Auto, 6=Closed, 7=Block Process
	ECDDThreshold               float64    `json:"ecdd_threshold"`
	ECDDReviewTrigger           int64      `json:"ecdd_review_trigger"`
	ECDDSuspensionDueDate       *time.Time `json:"ecdd_suspension_due_date,omitempty"`
	ECDDMultiplier              float64    `json:"ecdd_multiplier"`
	ECDDMultiplierRGFlag        bool       `json:"ecdd_multiplier_rg_flag"`
	UserLtEnggThresholdGBP      float64    `json:"user_lt_engg_threshold_gbp"`
	UserLtDepositThresholdGBP   float64    `json:"user_lt_deposit_threshold_gbp"`
	User12MonthDropThresholdGBP float64    `json:"user_12month_drop_threshold_gbp"`
	InfoSource                  int64      `json:"info_source"`
	SignOffStatus               int64      `json:"sign_off_status"`
	DateLastECDDSignOff         *time.Time `json:"date_last_ecdd_sign_off,omitempty"`
	ECDDRGReviewStatus          int64      `json:"ecdd_rg_review_status"`
	DateLastECDDRGSignOff       *time.Time `json:"date_last_ecdd_rg_sign_off,omitempty"`
	ECDDReportStatus            int64      `json:"ecdd_report_status"`
	ECDDReviewStatus            int64      `json:"ecdd_review_status"`
	ECDDDocumentStatus          int64      `json:"ecdd_document_status"`
	ECDDEscalationStatus        int64      `json:"ecdd_escalation_status"`
	UARStatus                   int64      `json:"uar_status"`
	LoggedAt                    time.Time  `json:"logged_at"`
	UpdatedBy                   string     `json:"updated_by"`
}
