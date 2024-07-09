package response

import (
	"api-social-network/internal/infrastructure/errors"
	"api-social-network/internal/infrastructure/validation"
	"encoding/json"
	"net/http"
)

// ErrorResponse represents a standardized error response.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// JSON writes the JSON response with a given status code.
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(response); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

// Error writes the error response with a given status code.
func Error(w http.ResponseWriter, err *errors.AppError) {
	JSON(w, err.Code, err)
}

// ValidationErrors writes the validation errors response.
func ValidationErrors(w http.ResponseWriter, validationErrors []validation.ValidationError) {
	JSON(w, http.StatusBadRequest, map[string]interface{}{
		"code":    http.StatusBadRequest,
		"error":   "Validation failed",
		"details": validationErrors,
	})
}
