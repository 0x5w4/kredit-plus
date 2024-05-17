package mongo

import (
	"context"
	"time"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout  = 30 * time.Second
	maxConnIdleTime = 3 * time.Minute
	minPoolSize     = 20
	maxPoolSize     = 300
)

func NewMongoDBConn(cfg Config, logger *logger.AppLogger) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.URI).
		SetAuth(options.Credential{Username: cfg.User, Password: cfg.Password}).
		SetConnectTimeout(connectTimeout).
		SetMaxConnIdleTime(maxConnIdleTime).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize))
	if err != nil {
		logger.SLogger.Fatalf("Failed to connect mongo: %v", err)
	}

	return client, nil
}
