package handlers

import (
	"context"
	"net/http"
	"strings"

	"api_gateway/internal/grpc_clients"
	"api_gateway/internal/models"

	pb "api_gateway/pkg/api/swapmeet"
	"api_gateway/pkg/logger"

	"github.com/gorilla/mux"
	"google.golang.org/grpc/metadata"
)

type SwapmeetHandlers struct {
	client *grpc_clients.SwapmeetClient
	logger logger.Logger
}

func NewSwapmeetHandlers(ctx context.Context, client *grpc_clients.SwapmeetClient) *SwapmeetHandlers {
	return &SwapmeetHandlers{
		client: client,
		logger: logger.GetLoggerFromCtx(ctx),
	}
}

func (h *SwapmeetHandlers) enrichContextWithAuth(ctx context.Context, r *http.Request) context.Context {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		md := metadata.Pairs("authorization", authHeader)
		return metadata.NewOutgoingContext(ctx, md)
	}
	return ctx
}

// @Summary Get categories
// @Description Retrieve a list of available advertisement categories
// @Tags Categories
// @Produce json
// @Success 200 {array} pb.Category "List of categories"
// @Failure 500 {string} string "Internal server error"
// @Router /categories [get]
func (h *SwapmeetHandlers) GetCategories(w http.ResponseWriter, r *http.Request) {
	grpcReq := &pb.GetCategoriesRequest{}
	resp, err := h.client.GetCategories(r.Context(), grpcReq)
	if err != nil {
		h.handleGRPCError(r.Context(), err, w)
		return
	}
	h.writeJSONResponse(w, r, http.StatusOK, resp.Categories)
}

// @Summary Create category
// @Description Create a new advertisement category (requires authentication)
// @Tags Categories
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Access Token"
// @Param request body models.CreateCategoryRequest true "Create Category Request"
// @Success 201 {object} pb.CreateCategoryResponse "Category created successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /categories [post]
func (h *SwapmeetHandlers) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var httpReq models.CreateCategoryRequest
	if err := h.decodeJSONBody(r, &httpReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	name := strings.TrimSpace(httpReq.Name)
	if name == "" {
		h.logger.Info(r.Context(), "Missing category name")
		http.Error(w, "Category name is required", http.StatusBadRequest)
		return
	}

	parent_id := strings.TrimSpace(httpReq.ParentId)
	if parent_id == "" {
		h.logger.Info(r.Context(), "Missing parent_id")
		http.Error(w, "ParentId is required", http.StatusBadRequest)
		return
	}

	grpcReq := &pb.CreateCategoryRequest{
		Name:     name,
		ParentId: parent_id,
	}

	ctx := h.enrichContextWithAuth(r.Context(), r)
	resp, err := h.client.CreateCategory(ctx, grpcReq)
	if err != nil {
		h.handleGRPCError(r.Context(), err, w)
		return
	}

	h.writeJSONResponse(w, r, http.StatusOK, resp)
}

// @Summary Get published advertisements
// @Description Retrieve a list of published advertisements
// @Tags Advertisements
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer Access Token"
// @Param cat query []int false "Category IDs to filter the advertisements (e.g., ?cat=3&cat=15)" style(form) explode(true)
// @Success 200 {array} pb.PublishedAdvertisement "List of published advertisements"
// @Failure 500 {string} string "Internal server error"
// @Router /advertisements [get]
func (h *SwapmeetHandlers) GetPublishedAdvertisements(w http.ResponseWriter, r *http.Request) {
	categoryIDs := r.URL.Query()["cat"]
	grpcReq := &pb.GetPublishedAdvertisementsRequest{
		CategoryIds: categoryIDs,
	}

	ctx := h.enrichContextWithAuth(r.Context(), r)
	resp, err := h.client.GetPublishedAdvertisements(ctx, grpcReq)
	if err != nil {
		h.handleGRPCError(r.Context(), err, w)
		return
	}

	h.writeJSONResponse(w, r, http.StatusOK, resp)
}

// @Summary Get published advertisement by ID
// @Description Retrieve a published advertisement by its ID
// @Tags Advertisements
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer Access Token"
// @Param id path string true "Advertisement ID"
// @Success 200 {object} pb.PublishedAdvertisement "Advertisement details"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Advertisement not found"
// @Failure 500 {string} string "Internal server error"
// @Router /advertisement/{id} [get]
func (h *SwapmeetHandlers) GetPublishedAdvertisementByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adID, ok := vars["id"]
	if !ok || adID == "" {
		http.Error(w, "Missing 'id' parameter in route", http.StatusBadRequest)
		return
	}

	grpcReq := &pb.GetPublishedAdvertisementByIDRequest{
		Id: adID,
	}

	ctx := h.enrichContextWithAuth(r.Context(), r)
	resp, err := h.client.GetPublishedAdvertisementByID(ctx, grpcReq)
	if err != nil {
		h.handleGRPCError(r.Context(), err, w)
		return
	}

	h.writeJSONResponse(w, r, http.StatusOK, resp)
}

// @Summary Get User Advertisements
// @Description Retrieve a list of advertisements created by the authenticated user
// @Tags Advertisements
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Access Token"
// @Success 200 {array} pb.UserAdvertisement "List of user advertisements"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /advertisements/user [get]
func (h *SwapmeetHandlers) GetUserAdvertisements(w http.ResponseWriter, r *http.Request) {
	grpcReq := &pb.GetUserAdvertisementsRequest{}

	ctx := h.enrichContextWithAuth(r.Context(), r)

	resp, err := h.client.GetUserAdvertisements(ctx, grpcReq)
	if err != nil {
		h.handleGRPCError(r.Context(), err, w)
		return
	}

	h.writeJSONResponse(w, r, http.StatusOK, resp)
}

// @Summary Create Advertisement
// @Description Create a new advertisement (requires authentication)
// @Tags Advertisements
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Access Token"
// @Param request body models.CreateAdvertisementRequest true "Create Advertisement Request"
// @Success 201 {object} pb.CreateAdvertisementResponse "Advertisement created successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /advertisements [post]
func (h *SwapmeetHandlers) CreateAdvertisement(w http.ResponseWriter, r *http.Request) {

	var httpReq models.CreateAdvertisementRequest
	if err := h.decodeJSONBody(r, &httpReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	grpcReq := &pb.CreateAdvertisementRequest{
		CategoryId:  httpReq.CategoryId,
		Title:       httpReq.Title,
		Description: httpReq.Description,
		Price:       httpReq.Price,
		ContactInfo: httpReq.ContactInfo,
	}

	ctx := h.enrichContextWithAuth(r.Context(), r)
	resp, err := h.client.CreateAdvertisement(ctx, grpcReq)
	if err != nil {
		h.handleGRPCError(r.Context(), err, w)
		return
	}

	h.writeJSONResponse(w, r, http.StatusOK, resp)
}

// @Summary Update advertisement
// @Description Update an existing advertisement (requires authentication)
// @Tags Advertisements
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Access Token"
// @Param request body models.UpdateAdvertisementRequest true "Update Advertisement Request"
// @Success 200 {object} pb.UpdateAdvertisementResponse "Advertisement updated successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Advertisement not found"
// @Failure 500 {string} string "Internal server error"
// @Router /advertisements [put]
func (h *SwapmeetHandlers) UpdateAdvertisement(w http.ResponseWriter, r *http.Request) {

	var httpReq models.UpdateAdvertisementRequest
	if err := h.decodeJSONBody(r, &httpReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	grpcReq := &pb.UpdateAdvertisementRequest{
		AdvertisementId: httpReq.AdvertisementId,
		Price:           httpReq.Price,
		Title:           httpReq.Title,
		Description:     httpReq.Description,
		ContactInfo:     httpReq.ContactInfo,
	}

	ctx := h.enrichContextWithAuth(r.Context(), r)
	resp, err := h.client.UpdateAdvertisement(ctx, grpcReq)
	if err != nil {
		h.handleGRPCError(r.Context(), err, w)
		return
	}

	h.writeJSONResponse(w, r, http.StatusOK, resp)
}

// @Summary Submit advertisement for moderation
// @Description Move an advertisement to the "moderation" status (requires authentication)
// @Tags Advertisements
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Access Token"
// @Param id path string true "Advertisement ID"
// @Success 200 {object} pb.SubmitAdvertisementForModerationResponse "Advertisement submitted for moderation successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Advertisement not found"
// @Failure 500 {string} string "Internal server error"
// @Router /advertisement/moderation/{id} [put]
func (h *SwapmeetHandlers) SubmitAdvertisementForModeration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adID, ok := vars["id"]
	if !ok || adID == "" {
		http.Error(w, "Missing 'id' parameter in route", http.StatusBadRequest)
		return
	}

	grpcReq := &pb.SubmitAdvertisementForModerationRequest{
		AdvertisementId: adID,
	}

	ctx := h.enrichContextWithAuth(r.Context(), r)
	resp, err := h.client.SubmitAdvertisementForModeration(ctx, grpcReq)
	if err != nil {
		h.handleGRPCError(r.Context(), err, w)
		return
	}

	h.writeJSONResponse(w, r, http.StatusOK, resp)
}
