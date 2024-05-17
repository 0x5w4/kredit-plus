package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	grpcServer "github.com/0x5w4/kredit-plus/pkg/grpc-server"
	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	loggerInterceptor "github.com/0x5w4/kredit-plus/pkg/logger-interceptor"
	"github.com/0x5w4/kredit-plus/pkg/postgres"
	"github.com/0x5w4/kredit-plus/reader-service/config"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
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

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.im = interceptors.NewInterceptorManager(s.log)
	s.metrics = metrics.NewWriterServiceMetrics(s.cfg)

	pgxConn, err := postgres.NewPgxConn(s.cfg.Postgresql)
	if err != nil {
		return errors.Wrap(err, "postgresql.NewPgxConn")
	}
	s.pgConn = pgxConn
	s.log.Infof("postgres connected: %v", pgxConn.Stat().TotalConns())
	defer pgxConn.Close()

	kafkaProducer := kafkaClient.NewProducer(s.log, s.cfg.Kafka.Brokers)
	defer kafkaProducer.Close() // nolint: errcheck

	productRepo := repository.NewProductRepository(s.log, s.cfg, pgxConn)
	s.ps = service.NewProductService(s.log, s.cfg, productRepo, kafkaProducer)
	productMessageProcessor := kafkaConsumer.NewProductMessageProcessor(s.log, s.cfg, s.v, s.ps, s.metrics)

	s.log.Info("Starting Writer Kafka consumers")
	cg := kafkaClient.NewConsumerGroup(s.cfg.Kafka.Brokers, s.cfg.Kafka.GroupID, s.log)
	go cg.ConsumeTopic(ctx, s.getConsumerGroupTopics(), kafkaConsumer.PoolSize, productMessageProcessor.ProcessMessages)

	closeGrpcServer, grpcServer, err := s.newWriterGrpcServer()
	if err != nil {
		return errors.Wrap(err, "NewScmGrpcServer")
	}
	defer closeGrpcServer() // nolint: errcheck

	if err := s.connectKafkaBrokers(ctx); err != nil {
		return errors.Wrap(err, "s.connectKafkaBrokers")
	}
	defer s.kafkaConn.Close() // nolint: errcheck

	if s.cfg.Kafka.InitTopics {
		s.initKafkaTopics(ctx)
	}

	s.runHealthCheck(ctx)
	s.runMetrics(cancel)

	if s.cfg.Jaeger.Enable {
		tracer, closer, err := tracing.NewJaegerTracer(s.cfg.Jaeger)
		if err != nil {
			return err
		}
		defer closer.Close() // nolint: errcheck
		opentracing.SetGlobalTracer(tracer)
	}

	<-ctx.Done()
	grpcServer.GracefulStop()

	return nil

	return nil
}
