package kafka

import (
	"context"
	"time"

	kafkaMessages "github.com/0x5w4/kredit-plus/proto/kafka"
	"github.com/0x5w4/kredit-plus/writer-service/internal/kredit/commands"
	"github.com/avast/retry-go"
	uuid "github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
)

func (mp *MessageProcessor) processCreateKonsumen(ctx context.Context, r *kafka.Reader, m kafka.Message) {

	var msg kafkaMessages.KonsumenCreate
	if err := proto.Unmarshal(m.Value, &msg); err != nil {
		mp.logger.SLogger.Warnf("proto.Unmarshal", err)
		mp.commitAndLogMsg(ctx, r, m)
		return
	}

	idKonsumen, err := uuid.Parse(msg.GetIdKonsumen())
	if err != nil {
		mp.logger.SLogger.Warnf("uuid.Parse", err)
		mp.commitAndLogMsg(ctx, r, m)
		return
	}

	command := commands.NewCreateKonsumenCommand(
		idKonsumen,
		msg.GetNik(),
		msg.GetFullName(),
		msg.GetLegalName(),
		msg.GetGaji(),
		msg.GetTempatLahir(),
		msg.GetTanggalLahir().AsTime(),
		msg.GetFotoKtp(),
		msg.GetFotoSelfie(),
		msg.GetEmail(),
		msg.GetPassword(),
	)
	if err := mp.v.StructCtx(ctx, command); err != nil {
		mp.logger.SLogger.Warnf("validate", err)
		mp.commitAndLogMsg(ctx, r, m)
		return
	}

	if err := retry.Do(func() error {
		return mp.ps.Commands.CreateKonsumen.Handle(ctx, command)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		mp.logger.SLogger.Warnf("CreateKonsumen.Handle", err)
		return
	}

	mp.commitAndLogMsg(ctx, r, m)
}
