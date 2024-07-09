package controllers

import (
	"api-social-network/internal/domain/service"
	"api-social-network/internal/infrastructure/errors"
	"api-social-network/internal/infrastructure/http/middlewares"
	"api-social-network/internal/infrastructure/http/response"
	"api-social-network/internal/infrastructure/http/types"
	"api-social-network/internal/infrastructure/security"
	"encoding/json"
	"net/http"
)

// AuthController handles user authentication.
type AuthController struct {
	UserService service.UserService
}

// Login godoc
// @Summary Login a user
// @Description Login a user and returns a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body types.AuthCredentials true "User credentials"
// @Success 200 {object} types.AuthResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /login [post]
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var credentials types.AuthCredentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		response.Error(w, errors.ErrInvalidPayload)
		return
	}

	user, err := c.UserService.GetUserByEmail(credentials.Email)
	if err != nil || security.VerifyPassword(user.Password, credentials.Password) != nil {
		response.Error(w, errors.ErrUnauthorized)
		return
	}

	token, err := middlewares.GenerateToken(user.ID)
	if err != nil {
		response.Error(w, errors.ErrInternal)
		return
	}

	response.JSON(w, http.StatusOK, types.AuthResponse{
		ID:    user.ID,
		Token: token,
	})
}
