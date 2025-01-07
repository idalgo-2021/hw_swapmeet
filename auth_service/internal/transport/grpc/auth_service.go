package grpc

import (
	"auth_service/internal/models"
	client "auth_service/pkg/api/auth_grpc"
	"auth_service/pkg/logger"
	"errors"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	GenerateToken(ctx context.Context, req *models.GenerateTokenRequest) (*models.GenerateTokenResponse, error)
	ValidateToken(ctx context.Context, req *models.ValidateTokenRequest) (*models.ValidateTokenResponse, error)
	RefreshToken(ctx context.Context, req *models.RefreshTokenRequest) (*models.RefreshTokenResponse, error)
	RegisterUser(ctx context.Context, req *models.RegisterUserRequest) (*models.RegisterUserResponse, error)
	// RevokeToken(ctx context.Context, req *models.RevokeTokenRequest) (*models.RevokeTokenResponse, error)
}

type AuthService struct {
	client.UnimplementedAuthServiceServer
	service Service
	logger  logger.Logger
}

func NewAuthService(ctx context.Context, srv Service) *AuthService {
	return &AuthService{
		service: srv,
		logger:  logger.GetLoggerFromCtx(ctx),
	}
}

func (s *AuthService) GenerateToken(ctx context.Context, req *client.GenerateTokenRequest) (*client.GenerateTokenResponse, error) {
	modelReq := grpcToModelGenerateTokenRequest(req)
	modelResp, err := s.service.GenerateToken(ctx, modelReq)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrInvalidCredentials):
			return nil, status.Errorf(codes.Unauthenticated, "invalid credentials: %v", err)
		case errors.Is(err, models.ErrUserNotFound):
			return nil, status.Errorf(codes.NotFound, "user not found")
		case errors.Is(err, models.ErrUnexpected):
			return nil, status.Errorf(codes.Internal, "unexpected error occurred")
		case errors.Is(err, models.ErrInvalidToken):
			return nil, status.Errorf(codes.Unauthenticated, "invalid token")
		case errors.Is(err, models.ErrInvalidRefreshToken):
			return nil, status.Errorf(codes.Unauthenticated, "invalid refresh token")
		default:
			return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
		}
	}
	return modelToGrpcGenerateTokenResponse(modelResp), nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *client.ValidateTokenRequest) (*client.ValidateTokenResponse, error) {
	modelReq := grpcToModelValidateTokenRequest(req)
	modelResp, err := s.service.ValidateToken(ctx, modelReq)
	if err != nil {
		if errors.Is(err, models.ErrInvalidToken) {
			return nil, status.Errorf(codes.Unauthenticated, "invalid or expired token")
		}
		return nil, status.Errorf(codes.Internal, "failed to validate token: %v", err)
	}
	return modelToGrpcValidateTokenResponse(modelResp), nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *client.RefreshTokenRequest) (*client.RefreshTokenResponse, error) {
	modelReq := grpcToModelRefreshTokenRequest(req)
	modelResp, err := s.service.RefreshToken(ctx, modelReq)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrInvalidRefreshToken):
			return nil, status.Errorf(codes.Unauthenticated, "invalid or expired refresh token")
		case errors.Is(err, models.ErrInvalidUserID):
			return nil, status.Errorf(codes.InvalidArgument, "invalid user ID format")
		default:
			return nil, status.Errorf(codes.Internal, "failed to validate refresh token: %v", err)
		}
	}
	return modelToGrpcRefreshTokenResponse(modelResp), nil
}

func (s *AuthService) RegisterUser(ctx context.Context, req *client.RegisterUserRequest) (*client.RegisterUserResponse, error) {
	modelReq := grpcToModelRegisterUserRequest(req)
	modelResp, err := s.service.RegisterUser(ctx, modelReq)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrInvalidCredentials):
			return nil, status.Errorf(codes.InvalidArgument, "invalid username or password")
		case errors.Is(err, models.ErrInvalidEmail):
			return nil, status.Errorf(codes.InvalidArgument, "email is required")
		case errors.Is(err, models.ErrUserAlreadyExists):
			return nil, status.Errorf(codes.AlreadyExists, "user already exists")
		case errors.Is(err, models.ErrEmailAlreadyExists):
			return nil, status.Errorf(codes.AlreadyExists, "user with this email already exists")
		case errors.Is(err, models.ErrUnexpected):
			return nil, status.Errorf(codes.Internal, "unexpected error occurred")
		default:
			return nil, status.Errorf(codes.Internal, "unknown error occurred: %v", err)
		}
	}
	return modelToGrpcRegisterUserResponse(modelResp), nil
}

// func (s *AuthService) RevokeToken(ctx context.Context, req *client.RevokeTokenRequest) (*client.RevokeTokenResponse, error) {
// 	modelReq := grpcToModelRevokeTokenRequest(req)
// 	modelResp, err := s.service.RevokeToken(ctx, modelReq)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, models.ErrInvalidToken):
// 			return nil, status.Errorf(codes.Unauthenticated, "invalid or expired token")
// 		case errors.Is(err, models.ErrInternal):
// 			return nil, status.Errorf(codes.Internal, "failed to save revoked token")
// 		default:
// 			return nil, status.Errorf(codes.Internal, "failed to revoke token: %v", err)
// 		}
// 	}
// 	return modelToGrpcRevokeTokenResponse(modelResp), nil
// }
