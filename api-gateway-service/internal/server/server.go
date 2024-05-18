package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/0x5w4/kredit-plus/api-gateway-service/config"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/kredit/service"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/middlewares"
	kafkaClient "github.com/0x5w4/kredit-plus/pkg/kafka"
	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	loggerInterceptor "github.com/0x5w4/kredit-plus/pkg/logger-interceptor"
	readerService "github.com/0x5w4/kredit-plus/reader-service/proto/reader"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type Server struct {
	cfg               *config.Config
	v                 *validator.Validate
	logger            *loggerClient.AppLogger
	echo              *echo.Echo
	mw                middlewares.MiddlewareManager
	loggerInterceptor loggerInterceptor.LoggerInterceptor
	service           *service.KreditService
	grpcClientConn    *grpc.ClientConn
	readerClient      readerService.ReaderServiceClient
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg:  cfg,
		v:    validator.New(),
		echo: echo.New(),
	}
}

func (s Server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.setupLogger()
	s.setupLoggerInterceptor()

	s.setupMiddleware()

	s.setupGrpcClient()
	defer s.grpcClientConn.Close()

	s.setupReaderServiceClient()

	kafkaProducer := kafkaClient.NewProducer(s.cfg.Kafka.Brokers)
	defer kafkaProducer.Close()

	s.setupService(kafkaProducer)

	s.setupHttpHandler(cancel)

	<-ctx.Done()
	if err := s.echo.Server.Shutdown(ctx); err != nil {
		s.logger.SLogger.Warn("echo.Server.Shutdown", err)
	}

	return nil
}
