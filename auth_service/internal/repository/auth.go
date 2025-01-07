package repository

import (
	"auth_service/internal/models"
	"auth_service/pkg/db/postgres"
	"auth_service/pkg/logger"

	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type DBAuthRepository struct {
	db     *postgres.DB
	logger logger.Logger
}

func NewAuthRepository(ctx context.Context, db *postgres.DB) *DBAuthRepository {
	return &DBAuthRepository{
		db:     db,
		logger: logger.GetLoggerFromCtx(ctx),
	}
}

func (r *DBAuthRepository) FindUserByCredentials(ctx context.Context, username string) (*models.User, error) {

	var user models.User
	query := `
		SELECT 
			u.id::text, 
			u.username,
			u.password_hash AS PasswordHash, 
			r.name AS role
		FROM users u
		LEFT JOIN roles r ON u.role_id = r.id
		WHERE u.username = $1
	`
	err := r.db.Db.GetContext(ctx, &user, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.logger.Info(ctx, fmt.Sprintf("User with username %s not found", username))
			return nil, models.ErrUserNotFound
		}

		r.logger.Info(ctx, fmt.Sprintf("Database query error: %v", err))
		return nil, fmt.Errorf("database query error: %w", err)
	}
	return &user, nil
}

func (r *DBAuthRepository) FindUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {

	var user models.User
	query := `
		SELECT 
			u.id::text, 
			u.username, 
			r.name AS role
		FROM users u
		LEFT JOIN roles r ON u.role_id = r.id
		WHERE u.id = $1
	`
	err := r.db.Db.GetContext(ctx, &user, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.logger.Info(ctx, fmt.Sprintf("User with userID %s not found", userID))
			return nil, models.ErrUserNotFound
		}

		r.logger.Info(ctx, fmt.Sprintf("Database query error: %v", err))
		return nil, fmt.Errorf("database query error: %w", err)
	}
	return &user, nil
}

func (r *DBAuthRepository) RegisterUser(ctx context.Context, req *models.RegisterUserRequest) (*models.RegisterUserResponse, error) {

	tx, err := r.db.Db.BeginTxx(ctx, nil)
	if err != nil {
		r.logger.Info(ctx, fmt.Sprintf("Failed to begin transaction: %v", err))
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Проверяем, существует ли пользователь
	var existingUser models.User
	query := `SELECT id FROM users WHERE username = $1 OR email = $2`
	err = tx.QueryRowContext(ctx, query, req.Username, req.Email).Scan(&existingUser.ID)
	if err == nil {
		r.logger.Info(ctx, fmt.Sprintf("User with username %s or email %s already exists", req.Username, req.Email))
		return nil, models.ErrUserAlreadyExists
	} else if err != sql.ErrNoRows {
		r.logger.Info(ctx, fmt.Sprintf("Unexpected error checking user existence: %v", err)) // Логируем ошибку
		return nil, fmt.Errorf("unexpected error checking user existence: %w", err)
	}

	// Вставляем нового пользователя
	query = `
        INSERT INTO users (username, password_hash, email)
        VALUES ($1, $2, $3)
        RETURNING id
    `
	var userID uuid.UUID
	err = tx.QueryRowContext(ctx, query, req.Username, req.Password, req.Email).Scan(&userID)
	if err != nil {
		r.logger.Info(ctx, fmt.Sprintf("Error inserting new user: %v", err)) // Логируем ошибку
		return nil, fmt.Errorf("error inserting new user: %w", err)
	}

	r.logger.Info(ctx, fmt.Sprintf("User %s successfully registered", req.Username)) // Логируем успех
	return &models.RegisterUserResponse{UserID: userID.String()}, nil
}

// func (r *DBAuthRepository) SaveRevokedToken(ctx context.Context, tokenHash string, expiresAt time.Time) error {
//
// 	query := `
// 		INSERT INTO revoked_tokens (token, expires_at)
// 		VALUES ($1, $2)
// 		ON CONFLICT (token) DO NOTHING;  -- Игнорируем конфликт по уникальному токену
// 	`
// 	_, err := r.db.Db.ExecContext(ctx, query, tokenHash, expiresAt)
// 	if err != nil {
// 		// log.Error(ctx, fmt.Sprintf("failed to save revoked token: %v", err))
// 		return fmt.Errorf("failed to save revoked token: %w", err)
// 	}
//
// 	return nil
// }

// func (r *DBAuthRepository) FindRevokedToken(ctx context.Context, tokenHash string) (*models.RevokedToken, error) {
// 	var revokedToken models.RevokedToken
// 	query := `
// 		SELECT token, expires_at
// 		FROM revoked_tokens
// 		WHERE token = $1
// 	`
// 	err := r.db.Db.GetContext(ctx, &revokedToken, query, tokenHash)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, nil // Токен не найден
// 		}
// 		return nil, fmt.Errorf("error checking revoked token: %w", err)
// 	}
// 	return &revokedToken, nil
// }
