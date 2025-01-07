package grpc_clients

import (
	"context"
	"fmt"

	pb "api_gateway/pkg/api/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	client pb.AuthServiceClient
}

func NewAuthClient(addr string) (*AuthClient, error) {

	clientOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(addr, clientOptions...)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to AuthService: %w", err)
	}

	return &AuthClient{
		client: pb.NewAuthServiceClient(conn),
	}, nil
}

func (c *AuthClient) GenerateToken(ctx context.Context, req *pb.GenerateTokenRequest, opts ...grpc.CallOption) (*pb.GenerateTokenResponse, error) {
	return c.client.GenerateToken(ctx, req, opts...)
}

func (c *AuthClient) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest, opts ...grpc.CallOption) (*pb.ValidateTokenResponse, error) {
	return c.client.ValidateToken(ctx, req, opts...)
}

func (c *AuthClient) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest, opts ...grpc.CallOption) (*pb.RefreshTokenResponse, error) {
	return c.client.RefreshToken(ctx, req, opts...)
}

func (c *AuthClient) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest, opts ...grpc.CallOption) (*pb.RegisterUserResponse, error) {
	return c.client.RegisterUser(ctx, req, opts...)
}
