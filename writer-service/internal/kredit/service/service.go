package service

import (
	kafkaClient "github.com/0x5w4/kredit-plus/pkg/kafka"
	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/0x5w4/kredit-plus/writer-service/config"
	"github.com/0x5w4/kredit-plus/writer-service/internal/kredit/commands"
	"github.com/0x5w4/kredit-plus/writer-service/internal/kredit/queries"
	"github.com/0x5w4/kredit-plus/writer-service/internal/kredit/repository"
)

type KreditService struct {
	Commands *commands.KreditCommands
	Queries  *queries.KreditQueries
}

func NewKreditService(logger *logger.AppLogger, cfg *config.Config, pgRepo repository.Repository, kafkaProducer kafkaClient.Producer) *KreditService {

	createKonsumenHandler := commands.NewCreateKonsumenHandler(logger, cfg, pgRepo, kafkaProducer)
	createLimitHandler := commands.NewCreateLimitHandler(logger, cfg, pgRepo, kafkaProducer)
	createTransaksiHandler := commands.NewCreateTransaksiHandler(logger, cfg, pgRepo, kafkaProducer)

	getLimitHandler := queries.NewGetLimitHandler(logger, cfg, pgRepo)
	getTransaksiHandler := queries.NewGetTransaksiHandler(logger, cfg, pgRepo)

	kreditCommands := commands.NewKreditCommands(createKonsumenHandler, createLimitHandler, createTransaksiHandler)
	kreditQueries := queries.NewKreditQueries(getLimitHandler, getTransaksiHandler)

	return &KreditService{Commands: kreditCommands, Queries: kreditQueries}
}
