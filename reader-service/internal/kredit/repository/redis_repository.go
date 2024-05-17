package repository

import (
	"context"
	"encoding/json"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/0x5w4/kredit-plus/reader-service/config"
	"github.com/0x5w4/kredit-plus/reader-service/internal/model"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

const (
	redisKreditPrefixKey = "reader:kredit"
)

type redisRepository struct {
	logger      *logger.AppLogger
	cfg         *config.Config
	redisClient redis.UniversalClient
}

func NewRedisRepository(logger *logger.AppLogger, cfg *config.Config, redisClient redis.UniversalClient) *redisRepository {
	return &redisRepository{
		logger:      logger,
		cfg:         cfg,
		redisClient: redisClient,
	}
}

func (r *redisRepository) PutKonsumen(ctx context.Context, key string, konsumen *model.Konsumen) {

	konsumenBytes, err := json.Marshal(konsumen)
	if err != nil {
		r.logger.SLogger.Warn("json.Marshal", err)
		return
	}

	if err := r.redisClient.HSetNX(ctx, r.getRedisKreditPrefixKey(), key, konsumenBytes).Err(); err != nil {
		r.logger.SLogger.Warn("redisClient.HSetNX", err)
		return
	}
	r.logger.SLogger.Debugf("HSetNX prefix: %s, key: %s", r.getRedisKreditPrefixKey(), key)
}

func (r *redisRepository) PutLimit(ctx context.Context, key string, limit *model.Limit) {

	limitBytes, err := json.Marshal(limit)
	if err != nil {
		r.logger.SLogger.Warn("json.Marshal", err)
		return
	}

	if err := r.redisClient.HSetNX(ctx, r.getRedisKreditPrefixKey(), key, limitBytes).Err(); err != nil {
		r.logger.SLogger.Warn("redisClient.HSetNX", err)
		return
	}
	r.logger.SLogger.Debugf("HSetNX prefix: %s, key: %s", r.getRedisKreditPrefixKey(), key)
}

func (r *redisRepository) PutTransaksi(ctx context.Context, key string, transaksi *model.Transaksi) {

	transaksiBytes, err := json.Marshal(transaksi)
	if err != nil {
		r.logger.SLogger.Warn("json.Marshal", err)
		return
	}

	if err := r.redisClient.HSetNX(ctx, r.getRedisKreditPrefixKey(), key, transaksiBytes).Err(); err != nil {
		r.logger.SLogger.Warn("redisClient.HSetNX", err)
		return
	}
	r.logger.SLogger.Debugf("HSetNX prefix: %s, key: %s", r.getRedisKreditPrefixKey(), key)
}

func (r *redisRepository) GetLimit(ctx context.Context, key string) (*model.Limit, error) {

	limitBytes, err := r.redisClient.HGet(ctx, r.getRedisKreditPrefixKey(), key).Bytes()
	if err != nil {
		if err != redis.Nil {
			r.logger.SLogger.Warn("redisClient.HGet", err)
		}
		return nil, errors.Wrap(err, "redisClient.HGet")
	}

	var limit model.Limit
	if err := json.Unmarshal(limitBytes, &limit); err != nil {
		return nil, err
	}

	r.logger.SLogger.Debugf("HGet prefix: %s, key: %s", r.getRedisKreditPrefixKey(), key)
	return &limit, nil
}

func (r *redisRepository) GetTransaksi(ctx context.Context, key string) (*model.Transaksi, error) {

	transaksiBytes, err := r.redisClient.HGet(ctx, r.getRedisKreditPrefixKey(), key).Bytes()
	if err != nil {
		if err != redis.Nil {
			r.logger.SLogger.Warn("redisClient.HGet", err)
		}
		return nil, errors.Wrap(err, "redisClient.HGet")
	}

	var transaksi model.Transaksi
	if err := json.Unmarshal(transaksiBytes, &transaksi); err != nil {
		return nil, err
	}

	r.logger.SLogger.Debugf("HGet prefix: %s, key: %s", r.getRedisKreditPrefixKey(), key)
	return &transaksi, nil
}

func (r *redisRepository) getRedisKreditPrefixKey() string {
	if r.cfg.ServiceSettings.RedisKreditPrefixKey != "" {
		return r.cfg.ServiceSettings.RedisKreditPrefixKey
	}

	return redisKreditPrefixKey
}
