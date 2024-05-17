package queries

import (
	"context"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/0x5w4/kredit-plus/writer-service/config"
	"github.com/0x5w4/kredit-plus/writer-service/internal/kredit/repository"
	"github.com/0x5w4/kredit-plus/writer-service/internal/model"
)

type GetTransaksiHandler interface {
	Handle(ctx context.Context, query *GetTransaksiQuery) (*model.Transaksi, error)
}

type getTransaksiHandler struct {
	logger *logger.AppLogger
	cfg    *config.Config
	pgRepo repository.Repository
}

func NewGetTransaksiHandler(logger *logger.AppLogger, cfg *config.Config, pgRepo repository.Repository) *getTransaksiHandler {
	return &getTransaksiHandler{
		logger: logger,
		cfg:    cfg,
		pgRepo: pgRepo,
	}
}

func (q *getTransaksiHandler) Handle(ctx context.Context, query *GetTransaksiQuery) (*model.Transaksi, error) {
	return q.pgRepo.GetTransaksi(ctx, query.IdTransaksi, query.IdKonsumen)
}
