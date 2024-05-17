package commands

import (
	"context"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/0x5w4/kredit-plus/reader-service/config"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/repository"
	"github.com/0x5w4/kredit-plus/reader-service/internal/model"
)

type CreateTransaksiCmdHandler interface {
	Handle(ctx context.Context, command *CreateTransaksiCommand) error
}

type createTransaksiHandler struct {
	logger    *logger.AppLogger
	cfg       *config.Config
	mongoRepo repository.Repository
	redisRepo repository.CacheRepository
}

func NewCreateTransaksiHandler(
	logger *logger.AppLogger,
	cfg *config.Config,
	mongoRepo repository.Repository,
	redisRepo repository.CacheRepository,
) *createTransaksiHandler {
	return &createTransaksiHandler{
		logger:    logger,
		cfg:       cfg,
		mongoRepo: mongoRepo,
		redisRepo: redisRepo,
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

	created, err := c.mongoRepo.CreateTransaksi(ctx, transaksiDto)
	if err != nil {
		return err
	}

	c.redisRepo.PutTransaksi(ctx, created.IdTransaksi.String()+created.IdKonsumen.String(), created)
	return nil
}
