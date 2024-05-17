package mapper

import (
	kafkaMessages "github.com/0x5w4/kredit-plus/proto/kafka"
	"github.com/0x5w4/kredit-plus/writer-service/internal/model"
	writerService "github.com/0x5w4/kredit-plus/writer-service/proto/writer"
	uuid "github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func LimitToGrpcMessage(limit *model.Limit) *kafkaMessages.Limit {
	return &kafkaMessages.Limit{
		IdLimit:     limit.IdLimit.String(),
		IdKonsumen:  limit.IdKonsumen.String(),
		Tenor:       limit.Tenor,
		BatasKredit: limit.BatasKredit,
		CreatedAt:   timestamppb.New(limit.CreatedAt),
		UpdatedAt:   timestamppb.New(limit.UpdatedAt),
	}
}

func LimitFromGrpcMessage(limit *kafkaMessages.Limit) (*model.Limit, error) {

	idLimit, err := uuid.Parse(limit.GetIdLimit())
	if err != nil {
		return nil, err
	}
	idKonsumen, err := uuid.Parse(limit.GetIdKonsumen())
	if err != nil {
		return nil, err
	}

	return &model.Limit{
		IdLimit:     idLimit,
		IdKonsumen:  idKonsumen,
		Tenor:       limit.GetTenor(),
		BatasKredit: limit.GetBatasKredit(),
		CreatedAt:   limit.GetCreatedAt().AsTime(),
		UpdatedAt:   limit.GetUpdatedAt().AsTime(),
	}, nil
}

func WriterLimitToGrpc(limit *model.Limit) *writerService.Limit {
	return &writerService.Limit{
		IdLimit:     limit.IdLimit.String(),
		IdKonsumen:  limit.IdKonsumen.String(),
		Tenor:       limit.Tenor,
		BatasKredit: limit.BatasKredit,
		CreatedAt:   timestamppb.New(limit.CreatedAt),
		UpdatedAt:   timestamppb.New(limit.UpdatedAt),
	}
}
