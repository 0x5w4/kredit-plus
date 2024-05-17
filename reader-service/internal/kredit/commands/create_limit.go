package commands

import (
	"context"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/0x5w4/kredit-plus/reader-service/config"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/repository"
	"github.com/0x5w4/kredit-plus/reader-service/internal/model"
)

type CreateLimitCmdHandler interface {
	Handle(ctx context.Context, command *CreateLimitCommand) error
}

type createLimitHandler struct {
	logger    *logger.AppLogger
	cfg       *config.Config
	mongoRepo repository.Repository
	redisRepo repository.CacheRepository
}

func NewCreateLimitHandler(
	logger *logger.AppLogger,
	cfg *config.Config,
	mongoRepo repository.Repository,
	redisRepo repository.CacheRepository,
) *createLimitHandler {
	return &createLimitHandler{
		logger:    logger,
		cfg:       cfg,
		mongoRepo: mongoRepo,
		redisRepo: redisRepo,
	}
}

func (c *createLimitHandler) Handle(ctx context.Context, command *CreateLimitCommand) error {
	limitDto := &model.Limit{
		IdLimit:     command.IdLimit,
		IdKonsumen:  command.IdKonsumen,
		Tenor:       command.Tenor,
		BatasKredit: command.BatasKredit,
	}

	created, err := c.mongoRepo.CreateLimit(ctx, limitDto)
	if err != nil {
		return err
	}

	c.redisRepo.PutLimit(ctx, created.IdLimit.String()+created.IdKonsumen.String(), created)
	return nil
}
