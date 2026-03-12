package region

import "amltoolmock/enums/country"

// Region represents a regulatory region for ECDD case management.
// There are 4 regions: Malta, Gibraltar, USA, Australia.
// UK is NOT a separate region — UK falls under Malta (licensed countries).
type Region string

const (
	Malta     Region = "MALTA"
	Gibraltar Region = "GIBRALTAR"
	USA       Region = "USA"
	Australia Region = "AUSTRALIA"
)

// AllRegions returns all valid regions
func AllRegions() []Region {
	return []Region{Malta, Gibraltar, USA, Australia}
}

// malteseRegulatedCountries contains Maltese regulated country IDs (from production enum)
var malteseRegulatedCountries = map[int]struct{}{
	country.Austria: {}, country.Croatia: {}, country.Finland: {}, country.Ghana: {},
	country.Hungary: {}, country.Iceland: {}, country.IrelandRepOf: {}, country.IvoryCoast: {},
	country.India: {}, country.Kenya: {}, country.Latvia: {}, country.Liechtenstein: {},
	country.Lithuania: {}, country.Luxembourg: {}, country.Malta: {}, country.NewZealand: {},
	country.Nicaragua: {}, country.Norway: {}, country.Slovakia: {}, country.Slovenia: {},
	country.Tanzania: {}, country.Tonga: {}, country.WesternSamoa: {},
}

// malteseLicensedCountries contains countries licensed under Malta jurisdiction.
// These countries show in the Malta case management area but are not "regulated" by Malta.
//
// European Regulated Markets (12):
//   UK, Spain, Italy, Denmark, Greenland, Sweden, Estonia, Bulgaria,
//   Cyprus, Greece, Germany, Netherlands, Czech Republic
// Americas (4):
//   Mexico, Canada Ontario, Argentina BAC, Argentina BAP
// Other Markets (3):
//   France, Switzerland, Japan
var malteseLicensedCountries = map[int]struct{}{
	// European Regulated Markets
	country.UK:              {},
	country.Spain:           {},
	country.Italy:           {},
	country.Denmark:         {},
	country.Greenland:       {},
	country.Sweden:          {},
	country.Estonia:         {},
	country.Bulgaria:        {},
	country.Cyprus:          {},
	country.Greece:          {},
	country.Germany:         {},
	country.Netherlands:     {},
	country.CzechRepublic:   {},
	// Americas
	country.Mexico:          {},
	country.CanadaOntario:   {},
	country.BuenosAiresCity: {},
	country.BuenosAiresProvince: {},
	// Other Markets
	country.France:          {},
	country.Switzerland:     {},
	country.Japan:           {},
}

// gibraltarRegulatedCountries contains all Gibraltar regulated country IDs
var gibraltarRegulatedCountries = map[int]struct{}{
	country.Albania: {}, country.Algeria: {}, country.Andorra: {}, country.Anguilla: {},
	country.AntiguaAndBarbuda: {}, country.Argentina: {}, country.Armenia: {}, country.Aruba: {},
	country.Azerbaijan: {}, country.Bahamas: {}, country.Bahrain: {}, country.Bangladesh: {},
	country.Barbados: {}, country.Belarus: {}, country.Belize: {}, country.Benin: {},
	country.Bermuda: {}, country.Bolivia: {}, country.BosniaHerzegovina: {}, country.Botswana: {},
	country.Brazil: {}, country.BritishVirginIslands: {}, country.BruneiDarussalam: {},
	country.BurkinaFaso: {}, country.Cameroon: {}, country.Canada: {}, country.CapeVerdeIslands: {},
	country.CaymanIslands: {}, country.CentralAfricanRepublic: {}, country.Chile: {},
	country.CookIslands: {}, country.CostaRica: {}, country.Djibouti: {}, country.Dominica: {},
	country.DominicanRepublic: {}, country.Ecuador: {}, country.Egypt: {}, country.ElSalvador: {},
	country.Ethiopia: {}, country.FaroeIslands: {}, country.Fiji: {}, country.FrenchPolynesia: {},
	country.Gabon: {}, country.Gambia: {}, country.Georgia: {}, country.Gibraltar: {},
	country.Grenada: {}, country.Guatemala: {}, country.Guinea: {}, country.Guyana: {},
	country.Honduras: {}, country.Indonesia: {}, country.Jamaica: {},
	country.Jordan: {}, country.Kazakhstan: {}, country.SouthKorea: {}, country.Kuwait: {},
	country.Kyrgyzstan: {}, country.Laos: {}, country.Lebanon: {}, country.Lesotho: {},
	country.Liberia: {}, country.Macedonia: {}, country.Madagascar: {}, country.Malawi: {},
	country.Malaysia: {}, country.Maldives: {}, country.Mali: {}, country.Mauritania: {},
	country.Mauritius: {}, country.Moldova: {}, country.Mongolia: {}, country.Montenegro: {},
	country.Montserrat: {}, country.Morocco: {}, country.Mozambique: {}, country.Namibia: {},
	country.Nepal: {}, country.NetherlandsAntilles: {}, country.NewCaledonia: {}, country.Niger: {},
	country.Nigeria: {}, country.Oman: {}, country.Pakistan: {}, country.Palestine: {},
	country.Panama: {}, country.PapuaNewGuinea: {}, country.Paraguay: {}, country.Peru: {},
	country.Qatar: {}, country.Rwanda: {}, country.SanMarino: {}, country.SaoTomeEPrincipe: {},
	country.SaudiArabia: {}, country.Senegal: {}, country.Serbia: {}, country.Seychelles: {},
	country.SierraLeone: {}, country.SolomonIslands: {}, country.SriLanka: {},
	country.StKittsAndNevis: {}, country.StLucia: {}, country.StVincentAndTheGrenadines: {},
	country.Suriname: {}, country.Swaziland: {}, country.Taiwan: {},
	country.Thailand: {}, country.Togo: {}, country.TrinidadAndTobago: {}, country.Tunisia: {},
	country.TurksAndCaicosIslands: {}, country.Uganda: {}, country.Ukraine: {},
	country.UnitedArabEmirates: {}, country.Uruguay: {}, country.Vanuatu: {}, country.Vietnam: {},
	country.Zambia: {},
}

// IsMalteseRegulatedCountry returns true if the countryID is regulated under Maltese jurisdiction
func IsMalteseRegulatedCountry(countryID int) bool {
	_, found := malteseRegulatedCountries[countryID]
	return found
}

// IsMalteseLicensedCountry returns true if the countryID is licensed under Maltese jurisdiction
func IsMalteseLicensedCountry(countryID int) bool {
	_, found := malteseLicensedCountries[countryID]
	return found
}

// IsMalteseCountry returns true if the countryID belongs to Malta region
// (either regulated or licensed)
func IsMalteseCountry(countryID int) bool {
	return IsMalteseRegulatedCountry(countryID) || IsMalteseLicensedCountry(countryID)
}

// IsGibraltarRegulatedCountry returns true if the countryID is regulated under Gibraltar jurisdiction
func IsGibraltarRegulatedCountry(countryID int) bool {
	_, found := gibraltarRegulatedCountries[countryID]
	return found
}

// IsUSRegulated returns true if the countryID is the US
func IsUSRegulated(countryID int) bool {
	return countryID == country.US
}

// IsAustraliaRegulated returns true if the countryID is Australia
func IsAustraliaRegulated(countryID int) bool {
	return countryID == country.Australia
}

// GetRegionForCountry returns the regulatory region for a given country ID.
// Returns empty string if the country does not belong to any known region.
func GetRegionForCountry(countryID int) Region {
	if IsUSRegulated(countryID) {
		return USA
	}
	if IsAustraliaRegulated(countryID) {
		return Australia
	}
	if IsMalteseCountry(countryID) {
		return Malta
	}
	if IsGibraltarRegulatedCountry(countryID) {
		return Gibraltar
	}
	return ""
}

// GetCountriesForRegion returns all country IDs that belong to a given region
func GetCountriesForRegion(r Region) []int {
	switch r {
	case USA:
		return []int{country.US}
	case Australia:
		return []int{country.Australia}
	case Malta:
		// Combine regulated + licensed countries
		countries := make([]int, 0, len(malteseRegulatedCountries)+len(malteseLicensedCountries))
		for id := range malteseRegulatedCountries {
			countries = append(countries, id)
		}
		for id := range malteseLicensedCountries {
			countries = append(countries, id)
		}
		return countries
	case Gibraltar:
		countries := make([]int, 0, len(gibraltarRegulatedCountries))
		for id := range gibraltarRegulatedCountries {
			countries = append(countries, id)
		}
		return countries
	default:
		return nil
	}
}
