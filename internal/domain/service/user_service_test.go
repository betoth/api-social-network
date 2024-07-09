package service

import (
	"api-social-network/internal/domain/entity"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// MockUserRepository is a mock implementation of the UserRepository interface.
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByID(id uint64) (entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) GetAllUsers() ([]entity.User, error) {
	args := m.Called()
	return args.Get(0).([]entity.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(id uint64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByEmail(email string) (entity.User, error) {
	args := m.Called(email)
	return args.Get(0).(entity.User), args.Error(1)
}

// senhaMatcher é um matcher personalizado que verifica se a senha foi hasheada corretamente.
func senhaMatcher(expectedPassword string) func(entity.User) bool {
	return func(user entity.User) bool {
		return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(expectedPassword)) == nil
	}
}

func TestUserService_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := UserService{UserRepository: mockRepo}

	user := entity.User{
		Name:     "John Doe",
		Nick:     "johndoe",
		Email:    "john@example.com",
		Password: "password123",
	}

	mockRepo.On("CreateUser", mock.MatchedBy(senhaMatcher(user.Password))).Return(nil)

	err := userService.CreateUser(user)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := UserService{UserRepository: mockRepo}

	user := entity.User{
		ID:    1,
		Name:  "John Doe",
		Nick:  "johndoe",
		Email: "john@example.com",
	}

	mockRepo.On("GetUserByID", uint64(1)).Return(user, nil)

	result, err := userService.GetUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := UserService{UserRepository: mockRepo}

	mockRepo.On("GetUserByID", uint64(1)).Return(entity.User{}, errors.New("not found"))

	result, err := userService.GetUserByID(1)
	assert.Error(t, err)
	assert.Equal(t, entity.User{}, result)
	mockRepo.AssertExpectations(t)
}
