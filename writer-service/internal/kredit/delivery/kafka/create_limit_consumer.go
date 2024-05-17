package kafka

import (
	"context"

	kafkaMessages "github.com/0x5w4/kredit-plus/proto/kafka"
	"github.com/0x5w4/kredit-plus/writer-service/internal/kredit/commands"
	"github.com/avast/retry-go"
	uuid "github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (mp *MessageProcessor) processCreateLimit(ctx context.Context, r *kafka.Reader, m kafka.Message) {

	var msg kafkaMessages.LimitCreate
	if err := proto.Unmarshal(m.Value, &msg); err != nil {
		mp.logger.SLogger.Warnf("proto.Unmarshal", err)
		mp.commitAndLogMsg(ctx, r, m)
		return
	}

	idLimit, err := uuid.Parse(msg.GetIdLimit())
	if err != nil {
		mp.logger.SLogger.Warnf("uuid.Parse", err)
		mp.commitAndLogMsg(ctx, r, m)
		return
	}

	idKonsumen, err := uuid.Parse(msg.GetIdKonsumen())
	if err != nil {
		mp.logger.SLogger.Warnf("uuid.Parse", err)
		mp.commitAndLogMsg(ctx, r, m)
		return
	}

	command := commands.NewCreateLimitCommand(
		idLimit,
		idKonsumen,
		msg.GetTenor(),
		msg.GetBatasKredit(),
	)
	if err := mp.v.StructCtx(ctx, command); err != nil {
		mp.logger.SLogger.Warnf("validate", err)
		mp.commitAndLogMsg(ctx, r, m)
		return
	}

	if err := retry.Do(func() error {
		return mp.ps.Commands.CreateLimit.Handle(ctx, command)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		mp.logger.SLogger.Warnf("CreateLimit.Handle", err)
		return
	}

	mp.commitAndLogMsg(ctx, r, m)
}
