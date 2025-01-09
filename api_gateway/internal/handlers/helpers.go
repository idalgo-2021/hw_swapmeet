package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func extractBearerTokenFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Authorization header is required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("Invalid Authorization format. Expected 'Bearer <token>'")
	}

	return strings.TrimSpace(parts[1]), nil
}

func (h *AuthHandlers) writeJSONResponse(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.logger.Error(r.Context(), "Failed to encode response: "+err.Error())
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *SwapmeetHandlers) writeJSONResponse(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.logger.Error(r.Context(), "Failed to encode response: "+err.Error())
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *AuthHandlers) decodeJSONBody(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		h.logger.Info(r.Context(), "Invalid request body: "+err.Error())
		return err
	}
	return nil
}

func (h *SwapmeetHandlers) decodeJSONBody(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		h.logger.Info(r.Context(), "Invalid request body: "+err.Error())
		return err
	}
	return nil
}

func (h *AuthHandlers) handleGRPCError(ctx context.Context, err error, w http.ResponseWriter) {
	h.logger.Info(ctx, "gRPC call failed: "+err.Error())
	if grpcErr, ok := status.FromError(err); ok {
		switch grpcErr.Code() {
		case codes.Unauthenticated:
			http.Error(w, grpcErr.Message(), http.StatusUnauthorized)
		case codes.InvalidArgument:
			http.Error(w, grpcErr.Message(), http.StatusBadRequest)
		case codes.NotFound:
			http.Error(w, grpcErr.Message(), http.StatusNotFound)
		case codes.AlreadyExists:
			http.Error(w, grpcErr.Message(), http.StatusConflict)
		case codes.Internal:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		default:
			http.Error(w, "Unexpected error", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Unknown error occurred", http.StatusInternalServerError)
	}
}

func (h *SwapmeetHandlers) handleGRPCError(ctx context.Context, err error, w http.ResponseWriter) {
	h.logger.Info(ctx, "gRPC call failed: "+err.Error())
	if grpcErr, ok := status.FromError(err); ok {
		switch grpcErr.Code() {
		case codes.Unauthenticated:
			http.Error(w, grpcErr.Message(), http.StatusUnauthorized)
		case codes.NotFound:
			http.Error(w, grpcErr.Message(), http.StatusNotFound)
		case codes.InvalidArgument:
			http.Error(w, grpcErr.Message(), http.StatusBadRequest)
		case codes.AlreadyExists:
			http.Error(w, grpcErr.Message(), http.StatusConflict)
		default:
			http.Error(w, grpcErr.Message(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
