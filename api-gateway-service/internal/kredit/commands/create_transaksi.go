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

type CreateTransaksiCmdHandler interface {
	Handle(ctx context.Context, command *CreateTransaksiCommand) error
}

type createTransaksiHandler struct {
	logger        *logger.AppLogger
	cfg           *config.Config
	kafkaProducer *kafkaClient.Producer
}

func NewCreateTransaksiHandler(
	logger *logger.AppLogger,
	cfg *config.Config,
	kafkaProducer *kafkaClient.Producer,
) *createTransaksiHandler {
	return &createTransaksiHandler{
		logger:        logger,
		cfg:           cfg,
		kafkaProducer: kafkaProducer,
	}
}

func (c *createTransaksiHandler) Handle(ctx context.Context, command *CreateTransaksiCommand) error {
	transaksiDto := &kafkaMessages.TransaksiCreate{
		IdTransaksi:      command.CreateTransaksiDto.IdTransaksi.String(),
		IdKonsumen:       command.CreateTransaksiDto.IdKonsumen.String(),
		NomorKontrak:     command.CreateTransaksiDto.NomorKontrak,
		TanggalTransaksi: timestamppb.New(command.CreateTransaksiDto.TanggalTransaksi),
		Otr:              command.CreateTransaksiDto.Otr,
		AdminFee:         command.CreateTransaksiDto.AdminFee,
		JumlahCicilan:    command.CreateTransaksiDto.JumlahCicilan,
		JumlahBunga:      command.CreateTransaksiDto.JumlahBunga,
		NamaAsset:        command.CreateTransaksiDto.NamaAsset,
		JenisTransaksi:   command.CreateTransaksiDto.JenisTransaksi,
	}

	msgBytes, err := proto.Marshal(transaksiDto)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Topic: c.cfg.KafkaTopics.TransaksiCreate.TopicName,
		Value: msgBytes,
		Time:  time.Now().UTC(),
	}

	return c.kafkaProducer.PublishMessage(ctx, message)
}
