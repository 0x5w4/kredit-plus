package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/0x5w4/kredit-plus/api-gateway-service/config"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/kredit/service"
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
	loggerInterceptor loggerInterceptor.LoggerInterceptor
	service           *service.KreditService
	grpcClientConn    *grpc.ClientConn
	readerClient      readerService.ReaderServiceClient
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
		v:   validator.New(),
	}
}

func (s Server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.setupLogger()

	s.setupLoggerInterceptor()
	s.setupGrpcClient()
	defer s.grpcClientConn.Close()

	s.setupReaderServiceClient()

	kafkaProducer := kafkaClient.NewProducer(s.cfg.Kafka.Brokers)
	defer kafkaProducer.Close()

	s.setupService(kafkaProducer)

	s.setupEcho()

	s.setupSwagger()

	s.setupHttpHandler()

	go s.runEcho(cancel)
	defer s.echo.Server.Shutdown(ctx)

	<-ctx.Done()

	return nil
}
