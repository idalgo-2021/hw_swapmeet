package repository

import (
	"auth_service/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestFindUserByCredentials(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	testUser := &models.User{
		ID:           uuid.New().String(),
		Username:     "testuser",
		PasswordHash: "hashedpassword",
	}

	mockRepo.On("FindUserByCredentials", mock.Anything, "testuser").Return(testUser, nil)

	user, err := mockRepo.FindUserByCredentials(context.Background(), "testuser")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, testUser.Username, user.Username)

	mockRepo.AssertExpectations(t)
}

func TestFindUserByID(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	testUser := &models.User{
		ID:       uuid.New().String(),
		Username: "testuser",
	}

	mockRepo.On("FindUserByID", mock.Anything, uuid.MustParse(testUser.ID)).Return(testUser, nil)

	user, err := mockRepo.FindUserByID(context.Background(), uuid.MustParse(testUser.ID))

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, testUser.Username, user.Username)

	mockRepo.On("FindUserByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, models.ErrUserNotFound)

	user, err = mockRepo.FindUserByID(context.Background(), uuid.New())

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, models.ErrUserNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestRegisterUser_Success(t *testing.T) {
	mockRepo := new(MockAuthRepository)

	req := &models.RegisterUserRequest{
		Username: "newuser",
		Password: "newpassword",
		Email:    "newemail@example.com",
	}

	testResponse := &models.RegisterUserResponse{
		UserID: uuid.New().String(),
	}

	mockRepo.On("RegisterUser", mock.Anything, req).Return(testResponse, nil)

	resp, err := mockRepo.RegisterUser(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, testResponse.UserID, resp.UserID)

	mockRepo.AssertExpectations(t)
}

func TestRegisterUser_ExistingUser(t *testing.T) {
	mockRepo := new(MockAuthRepository)

	req := &models.RegisterUserRequest{
		Username: "existinguser",
		Password: "newpassword",
		Email:    "newemail@example.com",
	}

	mockRepo.On("RegisterUser", mock.Anything, req).Return(nil, models.ErrUserAlreadyExists)

	resp, err := mockRepo.RegisterUser(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)

	mockRepo.AssertExpectations(t)
}
