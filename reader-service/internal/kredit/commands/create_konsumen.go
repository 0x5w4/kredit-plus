package commands

import (
	"context"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/0x5w4/kredit-plus/reader-service/config"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/repository"
	"github.com/0x5w4/kredit-plus/reader-service/internal/model"
)

type CreateKonsumenCmdHandler interface {
	Handle(ctx context.Context, command *CreateKonsumenCommand) error
}

type createKonsumenHandler struct {
	logger    *logger.AppLogger
	cfg       *config.Config
	mongoRepo repository.Repository
	redisRepo repository.CacheRepository
}

func NewCreateKonsumenHandler(
	logger *logger.AppLogger,
	cfg *config.Config,
	mongoRepo repository.Repository,
	redisRepo repository.CacheRepository,
) *createKonsumenHandler {
	return &createKonsumenHandler{
		logger:    logger,
		cfg:       cfg,
		mongoRepo: mongoRepo,
		redisRepo: redisRepo,
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

	created, err := c.mongoRepo.CreateKonsumen(ctx, konsumenDto)
	if err != nil {
		return err
	}

	c.redisRepo.PutKonsumen(ctx, created.IdKonsumen.String(), created)
	return nil
}
