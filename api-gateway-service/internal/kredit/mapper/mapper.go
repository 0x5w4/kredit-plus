package mapper

import (
	"time"

	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/dto"
	"github.com/google/uuid"
)

func CreateKonsumenHttpToDto(http dto.CreateKonsumenHttpRequest) (dto.CreateKonsumenRequestDto, error) {
	tglLahir, err := time.Parse("2006-01-02", http.TanggalLahir)
	if err != nil {
		return dto.CreateKonsumenRequestDto{}, err
	}

	return dto.CreateKonsumenRequestDto{
		Nik:          http.Nik,
		FullName:     http.FullName,
		LegalName:    http.LegalName,
		Gaji:         float64(http.Gaji),
		TempatLahir:  http.TempatLahir,
		TanggalLahir: tglLahir,
		FotoKtp:      http.FotoKtp,
		FotoSelfie:   http.FotoSelfie,
		Email:        http.Email,
		Password:     http.Password,
	}, nil
}

func CreateLimitHttpToDto(http dto.CreateLimitHttpRequest) (dto.CreateLimitRequestDto, error) {
	idKonsumen, err := uuid.Parse(http.IdKonsumen)
	if err != nil {
		return dto.CreateLimitRequestDto{}, err
	}

	return dto.CreateLimitRequestDto{
		IdKonsumen:  idKonsumen,
		Tenor:       uint32(http.Tenor),
		BatasKredit: float64(http.BatasKredit),
	}, nil
}

func CreateTransaksiHttpToDto(http dto.CreateTransaksiHttpRequest) (dto.CreateTransaksiRequestDto, error) {
	idKonsumen, err := uuid.Parse(http.IdKonsumen)
	if err != nil {
		return dto.CreateTransaksiRequestDto{}, err
	}

	tglTransaksi, err := time.Parse("2006-01-02", http.TanggalTransaksi)
	if err != nil {
		return dto.CreateTransaksiRequestDto{}, err
	}

	return dto.CreateTransaksiRequestDto{
		IdKonsumen:       idKonsumen,
		NomorKontrak:     http.NomorKontrak,
		TanggalTransaksi: tglTransaksi,
		Otr:              uint64(http.Otr),
		AdminFee:         float64(http.AdminFee),
		JumlahCicilan:    float64(http.JumlahCicilan),
		JumlahBunga:      float64(http.JumlahBunga),
		NamaAsset:        http.NamaAsset,
		JenisTransaksi:   http.JenisTransaksi,
	}, nil
}

func GetTransaksiHttpToDto(http dto.GetTransaksiHttpRequest) (dto.GetTransaksiRequestDto, error) {
	idTransaksi, err := uuid.Parse(http.IdTransaksi)
	if err != nil {
		return dto.GetTransaksiRequestDto{}, err
	}
	idKonsumen, err := uuid.Parse(http.IdKonsumen)
	if err != nil {
		return dto.GetTransaksiRequestDto{}, err
	}

	return dto.GetTransaksiRequestDto{
		IdTransaksi: idTransaksi,
		IdKonsumen:  idKonsumen,
	}, nil
}

func GetLimitHttpToDto(http dto.GetLimitHttpRequest) (dto.GetLimitRequestDto, error) {
	idLimit, err := uuid.Parse(http.IdLimit)
	if err != nil {
		return dto.GetLimitRequestDto{}, err
	}
	idKonsumen, err := uuid.Parse(http.IdKonsumen)
	if err != nil {
		return dto.GetLimitRequestDto{}, err
	}

	return dto.GetLimitRequestDto{
		IdLimit:    idLimit,
		IdKonsumen: idKonsumen,
	}, nil
}
