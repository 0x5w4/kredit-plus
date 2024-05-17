package queries

import (
	"context"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/0x5w4/kredit-plus/reader-service/config"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/repository"
	"github.com/0x5w4/kredit-plus/reader-service/internal/model"
)

type GetTransaksiHandler interface {
	Handle(ctx context.Context, query *GetTransaksiQuery) (*model.Transaksi, error)
}

type getTransaksiHandler struct {
	logger    *logger.AppLogger
	cfg       *config.Config
	mongoRepo repository.Repository
	redisRepo repository.CacheRepository
}

func NewGetTransaksiHandler(logger *logger.AppLogger, cfg *config.Config, mongoRepo repository.Repository, redisRepo repository.CacheRepository) *getTransaksiHandler {
	return &getTransaksiHandler{
		logger:    logger,
		cfg:       cfg,
		mongoRepo: mongoRepo,
		redisRepo: redisRepo,
	}
}

func (q *getTransaksiHandler) Handle(ctx context.Context, query *GetTransaksiQuery) (*model.Transaksi, error) {
	if transaksi, err := q.redisRepo.GetTransaksi(ctx, query.IdTransaksi.String()+query.IdKonsumen.String()); err == nil && transaksi != nil {
		return transaksi, nil
	}

	transaksi, err := q.mongoRepo.GetTransaksi(ctx, query.IdTransaksi, query.IdKonsumen)
	if err != nil {
		return nil, err
	}

	q.redisRepo.PutTransaksi(ctx, transaksi.IdTransaksi.String()+transaksi.IdKonsumen.String(), transaksi)
	return transaksi, nil
}
