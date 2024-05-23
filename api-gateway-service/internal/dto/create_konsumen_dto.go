package dto

import (
	"time"

	uuid "github.com/google/uuid"
)

type CreateKonsumenHttpRequest struct {
	Nik          string `json:"nik" validate:"required,gte=0,lte=255"`
	FullName     string `json:"full_name" validate:"required,gte=0,lte=255"`
	LegalName    string `json:"legal_name" validate:"required,gte=0,lte=255"`
	Gaji         uint64 `json:"gaji" validate:"required"`
	TempatLahir  string `json:"tempat_lahir" validate:"required,gte=0,lte=255"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	FotoKtp      string `json:"foto_ktp" validate:"required"`
	FotoSelfie   string `json:"foto_selfie" validate:"required"`
	Email        string `json:"email" validate:"required,gte=0,lte=255"`
	Password     string `json:"password" validate:"required,gte=0,lte=255"`
}

type CreateKonsumenRequestDto struct {
	IdKonsumen   uuid.UUID `json:"id_konsumen" validate:"required"`
	Nik          string    `json:"nik" validate:"required,gte=0,lte=255"`
	FullName     string    `json:"full_name" validate:"required,gte=0,lte=255"`
	LegalName    string    `json:"legal_name" validate:"required,gte=0,lte=255"`
	Gaji         float64   `json:"gaji" validate:"required"`
	TempatLahir  string    `json:"tempat_lahir" validate:"required,gte=0,lte=255"`
	TanggalLahir time.Time `json:"tanggal_lahir" validate:"required"`
	FotoKtp      string    `json:"foto_ktp" validate:"required"`
	FotoSelfie   string    `json:"foto_selfie" validate:"required"`
	Email        string    `json:"email" validate:"required,gte=0,lte=255"`
	Password     string    `json:"password" validate:"required,gte=0,lte=255"`
}

type CreateKonsumenResponseDto struct {
	IdKonsumen uuid.UUID `json:"id_konsumen" validate:"required"`
}
