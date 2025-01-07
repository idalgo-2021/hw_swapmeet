package grpc

import (
	"auth_service/pkg/logger"
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func ContextWithLogger(l logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		// Добавляем логгер в контекст
		ctx = context.WithValue(ctx, logger.LoggerKey, l)

		// Логируем начало запроса
		l.Info(ctx, "request started", zap.String("method", info.FullMethod))

		return handler(ctx, req)
	}
}
