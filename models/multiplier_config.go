package models

import "time"

// ECDDMultiplierConfig represents multiplier configuration for age and status-based threshold adjustments
type ECDDMultiplierConfig struct {
	ECDDMultiplierConfigPK string    `json:"ecddMultiplierConfigPk"`
	CountryID              int64     `json:"countryId"`
	StateID                *int64    `json:"stateId,omitempty"`
	AgeMultipliers         []int64   `json:"ageMultipliers"` // Array of ages where 0.5 multiplier applies
	StatusMultiplier       bool      `json:"statusMultiplier"`
	IsActive               bool      `json:"isActive"`
	LoggedAt               time.Time `json:"loggedAt"`
	UpdatedBy              string    `json:"updatedBy"`
}
