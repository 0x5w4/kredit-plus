package repository

import (
	"context"

	"github.com/0x5w4/kredit-plus/reader-service/internal/model"
	uuid "github.com/google/uuid"
)

type CacheRepository interface {
	PutKonsumen(ctx context.Context, key string, konsumen *model.Konsumen)
	PutLimit(ctx context.Context, key string, limit *model.Limit)
	PutTransaksi(ctx context.Context, key string, transaksi *model.Transaksi)
	GetLimit(ctx context.Context, key string) (*model.Limit, error)
	GetTransaksi(ctx context.Context, key string) (*model.Transaksi, error)
}

type Repository interface {
	CreateKonsumen(ctx context.Context, konsumen *model.Konsumen) (*model.Konsumen, error)
	CreateLimit(ctx context.Context, limit *model.Limit) (*model.Limit, error)
	CreateTransaksi(ctx context.Context, transaksi *model.Transaksi) (*model.Transaksi, error)

	GetLimit(ctx context.Context, idLimit uuid.UUID, idKonsumen uuid.UUID) (*model.Limit, error)
	GetTransaksi(ctx context.Context, idTransaksi uuid.UUID, idKonsumen uuid.UUID) (*model.Transaksi, error)
}
