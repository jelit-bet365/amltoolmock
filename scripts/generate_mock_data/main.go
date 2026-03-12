// Package main generates all mock data JSON files for amltoolmock.
// Run from the project root: go run scripts/generate_mock_data/main.go
package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/google/uuid"
)

// ---------------------------------------------------------------------------
// Country / state / region constants — sourced from enums packages.
// Hardcoded here to keep the script self-contained.
// ---------------------------------------------------------------------------

const (
	// --- Single-country regions ---
	countryUS        = 198
	countryAustralia = 13

	// --- Malta-REGULATED countries (23 total) ---
	countryAustria       = 14
	countryCroatia       = 50
	countryFinland       = 68
	countryGhana         = 76
	countryHungary       = 89
	countryIceland       = 90
	countryIreland       = 95
	countryIvoryCoast    = 48
	countryIndia         = 91 // correct: 91 (92 = Indonesia)
	countryKenya         = 102
	countryLatvia        = 108
	countryLiechtenstein = 113
	countryLithuania     = 114
	countryLuxembourg    = 115
	countryMalta         = 120
	countryNewZealand    = 138 // correct: 138 (140 = Niger)
	countryNicaragua     = 139
	countryNorway        = 143 // correct: 143
	countrySlovakia      = 172
	countrySlovenia      = 167
	countryTanzania      = 185
	countryTonga         = 188
	countryWesternSamoa  = 206

	// --- Malta-LICENSED countries (20 total) ---
	countryUK                  = 197
	countrySpain               = 171
	countryItaly               = 97
	countryDenmark             = 54
	countryGreenland           = 218
	countrySweden              = 181
	countryEstonia             = 64
	countryBulgaria            = 31
	countryCyprus              = 51
	countryGreece              = 78
	countryGermany             = 75
	countryNetherlands         = 135
	countryCzechRepublic       = 52
	countryMexico              = 126
	countryCanadaOntario       = 272
	countryBuenosAiresCity     = 270
	countryBuenosAiresProvince = 271
	countryFrance              = 70
	countrySwitzerland         = 174
	countryJapan               = 99

	// --- Gibraltar-regulated countries (subset for mock data) ---
	countryBrazil       = 28 // correct: 28 not 29
	countryCanada       = 36
	countryGibraltar    = 77
	countryUAE          = 196
	countryArgentina    = 10
	countrySouthKorea   = 104
	countryTurkey       = 190
	countryNigeria      = 141
	countrySouthAfrica  = 169
	countryChile        = 41
	countryPeru         = 150
	countryEgypt        = 59
	countryThailand     = 186
	countryMalaysia     = 121
	countryPakistan     = 145
	countryMorocco      = 130
	countryJordanGib    = 100
	countrySenegal      = 163
	countryBarbados     = 19
	countryTrinidadTob  = 193
)

// US state IDs (from countrystate enum)
const (
	stateNewYork    = 2
	stateNewJersey  = 3
	stateLouisiana  = 33
	stateColorado   = 16
	stateVirginia   = 66
	statePennsylv   = 56
	stateIllinois   = 28
	stateOhio       = 52
	stateMichigan   = 39
	stateConnectcut = 18
)

// Australian state IDs (from countrystate enum)
const (
	stateNSW = 93 // correct: 93 (not 91 which is ACT)
	stateVIC = 98 // correct: 98 (not 92)
	stateQLD = 95 // correct: 95 (not 93)
	stateSA  = 96
	stateWA  = 99
	stateTAS = 97
)

// Region strings
const (
	regionMalta     = "MALTA"
	regionUSA       = "USA"
	regionAustralia = "AUSTRALIA"
	regionGibraltar = "GIBRALTAR"
)

// Currency IDs
const (
	currGBP = 1
	currUSD = 2
	currAUD = 3
	currEUR = 4
	currCAD = 5
)

// ---------------------------------------------------------------------------
// Local struct definitions matching the JSON output format.
// ---------------------------------------------------------------------------

type CaseFolder struct {
	PK         string `json:"ecdd_case_management_folder_pk"`
	FolderName string `json:"folder_name"`
	Region     string `json:"region"`
	CountryID  *int64 `json:"country_id"`
	StateID    *int64 `json:"state_id"`
	LoggedAt   string `json:"logged_at"`
	UpdatedBy  string `json:"updated_by"`
}

type ThresholdConfig struct {
	PK                     string  `json:"ecdd_threshold_config_pk"`
	Title                  string  `json:"title"`
	IsActive               bool    `json:"is_active"`
	CountryID              int64   `json:"country_id"`
	StateID                *int64  `json:"state_id"`
	Type                   int64   `json:"type"`
	Reinvest               bool    `json:"reinvest"`
	Value                  float64 `json:"value"`
	CurrencyID             int64   `json:"currency_id"`
	Period                 int64   `json:"period"`
	UseMultipliers         bool    `json:"use_multipliers"`
	UseRGFlag              bool    `json:"use_rg_flag"`
	ApplyAllStatuses       bool    `json:"apply_all_statuses"`
	Backfill               bool    `json:"backfill"`
	Hierarchy              int64   `json:"hierarchy"`
	ECDDStatus             int64   `json:"ecdd_status"`
	ECDDReviewStatus       int64   `json:"ecdd_review_status"`
	ECDDReportStatus       int64   `json:"ecdd_report_status"`
	SignOffStatus          int64   `json:"sign_off_status"`
	CustomerRiskLevel      int64   `json:"customer_risk_level"`
	NDL28DayGBP            float64 `json:"ndl_28_day_gbp"`
	NDLMonthlyGBP          float64 `json:"ndl_monthly_gbp"`
	CaseManagementFolderPK *string `json:"case_management_folder_pk"`
	LoggedAt               string  `json:"logged_at"`
	UpdatedBy              string  `json:"updated_by"`
}

type UserStatus struct {
	PK                          string  `json:"ecdd_user_status_pk"`
	UserID                      int64   `json:"user_id"`
	UserName                    string  `json:"user_name"`
	CountryID                   int64   `json:"country_id"`
	StateID                     *int64  `json:"state_id"`
	Language                    string  `json:"language"`
	ECDDStatus                  int64   `json:"ecdd_status"`
	ECDDThreshold               float64 `json:"ecdd_threshold"`
	ECDDReviewTrigger           int64   `json:"ecdd_review_trigger"`
	ECDDSuspensionDueDate       *string `json:"ecdd_suspension_due_date"`
	ECDDMultiplier              float64 `json:"ecdd_multiplier"`
	ECDDMultiplierRGFlag        bool    `json:"ecdd_multiplier_rg_flag"`
	UserLtNetDepositThresholdGBP      float64 `json:"user_lt_net_deposit_threshold_gbp"`
	UserLtDepositThresholdGBP   float64 `json:"user_lt_deposit_threshold_gbp"`
	User12MonthNetDepositThresholdGBP float64 `json:"user_12month_net_deposit_threshold_gbp"`
	InfoSource                  int64   `json:"info_source"`
	SignOffStatus               int64   `json:"sign_off_status"`
	DateLastECDDSignOff         *string `json:"date_last_ecdd_sign_off"`
	ECDDRGReviewStatus          int64   `json:"ecdd_rg_review_status"`
	DateLastECDDRGSignOff       *string `json:"date_last_ecdd_rg_sign_off"`
	ECDDReportStatus            int64   `json:"ecdd_report_status"`
	ECDDReviewStatus            int64   `json:"ecdd_review_status"`
	ECDDDocumentStatus          int64   `json:"ecdd_document_status"`
	ECDDEscalationStatus        int64   `json:"ecdd_escalation_status"`
	UARStatus                   int64   `json:"uar_status"`
	LoggedAt                    string  `json:"logged_at"`
	UpdatedBy                   string  `json:"updated_by"`
}

type MultiplierConfig struct {
	PK               string  `json:"ecdd_multiplier_config_pk"`
	CountryID        int64   `json:"country_id"`
	StateID          *int64  `json:"state_id"`
	AgeMultipliers   []int64 `json:"age_multipliers"`
	StatusMultiplier bool    `json:"status_multiplier"`
	IsActive         bool    `json:"is_active"`
	LoggedAt         string  `json:"logged_at"`
	UpdatedBy        string  `json:"updated_by"`
}

type BusinessProfile struct {
	PK                string  `json:"ecdd_business_profile_pk"`
	CountryID         int64   `json:"country_id"`
	StateID           *int64  `json:"state_id"`
	RiskStatusID      int64   `json:"risk_status_id"`
	AverageDeposit    float64 `json:"average_deposit"`
	DepositMultiplier float64 `json:"deposit_multiplier"`
	TimePeriodDays    int64   `json:"time_period_days"`
	Enabled           bool    `json:"enabled"`
	LoggedAt          string  `json:"logged_at"`
	UpdatedBy         string  `json:"updated_by"`
}

type UserCaseFolder struct {
	PK                     string `json:"ecdd_user_case_management_folder_pk"`
	CaseManagementFolderPK string `json:"case_management_folder_pk"`
	UserStatusPK           string `json:"user_status_pk"`
	LoggedAt               string `json:"logged_at"`
	UpdatedBy              string `json:"updated_by"`
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func newUUID() string { return uuid.New().String() }

func i64(v int) *int64 {
	x := int64(v)
	return &x
}

func strPtr(s string) *string { return &s }

func ts(year, month, day, hour, min int) string {
	t := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.UTC)
	return t.Format(time.RFC3339)
}

func dateFmt(year, month, day int) string {
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}

// pick selects an element from a slice using a deterministic index.
func pick[T any](slice []T, i int) T { return slice[i%len(slice)] }

// regionForCountry maps country_id -> region string.
func regionForCountry(cid int) string {
	switch cid {
	case countryUS:
		return regionUSA
	case countryAustralia:
		return regionAustralia
	}

	maltaRegulated := map[int]bool{
		countryAustria: true, countryCroatia: true, countryFinland: true, countryGhana: true,
		countryHungary: true, countryIceland: true, countryIreland: true, countryIvoryCoast: true,
		countryIndia: true, countryKenya: true, countryLatvia: true, countryLiechtenstein: true,
		countryLithuania: true, countryLuxembourg: true, countryMalta: true, countryNewZealand: true,
		countryNicaragua: true, countryNorway: true, countrySlovakia: true, countrySlovenia: true,
		countryTanzania: true, countryTonga: true, countryWesternSamoa: true,
	}
	maltaLicensed := map[int]bool{
		countryUK: true, countrySpain: true, countryItaly: true, countryDenmark: true,
		countryGreenland: true, countrySweden: true, countryEstonia: true, countryBulgaria: true,
		countryCyprus: true, countryGreece: true, countryGermany: true, countryNetherlands: true,
		countryCzechRepublic: true, countryMexico: true, countryCanadaOntario: true,
		countryBuenosAiresCity: true, countryBuenosAiresProvince: true,
		countryFrance: true, countrySwitzerland: true, countryJapan: true,
	}
	if maltaRegulated[cid] || maltaLicensed[cid] {
		return regionMalta
	}
	return regionGibraltar
}

// languageForCountry returns a realistic ISO 639-1 language code.
func languageForCountry(cid int) string {
	m := map[int]string{
		countryUK: "EN", countryIreland: "EN", countryUS: "EN", countryAustralia: "EN",
		countryCanada: "EN", countryNewZealand: "EN", countryGibraltar: "EN",
		countryMalta: "MT", countryFinland: "FI", countrySweden: "SV",
		countryDenmark: "DA", countryNorway: "NO", countryIceland: "IS",
		countryGermany: "DE", countryAustria: "DE", countrySwitzerland: "DE",
		countryLiechtenstein: "DE", countryFrance: "FR", countryLuxembourg: "FR",
		countryItaly: "IT", countrySpain: "ES", countryGreece: "EL",
		countryCzechRepublic: "CS", countryHungary: "HU", countrySlovakia: "SK",
		countrySlovenia: "SL", countryCroatia: "HR", countryBulgaria: "BG",
		countryEstonia: "ET", countryLatvia: "LV", countryLithuania: "LT",
		countryCyprus: "EL", countryNetherlands: "NL", countryIndia: "EN",
		countryJapan: "JA", countryBrazil: "PT", countryArgentina: "ES",
		countryMexico: "ES", countryBuenosAiresCity: "ES", countryBuenosAiresProvince: "ES",
		countrySouthKorea: "KO", countryTurkey: "TR", countryUAE: "AR",
		countryNigeria: "EN", countrySouthAfrica: "EN", countryKenya: "EN",
		countryGhana: "EN", countryTanzania: "EN", countryEgypt: "AR",
		countryThailand: "TH", countryMalaysia: "MS", countryPakistan: "UR",
		countryMorocco: "AR", countryJordanGib: "AR", countrySenegal: "FR",
		countryBarbados: "EN", countryTrinidadTob: "EN", countryChile: "ES",
		countryPeru: "ES", countryIvoryCoast: "FR", countryNicaragua: "ES",
		countryCanadaOntario: "EN",
	}
	if lang, ok := m[cid]; ok {
		return lang
	}
	return "EN"
}

// ---------------------------------------------------------------------------
// 1. CASE FOLDERS — 67 Malta, 35 USA, 12 Australia, 4 Gibraltar = 118 total
// ---------------------------------------------------------------------------

func generateFolders() []CaseFolder {
	type fdef struct {
		name   string
		region string
	}
	var defs []fdef

	maltaNames := []string{
		"MT - Unusual Activity Reports",
		"MT - NEW Population - ECDD",
		"MT - NEW Population - OMT",
		"MT - Report Required - OMT",
		"MT - Report Required - ECDD",
		"MT - OMT - Monitoring",
		"MT - OMT - TL Review",
		"MT - OMT - Sup Review",
		"MT - OMT - MGMT Review",
		"MT - ROW - ISO Required (TL)",
		"MT - ROW - ISO Required (Sup)",
		"MT - ROW - ISO Required (MGMT)",
		"MT - ROW - MIC Complete",
		"MT - ROW - EV Pending Review",
		"MT - ROW - EV Satisfactory Review (TL)",
		"MT - ROW - EV Satisfactory Review (Sup)",
		"MT - ROW - EV Satisfactory Review (DM)",
		"MT - ROW - EV Satisfactory Review (MAN)",
		"MT - ROW - FSO Review (Specialist)",
		"MT - ROW - FSO Review (TL)",
		"MT - ROW - FSO Review (Sup)",
		"MT - ROW - FSO Review (MGMT)",
		"MT - ROW - Priority Review",
		"MT - UKO - ISO Required",
		"MT - UKO - EV Pending Review",
		"MT - UKO - EV Satisfactory Review (Sup)",
		"MT - UKO - EV Satisfactory Review (DM)",
		"MT - UKO - EV Satisfactory Review (MAN)",
		"MT - UKO - FSO Review (Sup)",
		"MT - UKO - FSO Review (MGMT)",
		"MT - UKO - Priority Review",
		"MT - UKO - 14 Day Ev Review",
		"MT - UKO - FSO Monitoring",
		"MT - CX - MIC Pending",
		"MT - CX - High Value",
		"MT - CX - Contact Monitoring",
		"MT - CX - Balance Review (Team Leader)",
		"MT - CX - Balance Review (Supervisor)",
		"MT - CX - Balance Review (MGMT)",
		"MT - CX - Expedited Review (Specialist)",
		"MT - CX - Expedited Review (Team Leader)",
		"MT - CX - Expedited Review (Supervisor)",
		"MT - CX - Expedited Review (MGMT)",
		"MT - CX - Initial Contact",
		"MT - CX - FIC Follow-Up",
		"MT - CX - Final Review Complete",
		"MT - CX - Documents Pending",
		"MT - CX - Admissions Pending",
		"MT - CX - Suspension",
		"MT - CX - Awaiting Evidence Summary (POI)",
		"MT - CX - Awaiting Evidence Summary (FOF)",
		"MT - CX - Awaiting Evidence Summary (Other)",
		"MT - CX - Call Back (Specialist)",
		"MT - CX - Call Back (TL and Above)",
		"MT - CDD ID Review",
		"MT - Senior MGMT Review",
		"MT - SCO Review",
		"MT - MLRO Review",
		"MT - MLRO Monitoring",
		"MT - Pending Closure",
		"MT - Reference",
		"MT - Pending QA",
		"MT - Process Review (Cat 1)",
		"MT - Process Review (Cat 2)",
		"ECDD Household",
		"ECDD Auto Suspend",
		"Details Changed",
		"NEW Population - AML Reports",
	}
	for _, n := range maltaNames {
		defs = append(defs, fdef{n, regionMalta})
	}

	usaNames := []string{
		"US - Unusual Activity Reports",
		"US - NEW Population - ECDD",
		"US - Report Required \u2013 ECDD",
		"US - ISO Required (Specialist)",
		"US - ISO Required (TL)",
		"US - ISO Required (Sup)",
		"US - ISO Required (MGMT)",
		"US - MIC Pending Customer",
		"US - MIC Complete",
		"US - Post Contact Review",
		"US - No Contact Review",
		"US - Awaiting Documents",
		"US - EV Pending Review",
		"US - EV Satisfactory Review (TL)",
		"US - EV Satisfactory Review (Sup)",
		"US - EV Satisfactory Review (DM)",
		"US - EV Satisfactory Review (MAN)",
		"US - FSO Review (Specialist)",
		"US - FSO Review (TL)",
		"US - FSO Review (Sup)",
		"US - FSO Review (MGMT)",
		"US - Process Review (Cat 1)",
		"US - Process Review (Cat 2)",
		"US - AML Report Pending",
		"US - AML Report Review",
		"US - TMP Report Pending",
		"US - TMP Report Review",
		"US - AML Specialist Review",
		"US - AML Team Leader Review",
		"US - SCO Review",
		"US - Pending Closure",
		"US - Details Changed",
		"US - ID Review",
		"ECDD Auto Suspend",
		"Details Changed",
		"NEW Population - AML Reports",
	}
	for _, n := range usaNames {
		defs = append(defs, fdef{n, regionUSA})
	}

	auNames := []string{
		"AU - Unusual Activity Reports",
		"AU - NEW Population - ECDD",
		"AU - Report Required - ECDD",
		"AU - ISO Required",
		"AU - Pending Contact",
		"AU - Evidence Review Required",
		"AU - FSO Required",
		"AU - Monitoring",
		"AU - CRO Review",
		"ECDD Auto Suspend",
		"Details Changed",
		"NEW Population - AML Reports",
	}
	for _, n := range auNames {
		defs = append(defs, fdef{n, regionAustralia})
	}

	gibNames := []string{
		"GIB - Unusual Activity Reports",
		"ECDD Auto Suspend",
		"Details Changed",
		"NEW Population - AML Reports",
	}
	for _, n := range gibNames {
		defs = append(defs, fdef{n, regionGibraltar})
	}

	updaters := []string{
		"admin@company.com", "compliance@company.com",
		"manager@company.com", "senior.manager@company.com",
	}
	baseDays := map[string]int{regionMalta: 1, regionUSA: 15, regionAustralia: 20, regionGibraltar: 25}
	regIdx := map[string]int{regionMalta: 0, regionUSA: 0, regionAustralia: 0, regionGibraltar: 0}

	folders := make([]CaseFolder, len(defs))
	for i, d := range defs {
		ri := regIdx[d.region]
		regIdx[d.region]++
		day := baseDays[d.region] + (ri / 4)
		month := 1
		if day > 31 {
			month = 2
			day -= 31
		}
		if day < 1 {
			day = 1
		}
		folders[i] = CaseFolder{
			PK:         newUUID(),
			FolderName: d.name,
			Region:     d.region,
			CountryID:  nil,
			StateID:    nil,
			LoggedAt:   ts(2026, month, day, 8+(ri%10), (ri*13)%60),
			UpdatedBy:  updaters[i%len(updaters)],
		}
	}
	return folders
}

// foldersByRegion groups folder PKs by region name.
func foldersByRegion(folders []CaseFolder) map[string][]string {
	m := make(map[string][]string)
	for _, f := range folders {
		m[f.Region] = append(m[f.Region], f.PK)
	}
	return m
}

// ---------------------------------------------------------------------------
// 2. THRESHOLD CONFIGS — 20 configs across all 4 regions
// ---------------------------------------------------------------------------

func generateThresholds(fbr map[string][]string) []ThresholdConfig {
	mtF := fbr[regionMalta]
	usF := fbr[regionUSA]
	auF := fbr[regionAustralia]
	giF := fbr[regionGibraltar]

	si := func(v int) *int64 { return i64(v) }

	return []ThresholdConfig{
		// ===== MALTA region (7 thresholds) =====
		{
			PK: newUUID(), Title: "UK 28-Day NDL 50K", IsActive: true,
			CountryID: countryUK, Type: 2, Value: 50000, CurrencyID: currGBP, Period: 2,
			UseMultipliers: true, UseRGFlag: true, Hierarchy: 1,
			ECDDStatus: 2, ECDDReviewStatus: 4, ECDDReportStatus: 2, SignOffStatus: 1,
			CustomerRiskLevel: 3, NDL28DayGBP: 5000, NDLMonthlyGBP: 20000,
			CaseManagementFolderPK: strPtr(mtF[0]),
			LoggedAt: ts(2026, 1, 5, 9, 0), UpdatedBy: "john.doe@company.com",
		},
		{
			PK: newUUID(), Title: "Ireland 28-Day NDL 20K", IsActive: true,
			CountryID: countryIreland, Type: 2, Reinvest: true, Value: 20000, CurrencyID: currEUR, Period: 2,
			UseMultipliers: true, UseRGFlag: true, ApplyAllStatuses: true, Backfill: true,
			Hierarchy: 2,
			ECDDStatus: 2, ECDDReviewStatus: 3, ECDDReportStatus: 2, SignOffStatus: 1,
			CustomerRiskLevel: 2, NDL28DayGBP: 4000, NDLMonthlyGBP: 15000,
			CaseManagementFolderPK: strPtr(mtF[5]),
			LoggedAt: ts(2026, 1, 6, 10, 30), UpdatedBy: "emma.wilson@company.com",
		},
		{
			PK: newUUID(), Title: "Finland 182-Day Stakes Monitor", IsActive: true,
			CountryID: countryFinland, Type: 3, Value: 35000, CurrencyID: currEUR, Period: 5,
			Hierarchy: 3,
			ECDDStatus: 1, ECDDReviewStatus: 2, ECDDReportStatus: 1, SignOffStatus: 1,
			CustomerRiskLevel: 2, NDL28DayGBP: 6000, NDLMonthlyGBP: 22000,
			CaseManagementFolderPK: strPtr(mtF[10]),
			LoggedAt: ts(2026, 1, 7, 11, 0), UpdatedBy: "senior.manager@company.com",
		},
		{
			PK: newUUID(), Title: "Malta Domestic 24hr Deposit 8K", IsActive: true,
			CountryID: countryMalta, Type: 1, Value: 8000, CurrencyID: currEUR, Period: 1,
			UseRGFlag: true, ApplyAllStatuses: true,
			Hierarchy: 4,
			ECDDStatus: 2, ECDDReviewStatus: 3, ECDDReportStatus: 2, SignOffStatus: 1,
			CustomerRiskLevel: 2, NDL28DayGBP: 2500, NDLMonthlyGBP: 9000,
			CaseManagementFolderPK: strPtr(mtF[3]),
			LoggedAt: ts(2026, 1, 8, 14, 0), UpdatedBy: "monitoring@company.com",
		},
		{
			PK: newUUID(), Title: "Sweden 91-Day Deposit Monitor", IsActive: true,
			CountryID: countrySweden, Type: 1, Reinvest: true, Value: 30000, CurrencyID: currEUR, Period: 4,
			Backfill: true, Hierarchy: 5,
			ECDDStatus: 2, ECDDReviewStatus: 3, ECDDReportStatus: 1, SignOffStatus: 1,
			CustomerRiskLevel: 2, NDL28DayGBP: 5000, NDLMonthlyGBP: 18000,
			CaseManagementFolderPK: strPtr(mtF[20]),
			LoggedAt: ts(2026, 1, 9, 15, 30), UpdatedBy: "compliance@company.com",
		},
		{
			PK: newUUID(), Title: "Germany 28-Day NDL 25K", IsActive: true,
			CountryID: countryGermany, Type: 2, Value: 25000, CurrencyID: currEUR, Period: 2,
			UseMultipliers: true, UseRGFlag: true, Hierarchy: 6,
			ECDDStatus: 2, ECDDReviewStatus: 3, ECDDReportStatus: 2, SignOffStatus: 1,
			CustomerRiskLevel: 2, NDL28DayGBP: 4500, NDLMonthlyGBP: 17000,
			CaseManagementFolderPK: strPtr(mtF[14]),
			LoggedAt: ts(2026, 1, 10, 9, 45), UpdatedBy: "compliance@company.com",
		},
		{
			PK: newUUID(), Title: "Italy 365-Day Stakes 60K", IsActive: true,
			CountryID: countryItaly, Type: 3, Value: 60000, CurrencyID: currEUR, Period: 6,
			Backfill: true, Hierarchy: 7,
			ECDDStatus: 2, ECDDReviewStatus: 2, ECDDReportStatus: 1, SignOffStatus: 1,
			CustomerRiskLevel: 2, NDL28DayGBP: 8000, NDLMonthlyGBP: 30000,
			CaseManagementFolderPK: strPtr(mtF[22]),
			LoggedAt: ts(2026, 1, 11, 10, 15), UpdatedBy: "senior.manager@company.com",
		},

		// ===== USA region (5 thresholds) =====
		{
			PK: newUUID(), Title: "US-NY 28-Day NDL 75K", IsActive: true,
			CountryID: countryUS, StateID: si(stateNewYork),
			Type: 2, Value: 75000, CurrencyID: currUSD, Period: 2,
			UseMultipliers: true, UseRGFlag: true,
			Hierarchy: 8,
			ECDDStatus: 3, ECDDReviewStatus: 4, ECDDReportStatus: 3, SignOffStatus: 2,
			CustomerRiskLevel: 4, NDL28DayGBP: 8000, NDLMonthlyGBP: 30000,
			CaseManagementFolderPK: strPtr(usF[0]),
			LoggedAt: ts(2026, 1, 12, 9, 0), UpdatedBy: "mike.johnson@company.com",
		},
		{
			PK: newUUID(), Title: "US-NJ 365-Day Stakes 90K", IsActive: true,
			CountryID: countryUS, StateID: si(stateNewJersey),
			Type: 3, Value: 90000, CurrencyID: currUSD, Period: 6,
			Backfill: true, Hierarchy: 9,
			ECDDStatus: 2, ECDDReviewStatus: 3, ECDDReportStatus: 2, SignOffStatus: 1,
			CustomerRiskLevel: 3, NDL28DayGBP: 10000, NDLMonthlyGBP: 40000,
			CaseManagementFolderPK: strPtr(usF[5]),
			LoggedAt: ts(2026, 1, 13, 14, 10), UpdatedBy: "compliance@company.com",
		},
		{
			PK: newUUID(), Title: "US-LA 24hr Deposit 5K", IsActive: true,
			CountryID: countryUS, StateID: si(stateLouisiana),
			Type: 1, Value: 5000, CurrencyID: currUSD, Period: 1,
			UseRGFlag: true, ApplyAllStatuses: true,
			Hierarchy: 10,
			ECDDStatus: 2, ECDDReviewStatus: 3, ECDDReportStatus: 1, SignOffStatus: 1,
			CustomerRiskLevel: 2, NDL28DayGBP: 2000, NDLMonthlyGBP: 8000,
			CaseManagementFolderPK: strPtr(usF[12]),
			LoggedAt: ts(2026, 1, 14, 11, 45), UpdatedBy: "risk.officer@company.com",
		},
		{
			PK: newUUID(), Title: "US AML Report Reinvest Monitor", IsActive: true,
			CountryID: countryUS,
			Type: 2, Reinvest: true, Value: 10000, CurrencyID: currUSD, Period: 3,
			ApplyAllStatuses: true, Hierarchy: 11,
			ECDDStatus: 2, ECDDReviewStatus: 2, ECDDReportStatus: 1, SignOffStatus: 1,
			CustomerRiskLevel: 2, NDL28DayGBP: 1500, NDLMonthlyGBP: 6000,
			CaseManagementFolderPK: strPtr(usF[23]),
			LoggedAt: ts(2026, 1, 15, 9, 0), UpdatedBy: "admin@company.com",
		},
		{
			PK: newUUID(), Title: "US-CO 84-Day NDL 40K", IsActive: true,
			CountryID: countryUS, StateID: si(stateColorado),
			Type: 2, Value: 40000, CurrencyID: currUSD, Period: 3,
			UseMultipliers: true, Hierarchy: 12,
			ECDDStatus: 2, ECDDReviewStatus: 3, ECDDReportStatus: 2, SignOffStatus: 1,
			CustomerRiskLevel: 3, NDL28DayGBP: 6000, NDLMonthlyGBP: 24000,
			CaseManagementFolderPK: strPtr(usF[3]),
			LoggedAt: ts(2026, 1, 16, 10, 30), UpdatedBy: "compliance@company.com",
		},

		// ===== AUSTRALIA region (4 thresholds) =====
		{
			PK: newUUID(), Title: "AUS-NSW 84-Day NDL 40K", IsActive: true,
			CountryID: countryAustralia, StateID: si(stateNSW),
			Type: 2, Value: 40000, CurrencyID: currAUD, Period: 3,
			Backfill: true, Hierarchy: 13,
			ECDDStatus: 2, ECDDReviewStatus: 3, ECDDReportStatus: 1, SignOffStatus: 1,
			CustomerRiskLevel: 2, NDL28DayGBP: 7000, NDLMonthlyGBP: 28000,
			CaseManagementFolderPK: strPtr(auF[0]),
			LoggedAt: ts(2026, 1, 17, 15, 30), UpdatedBy: "risk.officer@company.com",
		},
		{
			PK: newUUID(), Title: "AUS-VIC 28-Day Deposit 25K", IsActive: true,
			CountryID: countryAustralia, StateID: si(stateVIC),
			Type: 1, Value: 25000, CurrencyID: currAUD, Period: 2,
			UseRGFlag: true, ApplyAllStatuses: true,
			Hierarchy: 14,
			ECDDStatus: 2, ECDDReviewStatus: 4, ECDDReportStatus: 2, SignOffStatus: 1,
			CustomerRiskLevel: 3, NDL28DayGBP: 5000, NDLMonthlyGBP: 18000,
			CaseManagementFolderPK: strPtr(auF[4]),
			LoggedAt: ts(2026, 1, 18, 10, 0), UpdatedBy: "compliance@company.com",
		},
		{
			PK: newUUID(), Title: "AUS-QLD Reinvest 365-Day Monitor", IsActive: true,
			CountryID: countryAustralia, StateID: si(stateQLD),
			Type: 3, Reinvest: true, Value: 50000, CurrencyID: currAUD, Period: 6,
			Hierarchy: 15,
			ECDDStatus: 1, ECDDReviewStatus: 2, ECDDReportStatus: 1, SignOffStatus: 1,
			CustomerRiskLevel: 2, NDL28DayGBP: 9000, NDLMonthlyGBP: 35000,
			CaseManagementFolderPK: strPtr(auF[7]),
			LoggedAt: ts(2026, 1, 19, 14, 0), UpdatedBy: "senior.manager@company.com",
		},
		{
			PK: newUUID(), Title: "AUS 28-Day Deposit 15K National", IsActive: true,
			CountryID: countryAustralia,
			Type: 1, Value: 15000, CurrencyID: currAUD, Period: 2,
			ApplyAllStatuses: true, Hierarchy: 16,
			ECDDStatus: 2, ECDDReviewStatus: 3, ECDDReportStatus: 1, SignOffStatus: 1,
			CustomerRiskLevel: 2, NDL28DayGBP: 3500, NDLMonthlyGBP: 12000,
			CaseManagementFolderPK: strPtr(auF[1]),
			LoggedAt: ts(2026, 1, 20, 9, 30), UpdatedBy: "admin@company.com",
		},

		// ===== GIBRALTAR region (4 thresholds) =====
		{
			PK: newUUID(), Title: "Canada 182-Day Deposit Monitor", IsActive: true,
			CountryID: countryCanada, Type: 1, Value: 30000, CurrencyID: currCAD, Period: 5,
			UseMultipliers: true, UseRGFlag: true, Backfill: true,
			Hierarchy: 17,
			ECDDStatus: 2, ECDDReviewStatus: 4, ECDDReportStatus: 2, SignOffStatus: 1,
			CustomerRiskLevel: 3, NDL28DayGBP: 6000, NDLMonthlyGBP: 25000,
			CaseManagementFolderPK: strPtr(giF[0]),
			LoggedAt: ts(2026, 1, 21, 13, 55), UpdatedBy: "compliance@company.com",
		},
		{
			PK: newUUID(), Title: "Brazil 28-Day High Risk NDL", IsActive: true,
			CountryID: countryBrazil, Type: 2, Value: 25000, CurrencyID: currCAD, Period: 2,
			UseRGFlag: true, Hierarchy: 18,
			ECDDStatus: 3, ECDDReviewStatus: 4, ECDDReportStatus: 3, SignOffStatus: 2,
			CustomerRiskLevel: 4, NDL28DayGBP: 5000, NDLMonthlyGBP: 18000,
			CaseManagementFolderPK: strPtr(giF[1]),
			LoggedAt: ts(2026, 1, 22, 10, 15), UpdatedBy: "senior.manager@company.com",
		},
		{
			PK: newUUID(), Title: "Gibraltar Reinvest 365-Day Stakes", IsActive: true,
			CountryID: countryGibraltar, Type: 3, Reinvest: true, Value: 60000, CurrencyID: currGBP, Period: 6,
			ApplyAllStatuses: true, Backfill: true,
			Hierarchy: 19,
			ECDDStatus: 2, ECDDReviewStatus: 3, ECDDReportStatus: 2, SignOffStatus: 1,
			CustomerRiskLevel: 2, NDL28DayGBP: 10000, NDLMonthlyGBP: 35000,
			CaseManagementFolderPK: strPtr(giF[2]),
			LoggedAt: ts(2026, 1, 23, 12, 50), UpdatedBy: "monitoring@company.com",
		},
		{
			PK: newUUID(), Title: "UAE 91-Day Deposit 45K", IsActive: true,
			CountryID: countryUAE, Type: 1, Value: 45000, CurrencyID: currGBP, Period: 4,
			UseRGFlag: true, Hierarchy: 20,
			ECDDStatus: 2, ECDDReviewStatus: 3, ECDDReportStatus: 2, SignOffStatus: 1,
			CustomerRiskLevel: 3, NDL28DayGBP: 7000, NDLMonthlyGBP: 27000,
			CaseManagementFolderPK: strPtr(giF[0]),
			LoggedAt: ts(2026, 1, 24, 16, 0), UpdatedBy: "risk.officer@company.com",
		},
	}
}

// ---------------------------------------------------------------------------
// 3. MULTIPLIER CONFIGS — 12 configs
// ---------------------------------------------------------------------------

func generateMultipliers() []MultiplierConfig {
	return []MultiplierConfig{
		// MALTA region
		{PK: newUUID(), CountryID: countryUK, AgeMultipliers: []int64{18, 19, 20, 21, 22, 23, 24, 25}, StatusMultiplier: true, IsActive: true, LoggedAt: ts(2026, 1, 3, 10, 0), UpdatedBy: "admin@company.com"},
		{PK: newUUID(), CountryID: countryIreland, AgeMultipliers: []int64{18, 19, 20, 21, 22, 23, 24, 25}, StatusMultiplier: true, IsActive: true, LoggedAt: ts(2026, 1, 3, 10, 30), UpdatedBy: "admin@company.com"},
		{PK: newUUID(), CountryID: countryMalta, AgeMultipliers: []int64{18, 19, 20, 21}, StatusMultiplier: false, IsActive: true, LoggedAt: ts(2026, 1, 3, 11, 0), UpdatedBy: "compliance@company.com"},
		{PK: newUUID(), CountryID: countryGermany, AgeMultipliers: []int64{18, 19, 20, 21, 22, 23}, StatusMultiplier: true, IsActive: true, LoggedAt: ts(2026, 1, 3, 11, 30), UpdatedBy: "compliance@company.com"},
		{PK: newUUID(), CountryID: countryItaly, AgeMultipliers: []int64{18, 19, 20, 21, 22}, StatusMultiplier: false, IsActive: true, LoggedAt: ts(2026, 1, 3, 12, 0), UpdatedBy: "compliance@company.com"},
		// USA region
		{PK: newUUID(), CountryID: countryUS, StateID: i64(stateNewYork), AgeMultipliers: []int64{18, 19, 20, 21, 22, 23, 24, 25, 26}, StatusMultiplier: false, IsActive: true, LoggedAt: ts(2026, 1, 4, 9, 0), UpdatedBy: "admin@company.com"},
		{PK: newUUID(), CountryID: countryUS, StateID: i64(stateNewJersey), AgeMultipliers: []int64{18, 19, 20, 21, 22, 23, 24}, StatusMultiplier: true, IsActive: true, LoggedAt: ts(2026, 1, 4, 9, 30), UpdatedBy: "admin@company.com"},
		{PK: newUUID(), CountryID: countryUS, StateID: i64(stateLouisiana), AgeMultipliers: []int64{18, 19, 20, 21, 22}, StatusMultiplier: true, IsActive: true, LoggedAt: ts(2026, 1, 4, 10, 0), UpdatedBy: "admin@company.com"},
		{PK: newUUID(), CountryID: countryUS, StateID: i64(stateColorado), AgeMultipliers: []int64{18, 19, 20, 21, 22, 23}, StatusMultiplier: true, IsActive: true, LoggedAt: ts(2026, 1, 4, 10, 30), UpdatedBy: "admin@company.com"},
		// AUSTRALIA region
		{PK: newUUID(), CountryID: countryAustralia, AgeMultipliers: []int64{18, 19, 20, 21, 22}, StatusMultiplier: true, IsActive: true, LoggedAt: ts(2026, 1, 4, 11, 0), UpdatedBy: "compliance@company.com"},
		// GIBRALTAR region
		{PK: newUUID(), CountryID: countryCanada, AgeMultipliers: []int64{18, 19, 20, 21, 22, 23}, StatusMultiplier: true, IsActive: true, LoggedAt: ts(2026, 1, 4, 11, 30), UpdatedBy: "admin@company.com"},
		{PK: newUUID(), CountryID: countryBrazil, AgeMultipliers: []int64{18, 19, 20, 21}, StatusMultiplier: false, IsActive: true, LoggedAt: ts(2026, 1, 4, 12, 0), UpdatedBy: "risk.officer@company.com"},
	}
}

// ---------------------------------------------------------------------------
// 4. BUSINESS PROFILES — 18 profiles
// ---------------------------------------------------------------------------

func generateBusinessProfiles() []BusinessProfile {
	type bp struct {
		cid       int
		stateID   *int64
		risk      int64
		avgDep    float64
		mult      float64
		period    int64
		enabled   bool
		loggedAt  string
		updatedBy string
	}

	defs := []bp{
		// MALTA region
		{countryUK, nil, 2, 500, 1.5, 28, true, ts(2026, 1, 2, 9, 0), "admin@company.com"},
		{countryIreland, nil, 2, 450, 1.5, 28, true, ts(2026, 1, 2, 9, 30), "compliance@company.com"},
		{countryFinland, nil, 1, 350, 1.2, 182, true, ts(2026, 1, 2, 10, 0), "compliance@company.com"},
		{countryMalta, nil, 2, 400, 1.3, 28, true, ts(2026, 1, 2, 10, 30), "compliance@company.com"},
		{countrySweden, nil, 2, 420, 1.4, 91, true, ts(2026, 1, 2, 11, 0), "compliance@company.com"},
		{countryGermany, nil, 2, 550, 1.6, 28, true, ts(2026, 1, 2, 11, 30), "compliance@company.com"},
		{countryItaly, nil, 2, 480, 1.4, 28, true, ts(2026, 1, 2, 12, 0), "compliance@company.com"},
		{countrySpain, nil, 2, 460, 1.4, 28, true, ts(2026, 1, 2, 12, 30), "compliance@company.com"},
		// USA region
		{countryUS, i64(stateNewYork), 3, 800, 2.0, 28, true, ts(2026, 1, 2, 13, 0), "admin@company.com"},
		{countryUS, i64(stateNewJersey), 2, 600, 1.5, 28, true, ts(2026, 1, 2, 13, 30), "admin@company.com"},
		{countryUS, i64(stateLouisiana), 2, 400, 1.2, 28, true, ts(2026, 1, 2, 14, 0), "admin@company.com"},
		{countryUS, i64(stateColorado), 2, 550, 1.6, 28, true, ts(2026, 1, 2, 14, 30), "admin@company.com"},
		// AUSTRALIA region
		{countryAustralia, nil, 2, 700, 1.8, 84, true, ts(2026, 1, 2, 15, 0), "admin@company.com"},
		{countryAustralia, i64(stateNSW), 3, 750, 1.9, 28, true, ts(2026, 1, 2, 15, 30), "compliance@company.com"},
		// GIBRALTAR region
		{countryCanada, nil, 3, 650, 1.8, 182, true, ts(2026, 1, 2, 16, 0), "risk.officer@company.com"},
		{countryBrazil, nil, 4, 550, 1.6, 28, true, ts(2026, 1, 2, 16, 30), "risk.officer@company.com"},
		{countryGibraltar, nil, 2, 480, 1.3, 365, true, ts(2026, 1, 2, 17, 0), "compliance@company.com"},
		{countryUAE, nil, 3, 900, 2.2, 28, true, ts(2026, 1, 2, 17, 30), "risk.officer@company.com"},
	}

	profiles := make([]BusinessProfile, len(defs))
	for i, d := range defs {
		profiles[i] = BusinessProfile{
			PK:                newUUID(),
			CountryID:         int64(d.cid),
			StateID:           d.stateID,
			RiskStatusID:      d.risk,
			AverageDeposit:    d.avgDep,
			DepositMultiplier: d.mult,
			TimePeriodDays:    d.period,
			Enabled:           d.enabled,
			LoggedAt:          d.loggedAt,
			UpdatedBy:         d.updatedBy,
		}
	}
	return profiles
}

// ---------------------------------------------------------------------------
// 5. USER STATUSES — 200 total
//    80 Malta, 50 USA, 35 Australia, 35 Gibraltar
// ---------------------------------------------------------------------------

type userSpec struct {
	userID    int64
	userName  string
	countryID int64
	stateID   *int64
	language  string
}

func generateUsers() []UserStatus {
	specs := buildUserSpecs()
	users := make([]UserStatus, len(specs))

	// --- Enum value pools (see docs/ecdd_user_status_enums.md) ---

	// ecdd_status: 1-7
	// Weighted: In Progress & Not Required most common
	weightedStatuses := []int64{
		1, 1, 1, 1, // Not Required ~20%
		2, 2, 2, 2, 2, 2, // In Progress ~30%
		3, 3, 3, 3, // Complete ~20%
		4, 4, // Suspended - Manual ~10%
		5, 5, // Suspended - Auto ~10%
		6, // Closed ~5%
		7, // Block Deposits ~5%
	}

	// ecdd_review_trigger: 1-9
	reviewTriggers := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// info_source: 1-9
	infoSources := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// sign_off_status: 1-6
	signOffStatuses := []int64{1, 1, 2, 2, 3, 4, 5, 6}

	// ecdd_rg_review_status: 1-3
	rgReviewStatuses := []int64{1, 1, 2, 3}

	// ecdd_report_status: 1-3
	reportStatuses := []int64{1, 2, 3}

	// ecdd_review_status: 1-7
	ecddReviewStatuses := []int64{1, 2, 3, 4, 5, 6, 7}

	// ecdd_document_status: 1-5
	docStatuses := []int64{1, 2, 3, 4, 5}

	// ecdd_escalation_status: 1-3
	escalationStatuses := []int64{1, 1, 2, 3}

	// uar_status: 1-2
	uarStatuses := []int64{1, 1, 1, 2}

	thresholds := []float64{1000, 2000, 5000, 8000, 10000, 15000, 20000, 25000, 50000, 75000}
	multipliers := []float64{0.5, 1.0, 1.0, 1.0, 1.25, 1.5, 2.0}

	// Threshold tiers: each row = [LT Deposit, LT Net Deposit, 12-Month Net Deposit]
	// Invariant: LT Deposit >= LT Net Deposit >= 12-Month Net Deposit
	type thresholdTier struct {
		ltDeposit      float64 // user_lt_deposit_threshold_gbp (lifetime total deposits)
		ltNetDeposit   float64 // user_lt_net_deposit_threshold_gbp (lifetime net deposit = deposits - withdrawals)
		month12NetDep  float64 // user_12month_net_deposit_threshold_gbp (12-month net deposit)
	}
	tiers := []thresholdTier{
		{5000, 3000, 1500},
		{10000, 6000, 3000},
		{15000, 8000, 4000},
		{25000, 12000, 5000},
		{35000, 18000, 8000},
		{50000, 25000, 10000},
		{75000, 40000, 15000},
		{100000, 55000, 20000},
		{150000, 80000, 30000},
		{200000, 110000, 45000},
		{300000, 160000, 60000},
		{500000, 250000, 100000},
	}

	suspDates := []string{
		ts(2026, 3, 15, 0, 0), ts(2026, 4, 1, 0, 0), ts(2026, 4, 30, 0, 0),
		ts(2026, 5, 15, 0, 0), ts(2026, 6, 1, 0, 0), ts(2026, 6, 30, 0, 0),
	}
	signOffDates := []string{
		ts(2025, 11, 10, 14, 0), ts(2025, 12, 5, 10, 30), ts(2026, 1, 8, 9, 0),
		ts(2026, 1, 15, 11, 0), ts(2026, 1, 20, 15, 30), ts(2026, 2, 1, 10, 0),
		ts(2026, 2, 10, 14, 45), ts(2026, 2, 20, 9, 15),
	}
	updaters := []string{
		"admin@company.com", "compliance@company.com", "risk.officer@company.com",
		"senior.manager@company.com", "system@company.com", "monitoring@company.com",
	}

	for i, s := range specs {
		ecddStatus := pick(weightedStatuses, i*7+3)
		threshold := pick(thresholds, i*3+1)
		mult := pick(multipliers, i*5+2)
		reviewTrigger := pick(reviewTriggers, i*3+1)

		var suspDate *string
		if ecddStatus == 4 || ecddStatus == 5 {
			sd := pick(suspDates, i)
			suspDate = &sd
		}

		var signOffDate *string
		var rgSignOffDate *string
		if ecddStatus == 3 || ecddStatus == 6 {
			sd := pick(signOffDates, i)
			signOffDate = &sd
			rgd := pick(signOffDates, i+3)
			rgSignOffDate = &rgd
		}

		// Generate timestamp spread across Jan-Feb 2026
		day := 1 + (i % 59)
		month := 1
		if day > 31 {
			month = 2
			day -= 31
		}
		hour := 8 + (i % 12)
		min := (i * 7) % 60

		tier := pick(tiers, i*3+1)

		users[i] = UserStatus{
			PK:                          newUUID(),
			UserID:                      s.userID,
			UserName:                    s.userName,
			CountryID:                   s.countryID,
			StateID:                     s.stateID,
			Language:                    s.language,
			ECDDStatus:                  ecddStatus,
			ECDDThreshold:               threshold,
			ECDDReviewTrigger:           reviewTrigger,
			ECDDSuspensionDueDate:       suspDate,
			ECDDMultiplier:              mult,
			ECDDMultiplierRGFlag:        ((i+1)%3 == 0),
			UserLtNetDepositThresholdGBP:      tier.ltNetDeposit,
			UserLtDepositThresholdGBP:   tier.ltDeposit,
			User12MonthNetDepositThresholdGBP: tier.month12NetDep,
			InfoSource:                  pick(infoSources, i*2+1),
			SignOffStatus:               pick(signOffStatuses, i*2),
			DateLastECDDSignOff:         signOffDate,
			ECDDRGReviewStatus:          pick(rgReviewStatuses, i*3),
			DateLastECDDRGSignOff:       rgSignOffDate,
			ECDDReportStatus:            pick(reportStatuses, i),
			ECDDReviewStatus:            pick(ecddReviewStatuses, i*2+1),
			ECDDDocumentStatus:          pick(docStatuses, i*2),
			ECDDEscalationStatus:        pick(escalationStatuses, i),
			UARStatus:                   pick(uarStatuses, i),
			LoggedAt:                    ts(2026, month, day, hour, min),
			UpdatedBy:                   pick(updaters, i),
		}
	}
	return users
}

// buildUserSpecs creates 200 user definitions across all 4 regions.
func buildUserSpecs() []userSpec {
	var specs []userSpec
	baseID := int64(10001)

	// =========================================================================
	// MALTA region — 80 users
	// Country mix: UK=25, Ireland=10, Malta=8, Germany=8, Italy=5, Spain=5,
	//              Finland=5, Sweden=4, Denmark=3, Netherlands=3, Other=4
	// =========================================================================
	type mtUser struct {
		name      string
		countryID int
	}
	maltaUsers := []mtUser{
		// UK (25)
		{"Oliver Bennett", countryUK}, {"Sophie Clarke", countryUK}, {"James Harrison", countryUK},
		{"Emily Watson", countryUK}, {"William Foster", countryUK}, {"Charlotte Hughes", countryUK},
		{"Benjamin Taylor", countryUK}, {"Amelia Robinson", countryUK}, {"George Mitchell", countryUK},
		{"Isla Anderson", countryUK}, {"Henry Thompson", countryUK}, {"Poppy Williams", countryUK},
		{"Freddie Davies", countryUK}, {"Lily Evans", countryUK}, {"Archie Walker", countryUK},
		{"Edward Turner", countryUK}, {"Florence Baker", countryUK}, {"Arthur Cooper", countryUK},
		{"Grace Phillips", countryUK}, {"Oscar Richards", countryUK}, {"Rosie Jenkins", countryUK},
		{"Theodore Shaw", countryUK}, {"Evie Barnes", countryUK}, {"Alfie Morgan", countryUK},
		{"Millie Griffin", countryUK},
		// Ireland (10)
		{"Sionan Murphy", countryIreland}, {"Aoife O'Brien", countryIreland}, {"Patrick Galway", countryIreland},
		{"Siobhan Limerick", countryIreland}, {"Brendan Doyle", countryIreland}, {"Ciara Lynch", countryIreland},
		{"Declan Kelly", countryIreland}, {"Niamh Walsh", countryIreland}, {"Cormac Ryan", countryIreland},
		{"Roisin Brennan", countryIreland},
		// Malta (8)
		{"Luca Farrugia", countryMalta}, {"Maria Calleja", countryMalta}, {"Johann Borg", countryMalta},
		{"Anna Vella", countryMalta}, {"Mark Camilleri", countryMalta}, {"Laura Azzopardi", countryMalta},
		{"Karl Galea", countryMalta}, {"Sandra Zammit", countryMalta},
		// Germany (8)
		{"Maximilian Weber", countryGermany}, {"Anna Schmidt", countryGermany}, {"Lukas Bauer", countryGermany},
		{"Lena Fischer", countryGermany}, {"Felix Wagner", countryGermany}, {"Sophia Richter", countryGermany},
		{"Jonas Klein", countryGermany}, {"Mia Hoffmann", countryGermany},
		// Italy (5)
		{"Marco Rossi", countryItaly}, {"Giulia Bianchi", countryItaly}, {"Alessandro Ferrari", countryItaly},
		{"Chiara Esposito", countryItaly}, {"Lorenzo Romano", countryItaly},
		// Spain (5)
		{"Alejandro Garcia", countrySpain}, {"Sofia Martinez", countrySpain}, {"Pablo Hernandez", countrySpain},
		{"Lucia Lopez", countrySpain}, {"Diego Fernandez", countrySpain},
		// Finland (5)
		{"Erik Nordstrom", countryFinland}, {"Astrid Virtanen", countryFinland}, {"Matti Korhonen", countryFinland},
		{"Hanna Lehtinen", countryFinland}, {"Jukka Makinen", countryFinland},
		// Sweden (4)
		{"Lars Johansson", countrySweden}, {"Ingrid Andersson", countrySweden},
		{"Nils Eriksson", countrySweden}, {"Elsa Lindqvist", countrySweden},
		// Denmark (3)
		{"Mikkel Andersen", countryDenmark}, {"Freja Nielsen", countryDenmark}, {"Kasper Christensen", countryDenmark},
		// Netherlands (3)
		{"Daan de Vries", countryNetherlands}, {"Emma Bakker", countryNetherlands}, {"Bram Visser", countryNetherlands},
		// Other Malta-regulated (4)
		{"Lukas Kovacs", countryHungary}, {"Eva Novak", countryCzechRepublic},
		{"Antanas Kazlauskas", countryLithuania}, {"Janis Ozolins", countryLatvia},
	}

	for i, u := range maltaUsers {
		specs = append(specs, userSpec{
			userID:    baseID + int64(i),
			userName:  u.name,
			countryID: int64(u.countryID),
			language:  languageForCountry(u.countryID),
		})
	}

	// =========================================================================
	// USA region — 50 users
	// State mix: NY=15, NJ=12, LA=8, CO=6, VA=5, PA=4
	// =========================================================================
	usNames := []string{
		// NY (15)
		"Liam Johnson", "Emma Davis", "Noah Miller", "Olivia Wilson", "William Moore",
		"Ava Taylor", "James Anderson", "Isabella Thomas", "Oliver Jackson", "Sophia White",
		"Elijah Harris", "Mia Martin", "Lucas Thompson", "Charlotte Garcia", "Mason Martinez",
		// NJ (12)
		"Amelia Robinson", "Ethan Clark", "Harper Rodriguez", "Alexander Lewis", "Evelyn Lee",
		"Daniel Walker", "Abigail Hall", "Michael Allen", "Emily Young", "Owen Hernandez",
		"Elizabeth King", "Sebastian Wright",
		// LA (8)
		"Sofia Lopez", "Jack Hill", "Camila Scott", "Aiden Green", "Luna Adams",
		"Logan Nelson", "Scarlett Carter", "Jackson Mitchell",
		// CO (6)
		"Riley Brooks", "Chloe Rivera", "Henry Cooper", "Aurora Sanchez", "Grayson Price",
		"Zoe Rogers",
		// VA (5)
		"Levi Morgan", "Hannah Bailey", "Wyatt Patterson", "Addison Howard", "Lincoln Ward",
		// PA (4)
		"Stella Butler", "Mateo Barnes", "Nora Ross", "Leo Simmons",
	}
	usStates := []int{
		stateNewYork, stateNewYork, stateNewYork, stateNewYork, stateNewYork,
		stateNewYork, stateNewYork, stateNewYork, stateNewYork, stateNewYork,
		stateNewYork, stateNewYork, stateNewYork, stateNewYork, stateNewYork,
		stateNewJersey, stateNewJersey, stateNewJersey, stateNewJersey, stateNewJersey,
		stateNewJersey, stateNewJersey, stateNewJersey, stateNewJersey, stateNewJersey,
		stateNewJersey, stateNewJersey,
		stateLouisiana, stateLouisiana, stateLouisiana, stateLouisiana, stateLouisiana,
		stateLouisiana, stateLouisiana, stateLouisiana,
		stateColorado, stateColorado, stateColorado, stateColorado, stateColorado,
		stateColorado,
		stateVirginia, stateVirginia, stateVirginia, stateVirginia, stateVirginia,
		statePennsylv, statePennsylv, statePennsylv, statePennsylv,
	}

	for i := 0; i < 50; i++ {
		sv := int64(usStates[i])
		specs = append(specs, userSpec{
			userID:    baseID + 100 + int64(i),
			userName:  usNames[i],
			countryID: countryUS,
			stateID:   &sv,
			language:  "EN",
		})
	}

	// =========================================================================
	// AUSTRALIA region — 35 users
	// State mix: NSW=12, VIC=10, QLD=6, SA=4, WA=3
	// =========================================================================
	auNames := []string{
		// NSW (12)
		"Lachlan Smith", "Chloe Jones", "Jack Williams", "Ella Brown", "Noah Wilson",
		"Olivia Taylor", "Thomas Davis", "Ava Anderson", "Ethan Martin", "Sophie Thompson",
		"Lucas White", "Grace Johnson",
		// VIC (10)
		"Cooper Harris", "Zoe Lewis", "Joshua Robinson", "Emma Walker", "Riley Hall",
		"Mia Young", "Bailey King", "Isla Wright", "William Scott", "Amelia Green",
		// QLD (6)
		"Aiden Adams", "Charlotte Baker", "Nathan Carter", "Georgia Palmer", "Caleb Murray",
		"Phoebe Turner",
		// SA (4)
		"Hugo Maxwell", "Willow Simpson", "Archer Grant", "Ivy Coleman",
		// WA (3)
		"Finn Marshall", "Ruby Chapman", "Xavier Knight",
	}
	auStates := []int{
		stateNSW, stateNSW, stateNSW, stateNSW, stateNSW, stateNSW,
		stateNSW, stateNSW, stateNSW, stateNSW, stateNSW, stateNSW,
		stateVIC, stateVIC, stateVIC, stateVIC, stateVIC,
		stateVIC, stateVIC, stateVIC, stateVIC, stateVIC,
		stateQLD, stateQLD, stateQLD, stateQLD, stateQLD, stateQLD,
		stateSA, stateSA, stateSA, stateSA,
		stateWA, stateWA, stateWA,
	}

	for i := 0; i < 35; i++ {
		sv := int64(auStates[i])
		specs = append(specs, userSpec{
			userID:    baseID + 200 + int64(i),
			userName:  auNames[i],
			countryID: countryAustralia,
			stateID:   &sv,
			language:  "EN",
		})
	}

	// =========================================================================
	// GIBRALTAR region — 35 users
	// Country mix: Canada=8, Brazil=7, Gibraltar=4, UAE=4, SouthAfrica=3,
	//              Nigeria=3, Argentina=2, Turkey=2, Chile=2
	// =========================================================================
	type gibUser struct {
		name      string
		countryID int
	}
	gibUsers := []gibUser{
		// Canada (8)
		{"Carlos Rodriguez", countryCanada}, {"Jean-Pierre Dupont", countryCanada},
		{"Marie Lefebvre", countryCanada}, {"Robert Campbell", countryCanada},
		{"Sarah Mackenzie", countryCanada}, {"David Chen", countryCanada},
		{"Jennifer Okafor", countryCanada}, {"Michael Tremblay", countryCanada},
		// Brazil (7)
		{"Pedro Alves", countryBrazil}, {"Sofia Costa", countryBrazil},
		{"Rafael Santos", countryBrazil}, {"Ana Oliveira", countryBrazil},
		{"Lucas Ferreira", countryBrazil}, {"Isabela Lima", countryBrazil},
		{"Matheus Silva", countryBrazil},
		// Gibraltar (4)
		{"Sebastian Garcia", countryGibraltar}, {"Laura Fischer", countryGibraltar},
		{"Thomas Baglietto", countryGibraltar}, {"Claire Montegriffo", countryGibraltar},
		// UAE (4)
		{"Mohammed Al-Hassan", countryUAE}, {"Fatima Al-Rashid", countryUAE},
		{"Ahmed Al-Maktoum", countryUAE}, {"Layla Khalid", countryUAE},
		// South Africa (3)
		{"Thabo Molefe", countrySouthAfrica}, {"Zanele Dlamini", countrySouthAfrica},
		{"Johan van der Merwe", countrySouthAfrica},
		// Nigeria (3)
		{"Chukwuma Okonkwo", countryNigeria}, {"Aisha Bello", countryNigeria},
		{"Emeka Nwankwo", countryNigeria},
		// Argentina (2)
		{"Mateo Fernandez", countryArgentina}, {"Valentina Moreno", countryArgentina},
		// Turkey (2)
		{"Mehmet Yilmaz", countryTurkey}, {"Ayse Demir", countryTurkey},
		// Chile (2)
		{"Ignacio Vargas", countryChile}, {"Catalina Rojas", countryChile},
	}

	for i, u := range gibUsers {
		specs = append(specs, userSpec{
			userID:    baseID + 300 + int64(i),
			userName:  u.name,
			countryID: int64(u.countryID),
			language:  languageForCountry(u.countryID),
		})
	}

	return specs
}

// ---------------------------------------------------------------------------
// 6. USER-FOLDER ASSIGNMENTS — each user gets 1-3 folders from their region
// ---------------------------------------------------------------------------

func generateUserFolderAssignments(users []UserStatus, fbr map[string][]string) []UserCaseFolder {
	var assignments []UserCaseFolder
	assigned := make(map[string]struct{})

	updaters := []string{
		"admin@company.com", "compliance@company.com", "system@company.com",
	}

	for i, u := range users {
		reg := regionForCountry(int(u.CountryID))
		folders := fbr[reg]
		if len(folders) == 0 {
			continue
		}

		// 1-3 folders per user (deterministic)
		numFolders := 1 + (i % 3)
		if numFolders > len(folders) {
			numFolders = len(folders)
		}

		startIdx := (i * 7) % len(folders) // spread across folders more evenly
		for j := 0; j < numFolders; j++ {
			fIdx := (startIdx + j) % len(folders)
			folderPK := folders[fIdx]

			key := u.PK + "|" + folderPK
			if _, exists := assigned[key]; exists {
				continue
			}
			assigned[key] = struct{}{}

			day := 1 + ((i + j) % 59)
			month := 1
			if day > 31 {
				month = 2
				day -= 31
			}

			assignments = append(assignments, UserCaseFolder{
				PK:                     newUUID(),
				CaseManagementFolderPK: folderPK,
				UserStatusPK:           u.PK,
				LoggedAt:               ts(2026, month, day, 10+(i%8), ((i+j)*11)%60),
				UpdatedBy:              pick(updaters, i+j),
			})
		}
	}
	return assignments
}

// ---------------------------------------------------------------------------
// File writing
// ---------------------------------------------------------------------------

func writeJSON(path string, data interface{}) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal %s: %w", path, err)
	}
	if err := os.WriteFile(path, append(b, '\n'), 0644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

func main() {
	const dataDir = "data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "mkdir data: %v\n", err)
		os.Exit(1)
	}

	// 1. Case folders
	folders := generateFolders()
	must(writeJSON(dataDir+"/case_folders.json", folders))
	fmt.Printf("case_folders.json        — %d records\n", len(folders))

	fbr := foldersByRegion(folders)

	// 2. Threshold configs
	thresholds := generateThresholds(fbr)
	must(writeJSON(dataDir+"/threshold_configs.json", thresholds))
	fmt.Printf("threshold_configs.json   — %d records\n", len(thresholds))

	// 3. Multiplier configs
	multipliers := generateMultipliers()
	must(writeJSON(dataDir+"/multiplier_configs.json", multipliers))
	fmt.Printf("multiplier_configs.json  — %d records\n", len(multipliers))

	// 4. Business profiles
	profiles := generateBusinessProfiles()
	must(writeJSON(dataDir+"/business_profiles.json", profiles))
	fmt.Printf("business_profiles.json   — %d records\n", len(profiles))

	// 5. User statuses
	users := generateUsers()
	must(writeJSON(dataDir+"/user_statuses.json", users))
	fmt.Printf("user_statuses.json       — %d records\n", len(users))

	// 6. User-folder assignments
	assignments := generateUserFolderAssignments(users, fbr)
	must(writeJSON(dataDir+"/user_case_folders.json", assignments))
	fmt.Printf("user_case_folders.json   — %d records\n", len(assignments))

	fmt.Println()
	fmt.Printf("Folders by region: MALTA=%d, USA=%d, AUSTRALIA=%d, GIBRALTAR=%d\n",
		len(fbr[regionMalta]), len(fbr[regionUSA]), len(fbr[regionAustralia]), len(fbr[regionGibraltar]))
	fmt.Println()

	validate(folders, thresholds, multipliers, profiles, users, assignments)
}

func must(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// ---------------------------------------------------------------------------
// Validation — checks referential integrity and prints summary
// ---------------------------------------------------------------------------

func validate(
	folders []CaseFolder,
	thresholds []ThresholdConfig,
	multipliers []MultiplierConfig,
	profiles []BusinessProfile,
	users []UserStatus,
	assignments []UserCaseFolder,
) {
	fmt.Println("Validation:")
	errors := 0

	// Index folders
	folderPKs := make(map[string]string)
	for _, f := range folders {
		folderPKs[f.PK] = f.Region
	}

	// Index multipliers by country|state
	multIndex := make(map[string]bool)
	for _, m := range multipliers {
		key := fmt.Sprintf("%d|", m.CountryID)
		if m.StateID != nil {
			key = fmt.Sprintf("%d|%d", m.CountryID, *m.StateID)
		}
		multIndex[key] = true
	}

	// Index business profiles
	bpCountries := make(map[int64]bool)
	for _, p := range profiles {
		bpCountries[p.CountryID] = true
	}

	// Index users
	userRegions := make(map[string]string)
	for _, u := range users {
		userRegions[u.PK] = regionForCountry(int(u.CountryID))
	}

	// Check: all 4 regions have folders
	for _, reg := range []string{regionMalta, regionUSA, regionAustralia, regionGibraltar} {
		count := 0
		for _, f := range folders {
			if f.Region == reg {
				count++
			}
		}
		if count == 0 {
			fmt.Printf("  ERROR: region %s has no folders\n", reg)
			errors++
		} else {
			fmt.Printf("  OK: %-12s %d folders\n", reg, count)
		}
	}

	// Check: no UK region
	for _, f := range folders {
		if f.Region == "UK" {
			fmt.Printf("  ERROR: folder '%s' has UK region (should be MALTA)\n", f.FolderName)
			errors++
		}
	}

	// Check: threshold folder references are same-region
	for _, t := range thresholds {
		if t.CaseManagementFolderPK == nil {
			continue
		}
		fReg, ok := folderPKs[*t.CaseManagementFolderPK]
		if !ok {
			fmt.Printf("  ERROR: threshold '%s' references unknown folder PK\n", t.Title)
			errors++
			continue
		}
		tReg := regionForCountry(int(t.CountryID))
		if fReg != tReg {
			fmt.Printf("  ERROR: threshold '%s' (region %s) -> folder in region %s\n", t.Title, tReg, fReg)
			errors++
		}
	}

	// Check: hierarchy uniqueness
	hierarchies := make(map[int64]int)
	for _, t := range thresholds {
		hierarchies[t.Hierarchy]++
	}
	for h, count := range hierarchies {
		if count > 1 {
			fmt.Printf("  ERROR: hierarchy %d appears %d times\n", h, count)
			errors++
		}
	}

	// Check: user-folder assignments are same-region
	crossRegion := 0
	for _, a := range assignments {
		fReg, ok1 := folderPKs[a.CaseManagementFolderPK]
		uReg, ok2 := userRegions[a.UserStatusPK]
		if !ok1 {
			fmt.Printf("  ERROR: assignment references unknown folder PK\n")
			errors++
			continue
		}
		if !ok2 {
			fmt.Printf("  ERROR: assignment references unknown user PK\n")
			errors++
			continue
		}
		if fReg != uReg {
			crossRegion++
			errors++
		}
	}
	if crossRegion > 0 {
		fmt.Printf("  ERROR: %d cross-region user-folder assignments\n", crossRegion)
	}

	// Stats: ECDD status distribution
	statusDist := make(map[int64]int)
	for _, u := range users {
		statusDist[u.ECDDStatus]++
	}
	fmt.Printf("  ECDD status distribution:")
	for s := int64(1); s <= 7; s++ {
		pct := float64(statusDist[s]) / float64(len(users)) * 100
		fmt.Printf(" %d=%d(%.0f%%)", s, statusDist[s], math.Round(pct))
	}
	fmt.Println()

	// Stats: assignments per region
	aByReg := make(map[string]int)
	for _, a := range assignments {
		if r, ok := folderPKs[a.CaseManagementFolderPK]; ok {
			aByReg[r]++
		}
	}
	fmt.Printf("  Assignments by region: MALTA=%d, USA=%d, AUSTRALIA=%d, GIBRALTAR=%d (total=%d)\n",
		aByReg[regionMalta], aByReg[regionUSA], aByReg[regionAustralia], aByReg[regionGibraltar],
		len(assignments))

	// Stats: folder coverage
	usedFolders := make(map[string]int)
	for _, a := range assignments {
		usedFolders[a.CaseManagementFolderPK]++
	}
	emptyFolders := 0
	for _, f := range folders {
		if usedFolders[f.PK] == 0 {
			emptyFolders++
		}
	}
	fmt.Printf("  Folder coverage: %d/%d folders have assignments (%d empty)\n",
		len(usedFolders), len(folders), emptyFolders)

	// Stats: users per region
	usersByReg := make(map[string]int)
	for _, u := range users {
		usersByReg[regionForCountry(int(u.CountryID))]++
	}
	fmt.Printf("  Users by region: MALTA=%d, USA=%d, AUSTRALIA=%d, GIBRALTAR=%d\n",
		usersByReg[regionMalta], usersByReg[regionUSA], usersByReg[regionAustralia], usersByReg[regionGibraltar])

	if errors == 0 {
		fmt.Println("  All checks passed.")
	} else {
		fmt.Printf("  %d error(s) found.\n", errors)
		os.Exit(1)
	}
}
