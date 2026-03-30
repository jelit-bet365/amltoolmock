package handlers

import (
	"amltoolmock/models"
	regionpkg "amltoolmock/enums/region"
	"amltoolmock/services"
	"amltoolmock/utils"
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// FilterUsers applies optional language, country, and region filters to a
// slice of users and returns only those that match all provided criteria.
// Pass empty string / nil to skip a filter.
func FilterUsers(users []*models.ECDDUserStatus, language *int64, countryID *int64, region string) []*models.ECDDUserStatus {
	filtered := make([]*models.ECDDUserStatus, 0, len(users))
	for _, user := range users {
		if language != nil && user.Language != *language {
			continue
		}
		if countryID != nil && user.CountryID != *countryID {
			continue
		}
		if region != "" {
			countries := regionpkg.GetCountriesForRegion(regionpkg.Region(region))
			if len(countries) > 0 {
				allowed := make(map[int]struct{}, len(countries))
				for _, id := range countries {
					allowed[id] = struct{}{}
				}
				if _, ok := allowed[int(user.CountryID)]; !ok {
					continue
				}
			}
		}
		filtered = append(filtered, user)
	}
	return filtered
}

// GetFolderUsers handles GET /api/ecdd/usercasemanagement/folder/{id}/users with optional filters and pagination
func GetFolderUsers(w http.ResponseWriter, r *http.Request) {
	// Expect path format: /api/ecdd/usercasemanagement/folder/{id}/users
	path := strings.TrimPrefix(r.URL.Path, "/api/ecdd/usercasemanagement/folder/")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] != "users" {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid folder users path")
		return
	}
	folderID := parts[0]

	// Parse query parameters
	query := r.URL.Query()
	region := query.Get("region")

	var language *int64
	if langParam := query.Get("language"); langParam != "" {
		parsed, err := strconv.ParseInt(langParam, 10, 64)
		if err != nil {
			utils.WriteJSONError(w, http.StatusBadRequest, "Invalid language parameter")
			return
		}
		language = &parsed
	}

	var countryID *int64
	if countryParam := query.Get("countryId"); countryParam != "" {
		if parsed, err := strconv.ParseInt(countryParam, 10, 64); err == nil {
			countryID = &parsed
		} else {
			utils.WriteJSONError(w, http.StatusBadRequest, "Invalid country parameter")
			return
		}
	}

	ds := services.GetDataService()
	allUsers := ds.GetUsersByFolderPK(folderID)

	// Apply optional filters using the shared helper
	filtered := FilterUsers(allUsers, language, countryID, region)

	// Sort BEFORE pagination so page boundaries are consistent
	sp := utils.GetSortParams(r, "userId")
	sortFolderUsers(filtered, sp)

	total := len(filtered)

	// Use the standardised pagination utility
	pp := utils.GetPaginationParams(r)
	if !pp.Enabled {
		// Default pagination for this endpoint (backward compat)
		pp.Page = 1
		pp.PageSize = 10
		pp.Enabled = true

		if pageParam := query.Get("page"); pageParam != "" {
			if parsed, err := strconv.Atoi(pageParam); err == nil && parsed > 0 {
				pp.Page = parsed
			}
		}
		if pageSizeParam := query.Get("pageSize"); pageSizeParam != "" {
			if parsed, err := strconv.Atoi(pageSizeParam); err == nil && parsed > 0 {
				pp.PageSize = parsed
			}
		}
	}

	start, end := utils.CalculateOffset(pp.Page, pp.PageSize, total)
	pagedUsers := filtered[start:end]

	response := utils.Paginate(pagedUsers, pp, total)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// sortFolderUsers sorts a slice of ECDDUserStatus for the folder users endpoint.
func sortFolderUsers(users []*models.ECDDUserStatus, sp utils.SortParams) {
	desc := sp.IsDescending()
	sort.Slice(users, func(i, j int) bool {
		var less bool
		switch sp.SortBy {
		case "userId":
			less = users[i].UserID < users[j].UserID
		case "userName":
			less = strings.ToLower(users[i].UserName) < strings.ToLower(users[j].UserName)
		case "countryId":
			less = users[i].CountryID < users[j].CountryID
		case "ecddStatus":
			less = users[i].ECDDStatus < users[j].ECDDStatus
		case "ecddThreshold":
			less = users[i].ECDDThreshold < users[j].ECDDThreshold
		case "ecddMultiplier":
			less = users[i].ECDDMultiplier < users[j].ECDDMultiplier
		case "loggedAt":
			less = users[i].LoggedAt.Before(users[j].LoggedAt)
		case "ecddUserStatusPk":
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

// DeleteFolderUser handles DELETE /api/ecdd/usercasemanagement/folder/{id}/users/{user_id}
func DeleteFolderUser(w http.ResponseWriter, r *http.Request) {
	// Expect path format: /api/ecdd/usercasemanagement/folder/{folder_id}/users/{user_id}
	path := strings.TrimPrefix(r.URL.Path, "/api/ecdd/usercasemanagement/folder/")
	parts := strings.Split(path, "/")
	if len(parts) != 3 || parts[0] == "" || parts[1] != "users" || parts[2] == "" {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid path, expected /api/ecdd/usercasemanagement/folder/{id}/users/{user_id}")
		return
	}
	folderID := parts[0]
	userID := parts[2]

	ds := services.GetDataService()
	success := ds.DeleteUserCaseFolderByFolderPKAndUserStatusPK(folderID, userID)

	if !success {
		utils.WriteJSONError(w, http.StatusNotFound, "User-folder assignment not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User removed from folder successfully"})
}

// BulkDeleteFolderUsers handles POST /api/ecdd/usercasemanagement/folder/{id}/users/bulk-delete
func BulkDeleteFolderUsers(w http.ResponseWriter, r *http.Request) {
	// Expect path format: /api/ecdd/usercasemanagement/folder/{folder_id}/users/bulk-delete
	path := strings.TrimPrefix(r.URL.Path, "/api/ecdd/usercasemanagement/folder/")
	parts := strings.Split(path, "/")
	if len(parts) != 3 || parts[0] == "" || parts[1] != "users" || parts[2] != "bulk-delete" {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid path, expected /api/ecdd/usercasemanagement/folder/{id}/users/bulk-delete")
		return
	}
	folderID := parts[0]

	// Parse request body
	var reqBody struct {
		UserIDs   []string `json:"ecddUserStatusPks"`
		UpdatedBy string   `json:"updatedBy"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if len(reqBody.UserIDs) == 0 {
		utils.WriteJSONError(w, http.StatusBadRequest, "ecddUserStatusPks array is required and must not be empty")
		return
	}

	ds := services.GetDataService()
	deletedCount, failedIDs := ds.BulkDeleteUserCaseFoldersByFolderPKAndUserStatusPKs(folderID, reqBody.UserIDs)

	// Ensure failedEcddUserStatusPks is always an array in JSON, never null
	if failedIDs == nil {
		failedIDs = []string{}
	}

	message := "Users removed from folder successfully"
	if len(failedIDs) > 0 {
		message = "Some users could not be removed"
	}

	response := struct {
		Message       string   `json:"message"`
		DeletedCount  int      `json:"deletedCount"`
		FailedCount   int      `json:"failedCount"`
		FailedUserIDs []string `json:"failedEcddUserStatusPks"`
	}{
		Message:       message,
		DeletedCount:  deletedCount,
		FailedCount:   len(failedIDs),
		FailedUserIDs: failedIDs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// BulkAddFolderUsers handles POST /api/ecdd/usercasemanagement/folder/{id}/users/bulk-add
// Assigns multiple users to a folder in a single atomic request.
// Skips duplicates (users already assigned to the folder).
func BulkAddFolderUsers(w http.ResponseWriter, r *http.Request) {
	// Expect path format: /api/ecdd/usercasemanagement/folder/{folder_id}/users/bulk-add
	path := strings.TrimPrefix(r.URL.Path, "/api/ecdd/usercasemanagement/folder/")
	parts := strings.Split(path, "/")
	if len(parts) != 3 || parts[0] == "" || parts[1] != "users" || parts[2] != "bulk-add" {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid path, expected /api/ecdd/usercasemanagement/folder/{id}/users/bulk-add")
		return
	}
	folderID := parts[0]

	// Parse request body
	var reqBody struct {
		UserIDs   []string `json:"userIds"`
		UpdatedBy string   `json:"updatedBy"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if len(reqBody.UserIDs) == 0 {
		utils.WriteJSONError(w, http.StatusBadRequest, "userIds array is required and must not be empty")
		return
	}

	// Verify the folder exists
	ds := services.GetDataService()
	folder := ds.GetCaseFolderByID(folderID)
	if folder == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Folder not found")
		return
	}

	created, skipped := ds.BulkCreateUserCaseFolders(folderID, reqBody.UserIDs, reqBody.UpdatedBy)

	// Ensure skippedUserIds is always an array in JSON, never null
	if skipped == nil {
		skipped = []string{}
	}

	message := "Users assigned to folder successfully"
	if len(skipped) > 0 {
		message = "Some users were already assigned to the folder and were skipped"
	}

	response := struct {
		Message      string   `json:"message"`
		CreatedCount int      `json:"createdCount"`
		SkippedCount int      `json:"skippedCount"`
		SkippedIDs   []string `json:"skippedUserIds"`
	}{
		Message:      message,
		CreatedCount: len(created),
		SkippedCount: len(skipped),
		SkippedIDs:   skipped,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
