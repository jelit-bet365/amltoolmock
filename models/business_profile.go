package models

import "time"

// ECDDBusinessProfile represents business profile rules for risk-based deposit monitoring
type ECDDBusinessProfile struct {
	ECDDBusinessProfilePK string    `json:"ecdd_business_profile_pk"`
	CountryID             int64     `json:"country_id"`
	StateID               *int64    `json:"state_id,omitempty"`
	RiskStatusID          int64     `json:"risk_status_id"`
	AverageDeposit        float64   `json:"average_deposit"`
	DepositMultiplier     float64   `json:"deposit_multiplier"`
	TimePeriodDays        int64     `json:"time_period_days"`
	Enabled               bool      `json:"enabled"`
	LoggedAt              time.Time `json:"logged_at"`
	UpdatedBy             string    `json:"updated_by"`
}