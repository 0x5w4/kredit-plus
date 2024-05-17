package repository

import (
	"context"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/0x5w4/kredit-plus/writer-service/config"
	"github.com/0x5w4/kredit-plus/writer-service/internal/model"
	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type kreditRepository struct {
	logger *logger.AppLogger
	cfg    *config.Config
	db     *pgxpool.Pool
}

func NewKreditRepository(logger *logger.AppLogger, cfg *config.Config, db *pgxpool.Pool) *kreditRepository {
	return &kreditRepository{logger: logger, cfg: cfg, db: db}
}

func (r *kreditRepository) CreateKonsumen(ctx context.Context, konsumen *model.Konsumen) (*model.Konsumen, error) {

	var created model.Konsumen
	if err := r.db.QueryRow(
		ctx,
		createKonsumenQuery,
		&konsumen.IdKonsumen,
		&konsumen.Nik,
		&konsumen.FullName,
		&konsumen.LegalName,
		&konsumen.Gaji,
		&konsumen.TempatLahir,
		&konsumen.TanggalLahir,
		&konsumen.FotoKtp,
		&konsumen.FotoSelfie,
		&konsumen.Email,
		&konsumen.Password,
	).Scan(
		&created.IdKonsumen,
		&created.Nik,
		&created.FullName,
		&created.LegalName,
		&created.Gaji,
		&created.TempatLahir,
		&created.TanggalLahir,
		&created.FotoKtp,
		&created.FotoSelfie,
		&created.Email,
		&created.Password,
		&created.CreatedAt,
		&created.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "db.QueryRow")
	}

	return &created, nil
}

func (r *kreditRepository) CreateLimit(ctx context.Context, limit *model.Limit) (*model.Limit, error) {

	var created model.Limit
	if err := r.db.QueryRow(
		ctx,
		createLimitQuery,
		&limit.IdLimit,
		&limit.IdKonsumen,
		&limit.Tenor,
		&limit.BatasKredit,
	).Scan(
		&created.IdLimit,
		&created.IdKonsumen,
		&created.Tenor,
		&created.BatasKredit,
		&created.CreatedAt,
		&created.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "db.QueryRow")
	}

	return &created, nil
}

func (r *kreditRepository) CreateTransaksi(ctx context.Context, transaksi *model.Transaksi) (*model.Transaksi, error) {

	var created model.Transaksi
	if err := r.db.QueryRow(
		ctx,
		createTransaksiQuery,
		&transaksi.IdTransaksi,
		&transaksi.IdKonsumen,
		&transaksi.NomorKontrak,
		&transaksi.TanggalTransaksi,
		&transaksi.Otr,
		&transaksi.AdminFee,
		&transaksi.JumlahCicilan,
		&transaksi.JumlahBunga,
		&transaksi.NamaAsset,
		&transaksi.JenisTransaksi,
	).Scan(
		&created.IdTransaksi,
		&created.IdKonsumen,
		&created.NomorKontrak,
		&created.TanggalTransaksi,
		&created.Otr,
		&created.AdminFee,
		&created.JumlahCicilan,
		&created.JumlahBunga,
		&created.NamaAsset,
		&created.JenisTransaksi,
		&created.CreatedAt,
		&created.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "db.QueryRow")
	}

	return &created, nil
}

func (r *kreditRepository) GetLimit(ctx context.Context, idLimit uuid.UUID, idKonsumen uuid.UUID) (*model.Limit, error) {
	var limit model.Limit
	if err := r.db.QueryRow(ctx, getLimitQuery, idLimit, idKonsumen).Scan(
		&limit.IdLimit,
		&limit.IdKonsumen,
		&limit.Tenor,
		&limit.BatasKredit,
		&limit.CreatedAt,
		&limit.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "Scan")
	}

	return &limit, nil
}

func (r *kreditRepository) GetTransaksi(ctx context.Context, idTransaksi uuid.UUID, idKonsumen uuid.UUID) (*model.Transaksi, error) {
	var transaksi model.Transaksi
	if err := r.db.QueryRow(ctx, getTransaksiQuery, idTransaksi, idKonsumen).Scan(
		&transaksi.IdTransaksi,
		&transaksi.IdKonsumen,
		&transaksi.NomorKontrak,
		&transaksi.TanggalTransaksi,
		&transaksi.Otr,
		&transaksi.AdminFee,
		&transaksi.JumlahCicilan,
		&transaksi.JumlahBunga,
		&transaksi.NamaAsset,
		&transaksi.JenisTransaksi,
		&transaksi.CreatedAt,
		&transaksi.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "Scan")
	}

	return &transaksi, nil
}
