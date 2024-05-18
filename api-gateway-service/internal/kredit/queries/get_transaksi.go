package queries

import (
	"context"

	"github.com/0x5w4/kredit-plus/api-gateway-service/config"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/dto"
	"github.com/0x5w4/kredit-plus/pkg/logger"
	readerService "github.com/0x5w4/kredit-plus/reader-service/proto/reader"
)

type GetTransaksiHandler interface {
	Handle(ctx context.Context, query *GetTransaksiQuery) (*dto.GetTransaksiResponseDto, error)
}

type getTransaksiHandler struct {
	logger   *logger.AppLogger
	cfg      *config.Config
	rsClient readerService.ReaderServiceClient
}

func NewGetTransaksiHandler(logger *logger.AppLogger, cfg *config.Config, rsClient readerService.ReaderServiceClient) *getTransaksiHandler {
	return &getTransaksiHandler{
		logger:   logger,
		cfg:      cfg,
		rsClient: rsClient,
	}
}

func (q *getTransaksiHandler) Handle(ctx context.Context, query *GetTransaksiQuery) (*dto.GetTransaksiResponseDto, error) {
	res, err := q.rsClient.GetTransaksi(
		ctx,
		&readerService.GetTransaksiRequest{
			IdTransaksi: query.IdTransaksi.String(),
			IdKonsumen:  query.IdKonsumen.String(),
		})
	if err != nil {
		return nil, err
	}

	transaksi, err := dto.TransaksiResponseFromGrpc(res.GetTransaksi())
	if err != nil {
		return nil, err
	}

	return transaksi, nil
}
