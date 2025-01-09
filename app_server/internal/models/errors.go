package models

import "errors"

var (
	ErrMissingAuthToken = errors.New("auth token is missing")
	ErrEmptyUserID      = errors.New("userID cannot be empty")
	ErrInvalidToken     = errors.New("invalid or expired token")

	ErrUnexpected = errors.New("unexpected error occurred")

	ErrInternal = errors.New("internal server error")
	ErrDBQuery  = errors.New("database query error")

	ErrCategoriesNotFound          = errors.New("categories not found")
	ErrAdvertisementsNotFound      = errors.New("advertisements not found")
	ErrAdvertisementNotFound       = errors.New("advertisement not found")
	ErrAdvertisementNotPublishable = errors.New("advertisement cannot be published because of its current status")
	ErrAdvertisementAlreadyDraft   = errors.New("advertisement is already in draft status")
)
