package models

import "time"

// ECDDMultiplierConfig represents multiplier configuration for age and status-based threshold adjustments
type ECDDMultiplierConfig struct {
	ECDDMultiplierConfigPK string    `json:"ecdd_multiplier_config_pk"`
	CountryID              int64     `json:"country_id"`
	StateID                *int64    `json:"state_id,omitempty"`
	AgeMultipliers         []int64   `json:"age_multipliers"` // Array of ages where 0.5 multiplier applies
	StatusMultiplier       bool      `json:"status_multiplier"`
	IsActive               bool      `json:"is_active"`
	LoggedAt               time.Time `json:"logged_at"`
	UpdatedBy              string    `json:"updated_by"`
}