package utils

import (
	"math"
	"net/http"
	"strconv"
)

// PaginationParams holds pagination parameters from request
type PaginationParams struct {
	Page     int
	PageSize int
	// Enabled indicates whether pagination was explicitly requested.
	// When false the caller should return all results.
	Enabled bool
}

// PaginationResponse wraps paginated data with metadata.
type PaginationResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
	TotalPages int         `json:"totalPages"`
}

// GetPaginationParams extracts pagination parameters from HTTP request.
// If neither page nor pageSize is provided the returned Enabled field is
// false, signalling to the caller that all results should be returned.
func GetPaginationParams(r *http.Request) PaginationParams {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")

	// No pagination requested
	if pageStr == "" && pageSizeStr == "" {
		return PaginationParams{Enabled: false}
	}

	page := 1
	pageSize := 20 // default page size

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 && ps <= 100 {
			pageSize = ps
		}
	}

	return PaginationParams{
		Page:     page,
		PageSize: pageSize,
		Enabled:  true,
	}
}

// Paginate applies pagination to a slice and returns paginated response
func Paginate(data interface{}, params PaginationParams, totalCount int) PaginationResponse {
	totalPages := int(math.Ceil(float64(totalCount) / float64(params.PageSize)))

	return PaginationResponse{
		Data:       data,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}
}

// CalculateOffset calculates the offset for slicing based on page and page size.
// It clamps the values so they don't exceed totalCount.
func CalculateOffset(page, pageSize, totalCount int) (start, end int) {
	start = (page - 1) * pageSize
	if start > totalCount {
		start = totalCount
	}
	end = start + pageSize
	if end > totalCount {
		end = totalCount
	}
	return start, end
}
