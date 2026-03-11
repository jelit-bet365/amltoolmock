package models

import "time"

// ECDDThresholdConfig represents a threshold rule configuration
type ECDDThresholdConfig struct {
	ECDDThresholdConfigPK     string    `json:"ecdd_threshold_config_pk"`
	Title                     string    `json:"title"`
	IsActive                  bool      `json:"is_active"`
	CountryID                 int64     `json:"country_id"`
	StateID                   *int64    `json:"state_id,omitempty"`
	Type                      int64     `json:"type"` // 1=Deposit, 2=Net Deposit, 3=Stakes
	Value                     float64   `json:"value"`
	CurrencyID                int64     `json:"currency_id"`
	Period                    int64     `json:"period"` // 1=24hrs, 2=28days, 3=84days, 4=91days, 5=182days, 6=365days
	UseMultipliers            bool      `json:"use_multipliers"`
	UseRGFlag                 bool      `json:"use_rg_flag"`
	ApplyAllStatuses          bool      `json:"apply_all_statuses"`
	Backfill                  bool      `json:"backfill"`
	Hierarchy                 int64     `json:"hierarchy"`
	ECDDStatus                int64     `json:"ecdd_status"`
	ECDDReviewStatus          int64     `json:"ecdd_review_status"`
	ECDDReportStatus          int64     `json:"ecdd_report_status"`
	SignOffStatus             int64     `json:"sign_off_status"`
	CustomerRiskLevel         int64     `json:"customer_risk_level"` // 1=Low, 2=Medium, 3=Medium-High, 4=High
	NDL28DayGBP               float64   `json:"ndl_28_day_gbp"`
	NDLMonthlyGBP             float64   `json:"ndl_monthly_gbp"`
	CaseManagementFolderPK    *string   `json:"case_management_folder_pk,omitempty"`
	LoggedAt                  time.Time `json:"logged_at"`
	UpdatedBy                 string    `json:"updated_by"`
}