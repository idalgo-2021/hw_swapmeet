package service

import (
	"auth_service/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

type MockAuthRepository struct {
	mock.Mock
}

func (m *MockAuthRepository) FindUserByCredentials(ctx context.Context, username string) (*models.User, error) {
	args := m.Called(ctx, username)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockAuthRepository) FindUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	args := m.Called(ctx, userID)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockAuthRepository) RegisterUser(ctx context.Context, req *models.RegisterUserRequest) (*models.RegisterUserResponse, error) {
	args := m.Called(ctx, req)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.RegisterUserResponse), args.Error(1)
}

func TestRegisterUser_Success(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	authService := &AuthService{Repo: mockRepo}

	req := &models.RegisterUserRequest{
		Username: "newuser",
		Password: "newpassword",
		Email:    "newemail@example.com",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(hashedPassword)

	mockRepo.On("RegisterUser", mock.Anything, req).Return(&models.RegisterUserResponse{UserID: uuid.New().String()}, nil)

	resp, err := authService.RegisterUser(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.UserID)

	mockRepo.AssertExpectations(t)
}

func TestRegisterUser_InvalidInput(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	authService := &AuthService{Repo: mockRepo}

	req := &models.RegisterUserRequest{
		Username: "",
		Password: "newpassword",
		Email:    "newemail@example.com",
	}

	resp, err := authService.RegisterUser(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, models.ErrInvalidCredentials, err)

	// Test missing password
	req = &models.RegisterUserRequest{
		Username: "newuser",
		Password: "",
		Email:    "newemail@example.com",
	}

	resp, err = authService.RegisterUser(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, models.ErrInvalidCredentials, err)

	req = &models.RegisterUserRequest{
		Username: "newuser",
		Password: "newpassword",
		Email:    "",
	}

	resp, err = authService.RegisterUser(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, models.ErrInvalidEmail, err)
}

func TestRegisterUser_ExistingUser(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	authService := &AuthService{Repo: mockRepo}

	req := &models.RegisterUserRequest{
		Username: "existinguser",
		Password: "newpassword",
		Email:    "newemail@example.com",
	}

	mockRepo.On("RegisterUser", mock.Anything, req).Return(nil, models.ErrUserAlreadyExists)

	resp, err := authService.RegisterUser(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, models.ErrUserAlreadyExists, err)

	mockRepo.AssertExpectations(t)
}
