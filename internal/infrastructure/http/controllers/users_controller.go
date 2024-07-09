package controllers

import (
	"api-social-network/internal/domain/entity"
	"api-social-network/internal/domain/service"
	"api-social-network/internal/infrastructure/errors"
	"api-social-network/internal/infrastructure/http/response"
	"api-social-network/internal/infrastructure/http/types"
	"api-social-network/internal/infrastructure/validation"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// UserController handles user-related HTTP requests.
type UserController struct {
	UserService service.UserService
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body entity.User true "User"
// @Success 201 {object} types.UserResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /users [post]
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.Error(w, errors.ErrInvalidPayload)
		return
	}

	user.CreatedAt = time.Now()

	if validationErrors := validation.ValidateUser(user); validationErrors != nil {
		response.ValidationErrors(w, validationErrors)
		return
	}

	if err := c.UserService.CreateUser(user); err != nil {
		response.Error(w, errors.NewAppError(http.StatusInternalServerError, "Failed to create user", err.Error()))
		return
	}

	userResponse := types.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Nick:      user.Nick,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	response.JSON(w, http.StatusCreated, userResponse)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} types.UserResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /users/{id} [get]
func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		response.Error(w, errors.NewAppError(http.StatusBadRequest, "Invalid user ID", err.Error()))
		return
	}

	user, err := c.UserService.GetUserByID(id)
	if err != nil {
		response.Error(w, errors.NewAppError(http.StatusInternalServerError, "Failed to retrieve user", err.Error()))
		return
	}

	userResponse := types.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Nick:      user.Nick,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	response.JSON(w, http.StatusOK, userResponse)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} types.UserResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /users [get]
func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.UserService.GetAllUsers()
	if err != nil {
		response.Error(w, errors.NewAppError(http.StatusInternalServerError, "Failed to retrieve users", err.Error()))
		return
	}

	var usersResponse []types.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, types.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Nick:      user.Nick,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		})
	}

	response.JSON(w, http.StatusOK, usersResponse)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update an existing user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body entity.User true "User"
// @Success 200 {string} string "User updated successfully"
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /users/{id} [put]
func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.Error(w, errors.ErrInvalidPayload)
		return
	}

	if validationErrors := validation.ValidateUser(user); validationErrors != nil {
		response.ValidationErrors(w, validationErrors)
		return
	}

	if err := c.UserService.UpdateUser(user); err != nil {
		response.Error(w, errors.NewAppError(http.StatusInternalServerError, "Failed to update user", err.Error()))
		return
	}

	response.JSON(w, http.StatusOK, "User updated successfully")
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /users/{id} [delete]
func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		response.Error(w, errors.NewAppError(http.StatusBadRequest, "Invalid user ID", err.Error()))
		return
	}

	if err := c.UserService.DeleteUser(id); err != nil {
		response.Error(w, errors.NewAppError(http.StatusInternalServerError, "Failed to delete user", err.Error()))
		return
	}

	response.JSON(w, http.StatusOK, "User deleted successfully")
}
