package grpc

import (
	"app_server/pkg/logger"
	"context"
	"fmt"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ContextWithLogger(l logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctx = context.WithValue(ctx, logger.LoggerKey, l)

		l.Info(ctx, "request started", zap.String("method", info.FullMethod))

		return handler(ctx, req)
	}
}

func extractAuthToken(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, fmt.Errorf("missing metadata")
	}

	tokenStrs := md["authorization"]
	if len(tokenStrs) == 0 {
		return ctx, fmt.Errorf("missing authorization token")
	}

	tokenStr := strings.TrimPrefix(tokenStrs[0], "Bearer ")
	if tokenStr == tokenStrs[0] {
		return ctx, fmt.Errorf("invalid token format")
	}

	return context.WithValue(ctx, "authToken", tokenStr), nil
}
