package service

import (
	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/0x5w4/kredit-plus/reader-service/config"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/commands"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/queries"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/repository"
)

type KreditService struct {
	Commands *commands.KreditCommands
	Queries  *queries.KreditQueries
}

func NewKreditService(
	logger *logger.AppLogger,
	cfg *config.Config,
	mongoRepo repository.Repository,
	redisRepo repository.CacheRepository,
) *KreditService {

	createKonsumenHandler := commands.NewCreateKonsumenHandler(logger, cfg, mongoRepo, redisRepo)
	createLimitHandler := commands.NewCreateLimitHandler(logger, cfg, mongoRepo, redisRepo)
	createTransaksiHandler := commands.NewCreateTransaksiHandler(logger, cfg, mongoRepo, redisRepo)

	getLimitHandler := queries.NewGetLimitHandler(logger, cfg, mongoRepo, redisRepo)
	getTransaksiHandler := queries.NewGetTransaksiHandler(logger, cfg, mongoRepo, redisRepo)

	kreditCommands := commands.NewKreditCommands(createKonsumenHandler, createLimitHandler, createTransaksiHandler)
	kreditQueries := queries.NewKreditQueries(getLimitHandler, getTransaksiHandler)

	return &KreditService{Commands: kreditCommands, Queries: kreditQueries}
}
