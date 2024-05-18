package commands

import (
	"context"
	"time"

	"github.com/0x5w4/kredit-plus/api-gateway-service/config"
	kafkaClient "github.com/0x5w4/kredit-plus/pkg/kafka"
	"github.com/0x5w4/kredit-plus/pkg/logger"
	kafkaMessages "github.com/0x5w4/kredit-plus/proto/kafka"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type CreateLimitCmdHandler interface {
	Handle(ctx context.Context, command *CreateLimitCommand) error
}

type createLimitHandler struct {
	logger        *logger.AppLogger
	cfg           *config.Config
	kafkaProducer *kafkaClient.Producer
}

func NewCreateLimitHandler(
	logger *logger.AppLogger,
	cfg *config.Config,
	kafkaProducer *kafkaClient.Producer,
) *createLimitHandler {
	return &createLimitHandler{
		logger:        logger,
		cfg:           cfg,
		kafkaProducer: kafkaProducer,
	}
}

func (c *createLimitHandler) Handle(ctx context.Context, command *CreateLimitCommand) error {
	limitDto := &kafkaMessages.LimitCreate{
		IdLimit:     command.CreateLimitDto.IdLimit.String(),
		IdKonsumen:  command.CreateLimitDto.IdKonsumen.String(),
		Tenor:       command.CreateLimitDto.Tenor,
		BatasKredit: command.CreateLimitDto.BatasKredit,
	}
	msgBytes, err := proto.Marshal(limitDto)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Topic: c.cfg.KafkaTopics.LimitCreate.TopicName,
		Value: msgBytes,
		Time:  time.Now().UTC(),
	}

	return c.kafkaProducer.PublishMessage(ctx, message)
}
