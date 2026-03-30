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
	"time"
)

// FolderAndStats holds the computed statistics for a single folder.
type FolderAndStats struct {
	FolderPK       string     `json:"folderPk"`
	FolderName     string     `json:"folderName"`
	Region         string     `json:"region"`
	UserCount      int        `json:"userCount"`
	OldestUserDate *time.Time `json:"oldestUserDate"` // nil serialises as JSON null
}

// parseFilterParams extracts the language, country, and region query
// parameters from the request. It returns an error message string if the
// country parameter is present but not a valid int64. On success the error
// string is empty.
func parseFilterParams(r *http.Request) (language *int64, countryID *int64, region string, errMsg string) {
	query := r.URL.Query()
	region = query.Get("region")

	if langParam := query.Get("language"); langParam != "" {
		parsed, err := strconv.ParseInt(langParam, 10, 64)
		if err != nil {
			return nil, nil, "", "Invalid language parameter"
		}
		language = &parsed
	}

	if countryParam := query.Get("countryId"); countryParam != "" {
		parsed, err := strconv.ParseInt(countryParam, 10, 64)
		if err != nil {
			return nil, nil, "", "Invalid country parameter"
		}
		countryID = &parsed
	}

	return language, countryID, region, ""
}

// computeFolderAndStats calculates the user count and oldest assignment date for
// a single folder, applying the provided filters.
//
// It accepts pre-built lookup structures to avoid repeated O(M) scans:
//   - assignmentIndex: map from folder PK to its assignment records
//   - userMap: map from user PK to user status for O(1) user lookups
//
// If assignmentIndex or userMap is nil, the function falls back to the
// per-folder service calls (useful for single-folder lookups where building
// the full index is not worth it).
func computeFolderAndStats(
	ds *services.DataService,
	folder *models.ECDDCaseManagementFolder,
	language *int64, countryID *int64, region string,
	assignmentIndex map[string][]*models.ECDDUserCaseManagementFolder,
	userMap map[string]*models.ECDDUserStatus,
) FolderAndStats {
	folderID := folder.ECDDCaseManagementFolderPK

	// Get assignments for this folder — from pre-built index or service call
	var assignments []*models.ECDDUserCaseManagementFolder
	if assignmentIndex != nil {
		assignments = assignmentIndex[folderID]
	} else {
		assignments = ds.GetUserCaseFoldersByFolderPK(folderID)
	}

	// Resolve user objects from assignments and apply filters
	var allUsers []*models.ECDDUserStatus
	if userMap != nil {
		for _, a := range assignments {
			if u, ok := userMap[a.UserStatusPK]; ok {
				allUsers = append(allUsers, u)
			}
		}
	} else {
		allUsers = ds.GetUsersByFolderPK(folderID)
	}

	filtered := FilterUsers(allUsers, language, countryID, region)

	stats := FolderAndStats{
		FolderPK:   folderID,
		FolderName: folder.FolderName,
		Region:     folder.Region,
		UserCount:  len(filtered),
	}

	if len(filtered) == 0 {
		return stats
	}

	// Build a set of filtered user PKs for O(1) lookup when scanning
	// assignments to find the earliest LoggedAt.
	filteredSet := make(map[string]struct{}, len(filtered))
	for _, u := range filtered {
		filteredSet[u.ECDDUserStatusPK] = struct{}{}
	}

	// Walk the assignment records and find the earliest LoggedAt among
	// assignments whose user passed the filters.
	var oldest *time.Time
	for _, a := range assignments {
		if _, ok := filteredSet[a.UserStatusPK]; !ok {
			continue
		}
		t := a.LoggedAt // copy to avoid taking address of loop variable
		if oldest == nil || t.Before(*oldest) {
			oldest = &t
		}
	}
	stats.OldestUserDate = oldest

	return stats
}

// GetAllFolders handles GET /api/ecdd/casemanagementfolder
func GetAllFolders(w http.ResponseWriter, r *http.Request) {
	ds := services.GetDataService()
	allFolders := ds.GetAllCaseFolders()

	// Parse optional search query param
	search := r.URL.Query().Get("search")

	if search != "" {
		searchLower := strings.ToLower(search)
		filtered := make([]*models.ECDDCaseManagementFolder, 0, len(allFolders))
		for _, f := range allFolders {
			if strings.Contains(strings.ToLower(f.FolderName), searchLower) {
				filtered = append(filtered, f)
			}
		}
		allFolders = filtered
	}

	// Filter by region (case-insensitive)
	if regionParam := r.URL.Query().Get("region"); regionParam != "" {
		filtered := make([]*models.ECDDCaseManagementFolder, 0, len(allFolders))
		for _, f := range allFolders {
			if strings.EqualFold(f.Region, regionParam) {
				filtered = append(filtered, f)
			}
		}
		allFolders = filtered
	}

	// Sort the results
	sp := utils.GetSortParams(r, "folderName")
	sortFolders(allFolders, sp)

	// Optional pagination
	pp := utils.GetPaginationParams(r)
	if pp.Enabled {
		total := len(allFolders)
		start, end := utils.CalculateOffset(pp.Page, pp.PageSize, total)
		paged := allFolders[start:end]
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(utils.Paginate(paged, pp, total))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allFolders)
}

// sortFolders sorts a slice of ECDDCaseManagementFolder by the given sort parameters.
func sortFolders(folders []*models.ECDDCaseManagementFolder, sp utils.SortParams) {
	desc := sp.IsDescending()
	sort.Slice(folders, func(i, j int) bool {
		var less bool
		switch sp.SortBy {
		case "folderName":
			less = strings.ToLower(folders[i].FolderName) < strings.ToLower(folders[j].FolderName)
		case "ecddCaseManagementFolderPk":
			less = folders[i].ECDDCaseManagementFolderPK < folders[j].ECDDCaseManagementFolderPK
		case "loggedAt":
			less = folders[i].LoggedAt.Before(folders[j].LoggedAt)
		case "region":
			less = strings.ToLower(folders[i].Region) < strings.ToLower(folders[j].Region)
		default:
			less = strings.ToLower(folders[i].FolderName) < strings.ToLower(folders[j].FolderName)
		}
		if desc {
			return !less
		}
		return less
	})
}

// GetFolderByID handles GET /api/ecdd/casemanagementfolder/{id}
func GetFolderByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/casemanagementfolder/")
	ds := services.GetDataService()
	folder := ds.GetCaseFolderByID(id)

	if folder == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Folder not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folder)
}

// CreateFolder handles POST /api/ecdd/casemanagementfolder
func CreateFolder(w http.ResponseWriter, r *http.Request) {
	var folder models.ECDDCaseManagementFolder
	if err := json.NewDecoder(r.Body).Decode(&folder); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ds := services.GetDataService()
	createdFolder := ds.CreateCaseFolder(&folder)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdFolder)
}

// UpdateFolder handles PUT /api/ecdd/casemanagementfolder/{id}
func UpdateFolder(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/casemanagementfolder/")
	var folder models.ECDDCaseManagementFolder
	if err := json.NewDecoder(r.Body).Decode(&folder); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ds := services.GetDataService()
	updatedFolder := ds.UpdateCaseFolder(id, &folder)

	if updatedFolder == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Folder not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedFolder)
}

// GetFolderAndStats handles GET /api/ecdd/usercasemanagement/folder/{id}/stats
// Returns userCount and oldestUserDate for a single folder with optional
// language, country, and region filters.
func GetFolderAndStats(w http.ResponseWriter, r *http.Request) {
	// Extract folder ID from path: /api/ecdd/usercasemanagement/folder/{id}/stats
	path := strings.TrimPrefix(r.URL.Path, "/api/ecdd/usercasemanagement/folder/")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) < 1 || parts[0] == "" {
		utils.WriteJSONError(w, http.StatusBadRequest, "Folder ID required")
		return
	}
	folderID := parts[0]

	ds := services.GetDataService()
	folder := ds.GetCaseFolderByID(folderID)
	if folder == nil {
		utils.WriteJSONError(w, http.StatusNotFound, "Folder not found")
		return
	}

	language, countryID, region, errMsg := parseFilterParams(r)
	if errMsg != "" {
		utils.WriteJSONError(w, http.StatusBadRequest, errMsg)
		return
	}

	stats := computeFolderAndStats(ds, folder, language, countryID, region, nil, nil)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

// GetAllFolderAndStats handles GET /api/ecdd/usercasemanagement/stats
// Returns an array of stats objects for every folder, with optional
// language, country, and region filters applied to user counts.
func GetAllFolderAndStats(w http.ResponseWriter, r *http.Request) {
	language, countryID, region, errMsg := parseFilterParams(r)
	if errMsg != "" {
		utils.WriteJSONError(w, http.StatusBadRequest, errMsg)
		return
	}

	ds := services.GetDataService()
	allFolders := ds.GetAllCaseFolders()

	// Filter folders by region before computing stats
	if region != "" {
		filtered := make([]*models.ECDDCaseManagementFolder, 0, len(allFolders))
		for _, f := range allFolders {
			if strings.EqualFold(f.Region, region) {
				filtered = append(filtered, f)
			}
		}
		allFolders = filtered
	}

	// Pre-build index and user map once — O(M + U) instead of O(N × 2M)
	assignmentIndex := ds.GetFolderAssignmentIndex()
	userMap := ds.GetUserStatusMap()

	results := make([]FolderAndStats, 0, len(allFolders))
	for _, folder := range allFolders {
		results = append(results, computeFolderAndStats(ds, folder, language, countryID, region, assignmentIndex, userMap))
	}

	// Sort the results
	sp := utils.GetSortParams(r, "folderName")
	sortFolderAndStats(results, sp)

	// Optional pagination
	pp := utils.GetPaginationParams(r)
	if pp.Enabled {
		total := len(results)
		start, end := utils.CalculateOffset(pp.Page, pp.PageSize, total)
		paged := results[start:end]
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(utils.Paginate(paged, pp, total))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// sortFolderAndStats sorts a slice of FolderAndStats by the given sort parameters.
func sortFolderAndStats(stats []FolderAndStats, sp utils.SortParams) {
	desc := sp.IsDescending()
	sort.Slice(stats, func(i, j int) bool {
		var less bool
		switch sp.SortBy {
		case "folderName":
			less = strings.ToLower(stats[i].FolderName) < strings.ToLower(stats[j].FolderName)
		case "folderPk":
			less = stats[i].FolderPK < stats[j].FolderPK
		case "userCount":
			less = stats[i].UserCount < stats[j].UserCount
		case "region":
			less = strings.ToLower(stats[i].Region) < strings.ToLower(stats[j].Region)
		default:
			less = strings.ToLower(stats[i].FolderName) < strings.ToLower(stats[j].FolderName)
		}
		if desc {
			return !less
		}
		return less
	})
}

// DeleteFolder handles DELETE /api/ecdd/casemanagementfolder/{id}
// Hard deletes the folder and cascade-deletes all related user_case_folders
func DeleteFolder(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/ecdd/casemanagementfolder/")
	// Strip any trailing path segments (in case routed through handleFolderByID)
	id := strings.SplitN(path, "/", 2)[0]

	ds := services.GetDataService()
	success, cascadeCount := ds.DeleteCaseFolder(id)

	if !success {
		utils.WriteJSONError(w, http.StatusNotFound, "Folder not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Folder deleted successfully. %d user-folder assignment(s) cascade-deleted.", cascadeCount),
	})
}
