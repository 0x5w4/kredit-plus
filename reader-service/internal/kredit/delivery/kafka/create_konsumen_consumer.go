package kafka

import (
	"context"
	"time"

	kafkaMessages "github.com/0x5w4/kredit-plus/proto/kafka"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/commands"
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

func (mp *MessageProcessor) processKonsumenCreated(ctx context.Context, r *kafka.Reader, m kafka.Message) {

	var msg kafkaMessages.KonsumenCreated
	if err := proto.Unmarshal(m.Value, &msg); err != nil {
		mp.logger.SLogger.Warnf("proto.Unmarshal", err)
		mp.commitAndLogMsg(ctx, r, m)
		return
	}

	k := msg.GetKonsumen()
	idKonsumen, err := uuid.Parse(k.GetIdKonsumen())
	if err != nil {
		mp.logger.SLogger.Warnf("uuid.Parse", err)
		mp.commitAndLogMsg(ctx, r, m)
		return
	}

	command := commands.NewCreateKonsumenCommand(
		idKonsumen,
		k.GetNik(),
		k.GetFullName(),
		k.GetLegalName(),
		k.GetGaji(),
		k.GetTempatLahir(),
		k.GetTanggalLahir().AsTime(),
		k.GetFotoKtp(),
		k.GetFotoSelfie(),
		k.GetEmail(),
		k.GetPassword(),
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
