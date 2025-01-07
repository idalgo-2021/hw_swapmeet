package handlers

import (
	"context"
	"net/http"
	"strings"

	"api_gateway/internal/grpc_clients"
	"api_gateway/internal/models"

	pb "api_gateway/pkg/api/auth"
	"api_gateway/pkg/logger"
)

type AuthHandlers struct {
	client *grpc_clients.AuthClient
	logger logger.Logger
}

func NewAuthHandlers(ctx context.Context, client *grpc_clients.AuthClient) *AuthHandlers {
	return &AuthHandlers{
		client: client,
		logger: logger.GetLoggerFromCtx(ctx),
	}
}

// GenerateToken
// @Summary Generate a new access token
// @Description Generates a new JWT access token for authentication.
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.GenerateTokenRequest true "Login details"
// @Success 200 {object} pb.GenerateTokenResponse
// @Failure 400 {string} string "Invalid request format"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/token [post]
func (h *AuthHandlers) GenerateToken(w http.ResponseWriter, r *http.Request) {
	var httpReq models.GenerateTokenRequest
	if err := h.decodeJSONBody(r, &httpReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	username := strings.TrimSpace(httpReq.Username)
	password := strings.TrimSpace(httpReq.Password)
	if username == "" || password == "" {
		h.logger.Info(r.Context(), "Missing username or password")
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	grpcReq := &pb.GenerateTokenRequest{
		Username: username,
		Password: password,
	}

	resp, err := h.client.GenerateToken(r.Context(), grpcReq)
	if err != nil {
		h.handleGRPCError(r.Context(), err, w)
		return
	}

	h.writeJSONResponse(w, r, http.StatusOK, resp)
}

// ValidateToken
// @Summary Validate an access token
// @Description Validates the provided JWT access token.
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Access Token"
// @Success 200 {object} pb.ValidateTokenResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/validate [post]
func (h *AuthHandlers) ValidateToken(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header is required", http.StatusUnauthorized)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		http.Error(w, "Invalid Authorization format. Expected 'Bearer <token>'", http.StatusUnauthorized)
		return
	}

	accessToken := parts[1]

	req := &pb.ValidateTokenRequest{
		AccessToken: accessToken,
	}

	resp, err := h.client.ValidateToken(r.Context(), req)
	if err != nil {
		h.handleGRPCError(r.Context(), err, w)
		return
	}

	h.writeJSONResponse(w, r, http.StatusOK, resp)
}

// RefreshToken
// @Summary Refresh an existing JWT token
// @Description Refresh a JWT token using a refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} pb.RefreshTokenResponse
// @Failure 400 {string} string "Invalid request format"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/refresh [post]
func (h *AuthHandlers) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var httpReq models.RefreshTokenRequest
	if err := h.decodeJSONBody(r, &httpReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	refresh_token := strings.TrimSpace(httpReq.RefreshToken)
	if refresh_token == "" {
		h.logger.Info(r.Context(), "Missing refresh token")
		http.Error(w, "Refresh token is required", http.StatusBadRequest)
		return
	}

	grpcReq := &pb.RefreshTokenRequest{
		RefreshToken: httpReq.RefreshToken,
	}

	resp, err := h.client.RefreshToken(r.Context(), grpcReq)
	if err != nil {
		h.handleGRPCError(r.Context(), err, w)
		return
	}

	h.writeJSONResponse(w, r, http.StatusOK, resp)
}

// RegisterUser
// @Summary Register a new user
// @Description Register a user with username, password, and email
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.RegisterUserRequest true "User details"
// @Success 200 {object} pb.RegisterUserResponse
// @Failure 400 {string} string "Invalid request format"
// @Failure 409 {string} string "User already exists"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/register [post]
func (h *AuthHandlers) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var httpReq models.RegisterUserRequest
	if err := h.decodeJSONBody(r, &httpReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	username := strings.TrimSpace(httpReq.Username)
	password := strings.TrimSpace(httpReq.Password)
	if username == "" || password == "" {
		h.logger.Info(r.Context(), "Missing username or password")
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	email := strings.TrimSpace(httpReq.Email)
	if email == "" {
		h.logger.Info(r.Context(), "Missing email")
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	grpcReq := &pb.RegisterUserRequest{
		Username: username,
		Password: password,
		Email:    email,
	}

	resp, err := h.client.RegisterUser(r.Context(), grpcReq)
	if err != nil {
		h.handleGRPCError(r.Context(), err, w)
		return
	}

	h.writeJSONResponse(w, r, http.StatusOK, resp)
}
