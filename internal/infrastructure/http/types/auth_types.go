package types

// AuthCredentials represents the credentials for authentication.
type AuthCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// AuthResponse represents the response for a successful authentication.
type AuthResponse struct {
	ID    uint64 `json:"id"`
	Token string `json:"token"`
}
