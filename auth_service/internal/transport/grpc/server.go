package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	client "auth_service/pkg/api/auth_grpc"
	"auth_service/pkg/logger"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
	logger     logger.Logger
}

func New(ctx context.Context, port int, service Service) (*Server, error) {
	log := logger.GetLoggerFromCtx(ctx)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Error(ctx, "Failed to listen", zap.Error(err))
		return nil, err
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			ContextWithLogger(log),
		),
	}

	grpcServer := grpc.NewServer(opts...)
	client.RegisterAuthServiceServer(grpcServer, NewAuthService(ctx, service))

	// Регистрация reflection
	reflection.Register(grpcServer)

	return &Server{
		grpcServer,
		lis,
		log,
	}, nil
}

func (s *Server) Start(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		s.logger.Info(ctx, "starting gRPC server", zap.Int("port", s.listener.Addr().(*net.TCPAddr).Port))
		if err := s.grpcServer.Serve(s.listener); err != nil {
			return fmt.Errorf("gRPC server error: %w", err)
		}
		return nil
	})

	eg.Go(func() error {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-signalChan:
			s.logger.Info(ctx, "received shutdown signal", zap.String("signal", sig.String()))
			return s.Stop(ctx)
		}
	})

	return eg.Wait()
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info(ctx, "shutting down GRPC server(gracefull)...")
	s.grpcServer.GracefulStop()

	return nil
}
