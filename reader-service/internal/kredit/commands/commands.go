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

func NewKreditCommands(
	createKonsumen CreateKonsumenCmdHandler,
	createLimit CreateLimitCmdHandler,
	createTransaksi CreateTransaksiCmdHandler,
) *KreditCommands {
	return &KreditCommands{CreateKonsumen: createKonsumen, CreateLimit: createLimit, CreateTransaksi: createTransaksi}
}

type CreateKonsumenCommand struct {
	IdKonsumen   uuid.UUID `json:"id_konsumen" bson:"id_konsumen,omitempty"`
	Nik          string    `json:"nik,omitempty" bson:"nik,omitempty"`
	FullName     string    `json:"full_name,omitempty" bson:"full_name,omitempty"`
	LegalName    string    `json:"legal_name,omitempty" bson:"legal_name,omitempty"`
	Gaji         float64   `json:"gaji,omitempty" bson:"gaji,omitempty"`
	TempatLahir  string    `json:"tempat_lahir,omitempty" bson:"tempat_lahir,omitempty"`
	TanggalLahir time.Time `json:"tanggal_lahir,omitempty" bson:"tanggal_lahir,omitempty"`
	FotoKtp      string    `json:"foto_ktp,omitempty" bson:"foto_ktp,omitempty"`
	FotoSelfie   string    `json:"foto_selfie,omitempty" bson:"foto_selfie,omitempty"`
	Email        string    `json:"email,omitempty" bson:"email,omitempty"`
	Password     string    `json:"password,omitempty" bson:"password,omitempty"`
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
	IdLimit     uuid.UUID `json:"id_limit" bson:"id_limit,omitempty"`
	IdKonsumen  uuid.UUID `json:"id_konsumen" bson:"id_konsumen,omitempty"`
	Tenor       uint32    `json:"tenor,omitempty" bson:"tenor,omitempty"`
	BatasKredit float64   `json:"batas_kredit,omitempty" bson:"batas_kredit,omitempty"`
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
	IdTransaksi      uuid.UUID `json:"id_transaksi" bson:"id_transaksi,omitempty"`
	IdKonsumen       uuid.UUID `json:"id_konsumen" bson:"id_konsumen,omitempty"`
	NomorKontrak     string    `json:"nomor_kontrak,omitempty" bson:"nomor_kontrak,omitempty"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi,omitempty" bson:"tanggal_transaksi,omitempty"`
	Otr              uint64    `json:"otr,omitempty" bson:"otr,omitempty"`
	AdminFee         float64   `json:"admin_fee,omitempty" bson:"admin_fee,omitempty"`
	JumlahCicilan    float64   `json:"jumlah_cicilan,omitempty" bson:"jumlah_cicilan,omitempty"`
	JumlahBunga      float64   `json:"jumlah_bunga,omitempty" bson:"jumlah_bunga,omitempty"`
	NamaAsset        string    `json:"nama_asset,omitempty" bson:"nama_asset,omitempty"`
	JenisTransaksi   string    `json:"jenis_transaksi,omitempty" bson:"jenis_transaksi,omitempty"`
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
