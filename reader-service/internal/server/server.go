package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	grpcServer "github.com/0x5w4/kredit-plus/pkg/grpc-server"
	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	loggerInterceptor "github.com/0x5w4/kredit-plus/pkg/logger-interceptor"
	"github.com/0x5w4/kredit-plus/reader-service/config"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/service"
	"github.com/go-playground/validator"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	cfg               *config.Config
	v                 *validator.Validate
	logger            *loggerClient.AppLogger
	grpcServer        *grpcServer.GrpcServer
	mongoClient       *mongo.Client
	redisClient       redis.UniversalClient
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

	s.setupMongo()
	defer s.mongoClient.Disconnect(ctx)

	s.setupRedis()
	defer s.redisClient.Close()

	s.setupService()

	if err := s.setupKafka(ctx); err != nil {
		return fmt.Errorf("kafka setup failed: %w", err)
	}

	s.setupKafkaConsumers(ctx)

	s.setupGrpcServer()
	defer s.grpcServer.Stop(ctx)

	<-ctx.Done()

	return nil
}
