package repository

import (
	"context"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/0x5w4/kredit-plus/reader-service/config"
	"github.com/0x5w4/kredit-plus/reader-service/internal/model"
	uuid "github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	logger *logger.AppLogger
	cfg    *config.Config
	db     *mongo.Client
}

func NewMongoRepository(logger *logger.AppLogger, cfg *config.Config, db *mongo.Client) *mongoRepository {
	return &mongoRepository{
		logger: logger,
		cfg:    cfg,
		db:     db,
	}
}

func (p *mongoRepository) CreateKonsumen(ctx context.Context, konsumen *model.Konsumen) (*model.Konsumen, error) {

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Kredit)

	_, err := collection.InsertOne(ctx, konsumen, &options.InsertOneOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "InsertOne")
	}

	return konsumen, nil
}

func (p *mongoRepository) CreateLimit(ctx context.Context, limit *model.Limit) (*model.Limit, error) {

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Kredit)

	_, err := collection.InsertOne(ctx, limit, &options.InsertOneOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "InsertOne")
	}

	return limit, nil
}

func (p *mongoRepository) CreateTransaksi(ctx context.Context, transaksi *model.Transaksi) (*model.Transaksi, error) {

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Kredit)

	_, err := collection.InsertOne(ctx, transaksi, &options.InsertOneOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "InsertOne")
	}

	return transaksi, nil
}

func (p *mongoRepository) GetLimit(ctx context.Context, idLimit uuid.UUID, idKonsumen uuid.UUID) (*model.Limit, error) {

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Kredit)

	var limit model.Limit
	if err := collection.FindOne(ctx, bson.M{"id_limit": idLimit.String(), "id_konsumen": idKonsumen.String()}).Decode(&limit); err != nil {
		return nil, errors.Wrap(err, "Decode")
	}

	return &limit, nil
}

func (p *mongoRepository) GetTransaksi(ctx context.Context, idTransaksi uuid.UUID, idKonsumen uuid.UUID) (*model.Transaksi, error) {

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Kredit)

	var transaksi model.Transaksi
	if err := collection.FindOne(ctx, bson.M{"id_transaksi": idTransaksi.String(), "id_konsumen": idKonsumen.String()}).Decode(&transaksi); err != nil {
		return nil, errors.Wrap(err, "Decode")
	}

	return &transaksi, nil
}
