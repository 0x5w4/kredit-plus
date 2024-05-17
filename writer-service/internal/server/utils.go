package server

import (
	"context"
	"net"
	"strconv"

	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	loggerInterceptor "github.com/0x5w4/kredit-plus/pkg/logger-interceptor"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
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

const (
	stackSize = 1 << 10 // 1 KB
)

func (s *server) connectKafkaBrokers(ctx context.Context) error {
	kafkaConn, err := kafkaClient.NewKafkaConn(ctx, s.cfg.Kafka)
	if err != nil {
		return errors.Wrap(err, "kafka.NewKafkaCon")
	}

	s.kafkaConn = kafkaConn

	brokers, err := kafkaConn.Brokers()
	if err != nil {
		return errors.Wrap(err, "kafkaConn.Brokers")
	}

	s.log.Infof("kafka connected to brokers: %+v", brokers)

	return nil
}

func (s *server) initKafkaTopics(ctx context.Context) {
	controller, err := s.kafkaConn.Controller()
	if err != nil {
		s.log.WarnMsg("kafkaConn.Controller", err)
		return
	}

	controllerURI := net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port))
	s.log.Infof("kafka controller uri: %s", controllerURI)

	conn, err := kafka.DialContext(ctx, "tcp", controllerURI)
	if err != nil {
		s.log.WarnMsg("initKafkaTopics.DialContext", err)
		return
	}
	defer conn.Close() // nolint: errcheck

	s.log.Infof("established new kafka controller connection: %s", controllerURI)

	productCreateTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.ProductCreate.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.ProductCreate.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.ProductCreate.ReplicationFactor,
	}

	productCreatedTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.ProductCreated.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.ProductCreated.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.ProductCreated.ReplicationFactor,
	}

	productUpdateTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.ProductUpdate.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.ProductUpdate.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.ProductUpdate.ReplicationFactor,
	}

	productUpdatedTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.ProductUpdated.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.ProductUpdated.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.ProductUpdated.ReplicationFactor,
	}

	productDeleteTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.ProductDelete.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.ProductDelete.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.ProductDelete.ReplicationFactor,
	}

	productDeletedTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.ProductDeleted.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.ProductDeleted.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.ProductDeleted.ReplicationFactor,
	}

	if err := conn.CreateTopics(
		productCreateTopic,
		productUpdateTopic,
		productCreatedTopic,
		productUpdatedTopic,
		productDeleteTopic,
		productDeletedTopic,
	); err != nil {
		s.log.WarnMsg("kafkaConn.CreateTopics", err)
		return
	}

	s.log.Infof("kafka topics created or already exists: %+v", []kafka.TopicConfig{productCreateTopic, productUpdateTopic, productCreatedTopic, productUpdatedTopic, productDeleteTopic, productDeletedTopic})
}

func (s *server) getConsumerGroupTopics() []string {
	return []string{
		s.cfg.KafkaTopics.ProductCreate.TopicName,
		s.cfg.KafkaTopics.ProductUpdate.TopicName,
		s.cfg.KafkaTopics.ProductDelete.TopicName,
	}
}
