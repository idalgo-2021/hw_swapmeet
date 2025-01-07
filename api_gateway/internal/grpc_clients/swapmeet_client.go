package grpc_clients

import (
	"context"
	"fmt"

	pb "api_gateway/pkg/api/swapmeet"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SwapmeetClient struct {
	client pb.SwapmeetServiceClient
}

func NewSwapmeetClient(addr string) (*SwapmeetClient, error) {

	clientOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(addr, clientOptions...)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to Swapmeet service: %w", err)
	}

	return &SwapmeetClient{
		client: pb.NewSwapmeetServiceClient(conn),
	}, nil

}

func (c *SwapmeetClient) GetCategories(ctx context.Context, req *pb.GetCategoriesRequest) (*pb.GetCategoriesResponse, error) {
	return c.client.GetCategories(ctx, req)
}

func (c *SwapmeetClient) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	return c.client.CreateCategory(ctx, req)
}

func (c *SwapmeetClient) GetPublishedAdvertisements(ctx context.Context, req *pb.GetPublishedAdvertisementsRequest) (*pb.GetPublishedAdvertisementsResponse, error) {
	return c.client.GetPublishedAdvertisements(ctx, req)
}

func (c *SwapmeetClient) GetPublishedAdvertisementByID(ctx context.Context, req *pb.GetPublishedAdvertisementByIDRequest) (*pb.GetPublishedAdvertisementByIDResponse, error) {
	return c.client.GetPublishedAdvertisementByID(ctx, req)
}

func (c *SwapmeetClient) GetUserAdvertisements(ctx context.Context, req *pb.GetUserAdvertisementsRequest) (*pb.GetUserAdvertisementsResponse, error) {
	return c.client.GetUserAdvertisements(ctx, req)
}

func (c *SwapmeetClient) CreateAdvertisement(ctx context.Context, req *pb.CreateAdvertisementRequest) (*pb.CreateAdvertisementResponse, error) {
	return c.client.CreateAdvertisement(ctx, req)
}

func (c *SwapmeetClient) UpdateAdvertisement(ctx context.Context, req *pb.UpdateAdvertisementRequest) (*pb.UpdateAdvertisementResponse, error) {
	return c.client.UpdateAdvertisement(ctx, req)
}

func (c *SwapmeetClient) SubmitAdvertisementForModeration(ctx context.Context, req *pb.SubmitAdvertisementForModerationRequest) (*pb.SubmitAdvertisementForModerationResponse, error) {
	return c.client.SubmitAdvertisementForModeration(ctx, req)
}
