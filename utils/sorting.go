package utils

import "net/http"

// SortParams holds sorting parameters from request
type SortParams struct {
	SortBy  string
	SortDir string // "asc" or "desc"
}

// GetSortParams extracts sortBy and sortDir query parameters from the HTTP
// request. If sortBy is not provided, defaultSortBy is used. If sortDir is
// not provided or invalid, "asc" is used.
func GetSortParams(r *http.Request, defaultSortBy string) SortParams {
	sortBy := r.URL.Query().Get("sortBy")
	sortDir := r.URL.Query().Get("sortDir")

	if sortBy == "" {
		sortBy = defaultSortBy
	}
	if sortDir != "asc" && sortDir != "desc" {
		sortDir = "asc"
	}

	return SortParams{
		SortBy:  sortBy,
		SortDir: sortDir,
	}
}

// IsDescending returns true when the sort direction is descending.
func (s SortParams) IsDescending() bool {
	return s.SortDir == "desc"
}
