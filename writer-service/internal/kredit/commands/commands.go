package commands

import (
	"time"

	uuid "github.com/google/uuid"
)

type KreditCommands struct {
	CreateKonsumen  CreateKonsumenCmdHandler
	CreateLimit     CreateLimitCmdHandler
	CreateTransaksi CreateTransaksiCmdHandler
}

func NewKreditCommands(createKonsumen CreateKonsumenCmdHandler, createLimit CreateLimitCmdHandler, createTransaksi CreateTransaksiCmdHandler) *KreditCommands {
	return &KreditCommands{CreateKonsumen: createKonsumen, CreateLimit: createLimit, CreateTransaksi: createTransaksi}
}

type CreateKonsumenCommand struct {
	IdKonsumen   uuid.UUID `json:"id_konsumen" validate:"required"`
	Nik          string    `json:"nik" validate:"required,gte=0,lte=255"`
	FullName     string    `json:"full_name" validate:"required,gte=0,lte=255"`
	LegalName    string    `json:"legal_name" validate:"required,gte=0,lte=255"`
	Gaji         float64   `json:"gaji" validate:"required,gte=0"`
	TempatLahir  string    `json:"tempat_lahir" validate:"required,gte=0,lte=255"`
	TanggalLahir time.Time `json:"tanggal_lahir" validate:"required"`
	FotoKtp      string    `json:"foto_ktp" validate:"required"`
	FotoSelfie   string    `json:"foto_selfie" validate:"required"`
	Email        string    `json:"email" validate:"required,gte=0,lte=255"`
	Password     string    `json:"password" validate:"required,gte=0,lte=255"`
}

func NewCreateKonsumenCommand(
	idKonsumen uuid.UUID,
	nik string,
	fullName string,
	legalName string,
	gaji float64,
	tempatLahir string,
	tanggalLahir time.Time,
	fotoKtp string,
	fotoSelfie string,
	email string,
	password string,
) *CreateKonsumenCommand {
	return &CreateKonsumenCommand{
		IdKonsumen:   idKonsumen,
		Nik:          nik,
		FullName:     fullName,
		LegalName:    legalName,
		Gaji:         gaji,
		TempatLahir:  tempatLahir,
		TanggalLahir: tanggalLahir,
		FotoKtp:      fotoKtp,
		FotoSelfie:   fotoSelfie,
		Email:        email,
		Password:     password,
	}
}

type CreateLimitCommand struct {
	IdLimit     uuid.UUID `json:"id_limit" validate:"required"`
	IdKonsumen  uuid.UUID `json:"id_konsumen" validate:"required"`
	Tenor       uint32    `json:"tenor" validate:"required"`
	BatasKredit float64   `json:"batas_kredit" validate:"required"`
}

func NewCreateLimitCommand(
	idLimit uuid.UUID,
	idKonsumen uuid.UUID,
	tenor uint32,
	batasKredit float64,
) *CreateLimitCommand {
	return &CreateLimitCommand{
		IdLimit:     idLimit,
		IdKonsumen:  idKonsumen,
		Tenor:       tenor,
		BatasKredit: batasKredit,
	}
}

type CreateTransaksiCommand struct {
	IdTransaksi      uuid.UUID `json:"id_transaksi" validate:"required"`
	IdKonsumen       uuid.UUID `json:"id_konsumen" validate:"required"`
	NomorKontrak     string    `json:"nomor_kontrak" validate:"required"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi" validate:"required"`
	Otr              uint64    `json:"otr" validate:"required"`
	AdminFee         float64   `json:"admin_fee" validate:"required"`
	JumlahCicilan    float64   `json:"jumlah_cicilan" validate:"required"`
	JumlahBunga      float64   `json:"jumlah_bunga" validate:"required"`
	NamaAsset        string    `json:"nama_asset" validate:"required"`
	JenisTransaksi   string    `json:"jenis_transaksi" validate:"required"`
}

func NewCreateTransaksiCommand(
	idTransaksi uuid.UUID,
	idKonsumen uuid.UUID,
	nomorKontrak string,
	tanggalTransaksi time.Time,
	otr uint64,
	adminFee float64,
	jumlahCicilan float64,
	jumlahBunga float64,
	namaAsset string,
	jenisTransaksi string,
) *CreateTransaksiCommand {
	return &CreateTransaksiCommand{
		IdTransaksi:      idTransaksi,
		IdKonsumen:       idKonsumen,
		NomorKontrak:     nomorKontrak,
		TanggalTransaksi: tanggalTransaksi,
		Otr:              otr,
		AdminFee:         adminFee,
		JumlahCicilan:    jumlahCicilan,
		JumlahBunga:      jumlahBunga,
		NamaAsset:        namaAsset,
		JenisTransaksi:   jenisTransaksi,
	}
}
