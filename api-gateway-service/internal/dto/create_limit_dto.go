package dto

import uuid "github.com/google/uuid"

type CreateLimitHttpRequest struct {
	IdKonsumen  string `json:"id_konsumen" validate:"required"`
	Tenor       uint32 `json:"tenor" validate:"required"`
	BatasKredit uint64 `json:"batas_kredit" validate:"required"`
}

type CreateLimitRequestDto struct {
	IdLimit     uuid.UUID `json:"id_limit" validate:"required"`
	IdKonsumen  uuid.UUID `json:"id_konsumen" validate:"required"`
	Tenor       uint32    `json:"tenor" validate:"required"`
	BatasKredit float64   `json:"batas_kredit" validate:"required"`
}

type CreateLimitResponseDto struct {
	IdLimit uuid.UUID `json:"id_limit" validate:"required"`
}
