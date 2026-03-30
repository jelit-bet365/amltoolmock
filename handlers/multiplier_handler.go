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

// GetAllMultipliers handles GET /api/ecdd/multiplierconfig
func GetAllMultipliers(w http.ResponseWriter, r *http.Request) {
	ds := services.GetDataService()
	allMultipliers := ds.GetAllMultiplierConfigs()

	// Parse optional query params
	query := r.URL.Query()
	isActiveParam := query.Get("isActive")
	countryParam := query.Get("countryId")

	// Filter multipliers based on query params
	filtered := make([]*models.ECDDMultiplierConfig, 0, len(allMultipliers))
	for _, m := range allMultipliers {
		if isActiveParam != "" {
			isActive, err := strconv.ParseBool(isActiveParam)
			if err == nil && m.IsActive != isActive {
				continue
			}
		}
		if countryParam != "" {
			countryID, err := strconv.ParseInt(countryParam, 10, 64)
			if err == nil && m.CountryID != countryID {
				continue
			}
		}
		filtered = append(filtered, m)
	}

	// Sort the results
	sp := utils.GetSortParams(r, "ecddMultiplierConfigPk")
	sortMultipliers(filtered, sp)

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

// sortMultipliers sorts a slice of ECDDMultiplierConfig by the given sort parameters.
func sortMultipliers(multipliers []*models.ECDDMultiplierConfig, sp utils.SortParams) {
	desc := sp.IsDescending()
	sort.Slice(multipliers, func(i, j int) bool {
		var less bool
		switch sp.SortBy {
		case "ecddMultiplierConfigPk":
			less = multipliers[i].ECDDMultiplierConfigPK < multipliers[j].ECDDMultiplierConfigPK
		case "countryId":
			less = multipliers[i].CountryID < multipliers[j].CountryID
		case "loggedAt":
			less = multipliers[i].LoggedAt.Before(multipliers[j].LoggedAt)
		default:
			less = multipliers[i].ECDDMultiplierConfigPK < multipliers[j].ECDDMultiplierConfigPK
		}
		if desc {
			return !less
		}
		return less
	})
}

// GetMultiplierByID handles GET /api/ecdd/multiplierconfig/{id}
func GetMultiplierByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/multiplierconfig/")
	ds := services.GetDataService()
	multiplier := ds.GetMultiplierConfigByID(id)

	if multiplier == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Multiplier not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(multiplier)
}

// CreateMultiplier handles POST /api/ecdd/multiplierconfig
func CreateMultiplier(w http.ResponseWriter, r *http.Request) {
	var multiplier models.ECDDMultiplierConfig
	if err := json.NewDecoder(r.Body).Decode(&multiplier); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ds := services.GetDataService()
	createdMultiplier := ds.CreateMultiplierConfig(&multiplier)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdMultiplier)
}

// UpdateMultiplier handles PUT /api/ecdd/multiplierconfig/{id}
func UpdateMultiplier(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/multiplierconfig/")
	var multiplier models.ECDDMultiplierConfig
	if err := json.NewDecoder(r.Body).Decode(&multiplier); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ds := services.GetDataService()
	updatedMultiplier := ds.UpdateMultiplierConfig(id, &multiplier)

	if updatedMultiplier == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Multiplier not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedMultiplier)
}

// PatchMultiplier handles PATCH /api/ecdd/multiplierconfig/{id}
// Performs a partial update by merging the request body onto the existing record
func PatchMultiplier(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/multiplierconfig/")

	ds := services.GetDataService()
	existing := ds.GetMultiplierConfigByID(id)
	if existing == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Multiplier not found")
		return
	}

	// Decode partial update onto the existing record — only provided fields are overwritten
	if err := json.NewDecoder(r.Body).Decode(existing); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updated := ds.UpdateMultiplierConfig(id, existing)
	if updated == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Multiplier not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

// DeleteMultiplier handles DELETE /api/ecdd/multiplierconfig/{id}
func DeleteMultiplier(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/multiplierconfig/")

	ds := services.GetDataService()
	success := ds.DeleteMultiplierConfig(id)

	if !success {
		utils.WriteJSONError(w, http.StatusNotFound, "Multiplier not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Multiplier soft-deleted successfully"})
}
