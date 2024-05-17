package kafka

import (
	"context"

	kafkaMessages "github.com/0x5w4/kredit-plus/proto/kafka"
	"github.com/0x5w4/kredit-plus/reader-service/internal/kredit/commands"
	"github.com/avast/retry-go"
	uuid "github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (mp *MessageProcessor) processTransaksiCreated(ctx context.Context, r *kafka.Reader, m kafka.Message) {

	var msg kafkaMessages.TransaksiCreated
	if err := proto.Unmarshal(m.Value, &msg); err != nil {
		mp.logger.SLogger.Warnf("proto.Unmarshal", err)
		mp.commitAndLogMsg(ctx, r, m)
		return
	}

	t := msg.GetTransaksi()
	idTransaksi, err := uuid.Parse(t.GetIdTransaksi())
	if err != nil {
		mp.logger.SLogger.Warnf("uuid.Parse", err)
		mp.commitAndLogMsg(ctx, r, m)
		return
	}

	idKonsumen, err := uuid.Parse(t.GetIdKonsumen())
	if err != nil {
		mp.logger.SLogger.Warnf("uuid.Parse", err)
		mp.commitAndLogMsg(ctx, r, m)
		return
	}

	command := commands.NewCreateTransaksiCommand(
		idTransaksi,
		idKonsumen,
		t.GetNomorKontrak(),
		t.GetTanggalTransaksi().AsTime(),
		t.GetOtr(),
		t.GetAdminFee(),
		t.GetJumlahCicilan(),
		t.GetJumlahBunga(),
		t.GetNamaAsset(),
		t.GetJenisTransaksi(),
	)
	if err := mp.v.StructCtx(ctx, command); err != nil {
		mp.logger.SLogger.Warnf("validate", err)
		mp.commitAndLogMsg(ctx, r, m)
		return
	}

	if err := retry.Do(func() error {
		return mp.ps.Commands.CreateTransaksi.Handle(ctx, command)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		mp.logger.SLogger.Warnf("CreateTransaksi.Handle", err)
		return
	}

	mp.commitAndLogMsg(ctx, r, m)
}
