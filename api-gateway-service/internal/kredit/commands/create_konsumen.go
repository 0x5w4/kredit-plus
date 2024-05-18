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
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreateKonsumenCmdHandler interface {
	Handle(ctx context.Context, command *CreateKonsumenCommand) error
}

type createKonsumenHandler struct {
	logger        *logger.AppLogger
	cfg           *config.Config
	kafkaProducer *kafkaClient.Producer
}

func NewCreateKonsumenHandler(
	logger *logger.AppLogger,
	cfg *config.Config,
	kafkaProducer *kafkaClient.Producer,
) *createKonsumenHandler {
	return &createKonsumenHandler{
		logger:        logger,
		cfg:           cfg,
		kafkaProducer: kafkaProducer,
	}
}

func (c *createKonsumenHandler) Handle(ctx context.Context, command *CreateKonsumenCommand) error {
	konsumenDto := &kafkaMessages.KonsumenCreate{
		IdKonsumen:   command.CreateKonsumenDto.IdKonsumen.String(),
		Nik:          command.CreateKonsumenDto.Nik,
		FullName:     command.CreateKonsumenDto.FullName,
		LegalName:    command.CreateKonsumenDto.LegalName,
		Gaji:         command.CreateKonsumenDto.Gaji,
		TempatLahir:  command.CreateKonsumenDto.TempatLahir,
		TanggalLahir: timestamppb.New(command.CreateKonsumenDto.TanggalLahir),
		FotoKtp:      command.CreateKonsumenDto.FotoKtp,
		FotoSelfie:   command.CreateKonsumenDto.FotoSelfie,
		Email:        command.CreateKonsumenDto.Email,
		Password:     command.CreateKonsumenDto.Password,
	}

	msgBytes, err := proto.Marshal(konsumenDto)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Topic: c.cfg.KafkaTopics.KonsumenCreate.TopicName,
		Value: msgBytes,
		Time:  time.Now().UTC(),
	}

	return c.kafkaProducer.PublishMessage(ctx, message)
}
