package models

import "time"

type User struct {
	ID           string
	Username     string
	Role         string
	PasswordHash string
}

type GenerateTokenRequest struct {
	Username string
	Password string
}

type GenerateTokenResponse struct {
	AccessToken  string
	RefreshToken string
}

type ValidateTokenRequest struct {
	AccessToken string
}

type ValidateTokenResponse struct {
	UserID string
}

type RefreshTokenRequest struct {
	RefreshToken string
}

type RefreshTokenResponse struct {
	AccessToken  string
	RefreshToken string
}

type RegisterUserRequest struct {
	Username string
	Password string
	Email    string
}

type RegisterUserResponse struct {
	UserID string
}

type RevokeTokenRequest struct {
	AccessToken string
}

type RevokeTokenResponse struct {
	Success bool
}

type RevokedToken struct {
	Token     string    `db:"token"`
	ExpiresAt time.Time `db:"expires_at"`
}
