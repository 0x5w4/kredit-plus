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

type CreateKonsumenCmdHandler interface {
	Handle(ctx context.Context, command *CreateKonsumenCommand) error
}

type createKonsumenHandler struct {
	logger        *logger.AppLogger
	cfg           *config.Config
	pgRepo        repository.Repository
	kafkaProducer kafkaClient.Producer
}

func NewCreateKonsumenHandler(
	logger *logger.AppLogger,
	cfg *config.Config,
	pgRepo repository.Repository,
	kafkaProducer kafkaClient.Producer,
) *createKonsumenHandler {
	return &createKonsumenHandler{
		logger:        logger,
		cfg:           cfg,
		pgRepo:        pgRepo,
		kafkaProducer: kafkaProducer,
	}
}

func (c *createKonsumenHandler) Handle(ctx context.Context, command *CreateKonsumenCommand) error {
	konsumenDto := &model.Konsumen{
		IdKonsumen:   command.IdKonsumen,
		Nik:          command.Nik,
		FullName:     command.FullName,
		LegalName:    command.LegalName,
		Gaji:         command.Gaji,
		TempatLahir:  command.TempatLahir,
		TanggalLahir: command.TanggalLahir,
		FotoKtp:      command.FotoKtp,
		FotoSelfie:   command.FotoSelfie,
		Email:        command.Email,
		Password:     command.Password,
	}

	konsumen, err := c.pgRepo.CreateKonsumen(ctx, konsumenDto)
	if err != nil {
		return err
	}

	msg := &kafkaMessages.KonsumenCreated{Konsumen: mapper.KonsumenToGrpcMessage(konsumen)}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Topic: c.cfg.KafkaTopics.KonsumenCreated.TopicName,
		Value: msgBytes,
		Time:  time.Now().UTC(),
	}

	return c.kafkaProducer.PublishMessage(ctx, message)
}
