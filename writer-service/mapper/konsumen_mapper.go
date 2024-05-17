package mapper

import (
	kafkaMessages "github.com/0x5w4/kredit-plus/proto/kafka"
	"github.com/0x5w4/kredit-plus/writer-service/internal/model"
	writerService "github.com/0x5w4/kredit-plus/writer-service/proto/writer"
	uuid "github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func KonsumenToGrpcMessage(konsumen *model.Konsumen) *kafkaMessages.Konsumen {
	return &kafkaMessages.Konsumen{
		IdKonsumen:   konsumen.IdKonsumen.String(),
		Nik:          konsumen.Nik,
		FullName:     konsumen.FullName,
		LegalName:    konsumen.LegalName,
		Gaji:         konsumen.Gaji,
		TempatLahir:  konsumen.TempatLahir,
		TanggalLahir: timestamppb.New(konsumen.TanggalLahir),
		FotoKtp:      konsumen.FotoKtp,
		FotoSelfie:   konsumen.FotoSelfie,
		Email:        konsumen.Email,
		Password:     konsumen.Password,
		CreatedAt:    timestamppb.New(konsumen.CreatedAt),
		UpdatedAt:    timestamppb.New(konsumen.UpdatedAt),
	}
}

func KonsumenFromGrpcMessage(konsumen *kafkaMessages.Konsumen) (*model.Konsumen, error) {

	idKonsumen, err := uuid.Parse(konsumen.GetIdKonsumen())
	if err != nil {
		return nil, err
	}

	return &model.Konsumen{
		IdKonsumen:   idKonsumen,
		Nik:          konsumen.GetNik(),
		FullName:     konsumen.GetFullName(),
		LegalName:    konsumen.GetLegalName(),
		Gaji:         konsumen.GetGaji(),
		TempatLahir:  konsumen.GetTempatLahir(),
		TanggalLahir: konsumen.GetTanggalLahir().AsTime(),
		FotoKtp:      konsumen.GetFotoKtp(),
		FotoSelfie:   konsumen.GetFotoSelfie(),
		Email:        konsumen.GetEmail(),
		Password:     konsumen.GetPassword(),
		CreatedAt:    konsumen.GetCreatedAt().AsTime(),
		UpdatedAt:    konsumen.GetUpdatedAt().AsTime(),
	}, nil
}

func WriterKonsumenToGrpc(konsumen *model.Konsumen) *writerService.Konsumen {
	return &writerService.Konsumen{
		IdKonsumen:   konsumen.IdKonsumen.String(),
		Nik:          konsumen.Nik,
		FullName:     konsumen.FullName,
		LegalName:    konsumen.LegalName,
		Gaji:         konsumen.Gaji,
		TempatLahir:  konsumen.TempatLahir,
		TanggalLahir: timestamppb.New(konsumen.TanggalLahir),
		FotoKtp:      konsumen.FotoKtp,
		FotoSelfie:   konsumen.FotoSelfie,
		Email:        konsumen.Email,
		Password:     konsumen.Password,
		CreatedAt:    timestamppb.New(konsumen.CreatedAt),
		UpdatedAt:    timestamppb.New(konsumen.UpdatedAt),
	}
}
