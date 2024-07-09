package validation

import (
	"api-social-network/internal/domain/entity"

	"github.com/go-playground/validator/v10"
)

// ValidationError represents a detailed validation error.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidateUser validates the user entity.
func ValidateUser(user entity.User) []ValidationError {
	validate := validator.New()
	err := validate.Struct(user)
	if err == nil {
		return nil
	}

	var validationErrors []ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		validationError := ValidationError{
			Field:   err.Field(),
			Message: err.Tag(),
		}
		validationErrors = append(validationErrors, validationError)
	}

	return validationErrors
}
