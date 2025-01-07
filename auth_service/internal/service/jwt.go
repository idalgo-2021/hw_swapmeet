package service

import (
	"errors"
	"time"

	"auth_service/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretKey       string
	accessLifetime  int
	refreshLifetime int
}

func NewJWTService(secretKey string, accessLifetime int, refreshLifetime int) *JWTService {
	return &JWTService{secretKey: secretKey, accessLifetime: accessLifetime, refreshLifetime: refreshLifetime}
}

func (s *JWTService) GenerateTokens(user *models.User) (string, string, error) {
	accessToken, err := s.generateToken(user, time.Duration(s.accessLifetime)*time.Second)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.generateToken(user, time.Duration(s.refreshLifetime)*time.Second)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *JWTService) generateToken(user *models.User, expiryTime time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(expiryTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.secretKey))
}

func (s *JWTService) ValidateToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
