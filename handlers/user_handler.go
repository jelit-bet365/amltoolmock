package handlers

import (
	"amltoolmock/models"
	"amltoolmock/services"
	"amltoolmock/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// GetAllUsers handles GET /api/ecdd/userstatus
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ds := services.GetDataService()
	allUsers := ds.GetAllUserStatuses()

	// Parse optional query params
	query := r.URL.Query()
	countryParam := query.Get("country_ids")
	ecddStatusParam := query.Get("ecdd_status")
	searchParam := query.Get("search")

	// Build country set for O(1) lookups
	var countrySet map[int64]bool
	if countryParam != "" {
		countrySet = make(map[int64]bool)
		parts := strings.Split(countryParam, ",")
		for _, part := range parts {
			id, err := strconv.ParseInt(strings.TrimSpace(part), 10, 64)
			if err == nil {
				countrySet[id] = true
			}
		}
	}

	// Filter users based on query params
	filtered := make([]*models.ECDDUserStatus, 0, len(allUsers))
	for _, u := range allUsers {
		if countrySet != nil && !countrySet[u.CountryID] {
			continue
		}
		if ecddStatusParam != "" {
			ecddStatus, err := strconv.ParseInt(ecddStatusParam, 10, 64)
			if err == nil && u.ECDDStatus != ecddStatus {
				continue
			}
		}
		if searchParam != "" {
			searchLower := strings.ToLower(searchParam)
			if !strings.Contains(strings.ToLower(u.UserName), searchLower) {
				continue
			}
		}
		filtered = append(filtered, u)
	}

	// Sort the results
	sp := utils.GetSortParams(r, "user_id")
	sortUsers(filtered, sp)

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

// sortUsers sorts a slice of ECDDUserStatus by the given sort parameters.
func sortUsers(users []*models.ECDDUserStatus, sp utils.SortParams) {
	desc := sp.IsDescending()
	sort.Slice(users, func(i, j int) bool {
		var less bool
		switch sp.SortBy {
		case "user_id":
			less = users[i].UserID < users[j].UserID
		case "user_name":
			less = strings.ToLower(users[i].UserName) < strings.ToLower(users[j].UserName)
		case "country_id":
			less = users[i].CountryID < users[j].CountryID
		case "ecdd_status":
			less = users[i].ECDDStatus < users[j].ECDDStatus
		case "ecdd_threshold":
			less = users[i].ECDDThreshold < users[j].ECDDThreshold
		case "ecdd_multiplier":
			less = users[i].ECDDMultiplier < users[j].ECDDMultiplier
		case "logged_at":
			less = users[i].LoggedAt.Before(users[j].LoggedAt)
		case "ecdd_user_status_pk":
			less = users[i].ECDDUserStatusPK < users[j].ECDDUserStatusPK
		default:
			less = users[i].UserID < users[j].UserID
		}
		if desc {
			return !less
		}
		return less
	})
}

// GetUserByID handles GET /api/ecdd/userstatus/{id}
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/userstatus/")

	ds := services.GetDataService()
	user := ds.GetUserStatusByID(id)

	if user == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "User not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUser handles POST /api/ecdd/userstatus
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.ECDDUserStatus

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ds := services.GetDataService()
	createdUser := ds.CreateUserStatus(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// UpdateUser handles PUT /api/ecdd/userstatus/{id}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/userstatus/")

	var user models.ECDDUserStatus
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ds := services.GetDataService()
	updatedUser := ds.UpdateUserStatus(id, &user)

	if updatedUser == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "User not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

// PatchUser handles PATCH /api/ecdd/userstatus/{id}
// Performs a partial update by merging the request body onto the existing record
func PatchUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/userstatus/")

	ds := services.GetDataService()
	existing := ds.GetUserStatusByID(id)
	if existing == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "User not found")
		return
	}

	// Decode partial update onto the existing record — only provided fields are overwritten
	if err := json.NewDecoder(r.Body).Decode(existing); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updated := ds.UpdateUserStatus(id, existing)
	if updated == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "User not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

// DeleteUser handles DELETE /api/ecdd/userstatus/{id}
// Hard deletes the user and cascade-deletes all related user_case_folders
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/userstatus/")

	ds := services.GetDataService()
	success, cascadeCount := ds.DeleteUserStatus(id)

	if !success {
		utils.WriteJSONError(w, http.StatusNotFound, "User not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("User deleted successfully. %d folder assignment(s) cascade-deleted.", cascadeCount),
	})
}

// GetUserFolders handles GET /api/ecdd/userstatus/{id}/folders
// Returns the list of folders that a user is assigned to.
func GetUserFolders(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from path: /api/ecdd/userstatus/{id}/folders
	path := strings.TrimPrefix(r.URL.Path, "/api/ecdd/userstatus/")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) < 1 || parts[0] == "" {
		utils.WriteJSONError(w, http.StatusBadRequest, "User ID required")
		return
	}
	userID := parts[0]

	ds := services.GetDataService()

	// Verify user exists
	user := ds.GetUserStatusByID(userID)
	if user == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "User not found")
		return
	}

	folders := ds.GetFoldersByUserPK(userID)

	// Return empty array instead of null
	if folders == nil {
		folders = []*models.ECDDCaseManagementFolder{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folders)
}
