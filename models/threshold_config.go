package models

import "time"

// ECDDThresholdConfig represents a threshold rule configuration
type ECDDThresholdConfig struct {
	ECDDThresholdConfigPK     string    `json:"ecddThresholdConfigPk"`
	Title                     string    `json:"title"`
	IsActive                  bool      `json:"isActive"`
	CountryID                 int64     `json:"countryId"`
	StateID                   *int64    `json:"stateId,omitempty"`
	Type                      int64     `json:"type"` // 1=Deposit, 2=Net Deposit, 3=Stakes
	Reinvest                  bool      `json:"reinvest"`
	Value                     float64   `json:"value"`
	CurrencyID                int64     `json:"currencyId"`
	Period                    int64     `json:"period"` // 1=24hrs, 2=28days, 3=84days, 4=91days, 5=182days, 6=365days
	UseMultipliers            bool      `json:"useMultipliers"`
	UseRGFlag                 bool      `json:"useRgFlag"`
	ApplyAllStatuses          bool      `json:"applyAllStatuses"`
	Backfill                  bool      `json:"backfill"`
	Hierarchy                 int64     `json:"hierarchy"`
	ECDDStatus                int64     `json:"ecddStatus"`
	ECDDReviewStatus          int64     `json:"ecddReviewStatus"`
	ECDDReportStatus          int64     `json:"ecddReportStatus"`
	SignOffStatus             int64     `json:"signOffStatus"`
	CustomerRiskLevel         int64     `json:"customerRiskLevel"` // 1=Low, 2=Medium, 3=Medium-High, 4=High
	NDL28DayGBP               float64   `json:"ndl28DayGbp"`
	NDLMonthlyGBP             float64   `json:"ndlMonthlyGbp"`
	CaseManagementFolderPK    *string   `json:"ecddCaseManagementFolderPk,omitempty"`
	LoggedAt                  time.Time `json:"loggedAt"`
	UpdatedBy                 string    `json:"updatedBy"`
}
