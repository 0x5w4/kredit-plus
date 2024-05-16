package gorm

import (
	"context"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Provider struct {
	*gorm.DB
}

type ConfigConnProvider struct {
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
}

func NewPSQL(ctx context.Context, connString string, cfg *ConfigConnProvider, gormConfig *gorm.Config) (*Provider, error) {
	gormDb, err := gorm.Open(postgres.Open(connString), gormConfig)
	if err != nil {
		return nil, err
	}

	db, err := gormDb.WithContext(ctx).DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	return &Provider{gormDb}, nil
}
