package server

import (
	"context"

	grpcClient "github.com/0x5w4/kredit-plus/pkg/grpc-client"
	grpcServer "github.com/0x5w4/kredit-plus/pkg/grpc-server"
	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	loggerInterceptor "github.com/0x5w4/kredit-plus/pkg/logger-interceptor"
)

func (s *Server) setupLogger() error {
	var err error

	s.appLogger, err = loggerClient.NewAppLogger(s.cfg.Logger)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) setupLoggerInterceptor() {
	s.loggerInterceptor = loggerInterceptor.NewLoggerInterceptor(s.appLogger)
}

func (s *Server) setupGrpcServer() error {
	var err error

	s.grpcServer, err = grpcServer.NewGrpcServer(s.cfg.GrpcServer, s.loggerInterceptor, s.appLogger)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) setupGrpcClient(ctx context.Context) error {
	var err error

	s.grpcClient, err = grpcClient.NewGrpcClient(ctx, s.cfg.GrpcServer, s.loggerInterceptor, s.appLogger)
	if err != nil {
		return err
	}

	return nil
}
