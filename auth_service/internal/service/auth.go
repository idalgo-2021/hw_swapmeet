package service

import (
	"auth_service/internal/config"
	"auth_service/internal/models"
	"auth_service/pkg/logger"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository interface {
	FindUserByCredentials(ctx context.Context, username string) (*models.User, error)
	FindUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	RegisterUser(ctx context.Context, req *models.RegisterUserRequest) (*models.RegisterUserResponse, error)
	// SaveRevokedToken(ctx context.Context, tokenHash string, expiresAt time.Time) error
	// FindRevokedToken(ctx context.Context, tokenHash string) (*models.RevokedToken, error)
}

type AuthService struct {
	Repo       AuthRepository
	JWTService *JWTService
	logger     logger.Logger
}

func NewAuthService(ctx context.Context, repo AuthRepository, cfg *config.Config) (*AuthService, error) {
	if cfg.JWTSecretKey == "" {
		return nil, fmt.Errorf("missing JWT secret key in config")
	}

	if cfg.JWTAccessTokenLifetime <= 0 || cfg.JWTRefreshTokenLifetime <= 0 {
		return nil, fmt.Errorf("invalid token lifetimes")
	}

	jwtService := NewJWTService(cfg.JWTSecretKey, cfg.JWTAccessTokenLifetime, cfg.JWTRefreshTokenLifetime)

	return &AuthService{
		Repo:       repo,
		JWTService: jwtService,
		logger:     logger.GetLoggerFromCtx(ctx),
	}, nil
}

func (s *AuthService) GenerateToken(ctx context.Context, req *models.GenerateTokenRequest) (*models.GenerateTokenResponse, error) {

	username := strings.TrimSpace(req.Username)
	password := strings.TrimSpace(req.Password)
	if username == "" || password == "" {
		s.logger.Info(ctx, "Invalid credentials: missing username or password")
		return nil, models.ErrInvalidCredentials
	}

	// Проверка существования пользователя в БД
	user, err := s.Repo.FindUserByCredentials(ctx, username)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			s.logger.Info(ctx, fmt.Sprintf("User %s not found", username))
			return nil, models.ErrUserNotFound
		}
		s.logger.Info(ctx, fmt.Sprintf("Unexpected error: %v", err))
		return nil, models.ErrUnexpected
	}

	// Проверка пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		s.logger.Info(ctx, "Invalid credentials: password mismatch")
		return nil, models.ErrInvalidCredentials
	}

	accessToken, refreshToken, err := s.JWTService.GenerateTokens(user)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("Token generation failed: %v", err))
		return nil, models.ErrUnexpected
	}

	return &models.GenerateTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *models.ValidateTokenRequest) (*models.ValidateTokenResponse, error) {

	accessToken := strings.TrimSpace(req.AccessToken)
	if accessToken == "" {
		s.logger.Info(ctx, "invalid access_token: is empty")
		return nil, models.ErrInvalidToken
	}

	// Валидируем токен
	claims, err := s.JWTService.ValidateToken(accessToken)
	if err != nil {
		s.logger.Info(ctx, "invalid access_token: not valid")
		return nil, models.ErrInvalidToken
	}

	// // Проверяем, был ли токен отозван
	// tokenHash := hashToken(accessToken)
	// revokedToken, err := s.Repo.FindRevokedToken(ctx, tokenHash)
	// if err == nil && revokedToken != nil {
	// 	s.logger.Info(ctx, "token has been revoked")
	// 	return nil, models.ErrInvalidToken
	// }

	// Проверяем, что в пайлоаде есть userID
	userID, ok := claims["user_id"].(string)
	if !ok || userID == "" {
		s.logger.Info(ctx, "user_id missing or invalid in token")
		return nil, models.ErrInvalidToken
	}

	// Существование userID в БД не проверяем

	return &models.ValidateTokenResponse{
		UserID: userID,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *models.RefreshTokenRequest) (*models.RefreshTokenResponse, error) {

	refreshToken := strings.TrimSpace(req.RefreshToken)
	if refreshToken == "" {
		s.logger.Info(ctx, "invalid refresh_token: is empty")
		return nil, models.ErrInvalidRefreshToken
	}

	// Валидация токена
	claims, err := s.JWTService.ValidateToken(refreshToken)
	if err != nil {
		s.logger.Info(ctx, "invalid refresh_token: not valid")
		return nil, models.ErrInvalidRefreshToken
	}

	// // Проверяем, был ли токен отозван
	// tokenHash := hashToken(refreshToken)
	// revokedToken, err := s.Repo.FindRevokedToken(ctx, tokenHash)
	// if err == nil && revokedToken != nil {
	// 	s.logger.Info(ctx, "token has been revoked")
	// 	return nil, models.ErrInvalidToken
	// }

	// Проверяем, что в пайлоаде есть userID
	userIDStr, ok := claims["user_id"].(string)
	if !ok || userIDStr == "" {
		s.logger.Info(ctx, "user_id missing or invalid in token")
		return nil, models.ErrInvalidToken
	}

	// Проверяем существование в БД пользователя с userID(из payload токена)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("invalid user ID format: %v", err))
		return nil, models.ErrInvalidUserID
	}

	user, err := s.Repo.FindUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			return nil, models.ErrUserNotFound
		}
		s.logger.Info(ctx, fmt.Sprintf("unexpected error: %v", err))
		return nil, models.ErrUnexpected
	}

	accessToken, refreshToken, err := s.JWTService.GenerateTokens(user)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("token generation failed: %v", err))
		return nil, models.ErrUnexpected
	}

	return &models.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *AuthService) RegisterUser(ctx context.Context, req *models.RegisterUserRequest) (*models.RegisterUserResponse, error) {

	username := strings.TrimSpace(req.Username)
	password := strings.TrimSpace(req.Password)
	if username == "" || password == "" {
		s.logger.Info(ctx, "Invalid credentials: missing username or password")
		return nil, models.ErrInvalidCredentials
	}

	email := strings.TrimSpace(req.Email)
	if email == "" {
		s.logger.Info(ctx, "Invalid email: email is required")
		return nil, models.ErrInvalidEmail
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("Error hashing password: %v", err))
		return nil, models.ErrUnexpected
	}

	req.Password = hashedPassword
	return s.Repo.RegisterUser(ctx, req)
}

func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

// func (s *AuthService) RevokeToken(ctx context.Context, req *models.RevokeTokenRequest) (*models.RevokeTokenResponse, error) {
//
// 	accessToken := strings.TrimSpace(req.AccessToken)
// 	if accessToken == "" {
// 		s.logger.Info(ctx, "invalid access_token: is empty")
// 		return nil, models.ErrInvalidToken
// 	}
//
// 	claims, err := s.JWTService.ValidateToken(accessToken)
// 	if err != nil {
// 		s.logger.Info(ctx, "invalid access_token: not valid")
// 		return nil, models.ErrInvalidToken
//
// 	}
//
// 	tokenHash := hashToken(accessToken)
// 	expiresAt := time.Unix(int64(claims["exp"].(float64)), 0)
//
// 	revokedToken, err := s.Repo.FindRevokedToken(ctx, tokenHash)
// 	if err != nil {
// 		s.logger.Info(ctx, fmt.Sprintf("failed to find revoked token: %v", err))
// 		return nil, models.ErrUnexpected
// 	}
// 	if revokedToken != nil {
// 		s.logger.Info(ctx, "token has already been revoked")
// 		return nil, models.ErrInvalidToken
// 	}
//
// 	if err := s.Repo.SaveRevokedToken(ctx, tokenHash, expiresAt); err != nil {
// 		s.logger.Info(ctx, fmt.Sprintf("failed to save revoked token: %v", err))
// 		return nil, models.ErrInternal
// 	}
//
// 	return &models.RevokeTokenResponse{Success: true}, nil
// }
