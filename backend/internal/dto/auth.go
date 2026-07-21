package dto

type RegisterRequest struct {
	FirstName string `json:"first_name" binding:"required"`

	LastName string `json:"last_name" binding:"required"`

	Email string `json:"email" binding:"required,email"`

	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`

	Password string `json:"password" binding:"required"`
}

// LoginResponse is returned after successful authentication
type LoginResponse struct {
	AccessToken string `json:"access_token"`

	RefreshToken string `json:"refresh_token"`
}

// TokenResponse is used when only returning a new access token
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// RefreshRequest is used for refresh and logout operations
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
