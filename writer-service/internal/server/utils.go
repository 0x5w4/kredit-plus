package server

import (
	"context"
	"fmt"

	grpcServer "github.com/0x5w4/kredit-plus/pkg/grpc-server"
	kafkaClient "github.com/0x5w4/kredit-plus/pkg/kafka"
	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	loggerInterceptor "github.com/0x5w4/kredit-plus/pkg/logger-interceptor"
	postgresClient "github.com/0x5w4/kredit-plus/pkg/postgres"
	writerGrpc "github.com/0x5w4/kredit-plus/writer-service/internal/kredit/delivery/grpc"
	messageProcessor "github.com/0x5w4/kredit-plus/writer-service/internal/kredit/delivery/kafka"
	"github.com/0x5w4/kredit-plus/writer-service/internal/kredit/repository"
	"github.com/0x5w4/kredit-plus/writer-service/internal/kredit/service"
	writerPb "github.com/0x5w4/kredit-plus/writer-service/proto/writer"
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

func (s *Server) setupPgConn() error {
	var err error
	s.pgxConn, err = postgresClient.NewPgxConn(s.cfg.Postgresql, s.logger)
	if err != nil {
		return err
	}

	return nil
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
		Topic:             s.cfg.KafkaTopics.KonsumenCreate.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.KonsumenCreate.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.KonsumenCreate.ReplicationFactor,
	})
	topicConfigs = append(topicConfigs, kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.KonsumenCreated.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.KonsumenCreated.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.KonsumenCreated.ReplicationFactor,
	})
	topicConfigs = append(topicConfigs, kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.LimitCreate.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.LimitCreate.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.LimitCreate.ReplicationFactor,
	})
	topicConfigs = append(topicConfigs, kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.LimitCreated.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.LimitCreated.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.LimitCreated.ReplicationFactor,
	})
	topicConfigs = append(topicConfigs, kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.TransaksiCreate.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.TransaksiCreate.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.TransaksiCreate.ReplicationFactor,
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

func (s *Server) setupService(kafkaProducer *kafkaClient.Producer) {
	pgRepo := repository.NewKreditRepository(s.logger, s.cfg, s.pgxConn)

	s.service = service.NewKreditService(s.logger, s.cfg, pgRepo, kafkaProducer)
}

func (s *Server) setupKafkaConsumers(ctx context.Context) {
	messageProcessor := messageProcessor.NewMessageProcessor(s.logger, s.cfg, s.v, s.service)

	s.logger.Logger.Info("Starting Writer Kafka consumers")

	consumerTopics := []string{
		s.cfg.KafkaTopics.KonsumenCreate.TopicName,
		s.cfg.KafkaTopics.LimitCreate.TopicName,
		s.cfg.KafkaTopics.TransaksiCreate.TopicName,
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

	writerGrpc := writerGrpc.NewWriterGrpcService(s.logger, s.cfg, s.v, s.service)
	writerPb.RegisterWriterServiceServer(s.grpcServer.Server, writerGrpc)

	return nil
}
