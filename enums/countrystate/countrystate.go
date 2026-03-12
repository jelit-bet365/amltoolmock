package countrystate

import "fmt"

const (
	Alabama                           = 1
	NewYork                           = 2
	NewJersey                         = 3
	Alaska                            = 7
	AmericanSamoa                     = 8
	Arizona                           = 9
	Arkansas                          = 10
	ArmedForcesEurope                 = 11
	ArmedForcesPacific                = 13
	ArmedForcesTheAmericas            = 14
	California                        = 15
	Colorado                          = 16
	Connecticut                       = 18
	Delaware                          = 19
	DistrictofColumbia                = 20
	FederatedStatesofMicronesia       = 21
	Florida                           = 22
	Georgia                           = 23
	Guam                              = 24
	Hawaii                            = 26
	Idaho                             = 27
	Illinois                          = 28
	Indiana                           = 29
	Iowa                              = 30
	Kansas                            = 31
	Kentucky                          = 32
	Louisiana                         = 33
	Maine                             = 34
	MarshallIslands                   = 35
	Maryland                          = 36
	Massachusetts                     = 38
	Michigan                          = 39
	Minnesota                         = 40
	Mississippi                       = 41
	Missouri                          = 43
	Montana                           = 44
	Nebraska                          = 45
	Nevada                            = 46
	NewHampshire                      = 47
	NewMexico                         = 48
	NorthCarolina                     = 49
	NorthDakota                       = 50
	NorthernMarianaIslands            = 51
	Ohio                              = 52
	Oklahoma                          = 53
	Oregon                            = 54
	Palau                             = 55
	Pennsylvania                      = 56
	PuertoRico                        = 57
	RhodeIsland                       = 58
	SouthCarolina                     = 59
	SouthDakota                       = 60
	Tennessee                         = 61
	Texas                             = 62
	Utah                              = 63
	Vermont                           = 64
	VirginIslands                     = 65
	Virginia                          = 66
	Wisconsin                         = 67
	Wyoming                           = 68
	Alberta                           = 69
	BritishColumbia                   = 70
	Manitoba                          = 71
	NewBrunswick                      = 72
	NewFoundland                      = 73
	NovaScotia                        = 74
	NorthwestTerritories              = 75
	Nunavut                           = 76
	OntarioState                      = 77
	PrinceEdwardIsland                = 78
	Quebec                            = 79
	Saskatchewan                      = 80
	Yukon                             = 81
	EasternCape                       = 82
	FreeState                         = 83
	Guateng                           = 84
	KwaZuluNatal                      = 85
	Mpumalamanga                      = 86
	NorthCape                         = 87
	NorthernProvince                  = 88
	NorthWest                         = 89
	WesternCape                       = 90
	AustralianCapitalTerritory        = 91
	NewSouthWales                     = 93
	NorthernTerritory                 = 94
	Queensland                        = 95
	SouthAustralia                    = 96
	Tasmania                          = 97
	Victoria                          = 98
	WesternAustralia                  = 99
	Washington                        = 100
	WestVirginia                      = 101
	Ontario                           = 510
)

// US states
var stateNames = map[int]string{
	1: "Alabama", 2: "New York", 3: "New Jersey", 7: "Alaska", 8: "American Samoa",
	9: "Arizona", 10: "Arkansas", 11: "Armed Forces - Europe",
	13: "Armed Forces - Pacific", 14: "Armed Forces - The Americas",
	15: "California", 16: "Colorado", 18: "Connecticut", 19: "Delaware",
	20: "District of Columbia", 21: "Federated States of Micronesia",
	22: "Florida", 23: "Georgia", 24: "Guam", 26: "Hawaii", 27: "Idaho",
	28: "Illinois", 29: "Indiana", 30: "Iowa", 31: "Kansas", 32: "Kentucky",
	33: "Louisiana", 34: "Maine", 35: "Marshall Islands", 36: "Maryland",
	38: "Massachusetts", 39: "Michigan", 40: "Minnesota", 41: "Mississippi",
	43: "Missouri", 44: "Montana", 45: "Nebraska", 46: "Nevada",
	47: "New Hampshire", 48: "New Mexico", 49: "North Carolina",
	50: "North Dakota", 51: "Northern Mariana Islands", 52: "Ohio",
	53: "Oklahoma", 54: "Oregon", 55: "Palau", 56: "Pennsylvania",
	57: "Puerto Rico", 58: "Rhode Island", 59: "South Carolina",
	60: "South Dakota", 61: "Tennessee", 62: "Texas", 63: "Utah",
	64: "Vermont", 65: "Virgin Islands", 66: "Virginia", 67: "Wisconsin",
	68: "Wyoming", 100: "Washington", 101: "West Virginia",
}

var shortStateNames = map[int]string{
	1: "AL", 2: "NY", 3: "NJ", 7: "AK", 8: "AS", 9: "AZ", 10: "AR", 11: "AE",
	13: "AP", 14: "AA", 15: "CA", 16: "CO", 18: "CT", 19: "DE", 20: "DC",
	21: "FM", 22: "FL", 23: "GA", 24: "GU", 26: "HI", 27: "ID", 28: "IL",
	29: "IN", 30: "IA", 31: "KS", 32: "KY", 33: "LA", 34: "ME", 35: "MH",
	36: "MD", 38: "MA", 39: "MI", 40: "MN", 41: "MS", 43: "MO", 44: "MT",
	45: "NE", 46: "NV", 47: "NH", 48: "NM", 49: "NC", 50: "ND", 51: "MP",
	52: "OH", 53: "OK", 54: "OR", 55: "PW", 56: "PA", 57: "PR", 58: "RI",
	59: "SC", 60: "SD", 61: "TN", 62: "TX", 63: "UT", 64: "VT", 65: "VI",
	66: "VA", 67: "WI", 68: "WY", 100: "WA", 101: "WV",
}

// GetName returns the user-friendly state name
func GetName(stateID int) (string, error) {
	name, ok := stateNames[stateID]
	if ok {
		return name, nil
	}
	return "", fmt.Errorf("message='unable to find state name for %v'", stateID)
}

// GetShortName returns the 2 character state name
func GetShortName(stateID int) (string, error) {
	name, ok := shortStateNames[stateID]
	if ok {
		return name, nil
	}
	return "", fmt.Errorf("message='unable to find short state name for %v'", stateID)
}
