package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	maxConn           = 50
	healthCheckPeriod = 1 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifetime   = 3 * time.Minute
	minConns          = 10
	lazyConnect       = false
)

func NewPgxConn(cfg Config, logger logger.AppLogger) (*pgxpool.Pool, error) {
	dataSourceName := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%v",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)

	poolCfg, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		logger.SLogger.Fatalf("Failed to parse postgres config: %v", err)
	}

	poolCfg.MaxConns = maxConn
	poolCfg.HealthCheckPeriod = healthCheckPeriod
	poolCfg.MaxConnIdleTime = maxConnIdleTime
	poolCfg.MaxConnLifetime = maxConnLifetime
	poolCfg.MinConns = minConns

	connPool, err := pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		logger.SLogger.Fatalf("Failed to create postgres pool: %v", err)
	}

	return connPool, nil
}
