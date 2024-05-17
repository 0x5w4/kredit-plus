package queries

import (
	"context"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/0x5w4/kredit-plus/reader-service/config"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/repository"
	"github.com/0x5w4/kredit-plus/reader-service/internal/model"
)

type GetLimitHandler interface {
	Handle(ctx context.Context, query *GetLimitQuery) (*model.Limit, error)
}

type getLimitHandler struct {
	logger    *logger.AppLogger
	cfg       *config.Config
	mongoRepo repository.Repository
	redisRepo repository.CacheRepository
}

func NewGetLimitHandler(logger *logger.AppLogger, cfg *config.Config, mongoRepo repository.Repository, redisRepo repository.CacheRepository) *getLimitHandler {
	return &getLimitHandler{
		logger:    logger,
		cfg:       cfg,
		mongoRepo: mongoRepo,
		redisRepo: redisRepo,
	}
}

func (q *getLimitHandler) Handle(ctx context.Context, query *GetLimitQuery) (*model.Limit, error) {
	if limit, err := q.redisRepo.GetLimit(ctx, query.IdLimit.String()+query.IdKonsumen.String()); err == nil && limit != nil {
		return limit, nil
	}

	limit, err := q.mongoRepo.GetLimit(ctx, query.IdLimit, query.IdKonsumen)
	if err != nil {
		return nil, err
	}

	q.redisRepo.PutLimit(ctx, limit.IdLimit.String()+limit.IdKonsumen.String(), limit)
	return limit, nil
}
