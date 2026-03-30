package models

import "time"

// ECDDBusinessProfile represents business profile rules for risk-based deposit monitoring
type ECDDBusinessProfile struct {
	ECDDBusinessProfilePK string    `json:"ecddBusinessProfilePk"`
	CountryID             int64     `json:"countryId"`
	StateID               *int64    `json:"stateId,omitempty"`
	RiskStatusID          int64     `json:"riskStatusId"`
	AverageDeposit        float64   `json:"averageDeposit"`
	DepositMultiplier     float64   `json:"depositMultiplier"`
	TimePeriodDays        int64     `json:"timePeriodDays"`
	Enabled               bool      `json:"enabled"`
	LoggedAt              time.Time `json:"loggedAt"`
	UpdatedBy             string    `json:"updatedBy"`
}
