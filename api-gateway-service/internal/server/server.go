package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/0x5w4/kredit-plus/api-gateway-service/config"
	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	loggerInterceptor "github.com/0x5w4/kredit-plus/pkg/logger-interceptor"
	"google.golang.org/grpc"
)

type Server struct {
	cfg               *config.Config
	appLogger         *loggerClient.AppLogger
	readerServiceConn *grpc.ClientConn
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

	s.setupReaderServiceConn()
	defer func() {
		if err := s.readerServiceConn.Close(); err != nil {
			s.appLogger.SLogger.Fatalf("Failed to close reader service connection:%v", err)
		}
	}()

	return nil
}
