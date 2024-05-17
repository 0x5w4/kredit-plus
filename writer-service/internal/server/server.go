package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	grpcServer "github.com/0x5w4/kredit-plus/pkg/grpc-server"
	kafkaClient "github.com/0x5w4/kredit-plus/pkg/kafka"
	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	loggerInterceptor "github.com/0x5w4/kredit-plus/pkg/logger-interceptor"
	"github.com/0x5w4/kredit-plus/writer-service/config"
	"github.com/0x5w4/kredit-plus/writer-service/internal/kredit/service"
	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	cfg               *config.Config
	v                 *validator.Validate
	logger            *loggerClient.AppLogger
	grpcServer        *grpcServer.GrpcServer
	pgxConn           *pgxpool.Pool
	loggerInterceptor loggerInterceptor.LoggerInterceptor
	service           *service.KreditService
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

	s.setupPgConn()
	defer s.pgxConn.Close()

	if err := s.setupKafka(ctx); err != nil {
		return fmt.Errorf("kafka setup failed: %w", err)
	}

	kafkaProducer := kafkaClient.NewProducer(s.cfg.Kafka.Brokers)
	defer kafkaProducer.Close()

	s.setupService(kafkaProducer)

	s.setupKafkaConsumers(ctx)

	s.setupGrpcServer()
	defer s.grpcServer.Stop(ctx)

	<-ctx.Done()

	return nil
}
