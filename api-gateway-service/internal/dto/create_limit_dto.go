package dto

import uuid "github.com/google/uuid"

type CreateLimitRequestDto struct {
	IdLimit     uuid.UUID `json:"id_limit" validate:"required"`
	IdKonsumen  uuid.UUID `json:"id_konsumen" validate:"required"`
	Tenor       uint32    `json:"tenor" validate:"required,gte=0"`
	BatasKredit float64   `json:"batas_kredit" validate:"required,gte=0"`
}

type CreateLimitResponseDto struct {
	IdLimit uuid.UUID `json:"id_limit" validate:"required"`
}
