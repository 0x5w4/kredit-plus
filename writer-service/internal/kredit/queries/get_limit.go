package queries

import (
	"context"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/0x5w4/kredit-plus/writer-service/config"
	"github.com/0x5w4/kredit-plus/writer-service/internal/kredit/repository"
	"github.com/0x5w4/kredit-plus/writer-service/internal/model"
)

type GetLimitHandler interface {
	Handle(ctx context.Context, query *GetLimitQuery) (*model.Limit, error)
}

type getLimitHandler struct {
	logger *logger.AppLogger
	cfg    *config.Config
	pgRepo repository.Repository
}

func NewGetLimitHandler(logger *logger.AppLogger, cfg *config.Config, pgRepo repository.Repository) *getLimitHandler {
	return &getLimitHandler{
		logger: logger,
		cfg:    cfg,
		pgRepo: pgRepo,
	}
}

func (q *getLimitHandler) Handle(ctx context.Context, query *GetLimitQuery) (*model.Limit, error) {
	return q.pgRepo.GetLimit(ctx, query.IdLimit, query.IdKonsumen)
}
