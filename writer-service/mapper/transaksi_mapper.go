package mapper

import (
	kafkaMessages "github.com/0x5w4/kredit-plus/proto/kafka"
	"github.com/0x5w4/kredit-plus/writer-service/internal/model"
	writerService "github.com/0x5w4/kredit-plus/writer-service/proto/writer"
	uuid "github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TransaksiToGrpcMessage(transaksi *model.Transaksi) *kafkaMessages.Transaksi {
	return &kafkaMessages.Transaksi{
		IdTransaksi:      transaksi.IdTransaksi.String(),
		IdKonsumen:       transaksi.IdKonsumen.String(),
		NomorKontrak:     transaksi.NomorKontrak,
		TanggalTransaksi: timestamppb.New(transaksi.TanggalTransaksi),
		Otr:              transaksi.Otr,
		AdminFee:         transaksi.AdminFee,
		JumlahCicilan:    transaksi.JumlahCicilan,
		JumlahBunga:      transaksi.JumlahBunga,
		NamaAsset:        transaksi.NamaAsset,
		JenisTransaksi:   transaksi.JenisTransaksi,
		CreatedAt:        timestamppb.New(transaksi.CreatedAt),
		UpdatedAt:        timestamppb.New(transaksi.UpdatedAt),
	}
}

func TransaksiFromGrpcMessage(transaksi *kafkaMessages.Transaksi) (*model.Transaksi, error) {

	idTransaksi, err := uuid.Parse(transaksi.GetIdTransaksi())
	if err != nil {
		return nil, err
	}

	idKonsumen, err := uuid.Parse(transaksi.GetIdKonsumen())
	if err != nil {
		return nil, err
	}

	return &model.Transaksi{
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
		CreatedAt:        transaksi.GetCreatedAt().AsTime(),
		UpdatedAt:        transaksi.GetUpdatedAt().AsTime(),
	}, nil
}

func WriterTransaksiToGrpc(transaksi *model.Transaksi) *writerService.Transaksi {
	return &writerService.Transaksi{
		IdTransaksi:      transaksi.IdTransaksi.String(),
		IdKonsumen:       transaksi.IdKonsumen.String(),
		NomorKontrak:     transaksi.NomorKontrak,
		TanggalTransaksi: timestamppb.New(transaksi.TanggalTransaksi),
		Otr:              transaksi.Otr,
		AdminFee:         transaksi.AdminFee,
		JumlahCicilan:    transaksi.JumlahCicilan,
		JumlahBunga:      transaksi.JumlahBunga,
		NamaAsset:        transaksi.NamaAsset,
		JenisTransaksi:   transaksi.JenisTransaksi,
		CreatedAt:        timestamppb.New(transaksi.CreatedAt),
		UpdatedAt:        timestamppb.New(transaksi.UpdatedAt),
	}
}
