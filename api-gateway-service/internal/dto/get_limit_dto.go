package dto

import (
	uuid "github.com/google/uuid"

	readerService "github.com/0x5w4/kredit-plus/reader-service/proto/reader"
)

type GetLimitHttpRequest struct {
	IdLimit    string `json:"id_limit" validate:"required"`
	IdKonsumen string `json:"id_konsumen" validate:"required"`
}

type GetLimitRequestDto struct {
	IdLimit    uuid.UUID `json:"id_limit" validate:"required"`
	IdKonsumen uuid.UUID `json:"id_konsumen" validate:"required"`
}

type GetLimitResponseDto struct {
	IdLimit     uuid.UUID `json:"id_limit" validate:"required"`
	IdKonsumen  uuid.UUID `json:"id_konsumen" validate:"required"`
	Tenor       uint32    `json:"tenor" validate:"required,gte=0"`
	BatasKredit float64   `json:"batas_kredit" validate:"required,gte=0"`
}

func LimitResponseFromGrpc(limit *readerService.Limit) (*GetLimitResponseDto, error) {
	idLimit, err := uuid.Parse(limit.GetIdLimit())
	if err != nil {
		return nil, err
	}

	idKonsumen, err := uuid.Parse(limit.GetIdKonsumen())
	if err != nil {
		return nil, err
	}

	return &GetLimitResponseDto{
		IdLimit:     idLimit,
		IdKonsumen:  idKonsumen,
		Tenor:       limit.GetTenor(),
		BatasKredit: limit.GetBatasKredit(),
	}, nil
}
