package grpc

import (
	"auth_service/internal/models"
	client "auth_service/pkg/api/auth_grpc"
)

func grpcToModelGenerateTokenRequest(req *client.GenerateTokenRequest) *models.GenerateTokenRequest {
	return &models.GenerateTokenRequest{
		Username: req.Username,
		Password: req.Password,
	}
}

func modelToGrpcGenerateTokenResponse(resp *models.GenerateTokenResponse) *client.GenerateTokenResponse {
	return &client.GenerateTokenResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}
}

func grpcToModelValidateTokenRequest(req *client.ValidateTokenRequest) *models.ValidateTokenRequest {
	return &models.ValidateTokenRequest{
		AccessToken: req.AccessToken,
	}
}

func modelToGrpcValidateTokenResponse(resp *models.ValidateTokenResponse) *client.ValidateTokenResponse {
	return &client.ValidateTokenResponse{
		UserId: resp.UserID,
	}
}

func grpcToModelRefreshTokenRequest(req *client.RefreshTokenRequest) *models.RefreshTokenRequest {
	return &models.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	}
}

func modelToGrpcRefreshTokenResponse(resp *models.RefreshTokenResponse) *client.RefreshTokenResponse {
	return &client.RefreshTokenResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}
}

func grpcToModelRegisterUserRequest(req *client.RegisterUserRequest) *models.RegisterUserRequest {
	return &models.RegisterUserRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
}

func modelToGrpcRegisterUserResponse(resp *models.RegisterUserResponse) *client.RegisterUserResponse {
	return &client.RegisterUserResponse{
		UserId: resp.UserID,
	}
}

// func grpcToModelRevokeTokenRequest(req *client.RevokeTokenRequest) *models.RevokeTokenRequest {
// 	return &models.RevokeTokenRequest{
// 		AccessToken: req.AccessToken,
// 	}
// }

// func modelToGrpcRevokeTokenResponse(resp *models.RevokeTokenResponse) *client.RevokeTokenResponse {
// 	return &client.RevokeTokenResponse{
// 		Success: resp.Success,
// 	}
// }
