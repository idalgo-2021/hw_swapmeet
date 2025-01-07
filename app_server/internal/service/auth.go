package service

import (
	"app_server/internal/models"
	"context"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func (s *SwapmeetService) authenticate(ctx context.Context) (*models.TokenClaims, error) {
	token, ok := ctx.Value("authToken").(string)
	if !ok || strings.TrimSpace(token) == "" {
		return nil, models.ErrMissingAuthToken
	}
	claims, err := s.validateToken(token)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErrInvalidToken, err)
	}

	if strings.TrimSpace(claims.UserID) == "" {
		return nil, models.ErrEmptyUserID
	}

	return claims, nil
}

func (s *SwapmeetService) validateToken(token string) (*models.TokenClaims, error) {
	claims := &models.TokenClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(s.JWTSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	return claims, nil
}
