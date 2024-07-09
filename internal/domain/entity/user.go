package entity

import "time"

// User represents a user in social network.
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty" validate:"required,min=3,max=50"`
	Nick      string    `json:"nick,omitempty" validate:"required,min=3,max=50"`
	Email     string    `json:"email,omitempty" validate:"required,email"`
	Password  string    `json:"password,omitempty" validate:"required,min=1"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
