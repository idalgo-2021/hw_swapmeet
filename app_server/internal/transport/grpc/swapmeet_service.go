package grpc

import (
	"app_server/internal/models"
	client "app_server/pkg/api/swapmeet_grpc"
	"app_server/pkg/logger"
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	GetCategories(ctx context.Context) ([]models.Category, error)
	CreateCategory(ctx context.Context, name string, parentID string) (*models.Category, error)

	GetPublishedAdvertisements(ctx context.Context, categoryIDs []string) ([]models.UserAdvertisement, error)
	GetPublishedAdvertisementByID(ctx context.Context, advertisementID string) (*models.UserAdvertisement, error)

	GetUserAdvertisements(ctx context.Context) ([]models.UserAdvertisement, error)
	CreateAdvertisement(ctx context.Context, categoryID, title, description, price, contactInfo string) (*models.UserAdvertisement, error)
	UpdateAdvertisement(ctx context.Context, advertisementID, title, description, price, contactInfo string) (*models.UserAdvertisement, error)
	SubmitAdvertisementForModeration(ctx context.Context, advertisementID string) (*models.UserAdvertisement, error)

	GetModerationAdvertisements(ctx context.Context, statuses, categoryIDs []string) ([]models.UserAdvertisement, error)
	PublishAdvertisement(ctx context.Context, advertisementID string) (*models.UserAdvertisement, error)
	ReturnAdvertisementToDraft(ctx context.Context, advertisementID string) (*models.UserAdvertisement, error)
}

type SwapmeetService struct {
	client.UnimplementedSwapmeetServiceServer
	service Service
	logger  logger.Logger
}

func NewSwapmeetService(ctx context.Context, srv Service) *SwapmeetService {
	return &SwapmeetService{
		service: srv,
		logger:  logger.GetLoggerFromCtx(ctx),
	}
}

////

func (s *SwapmeetService) GetCategories(ctx context.Context, req *client.GetCategoriesRequest) (*client.GetCategoriesResponse, error) {
	categories, err := s.service.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	resp := &client.GetCategoriesResponse{
		Categories: toGrpcCategories(categories),
	}

	return resp, nil
}

func (s *SwapmeetService) CreateCategory(ctx context.Context, req *client.CreateCategoryRequest) (*client.CreateCategoryResponse, error) {
	ctx, err := extractAuthToken(ctx)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("failed to extract auth token: %v", err))
	}

	newCategory, err := s.service.CreateCategory(ctx, req.Name, req.ParentId)
	if err != nil {
		return nil, err
	}

	resp := &client.CreateCategoryResponse{
		Category: toGrpcCategory(newCategory),
	}

	return resp, nil
}

////

func (s *SwapmeetService) GetPublishedAdvertisements(ctx context.Context, req *client.GetPublishedAdvertisementsRequest) (*client.GetPublishedAdvertisementsResponse, error) {
	ctx, err := extractAuthToken(ctx)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("failed to extract auth token: %v", err))
	}

	categoryIDs := req.CategoryIds

	advertisements, err := s.service.GetPublishedAdvertisements(ctx, categoryIDs)
	if err != nil {
		if errors.Is(err, models.ErrAdvertisementsNotFound) {
			return &client.GetPublishedAdvertisementsResponse{Advertisements: []*client.UserAdvertisement{}}, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to get advertisements: %v", err)
	}

	resp := &client.GetPublishedAdvertisementsResponse{
		Advertisements: toGrpcUserAdvertisements(advertisements),
	}

	return resp, nil
}

func (s *SwapmeetService) GetPublishedAdvertisementByID(ctx context.Context, req *client.GetPublishedAdvertisementByIDRequest) (*client.GetPublishedAdvertisementByIDResponse, error) {
	ctx, err := extractAuthToken(ctx)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("failed to extract auth token: %v", err))
	}

	advertisement, err := s.service.GetPublishedAdvertisementByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, models.ErrAdvertisementNotFound) {
			return nil, status.Errorf(codes.NotFound, "could not fetch advertisement: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "could not fetch advertisement: %v", err)
	}

	resp := &client.GetPublishedAdvertisementByIDResponse{
		Advertisement: toGrpcUserAdvertisement(advertisement),
	}

	return resp, nil
}

////

func (s *SwapmeetService) GetUserAdvertisements(ctx context.Context, req *client.GetUserAdvertisementsRequest) (*client.GetUserAdvertisementsResponse, error) {
	ctx, err := extractAuthToken(ctx)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("failed to extract auth token: %v", err))
	}

	advertisements, err := s.service.GetUserAdvertisements(ctx)
	if err != nil {
		if errors.Is(err, models.ErrAdvertisementsNotFound) {
			return &client.GetUserAdvertisementsResponse{Advertisements: []*client.UserAdvertisement{}}, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to get advertisements: %v", err)
	}

	resp := &client.GetUserAdvertisementsResponse{
		Advertisements: toGrpcUserAdvertisements(advertisements),
	}

	return resp, nil
}

func (s *SwapmeetService) CreateAdvertisement(ctx context.Context, req *client.CreateAdvertisementRequest) (*client.CreateAdvertisementResponse, error) {
	ctx, err := extractAuthToken(ctx)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("failed to extract auth token: %v", err))
	}

	advertisement, err := s.service.CreateAdvertisement(ctx, req.CategoryId, req.Title, req.Description, req.Price, req.ContactInfo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create advertisement: %v", err)
	}

	resp := &client.CreateAdvertisementResponse{
		Advertisement: toGrpcUserAdvertisement(advertisement),
	}

	return resp, nil
}

func (s *SwapmeetService) UpdateAdvertisement(ctx context.Context, req *client.UpdateAdvertisementRequest) (*client.UpdateAdvertisementResponse, error) {
	ctx, err := extractAuthToken(ctx)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("failed to extract auth token: %v", err))
	}

	advertisement, err := s.service.UpdateAdvertisement(ctx, req.AdvertisementId, req.Title, req.Description, req.Price, req.ContactInfo)
	if err != nil {
		if errors.Is(err, models.ErrAdvertisementNotFound) {
			return nil, status.Errorf(codes.NotFound, "could not fetch advertisement: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to update advertisement: %v", err)
	}

	resp := &client.UpdateAdvertisementResponse{
		Advertisement: toGrpcUserAdvertisement(advertisement),
	}

	return resp, nil
}

func (s *SwapmeetService) SubmitAdvertisementForModeration(ctx context.Context, req *client.SubmitAdvertisementForModerationRequest) (*client.SubmitAdvertisementForModerationResponse, error) {
	ctx, err := extractAuthToken(ctx)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("failed to extract auth token: %v", err))
	}

	advertisement, err := s.service.SubmitAdvertisementForModeration(ctx, req.AdvertisementId)
	if err != nil {
		if errors.Is(err, models.ErrAdvertisementNotFound) {
			return nil, status.Errorf(codes.NotFound, "could not fetch advertisement: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to submit advertisement for moderation: %v", err)
	}

	resp := &client.SubmitAdvertisementForModerationResponse{
		Advertisement: toGrpcUserAdvertisement(advertisement),
	}

	return resp, nil
}

////

func (s *SwapmeetService) GetModerationAdvertisements(ctx context.Context, req *client.GetModerationAdvertisementsRequest) (*client.GetModerationAdvertisementsResponse, error) {
	ctx, err := extractAuthToken(ctx)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("failed to extract auth token: %v", err))
	}

	statuses := req.Statuses
	categoryIDs := req.CategoryIds

	advertisements, err := s.service.GetModerationAdvertisements(ctx, statuses, categoryIDs)
	if err != nil {
		if errors.Is(err, models.ErrAdvertisementNotFound) {
			return nil, status.Errorf(codes.NotFound, "could not fetch advertisement: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get advertisement for moderation: %v", err)
	}

	resp := &client.GetModerationAdvertisementsResponse{
		Advertisements: toGrpcUserAdvertisements(advertisements),
	}

	return resp, nil
}

func (s *SwapmeetService) PublishAdvertisement(ctx context.Context, req *client.PublishAdvertisementRequest) (*client.PublishAdvertisementResponse, error) {
	ctx, err := extractAuthToken(ctx)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("failed to extract auth token: %v", err))
	}

	advertisement, err := s.service.PublishAdvertisement(ctx, req.AdvertisementId)
	if err != nil {
		if errors.Is(err, models.ErrAdvertisementNotFound) {
			return nil, status.Errorf(codes.NotFound, "could not fetch advertisement: %v", err)
		}
		if errors.Is(err, models.ErrAdvertisementNotPublishable) {
			return nil, status.Errorf(codes.AlreadyExists, "could not fetch advertisement: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to publish advertisement: %v", err)
	}

	resp := &client.PublishAdvertisementResponse{
		Advertisement: toGrpcUserAdvertisement(advertisement),
	}

	return resp, nil
}

func (s *SwapmeetService) ReturnAdvertisementToDraft(ctx context.Context, req *client.ReturnAdvertisementToDraftRequest) (*client.ReturnAdvertisementToDraftResponse, error) {
	ctx, err := extractAuthToken(ctx)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("failed to extract auth token: %v", err))
	}

	advertisement, err := s.service.ReturnAdvertisementToDraft(ctx, req.AdvertisementId)
	if err != nil {
		if errors.Is(err, models.ErrAdvertisementNotFound) {
			return nil, status.Errorf(codes.NotFound, "could not fetch advertisement: %v", err)
		}
		if errors.Is(err, models.ErrAdvertisementAlreadyDraft) {
			return nil, status.Errorf(codes.AlreadyExists, "could not fetch advertisement: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to publish advertisement: %v", err)
	}

	resp := &client.ReturnAdvertisementToDraftResponse{
		Advertisement: toGrpcUserAdvertisement(advertisement),
	}

	return resp, nil
}
