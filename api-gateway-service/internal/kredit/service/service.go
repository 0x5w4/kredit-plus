package service

import (
	"github.com/0x5w4/kredit-plus/api-gateway-service/config"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/kredit/commands"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/kredit/queries"
	kafkaClient "github.com/0x5w4/kredit-plus/pkg/kafka"
	"github.com/0x5w4/kredit-plus/pkg/logger"
	readerService "github.com/0x5w4/kredit-plus/reader-service/proto/reader"
)

type KreditService struct {
	Commands *commands.KreditCommands
	Queries  *queries.KreditQueries
}

func NewKreditService(
	logger *logger.AppLogger,
	cfg *config.Config,
	kafkaProducer *kafkaClient.Producer,
	rsClient readerService.ReaderServiceClient,
) *KreditService {

	createKonsumenHandler := commands.NewCreateKonsumenHandler(logger, cfg, kafkaProducer)
	createLimitHandler := commands.NewCreateLimitHandler(logger, cfg, kafkaProducer)
	createTransaksiHandler := commands.NewCreateTransaksiHandler(logger, cfg, kafkaProducer)

	getLimitHandler := queries.NewGetLimitHandler(logger, cfg, rsClient)
	getTransaksiHandler := queries.NewGetTransaksiHandler(logger, cfg, rsClient)

	kreditCommands := commands.NewKreditCommands(createKonsumenHandler, createLimitHandler, createTransaksiHandler)
	kreditQueries := queries.NewKreditQueries(getLimitHandler, getTransaksiHandler)

	return &KreditService{Commands: kreditCommands, Queries: kreditQueries}
}
