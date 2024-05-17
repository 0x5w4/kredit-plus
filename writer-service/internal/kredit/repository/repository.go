package repository

import (
	"context"

	"github.com/0x5w4/kredit-plus/writer-service/internal/model"
	uuid "github.com/google/uuid"
)

type Repository interface {
	CreateKonsumen(ctx context.Context, konsumen *model.Konsumen) (*model.Konsumen, error)
	CreateLimit(ctx context.Context, limit *model.Limit) (*model.Limit, error)
	CreateTransaksi(ctx context.Context, transaksi *model.Transaksi) (*model.Transaksi, error)

	GetLimit(ctx context.Context, idLimit uuid.UUID, idKonsumen uuid.UUID) (*model.Limit, error)
	GetTransaksi(ctx context.Context, idTransaksi uuid.UUID, idKonsumen uuid.UUID) (*model.Transaksi, error)
}
