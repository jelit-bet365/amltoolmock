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

// GetAllThresholds handles GET /api/ecdd/thresholdconfig
func GetAllThresholds(w http.ResponseWriter, r *http.Request) {
	ds := services.GetDataService()
	allThresholds := ds.GetAllThresholdConfigs()

	// Parse optional query params
	query := r.URL.Query()
	isActiveParam := query.Get("isActive")
	countryParam := query.Get("countryId")
	typeParam := query.Get("type")

	// Filter thresholds based on query params
	filtered := make([]*models.ECDDThresholdConfig, 0, len(allThresholds))
	for _, t := range allThresholds {
		if isActiveParam != "" {
			isActive, err := strconv.ParseBool(isActiveParam)
			if err == nil && t.IsActive != isActive {
				continue
			}
		}
		if countryParam != "" {
			countryID, err := strconv.ParseInt(countryParam, 10, 64)
			if err == nil && t.CountryID != countryID {
				continue
			}
		}
		if typeParam != "" {
			typeVal, err := strconv.ParseInt(typeParam, 10, 64)
			if err == nil && t.Type != typeVal {
				continue
			}
		}
		filtered = append(filtered, t)
	}

	// Sort the results
	sp := utils.GetSortParams(r, "ecddThresholdConfigPk")
	sortThresholds(filtered, sp)

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

// sortThresholds sorts a slice of ECDDThresholdConfig by the given sort parameters.
func sortThresholds(thresholds []*models.ECDDThresholdConfig, sp utils.SortParams) {
	desc := sp.IsDescending()
	sort.Slice(thresholds, func(i, j int) bool {
		var less bool
		switch sp.SortBy {
		case "ecddThresholdConfigPk":
			less = thresholds[i].ECDDThresholdConfigPK < thresholds[j].ECDDThresholdConfigPK
		case "title":
			less = strings.ToLower(thresholds[i].Title) < strings.ToLower(thresholds[j].Title)
		case "countryId":
			less = thresholds[i].CountryID < thresholds[j].CountryID
		case "type":
			less = thresholds[i].Type < thresholds[j].Type
		case "value":
			less = thresholds[i].Value < thresholds[j].Value
		case "hierarchy":
			less = thresholds[i].Hierarchy < thresholds[j].Hierarchy
		case "loggedAt":
			less = thresholds[i].LoggedAt.Before(thresholds[j].LoggedAt)
		default:
			less = thresholds[i].ECDDThresholdConfigPK < thresholds[j].ECDDThresholdConfigPK
		}
		if desc {
			return !less
		}
		return less
	})
}

// GetThresholdByID handles GET /api/ecdd/thresholdconfig/{id}
func GetThresholdByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/thresholdconfig/")

	ds := services.GetDataService()
	threshold := ds.GetThresholdConfigByID(id)

	if threshold == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Threshold not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(threshold)
}

// CreateThreshold handles POST /api/ecdd/thresholdconfig
func CreateThreshold(w http.ResponseWriter, r *http.Request) {
	var threshold models.ECDDThresholdConfig

	if err := json.NewDecoder(r.Body).Decode(&threshold); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ds := services.GetDataService()
	createdThreshold := ds.CreateThresholdConfig(&threshold)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdThreshold)
}

// UpdateThreshold handles PUT /api/ecdd/thresholdconfig/{id}
func UpdateThreshold(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/thresholdconfig/")

	var threshold models.ECDDThresholdConfig
	if err := json.NewDecoder(r.Body).Decode(&threshold); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ds := services.GetDataService()
	updatedThreshold := ds.UpdateThresholdConfig(id, &threshold)

	if updatedThreshold == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Threshold not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedThreshold)
}

// PatchThreshold handles PATCH /api/ecdd/thresholdconfig/{id}
// Performs a partial update by merging the request body onto the existing record
func PatchThreshold(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/thresholdconfig/")

	ds := services.GetDataService()
	existing := ds.GetThresholdConfigByID(id)
	if existing == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Threshold not found")
		return
	}

	// Decode partial update onto the existing record — only provided fields are overwritten
	if err := json.NewDecoder(r.Body).Decode(existing); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updated := ds.UpdateThresholdConfig(id, existing)
	if updated == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Threshold not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

// DeleteThreshold handles DELETE /api/ecdd/thresholdconfig/{id}
func DeleteThreshold(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/thresholdconfig/")

	ds := services.GetDataService()
	success := ds.DeleteThresholdConfig(id)

	if !success {
		utils.WriteJSONError(w, http.StatusNotFound, "Threshold not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Threshold soft-deleted successfully"})
}
