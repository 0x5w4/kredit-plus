package grpc

import (
	"context"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/0x5w4/kredit-plus/reader-service/config"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/commands"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/queries"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/service"
	"github.com/0x5w4/kredit-plus/reader-service/mapper"
	readerService "github.com/0x5w4/kredit-plus/reader-service/proto/reader"
	"github.com/go-playground/validator"
	uuid "github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcService struct {
	logger *logger.AppLogger
	cfg    *config.Config
	v      *validator.Validate
	ps     *service.KreditService
}

func NewReaderGrpcService(logger *logger.AppLogger, cfg *config.Config, v *validator.Validate, ps *service.KreditService) *grpcService {
	return &grpcService{logger: logger, cfg: cfg, v: v, ps: ps}
}

func (s *grpcService) CreateKonsumen(ctx context.Context, req *readerService.CreateKonsumenRequest) (*readerService.CreateKonsumenResponse, error) {
	idKonsumen, err := uuid.Parse(req.GetIdKonsumen())
	if err != nil {
		s.logger.SLogger.Warn("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	command := commands.NewCreateKonsumenCommand(
		idKonsumen,
		req.GetNik(),
		req.GetFullName(),
		req.GetLegalName(),
		req.GetGaji(),
		req.GetTempatLahir(),
		req.GetTanggalLahir().AsTime(),
		req.GetFotoKtp(),
		req.GetFotoSelfie(),
		req.GetEmail(),
		req.GetPassword(),
	)
	if err := s.v.StructCtx(ctx, command); err != nil {
		s.logger.SLogger.Warn("validate", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	err = s.ps.Commands.CreateKonsumen.Handle(ctx, command)
	if err != nil {
		s.logger.SLogger.Warn("CreateKonsumen.Handle", err)
		return nil, s.errResponse(codes.Internal, err)
	}

	return &readerService.CreateKonsumenResponse{IdKonsumen: idKonsumen.String()}, nil
}

func (s *grpcService) CreateLimit(ctx context.Context, req *readerService.CreateLimitRequest) (*readerService.CreateLimitResponse, error) {
	idLimit, err := uuid.Parse(req.GetIdLimit())
	if err != nil {
		s.logger.SLogger.Warn("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	idKonsumen, err := uuid.Parse(req.GetIdKonsumen())
	if err != nil {
		s.logger.SLogger.Warn("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	command := commands.NewCreateLimitCommand(
		idLimit,
		idKonsumen,
		req.GetTenor(),
		req.GetBatasKredit(),
	)
	if err := s.v.StructCtx(ctx, command); err != nil {
		s.logger.SLogger.Warn("validate", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	err = s.ps.Commands.CreateLimit.Handle(ctx, command)
	if err != nil {
		s.logger.SLogger.Warn("CreateProduct.Handle", err)
		return nil, s.errResponse(codes.Internal, err)
	}

	return &readerService.CreateLimitResponse{IdLimit: idLimit.String()}, nil
}

func (s *grpcService) CreateTransaksi(ctx context.Context, req *readerService.CreateTransaksiRequest) (*readerService.CreateTransaksiResponse, error) {
	idTransaksi, err := uuid.Parse(req.GetIdTransaksi())
	if err != nil {
		s.logger.SLogger.Warn("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	idKonsumen, err := uuid.Parse(req.GetIdKonsumen())
	if err != nil {
		s.logger.SLogger.Warn("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	command := commands.NewCreateTransaksiCommand(
		idTransaksi,
		idKonsumen,
		req.GetNomorKontrak(),
		req.GetTanggalTransaksi().AsTime(),
		req.GetOtr(),
		req.GetAdminFee(),
		req.GetJumlahCicilan(),
		req.GetJumlahBunga(),
		req.GetNamaAsset(),
		req.GetJenisTransaksi(),
	)
	if err := s.v.StructCtx(ctx, command); err != nil {
		s.logger.SLogger.Warn("validate", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	err = s.ps.Commands.CreateTransaksi.Handle(ctx, command)
	if err != nil {
		s.logger.SLogger.Warn("CreateProduct.Handle", err)
		return nil, s.errResponse(codes.Internal, err)
	}

	return &readerService.CreateTransaksiResponse{IdTransaksi: idTransaksi.String()}, nil
}

func (s *grpcService) GetLimit(ctx context.Context, req *readerService.GetLimitRequest) (*readerService.GetLimitResponse, error) {
	idLimit, err := uuid.Parse(req.GetIdLimit())
	if err != nil {
		s.logger.SLogger.Warn("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	idKonsumen, err := uuid.Parse(req.GetIdKonsumen())
	if err != nil {
		s.logger.SLogger.Warn("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	query := queries.NewGetLimitQuery(idLimit, idKonsumen)
	if err := s.v.StructCtx(ctx, query); err != nil {
		s.logger.SLogger.Warn("validate", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	limit, err := s.ps.Queries.GetLimit.Handle(ctx, query)
	if err != nil {
		s.logger.SLogger.Warn("GetLimit.Handle", err)
		return nil, s.errResponse(codes.Internal, err)
	}

	return &readerService.GetLimitResponse{Limit: mapper.ReaderLimitToGrpc(limit)}, nil
}

func (s *grpcService) GetTransaksi(ctx context.Context, req *readerService.GetTransaksiRequest) (*readerService.GetTransaksiResponse, error) {
	idTransaksi, err := uuid.Parse(req.GetIdTransaksi())
	if err != nil {
		s.logger.SLogger.Warn("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	idKonsumen, err := uuid.Parse(req.GetIdKonsumen())
	if err != nil {
		s.logger.SLogger.Warn("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	query := queries.NewGetTransaksiQuery(idTransaksi, idKonsumen)
	if err := s.v.StructCtx(ctx, query); err != nil {
		s.logger.SLogger.Warn("validate", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	limit, err := s.ps.Queries.GetTransaksi.Handle(ctx, query)
	if err != nil {
		s.logger.SLogger.Warn("GetTransaksi.Handle", err)
		return nil, s.errResponse(codes.Internal, err)
	}

	return &readerService.GetTransaksiResponse{Transaksi: mapper.ReaderTransaksiToGrpc(limit)}, nil
}

func (s *grpcService) errResponse(c codes.Code, err error) error {
	return status.Error(c, err.Error())
}
