package queries

import (
	"context"

	"github.com/0x5w4/kredit-plus/api-gateway-service/config"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/dto"
	"github.com/0x5w4/kredit-plus/pkg/logger"
	readerService "github.com/0x5w4/kredit-plus/reader-service/proto/reader"
)

type GetLimitHandler interface {
	Handle(ctx context.Context, query *GetLimitQuery) (*dto.GetLimitResponseDto, error)
}

type getLimitHandler struct {
	logger   *logger.AppLogger
	cfg      *config.Config
	rsClient readerService.ReaderServiceClient
}

func NewGetLimitHandler(logger *logger.AppLogger, cfg *config.Config, rsClient readerService.ReaderServiceClient) *getLimitHandler {
	return &getLimitHandler{
		logger:   logger,
		cfg:      cfg,
		rsClient: rsClient,
	}
}

func (q *getLimitHandler) Handle(ctx context.Context, query *GetLimitQuery) (*dto.GetLimitResponseDto, error) {
	res, err := q.rsClient.GetLimit(
		ctx,
		&readerService.GetLimitRequest{
			IdLimit:    query.IdLimit.String(),
			IdKonsumen: query.IdKonsumen.String(),
		})
	if err != nil {
		return nil, err
	}

	limit, err := dto.LimitResponseFromGrpc(res.GetLimit())
	if err != nil {
		return nil, err
	}

	return limit, nil
}
