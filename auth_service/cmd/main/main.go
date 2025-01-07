package main

import (
	"auth_service/internal/config"
	"context"
	"os"
	"os/signal"
	"syscall"

	"auth_service/internal/repository"
	service "auth_service/internal/service"
	"auth_service/internal/transport/grpc"
	"auth_service/pkg/db/postgres"
	"auth_service/pkg/logger"
)

const (
	serviceName = "AUTH_SERVICE"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mainLogger := logger.New(serviceName)
	if mainLogger == nil {
		panic("failed to create logger")
	}

	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)

	cfg := config.New()
	if cfg == nil {
		mainLogger.Info(ctx, "failed to load config")
		return
	}

	db, err := postgres.New(cfg.Config)
	if err != nil {
		mainLogger.Error(ctx, "failed to connect to database: "+err.Error())
		return
	}
	authRepo := repository.NewAuthRepository(ctx, db)
	authService, err := service.NewAuthService(ctx, authRepo, cfg)
	if err != nil {
		mainLogger.Info(ctx, "failed to initialize auth service: "+err.Error())
		return
	}

	grpcserver, err := grpc.New(ctx, cfg.GRPCServerPort, authService)
	if err != nil {
		mainLogger.Info(ctx, "failed to create GRPC server: "+err.Error())
		return
	}

	graceCh := make(chan os.Signal, 1)
	signal.Notify(graceCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := grpcserver.Start(ctx); err != nil {
			mainLogger.Info(ctx, "failed to run GRPC server: "+err.Error())
			cancel()
		} else {
			mainLogger.Info(ctx, "Server is runned succefully")
		}
	}()

	<-graceCh
	mainLogger.Info(ctx, "Received termination signal, shutting down the server......")

	if err := grpcserver.Stop(ctx); err != nil {
		mainLogger.Info(ctx, "Error while stopping the gRPC server: "+err.Error())
	} else {
		mainLogger.Info(ctx, "Server is stopped")
	}

	mainLogger.Info(ctx, "Server succefully stopped")

}
