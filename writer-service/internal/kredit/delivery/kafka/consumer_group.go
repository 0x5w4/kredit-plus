package kafka

import (
	"context"
	"sync"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/0x5w4/kredit-plus/writer-service/config"
	"github.com/0x5w4/kredit-plus/writer-service/internal/kredit/service"
	"github.com/go-playground/validator"
	"github.com/segmentio/kafka-go"
)

const (
	PoolSize = 10
)

type MessageProcessor struct {
	logger *logger.AppLogger
	cfg    *config.Config
	v      *validator.Validate
	ps     *service.KreditService
}

func NewProductMessageProcessor(logger *logger.AppLogger, cfg *config.Config, v *validator.Validate, ps *service.KreditService) *MessageProcessor {
	return &MessageProcessor{logger: logger, cfg: cfg, v: v, ps: ps}
}

func (mp *MessageProcessor) ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		m, err := r.FetchMessage(ctx)
		if err != nil {
			mp.logger.SLogger.Warnf("workerID: %v, err: %v", workerID, err)
			continue
		}

		mp.logKafkaMessage(true, nil, "Kafka message received and is being processed")

		switch m.Topic {
		case mp.cfg.KafkaTopics.KonsumenCreate.TopicName:
			mp.processCreateKonsumen(ctx, r, m)
		case mp.cfg.KafkaTopics.LimitCreate.TopicName:
			mp.processCreateLimit(ctx, r, m)
		case mp.cfg.KafkaTopics.TransaksiCreate.TopicName:
			mp.processCreateTransaksi(ctx, r, m)
		}
	}
}
