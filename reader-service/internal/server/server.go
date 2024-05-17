package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	grpcServer "github.com/0x5w4/kredit-plus/pkg/grpc-server"
	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	loggerInterceptor "github.com/0x5w4/kredit-plus/pkg/logger-interceptor"
	"github.com/0x5w4/kredit-plus/reader-service/config"
)

type Server struct {
	cfg               *config.Config
	appLogger         *loggerClient.AppLogger
	readerService     *grpcServer.GrpcServer
	loggerInterceptor loggerInterceptor.LoggerInterceptor
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s Server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.setupLogger()
	s.setupLoggerInterceptor()

	return nil
}
