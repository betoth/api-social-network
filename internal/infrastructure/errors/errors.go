package errors

import (
	"fmt"
	"net/http"
)

// AppError represents a standardized error structure.
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

// NewAppError creates a new AppError.
func NewAppError(code int, message, detail string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
}

// Error implements the error interface for AppError.
func (e *AppError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, detail: %s", e.Code, e.Message, e.Detail)
}

// Predefined errors
var (
	ErrInvalidPayload = NewAppError(http.StatusBadRequest, "Invalid request payload", "")
	ErrUnauthorized   = NewAppError(http.StatusUnauthorized, "Unauthorized", "")
	ErrNotFound       = NewAppError(http.StatusNotFound, "Resource not found", "")
	ErrInternal       = NewAppError(http.StatusInternalServerError, "Internal server error", "")
)
