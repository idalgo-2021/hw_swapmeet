package models

import "errors"

var (
	ErrInvalidCredentials  = errors.New("invalid username or password")
	ErrInvalidToken        = errors.New("invalid or expired token")
	ErrInvalidRefreshToken = errors.New("invalid or expired refresh token")
	ErrInvalidUserID       = errors.New("invalid user ID format")

	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrUnexpected         = errors.New("unexpected error occurred")
	ErrEmailAlreadyExists = errors.New("user with this email already exists")
	ErrInvalidEmail       = errors.New("email is required")
	ErrInternal           = errors.New("internal server error")
)
