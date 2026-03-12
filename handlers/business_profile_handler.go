package handlers

import (
	"amltoolmock/models"
	"amltoolmock/services"
	"amltoolmock/utils"
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// GetAllBusinessProfiles handles GET /api/ecdd/businessprofile
func GetAllBusinessProfiles(w http.ResponseWriter, r *http.Request) {
	ds := services.GetDataService()
	allProfiles := ds.GetAllBusinessProfiles()

	// Parse optional query params
	query := r.URL.Query()
	enabledParam := query.Get("enabled")
	countryParam := query.Get("country_id")
	riskStatusParam := query.Get("risk_status_id")

	// Filter business profiles based on query params
	filtered := make([]*models.ECDDBusinessProfile, 0, len(allProfiles))
	for _, p := range allProfiles {
		if enabledParam != "" {
			enabled, err := strconv.ParseBool(enabledParam)
			if err == nil && p.Enabled != enabled {
				continue
			}
		}
		if countryParam != "" {
			countryID, err := strconv.ParseInt(countryParam, 10, 64)
			if err == nil && p.CountryID != countryID {
				continue
			}
		}
		if riskStatusParam != "" {
			riskStatusID, err := strconv.ParseInt(riskStatusParam, 10, 64)
			if err == nil && p.RiskStatusID != riskStatusID {
				continue
			}
		}
		filtered = append(filtered, p)
	}

	// Sort the results
	sp := utils.GetSortParams(r, "ecdd_business_profile_pk")
	sortBusinessProfiles(filtered, sp)

	// Optional pagination
	pp := utils.GetPaginationParams(r)
	if pp.Enabled {
		total := len(filtered)
		start, end := utils.CalculateOffset(pp.Page, pp.PageSize, total)
		paged := filtered[start:end]
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(utils.Paginate(paged, pp, total))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

// sortBusinessProfiles sorts a slice of ECDDBusinessProfile by the given sort parameters.
func sortBusinessProfiles(profiles []*models.ECDDBusinessProfile, sp utils.SortParams) {
	desc := sp.IsDescending()
	sort.Slice(profiles, func(i, j int) bool {
		var less bool
		switch sp.SortBy {
		case "ecdd_business_profile_pk":
			less = profiles[i].ECDDBusinessProfilePK < profiles[j].ECDDBusinessProfilePK
		case "country_id":
			less = profiles[i].CountryID < profiles[j].CountryID
		case "risk_status_id":
			less = profiles[i].RiskStatusID < profiles[j].RiskStatusID
		case "average_deposit":
			less = profiles[i].AverageDeposit < profiles[j].AverageDeposit
		case "deposit_multiplier":
			less = profiles[i].DepositMultiplier < profiles[j].DepositMultiplier
		case "logged_at":
			less = profiles[i].LoggedAt.Before(profiles[j].LoggedAt)
		default:
			less = profiles[i].ECDDBusinessProfilePK < profiles[j].ECDDBusinessProfilePK
		}
		if desc {
			return !less
		}
		return less
	})
}

// GetBusinessProfileByID handles GET /api/ecdd/businessprofile/{id}
func GetBusinessProfileByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/businessprofile/")
	ds := services.GetDataService()
	profile := ds.GetBusinessProfileByID(id)

	if profile == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Business profile not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// CreateBusinessProfile handles POST /api/ecdd/businessprofile
func CreateBusinessProfile(w http.ResponseWriter, r *http.Request) {
	var profile models.ECDDBusinessProfile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ds := services.GetDataService()
	createdProfile := ds.CreateBusinessProfile(&profile)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProfile)
}

// UpdateBusinessProfile handles PUT /api/ecdd/businessprofile/{id}
func UpdateBusinessProfile(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/businessprofile/")
	var profile models.ECDDBusinessProfile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ds := services.GetDataService()
	updatedProfile := ds.UpdateBusinessProfile(id, &profile)

	if updatedProfile == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Business profile not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProfile)
}

// PatchBusinessProfile handles PATCH /api/ecdd/businessprofile/{id}
// Performs a partial update by merging the request body onto the existing record
func PatchBusinessProfile(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/businessprofile/")

	ds := services.GetDataService()
	existing := ds.GetBusinessProfileByID(id)
	if existing == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Business profile not found")
		return
	}

	// Decode partial update onto the existing record — only provided fields are overwritten
	if err := json.NewDecoder(r.Body).Decode(existing); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updated := ds.UpdateBusinessProfile(id, existing)
	if updated == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Business profile not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

// DeleteBusinessProfile handles DELETE /api/ecdd/businessprofile/{id}
func DeleteBusinessProfile(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/businessprofile/")

	ds := services.GetDataService()
	success := ds.DeleteBusinessProfile(id)

	if !success {
		utils.WriteJSONError(w, http.StatusNotFound, "Business profile not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Business profile soft-deleted successfully"})
}
