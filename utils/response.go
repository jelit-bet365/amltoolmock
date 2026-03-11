package utils

import (
	"encoding/json"
	"net/http"
)

// WriteJSONError writes a JSON-formatted error response with the correct
// Content-Type header. Using json.NewEncoder ensures proper JSON encoding,
// preventing injection via untrusted error messages.
func WriteJSONError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
