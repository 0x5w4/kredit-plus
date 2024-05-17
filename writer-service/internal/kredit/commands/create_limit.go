package commands

import (
	"context"
	"time"

	kafkaClient "github.com/0x5w4/kredit-plus/pkg/kafka"
	"github.com/0x5w4/kredit-plus/pkg/logger"
	kafkaMessages "github.com/0x5w4/kredit-plus/proto/kafka"
	"github.com/0x5w4/kredit-plus/writer-service/config"
	"github.com/0x5w4/kredit-plus/writer-service/internal/kredit/repository"
	"github.com/0x5w4/kredit-plus/writer-service/internal/model"
	"github.com/0x5w4/kredit-plus/writer-service/mapper"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type CreateLimitCmdHandler interface {
	Handle(ctx context.Context, command *CreateLimitCommand) error
}

type createLimitHandler struct {
	logger        *logger.AppLogger
	cfg           *config.Config
	pgRepo        repository.Repository
	kafkaProducer kafkaClient.Producer
}

func NewCreateLimitHandler(
	logger *logger.AppLogger,
	cfg *config.Config,
	pgRepo repository.Repository,
	kafkaProducer kafkaClient.Producer,
) *createLimitHandler {
	return &createLimitHandler{
		logger:        logger,
		cfg:           cfg,
		pgRepo:        pgRepo,
		kafkaProducer: kafkaProducer,
	}
}

func (c *createLimitHandler) Handle(ctx context.Context, command *CreateLimitCommand) error {
	limitDto := &model.Limit{
		IdLimit:     command.IdLimit,
		IdKonsumen:  command.IdKonsumen,
		Tenor:       command.Tenor,
		BatasKredit: command.BatasKredit,
	}

	limit, err := c.pgRepo.CreateLimit(ctx, limitDto)
	if err != nil {
		return err
	}

	msg := &kafkaMessages.LimitCreated{Limit: mapper.LimitToGrpcMessage(limit)}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Topic: c.cfg.KafkaTopics.LimitCreated.TopicName,
		Value: msgBytes,
		Time:  time.Now().UTC(),
	}

	return c.kafkaProducer.PublishMessage(ctx, message)
}
