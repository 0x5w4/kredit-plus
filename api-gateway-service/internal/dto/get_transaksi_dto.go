package dto

import (
	"time"

	readerService "github.com/0x5w4/kredit-plus/reader-service/proto/reader"
	uuid "github.com/google/uuid"
)

type GetTransaksiRequestDto struct {
	IdTransaksi uuid.UUID `json:"id_transaksi" validate:"required"`
	IdKonsumen  uuid.UUID `json:"id_konsumen" validate:"required"`
}

type GetTransaksiResponseDto struct {
	IdTransaksi      uuid.UUID `json:"id_transaksi" validate:"required"`
	IdKonsumen       uuid.UUID `json:"id_konsumen" validate:"required"`
	NomorKontrak     string    `json:"nomor_kontrak" validate:"required,gte=0,lte=255"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi" validate:"required"`
	Otr              uint64    `json:"otr" validate:"required"`
	AdminFee         float64   `json:"admin_fee" validate:"required,gte=0"`
	JumlahCicilan    float64   `json:"jumlah_cicilan" validate:"required,gte=0"`
	JumlahBunga      float64   `json:"jumlah_bunga" validate:"required,gte=0"`
	NamaAsset        string    `json:"nama_asset" validate:"required,gte=0,lte=255"`
	JenisTransaksi   string    `json:"jenis_transaksi" validate:"required,gte=0,lte=255"`
}

func TransaksiResponseFromGrpc(transaksi *readerService.Transaksi) (*GetTransaksiResponseDto, error) {
	idTransaksi, err := uuid.Parse(transaksi.GetIdTransaksi())
	if err != nil {
		return nil, err
	}

	idKonsumen, err := uuid.Parse(transaksi.GetIdKonsumen())
	if err != nil {
		return nil, err
	}

	return &GetTransaksiResponseDto{
		IdTransaksi:      idTransaksi,
		IdKonsumen:       idKonsumen,
		NomorKontrak:     transaksi.GetNomorKontrak(),
		TanggalTransaksi: transaksi.GetTanggalTransaksi().AsTime(),
		Otr:              transaksi.GetOtr(),
		AdminFee:         transaksi.GetAdminFee(),
		JumlahCicilan:    transaksi.GetJumlahCicilan(),
		JumlahBunga:      transaksi.GetJumlahBunga(),
		NamaAsset:        transaksi.GetNamaAsset(),
		JenisTransaksi:   transaksi.GetJenisTransaksi(),
	}, nil
}
