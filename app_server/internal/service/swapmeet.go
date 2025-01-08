package service

import (
	"app_server/internal/config"
	"app_server/internal/models"
	"app_server/pkg/logger"
	"context"
	"fmt"
	"strconv"
	"strings"
)

type SwapmeetRepo interface {
	GetCategories(ctx context.Context) ([]models.Category, error)                                           // for all
	CreateCategory(ctx context.Context, userID string, name string, parentID int) (*models.Category, error) // for admins and moderators only

	GetPublishedAdvertisements(ctx context.Context, categoryIDs []string) ([]models.UserAdvertisement, error)     // for all
	GetPublishedAdvertisementByID(ctx context.Context, advertisementID string) (*models.UserAdvertisement, error) // for all

	GetUserAdvertisements(ctx context.Context, userID string) ([]models.UserAdvertisement, error)                                                       // for authorized users
	CreateAdvertisement(ctx context.Context, userID, categoryID, title, description, price, contactInfo string) (*models.UserAdvertisement, error)      // for authorized users
	UpdateAdvertisement(ctx context.Context, userID, advertisementID, title, description, price, contactInfo string) (*models.UserAdvertisement, error) // for authorized users
	SetModerationStatusForAdvertisement(ctx context.Context, userID, advertisementID string) (*models.UserAdvertisement, error)                         // for authorized users

	EnsureUserExists(ctx context.Context, userID, userName string) error

	GetAdvertisements(ctx context.Context, statuses, categoryIDs []string) ([]models.UserAdvertisement, error) // for admins and moderators only
	PublishAdvertisement(ctx context.Context, advertisementID string) (*models.UserAdvertisement, error)       // for admins and moderators only
	ReturnAdvertisementToDraft(ctx context.Context, advertisementID string) (*models.UserAdvertisement, error) // for admins and moderators only

}

type SwapmeetService struct {
	Repo      SwapmeetRepo
	JWTSecret string
	logger    logger.Logger
}

func NewSwapmeetService(ctx context.Context, repo SwapmeetRepo, cfg *config.Config) (*SwapmeetService, error) {

	// Проверяем наличие ключа для валидации токенов
	if strings.TrimSpace(cfg.JWTSecretKey) == "" {
		return nil, fmt.Errorf("missing JWT secret key in config")
	}

	return &SwapmeetService{
		Repo:      repo,
		JWTSecret: cfg.JWTSecretKey,
		logger:    logger.GetLoggerFromCtx(ctx),
	}, nil
}

////

func (s *SwapmeetService) GetCategories(ctx context.Context) ([]models.Category, error) {
	categories, err := s.Repo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *SwapmeetService) CreateCategory(ctx context.Context, name string, parentID string) (*models.Category, error) {
	claims, err := s.authenticate(ctx)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("failed to authenticate token: %v", err))
		return nil, err
	}

	if claims.Role != "admin" && claims.Role != "moderator" {
		s.logger.Info(ctx, fmt.Sprintf("permission denied: insufficient role"))
		return nil, fmt.Errorf("permission denied: insufficient role")
	}

	// Проверка названия категории
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	// Преобразование parentID в целое число
	parentIDInt, err := strconv.Atoi(parentID)
	if err != nil {
		return nil, fmt.Errorf("invalid parent ID: %v", err)
	}

	userID := strings.TrimSpace(claims.UserID)
	if userID == "" {
		return nil, models.ErrEmptyUserID
	}

	err = s.Repo.EnsureUserExists(ctx, userID, claims.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure user exists: %v", err)
	}

	return s.Repo.CreateCategory(ctx, userID, name, parentIDInt)
}

////

func (s *SwapmeetService) GetPublishedAdvertisements(ctx context.Context, categoryIDs []string) ([]models.UserAdvertisement, error) {

	_, err := s.authenticate(ctx)
	isAuthed := err == nil

	advertisements, err := s.Repo.GetPublishedAdvertisements(ctx, categoryIDs)
	if err != nil {
		return nil, err
	}

	// Маскируем часть полей
	if !isAuthed {
		maskContactInfo(advertisements)
	}

	return advertisements, nil
}

func (s *SwapmeetService) GetPublishedAdvertisementByID(ctx context.Context, advertisementID string) (*models.UserAdvertisement, error) {

	_, err := s.authenticate(ctx)
	isAuthed := err == nil

	advertisement, err := s.Repo.GetPublishedAdvertisementByID(ctx, advertisementID)
	if err != nil {
		return nil, err
	}

	// Маскируем часть полей
	if !isAuthed {
		advertisement.ContactInfo = "Для просмотра необходимо авторизоваться"
	}

	return advertisement, nil
}

func maskContactInfo(ads []models.UserAdvertisement) {
	for i := range ads {
		ads[i].ContactInfo = "Для просмотра необходимо авторизоваться"
	}
}

////

func (s *SwapmeetService) GetUserAdvertisements(ctx context.Context) ([]models.UserAdvertisement, error) {
	claims, err := s.authenticate(ctx)
	if err != nil {
		return nil, err
	}

	userID := strings.TrimSpace(claims.UserID)
	if userID == "" {
		return nil, models.ErrEmptyUserID
	}

	return s.Repo.GetUserAdvertisements(ctx, userID)
}

func (s *SwapmeetService) CreateAdvertisement(ctx context.Context, categoryID, title, description, price, contactInfo string) (*models.UserAdvertisement, error) {
	claims, err := s.authenticate(ctx)
	if err != nil {
		return nil, err
	}

	userID := strings.TrimSpace(claims.UserID)
	if userID == "" {
		return nil, models.ErrEmptyUserID
	}

	err = s.Repo.EnsureUserExists(ctx, userID, claims.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure user exists: %v", err)
	}

	return s.Repo.CreateAdvertisement(ctx, userID, categoryID, title, description, price, contactInfo)
}

func (s *SwapmeetService) UpdateAdvertisement(ctx context.Context, advertisementID, title, description, price, contactInfo string) (*models.UserAdvertisement, error) {
	claims, err := s.authenticate(ctx)
	if err != nil {
		return nil, err
	}

	userID := strings.TrimSpace(claims.UserID)
	if userID == "" {
		return nil, models.ErrEmptyUserID
	}

	return s.Repo.UpdateAdvertisement(ctx, userID, advertisementID, title, description, price, contactInfo)
}

////

func (s *SwapmeetService) SubmitAdvertisementForModeration(ctx context.Context, advertisementID string) (*models.UserAdvertisement, error) {
	claims, err := s.authenticate(ctx)
	if err != nil {
		return nil, err
	}

	userID := strings.TrimSpace(claims.UserID)
	if userID == "" {
		return nil, models.ErrEmptyUserID
	}

	return s.Repo.SetModerationStatusForAdvertisement(ctx, userID, advertisementID)
}

////

func (s *SwapmeetService) GetModerationAdvertisements(ctx context.Context, statuses, categoryIDs []string) ([]models.UserAdvertisement, error) {
	claims, err := s.authenticate(ctx)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("failed to authenticate token: %v", err))
		return nil, err
	}

	if claims.Role != "admin" && claims.Role != "moderator" {
		s.logger.Info(ctx, fmt.Sprintf("permission denied: insufficient role"))
		return nil, fmt.Errorf("permission denied: insufficient role")
	}

	advertisements, err := s.Repo.GetAdvertisements(ctx, statuses, categoryIDs)
	if err != nil {
		return nil, err
	}

	return advertisements, nil
}

func (s *SwapmeetService) PublishAdvertisement(ctx context.Context, advertisementID string) (*models.UserAdvertisement, error) {
	claims, err := s.authenticate(ctx)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("failed to authenticate token: %v", err))
		return nil, err
	}

	if claims.Role != "admin" && claims.Role != "moderator" {
		s.logger.Info(ctx, fmt.Sprintf("permission denied: insufficient role"))
		return nil, fmt.Errorf("permission denied: insufficient role")
	}

	return s.Repo.PublishAdvertisement(ctx, advertisementID)
}

func (s *SwapmeetService) ReturnAdvertisementToDraft(ctx context.Context, advertisementID string) (*models.UserAdvertisement, error) {
	claims, err := s.authenticate(ctx)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("failed to authenticate token: %v", err))
		return nil, err
	}

	if claims.Role != "admin" && claims.Role != "moderator" {
		s.logger.Info(ctx, fmt.Sprintf("permission denied: insufficient role"))
		return nil, fmt.Errorf("permission denied: insufficient role")
	}

	return s.Repo.ReturnAdvertisementToDraft(ctx, advertisementID)
}
