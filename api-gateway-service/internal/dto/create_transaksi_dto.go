package dto

import (
	"time"

	uuid "github.com/google/uuid"
)

type CreateTransaksiHttpRequest struct {
	IdKonsumen       string `json:"id_konsumen" validate:"required"`
	NomorKontrak     string `json:"nomor_kontrak" validate:"required,gte=0,lte=255"`
	TanggalTransaksi string `json:"tanggal_transaksi" validate:"required"`
	Otr              uint64 `json:"otr" validate:"required"`
	AdminFee         uint64 `json:"admin_fee" validate:"required"`
	JumlahCicilan    uint64 `json:"jumlah_cicilan" validate:"required"`
	JumlahBunga      uint64 `json:"jumlah_bunga" validate:"required"`
	NamaAsset        string `json:"nama_asset" validate:"required,gte=0,lte=255"`
	JenisTransaksi   string `json:"jenis_transaksi" validate:"required,gte=0,lte=255"`
}

type CreateTransaksiRequestDto struct {
	IdTransaksi      uuid.UUID `json:"id_transaksi" validate:"required"`
	IdKonsumen       uuid.UUID `json:"id_konsumen" validate:"required"`
	NomorKontrak     string    `json:"nomor_kontrak" validate:"required,gte=0,lte=255"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi" validate:"required"`
	Otr              uint64    `json:"otr" validate:"required"`
	AdminFee         float64   `json:"admin_fee" validate:"required"`
	JumlahCicilan    float64   `json:"jumlah_cicilan" validate:"required"`
	JumlahBunga      float64   `json:"jumlah_bunga" validate:"required"`
	NamaAsset        string    `json:"nama_asset" validate:"required,gte=0,lte=255"`
	JenisTransaksi   string    `json:"jenis_transaksi" validate:"required,gte=0,lte=255"`
}

type CreateTransaksiResponseDto struct {
	IdTransaksi uuid.UUID `json:"id_transaksi" validate:"required"`
}
