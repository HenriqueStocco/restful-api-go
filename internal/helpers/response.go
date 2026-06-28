package helpers

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type APIResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    any       `json:"data,omitempty"`
	Error   *APIError `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, statusCode int, response APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func WriteSuccess(w http.ResponseWriter, statusCode int, message string, data any) {
	writeJSON(w, statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}
func WriteError(w http.ResponseWriter, statusCode int, code string, message string) {
	writeJSON(w, statusCode, APIResponse{
		Success: false,
		Message: message,
		Error: &APIError{
			Code:    code,
			Message: message,
		},
	})
}
