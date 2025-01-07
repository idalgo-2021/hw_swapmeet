package main

import (
	"app_server/internal/config"
	"sync"

	"context"
	"os"
	"os/signal"
	"syscall"

	"app_server/internal/repository"
	service "app_server/internal/service"
	"app_server/internal/transport/grpc"
	"app_server/pkg/db/cache"
	"app_server/pkg/db/postgres"
	"app_server/pkg/logger"
)

const (
	serviceName = "APP_SERVER"
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
		mainLogger.Error(ctx, "failed to load config")
		return
	}

	db, err := postgres.New(cfg.PGConfig)
	if err != nil {
		mainLogger.Error(ctx, "failed to connect to database: "+err.Error())
		return
	}

	redis, err := cache.New(ctx, cfg.RedisConfig)
	if err != nil {
		mainLogger.Error(ctx, "failed to connect to Redis: "+err.Error())
		return
	}

	Repo := repository.NewSwapmeetRepo(ctx, db, redis)
	Serv, err := service.NewSwapmeetService(ctx, Repo, cfg)
	if err != nil {
		mainLogger.Info(ctx, "failed to initialize SwapMeet service: "+err.Error())
		return
	}

	// Очистка кэша
	mainLogger.Info(ctx, "Flushing Redis cache...")
	err = redis.FlushAll(ctx).Err()
	if err != nil {
		mainLogger.Error(ctx, "Failed to flush Redis cache: "+err.Error())
		return
	}
	mainLogger.Info(ctx, "Redis cache cleared successfully")

	// Прогрев кэша
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		mainLogger.Info(ctx, "Starting cache warm-up...")
		err := Repo.WarmUpCache(ctx)
		if err != nil {
			mainLogger.Error(ctx, "Cache warm-up failed: "+err.Error())
		} else {
			mainLogger.Info(ctx, "Cache warm-up completed successfully")
		}
	}()
	wg.Wait()

	grpcserver, err := grpc.New(ctx, cfg.GRPCServerPort, Serv)
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
