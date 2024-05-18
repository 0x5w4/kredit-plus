package commands

import (
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/dto"
)

type KreditCommands struct {
	CreateKonsumen  CreateKonsumenCmdHandler
	CreateLimit     CreateLimitCmdHandler
	CreateTransaksi CreateTransaksiCmdHandler
}

func NewKreditCommands(createKonsumen CreateKonsumenCmdHandler, createLimit CreateLimitCmdHandler, createTransaksi CreateTransaksiCmdHandler) *KreditCommands {
	return &KreditCommands{CreateKonsumen: createKonsumen, CreateLimit: createLimit, CreateTransaksi: createTransaksi}
}

type CreateKonsumenCommand struct {
	CreateKonsumenDto *dto.CreateKonsumenRequestDto
}

func NewCreateKonsumenCommand(createKonsumenDto *dto.CreateKonsumenRequestDto) *CreateKonsumenCommand {
	return &CreateKonsumenCommand{
		CreateKonsumenDto: createKonsumenDto,
	}
}

type CreateLimitCommand struct {
	CreateLimitDto *dto.CreateLimitRequestDto
}

func NewCreateLimitCommand(createLimitDto *dto.CreateLimitRequestDto) *CreateLimitCommand {
	return &CreateLimitCommand{
		CreateLimitDto: createLimitDto,
	}
}

type CreateTransaksiCommand struct {
	CreateTransaksiDto *dto.CreateTransaksiRequestDto
}

func NewCreateTransaksiCommand(createTransaksiDto *dto.CreateTransaksiRequestDto) *CreateTransaksiCommand {
	return &CreateTransaksiCommand{
		CreateTransaksiDto: createTransaksiDto,
	}
}
