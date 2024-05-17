package model

import (
	"time"

	uuid "github.com/google/uuid"
)

type Konsumen struct {
	IdKonsumen   uuid.UUID `json:"id_konsumen"`
	Nik          string    `json:"nik"`
	FullName     string    `json:"full_name"`
	LegalName    string    `json:"legal_name"`
	Gaji         float64   `json:"gaji"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	FotoKtp      string    `json:"foto_ktp"`
	FotoSelfie   string    `json:"foto_selfie"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Limit struct {
	IdLimit     uuid.UUID `json:"id_limit"`
	IdKonsumen  uuid.UUID `json:"id_konsumen"`
	Tenor       uint32    `json:"tenor"`
	BatasKredit float64   `json:"batas_kredit"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Transaksi struct {
	IdTransaksi      uuid.UUID `json:"id_transaksi"`
	IdKonsumen       uuid.UUID `json:"id_konsumen"`
	NomorKontrak     string    `json:"nomor_kontrak"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi"`
	Otr              uint64    `json:"otr"`
	AdminFee         float64   `json:"admin_fee"`
	JumlahCicilan    float64   `json:"jumlah_cicilan"`
	JumlahBunga      float64   `json:"jumlah_bunga"`
	NamaAsset        string    `json:"nama_asset"`
	JenisTransaksi   string    `json:"jenis_transaksi"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
