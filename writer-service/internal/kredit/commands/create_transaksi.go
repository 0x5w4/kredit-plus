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

type CreateTransaksiCmdHandler interface {
	Handle(ctx context.Context, command *CreateTransaksiCommand) error
}

type createTransaksiHandler struct {
	logger        *logger.AppLogger
	cfg           *config.Config
	pgRepo        repository.Repository
	kafkaProducer *kafkaClient.Producer
}

func NewCreateTransaksiHandler(
	logger *logger.AppLogger,
	cfg *config.Config,
	pgRepo repository.Repository,
	kafkaProducer *kafkaClient.Producer,
) *createTransaksiHandler {
	return &createTransaksiHandler{
		logger:        logger,
		cfg:           cfg,
		pgRepo:        pgRepo,
		kafkaProducer: kafkaProducer,
	}
}

func (c *createTransaksiHandler) Handle(ctx context.Context, command *CreateTransaksiCommand) error {
	transaksiDto := &model.Transaksi{
		IdTransaksi:      command.IdTransaksi,
		IdKonsumen:       command.IdKonsumen,
		NomorKontrak:     command.NomorKontrak,
		TanggalTransaksi: command.TanggalTransaksi,
		Otr:              command.Otr,
		AdminFee:         command.AdminFee,
		JumlahCicilan:    command.JumlahCicilan,
		JumlahBunga:      command.JumlahBunga,
		NamaAsset:        command.NamaAsset,
		JenisTransaksi:   command.JenisTransaksi,
	}

	transaksi, err := c.pgRepo.CreateTransaksi(ctx, transaksiDto)
	if err != nil {
		return err
	}

	msg := &kafkaMessages.TransaksiCreated{Transaksi: mapper.TransaksiToGrpcMessage(transaksi)}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Topic: c.cfg.KafkaTopics.TransaksiCreated.TopicName,
		Value: msgBytes,
		Time:  time.Now().UTC(),
	}

	return c.kafkaProducer.PublishMessage(ctx, message)
}
