package main

import (
	_ "api_gateway/docs"
	"api_gateway/internal/config"
	"api_gateway/internal/gateway"
	"api_gateway/pkg/logger"
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	serviceName = "GATEWAY"
)

// @title Gateway HTTP API Swapmeet
// @version 1.0
// @description Gateway HTTP API Swapmeet
// @BasePath /
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

	gtw, err := gateway.New(ctx, cfg)
	if err != nil {
		mainLogger.Error(ctx, "failed to init gateway: "+err.Error())
		return
	}

	graceCh := make(chan os.Signal, 1)
	signal.Notify(graceCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		mainLogger.Info(ctx, "Starting gateway...")
		mainLogger.Info(ctx, "Listening on "+cfg.HTTPServerAddress+":"+strconv.Itoa(cfg.HTTPServerPort))
		if err := gtw.Run(); err != nil {
			mainLogger.Error(ctx, "gateway stopped with error: "+err.Error())
		}
		cancel()
	}()

	<-graceCh
	mainLogger.Info(ctx, "Shutting down gracefully...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := gtw.Shutdown(shutdownCtx); err != nil {
		mainLogger.Error(ctx, "failed to shutdown gateway gracefully: "+err.Error())
	} else {
		mainLogger.Info(ctx, "Gateway shutdown completed.")
	}
}
