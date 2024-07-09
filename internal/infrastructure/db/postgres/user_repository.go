package postgres

import (
	"api-social-network/internal/domain/entity"

	"gorm.io/gorm"
)

// UserRepositoryPostgres implements the ports.UserRepository interface.
type UserRepositoryPostgres struct {
	DB *gorm.DB
}

// CreateUser creates a new user in the database.
func (repo *UserRepositoryPostgres) CreateUser(user entity.User) error {
	return repo.DB.Create(&user).Error
}

// GetUserByID retrieves a user by its ID.
func (repo *UserRepositoryPostgres) GetUserByID(id uint64) (entity.User, error) {
	var user entity.User
	err := repo.DB.First(&user, id).Error
	return user, err
}

// GetAllUsers retrieves all users from the database.
func (repo *UserRepositoryPostgres) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	err := repo.DB.Find(&users).Error
	return users, err
}

// UpdateUser updates an existing user in the database.
func (repo *UserRepositoryPostgres) UpdateUser(user entity.User) error {
	return repo.DB.Save(&user).Error
}

// DeleteUser deletes a user by its ID.
func (repo *UserRepositoryPostgres) DeleteUser(id uint64) error {
	return repo.DB.Delete(&entity.User{}, id).Error
}

// GetUserByEmail retrieves a user by its email.
func (repo *UserRepositoryPostgres) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	err := repo.DB.Where("email = ?", email).First(&user).Error
	return user, err
}
