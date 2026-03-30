package handlers

import (
	"amltoolmock/models"
	"amltoolmock/services"
	"amltoolmock/utils"
	"encoding/json"
	"net/http"
	"sort"
	"strings"
)

// GetAllFolderAssignments handles GET /api/ecdd/usercasemanagement
func GetAllFolderAssignments(w http.ResponseWriter, r *http.Request) {
	ds := services.GetDataService()

	// Parse optional query params
	query := r.URL.Query()
	folderID := query.Get("ecddCaseManagementFolderPk")
	userID := query.Get("ecddUserStatusPk")

	var assignments []*models.ECDDUserCaseManagementFolder

	if folderID != "" && userID != "" {
		// Both provided: get by folder, then filter by user (AND logic)
		byFolder := ds.GetUserCaseFoldersByFolderPK(folderID)
		assignments = make([]*models.ECDDUserCaseManagementFolder, 0, len(byFolder))
		for _, a := range byFolder {
			if a.UserStatusPK == userID {
				assignments = append(assignments, a)
			}
		}
	} else if folderID != "" {
		assignments = ds.GetUserCaseFoldersByFolderPK(folderID)
	} else if userID != "" {
		assignments = ds.GetUserCaseFoldersByUserStatusPK(userID)
	} else {
		assignments = ds.GetAllUserCaseFolders()
	}

	// Ensure JSON [] not null
	if assignments == nil {
		assignments = make([]*models.ECDDUserCaseManagementFolder, 0)
	}

	// Sort the results
	sp := utils.GetSortParams(r, "ecddUserCaseManagementFolderPk")
	sortFolderAssignments(assignments, sp)

	// Optional pagination
	pp := utils.GetPaginationParams(r)
	if pp.Enabled {
		total := len(assignments)
		start, end := utils.CalculateOffset(pp.Page, pp.PageSize, total)
		paged := assignments[start:end]
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(utils.Paginate(paged, pp, total))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assignments)
}

// sortFolderAssignments sorts a slice of ECDDUserCaseManagementFolder by the given sort parameters.
func sortFolderAssignments(assignments []*models.ECDDUserCaseManagementFolder, sp utils.SortParams) {
	desc := sp.IsDescending()
	sort.Slice(assignments, func(i, j int) bool {
		var less bool
		switch sp.SortBy {
		case "ecddUserCaseManagementFolderPk":
			less = assignments[i].ECDDUserCaseManagementFolderPK < assignments[j].ECDDUserCaseManagementFolderPK
		case "ecddCaseManagementFolderPk":
			less = assignments[i].FolderPK < assignments[j].FolderPK
		case "ecddUserStatusPk":
			less = assignments[i].UserStatusPK < assignments[j].UserStatusPK
		case "loggedAt":
			less = assignments[i].LoggedAt.Before(assignments[j].LoggedAt)
		default:
			less = assignments[i].ECDDUserCaseManagementFolderPK < assignments[j].ECDDUserCaseManagementFolderPK
		}
		if desc {
			return !less
		}
		return less
	})
}

// CreateFolderAssignment handles POST /api/ecdd/usercasemanagement
func CreateFolderAssignment(w http.ResponseWriter, r *http.Request) {
	var assignment models.ECDDUserCaseManagementFolder
	if err := json.NewDecoder(r.Body).Decode(&assignment); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ds := services.GetDataService()
	createdAssignment := ds.CreateUserCaseFolder(&assignment)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdAssignment)
}

// DeleteFolderAssignment handles DELETE /api/ecdd/usercasemanagement/{id}
func DeleteFolderAssignment(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/ecdd/usercasemanagement/")
	ds := services.GetDataService()
	success := ds.DeleteUserCaseFolder(id)

	if !success {
		utils.WriteJSONError(w, http.StatusNotFound, "Assignment not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Assignment deleted successfully"})
}

// DeleteFolderAssignmentsByFolder handles DELETE /api/ecdd/usercasemanagement?ecddCaseManagementFolderPk=
func DeleteFolderAssignmentsByFolder(w http.ResponseWriter, r *http.Request) {
	folderID := r.URL.Query().Get("ecddCaseManagementFolderPk")
	if folderID == "" {
		utils.WriteJSONError(w, http.StatusBadRequest, "ecddCaseManagementFolderPk query parameter is required")
		return
	}

	ds := services.GetDataService()
	count := ds.DeleteUserCaseFoldersByFolderPK(folderID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Assignments deleted successfully",
		"count":   count,
	})
}
