package ports

import "api-social-network/internal/domain/entity"

// UserRepository defines the methods that any data storage provider needs to implement to get and store users.
type UserRepository interface {
	CreateUser(user entity.User) error
	GetUserByID(id uint64) (entity.User, error)
	GetAllUsers() ([]entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id uint64) error
	GetUserByEmail(email string) (entity.User, error)
}
