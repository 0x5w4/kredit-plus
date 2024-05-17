package server

import (
	"context"
	"fmt"

	grpcServer "github.com/0x5w4/kredit-plus/pkg/grpc-server"
	kafkaClient "github.com/0x5w4/kredit-plus/pkg/kafka"
	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	loggerInterceptor "github.com/0x5w4/kredit-plus/pkg/logger-interceptor"
	mongoClient "github.com/0x5w4/kredit-plus/pkg/mongo"
	redisClient "github.com/0x5w4/kredit-plus/pkg/redis"
	readerGrpc "github.com/0x5w4/kredit-plus/reader-service/internal/kredit/delivery/grpc"
	messageProcessor "github.com/0x5w4/kredit-plus/reader-service/internal/kredit/delivery/kafka"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/repository"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/service"
	readerPb "github.com/0x5w4/kredit-plus/reader-service/proto/reader"
	"github.com/segmentio/kafka-go"
)

func (s *Server) setupLogger() error {
	var err error
	s.logger, err = loggerClient.NewAppLogger(s.cfg.Logger)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) setupLoggerInterceptor() {
	s.loggerInterceptor = loggerInterceptor.NewLoggerInterceptor(s.logger)
}

func (s *Server) setupMongo() error {
	var err error
	s.mongoClient, err = mongoClient.NewMongoDBConn(s.cfg.Mongo, s.logger)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) setupRedis() {
	s.redisClient = redisClient.NewUniversalRedisClient(s.cfg.Redis)
}

func (s *Server) setupService() {
	mongoRepo := repository.NewMongoRepository(s.logger, s.cfg, s.mongoClient)
	redisRepo := repository.NewRedisRepository(s.logger, s.cfg, s.redisClient)

	s.service = service.NewKreditService(s.logger, s.cfg, mongoRepo, redisRepo)
}

func (s *Server) setupKafka(ctx context.Context) error {
	if err := s.initKafkaTopics(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Server) initKafkaTopics(ctx context.Context) error {
	conn, err := kafkaClient.NewKafkaConn(ctx, s.cfg.Kafka)
	if err != nil {
		return err
	}

	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			fmt.Println("Failed to close Kafka connection:", closeErr)
		}
	}()

	var topicConfigs []kafka.TopicConfig
	topicConfigs = append(topicConfigs, kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.KonsumenCreated.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.KonsumenCreated.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.KonsumenCreated.ReplicationFactor,
	})
	topicConfigs = append(topicConfigs, kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.LimitCreated.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.LimitCreated.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.LimitCreated.ReplicationFactor,
	})
	topicConfigs = append(topicConfigs, kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.TransaksiCreated.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.TransaksiCreated.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.TransaksiCreated.ReplicationFactor,
	})

	if err := conn.CreateTopics(topicConfigs...); err != nil {
		return err
	}

	return nil
}

func (s *Server) setupKafkaConsumers(ctx context.Context) {
	messageProcessor := messageProcessor.NewMessageProcessor(s.logger, s.cfg, s.v, s.service)

	s.logger.Logger.Info("Starting Writer Kafka consumers")

	consumerTopics := []string{
		s.cfg.KafkaTopics.KonsumenCreated.TopicName,
		s.cfg.KafkaTopics.LimitCreated.TopicName,
		s.cfg.KafkaTopics.TransaksiCreated.TopicName,
	}
	s.startKafkaConsumers(ctx, consumerTopics, messageProcessor)
}

func (s *Server) startKafkaConsumers(ctx context.Context, consumerTopics []string, messageProcessor *messageProcessor.MessageProcessor) error {
	consumer := kafkaClient.NewConsumer(s.cfg.Kafka.Brokers)
	const poolSize = 10

	for _, topic := range consumerTopics {
		go consumer.StartWorkers(ctx, s.cfg.Kafka.GroupID, topic, poolSize, messageProcessor.ProcessMessages)
	}

	return nil
}

func (s *Server) setupGrpcServer() error {
	var err error
	s.grpcServer, err = grpcServer.NewGrpcServer(s.cfg.GrpcServer, s.loggerInterceptor, s.logger)
	if err != nil {
		return fmt.Errorf("new grpc server failed: %w", err)
	}

	readerGrpc := readerGrpc.NewReaderGrpcService(s.logger, s.cfg, s.v, s.service)
	readerPb.RegisterReaderServiceServer(s.grpcServer.Server, readerGrpc)

	return nil
}
