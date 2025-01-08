package router

import (
	"api_gateway/internal/grpc_clients"
	"api_gateway/internal/handlers"
	"api_gateway/pkg/logger"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(ctx context.Context, authClient *grpc_clients.AuthClient, swapmeetClient *grpc_clients.SwapmeetClient) *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Auth-service routes
	authHandlers := handlers.NewAuthHandlers(ctx, authClient)
	r.HandleFunc("/auth/token", authHandlers.GenerateToken).Methods(http.MethodPost)
	r.HandleFunc("/auth/validate", authHandlers.ValidateToken).Methods(http.MethodPost)
	r.HandleFunc("/auth/refresh", authHandlers.RefreshToken).Methods(http.MethodPost)
	r.HandleFunc("/auth/register", authHandlers.RegisterUser).Methods(http.MethodPost)

	// SwapMeet-service routes
	swapmeetHandlers := handlers.NewSwapmeetHandlers(ctx, swapmeetClient)
	r.HandleFunc("/categories", swapmeetHandlers.GetCategories).Methods(http.MethodGet)
	r.Handle("/categories", handlers.AuthMiddleware(authClient)(http.HandlerFunc(swapmeetHandlers.CreateCategory))).Methods(http.MethodPost)

	r.HandleFunc("/advertisements", swapmeetHandlers.GetPublishedAdvertisements).Methods(http.MethodGet)
	r.HandleFunc("/advertisement/{id:[0-9]+}", swapmeetHandlers.GetPublishedAdvertisementByID).Methods(http.MethodGet)

	r.Handle("/advertisements/user", handlers.AuthMiddleware(authClient)(http.HandlerFunc(swapmeetHandlers.GetUserAdvertisements))).Methods(http.MethodGet)
	r.Handle("/advertisements", handlers.AuthMiddleware(authClient)(http.HandlerFunc(swapmeetHandlers.CreateAdvertisement))).Methods(http.MethodPost)
	r.Handle("/advertisements", handlers.AuthMiddleware(authClient)(http.HandlerFunc(swapmeetHandlers.UpdateAdvertisement))).Methods(http.MethodPut)
	r.Handle("/advertisement/{id}/submit-for-moderation", handlers.AuthMiddleware(authClient)(http.HandlerFunc(swapmeetHandlers.SubmitAdvertisementForModeration))).Methods(http.MethodPut)

	r.Handle("/advertisements/moderation", handlers.AuthMiddleware(authClient)(http.HandlerFunc(swapmeetHandlers.GetModerationAdvertisements))).Methods(http.MethodGet)
	r.Handle("/advertisement/{id:[0-9]+}/publish", handlers.AuthMiddleware(authClient)(http.HandlerFunc(swapmeetHandlers.PublishAdvertisement))).Methods(http.MethodPut)
	r.Handle("/advertisement/{id:[0-9]+}/return-to-draft", handlers.AuthMiddleware(authClient)(http.HandlerFunc(swapmeetHandlers.ReturnAdvertisementToDraft))).Methods(http.MethodPut)

	r.Use(loggingMiddleware(ctx))

	return r
}

func loggingMiddleware(ctx context.Context) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := logger.GetLoggerFromCtx(ctx)
			logger.Info(ctx, "Incoming request: "+r.Method+" "+r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}
}
