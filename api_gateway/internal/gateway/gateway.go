package gateway

import (
	"api_gateway/internal/config"
	"api_gateway/internal/grpc_clients"
	"api_gateway/internal/router"
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type Gateway struct {
	httpServer     *http.Server
	authClient     *grpc_clients.AuthClient
	swapmeetClient *grpc_clients.SwapmeetClient
}

func New(ctx context.Context, cfg *config.Config) (*Gateway, error) {

	authClient, err := grpc_clients.NewAuthClient(cfg.AuthServiceURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth client: %w", err)
	}

	swapmeetClient, err := grpc_clients.NewSwapmeetClient(cfg.SwapmeetServiceURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create swapmeet client: %w", err)
	}

	r := router.NewRouter(ctx, authClient, swapmeetClient)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.HTTPServerAddress, strconv.Itoa(cfg.HTTPServerPort)),
		Handler: r,
	}

	return &Gateway{
		httpServer:     httpServer,
		authClient:     authClient,
		swapmeetClient: swapmeetClient,
	}, nil
}

func (gtw *Gateway) Run() error {
	if err := gtw.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to run HTTP server: %w", err)
	}
	return nil
}

func (g *Gateway) Shutdown(ctx context.Context) error {
	if g.httpServer != nil {
		return g.httpServer.Shutdown(ctx)
	}
	return nil
}
