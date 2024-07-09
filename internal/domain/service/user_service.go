package service

import (
	"api-social-network/internal/application/ports"
	"api-social-network/internal/domain/entity"
	"api-social-network/internal/infrastructure/security"
	"errors"
)

// UserService provides user-related services.
type UserService struct {
	UserRepository ports.UserRepository
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(user entity.User) error {
	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = hashedPassword
	return s.UserRepository.CreateUser(user)
}

// GetUserByID retrieves a user by its ID.
func (s *UserService) GetUserByID(id uint64) (entity.User, error) {
	return s.UserRepository.GetUserByID(id)
}

// GetAllUsers retrieves all users.
func (s *UserService) GetAllUsers() ([]entity.User, error) {
	return s.UserRepository.GetAllUsers()
}

// UpdateUser updates an existing user.
func (s *UserService) UpdateUser(user entity.User) error {
	return s.UserRepository.UpdateUser(user)
}

// DeleteUser deletes a user by its ID.
func (s *UserService) DeleteUser(id uint64) error {
	return s.UserRepository.DeleteUser(id)
}

// GetUserByEmail retrieves a user by its email.
func (s *UserService) GetUserByEmail(email string) (entity.User, error) {
	return s.UserRepository.GetUserByEmail(email)
}
