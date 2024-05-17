package queries

import uuid "github.com/google/uuid"

type KreditQueries struct {
	GetLimit     GetLimitHandler
	GetTransaksi GetTransaksiHandler
}

func NewKreditQueries(getLimit GetLimitHandler, getTransaksi GetTransaksiHandler) *KreditQueries {
	return &KreditQueries{
		GetLimit:     getLimit,
		GetTransaksi: getTransaksi,
	}
}

type GetLimitQuery struct {
	IdLimit    uuid.UUID `json:"id_limit" bson:"id_limit,omitempty"`
	IdKonsumen uuid.UUID `json:"id_konsumen" bson:"id_konsumen,omitempty"`
}

func NewGetLimitQuery(idLimit uuid.UUID, idKonsumen uuid.UUID) *GetLimitQuery {
	return &GetLimitQuery{IdLimit: idLimit, IdKonsumen: idKonsumen}
}

type GetTransaksiQuery struct {
	IdTransaksi uuid.UUID `json:"id_transaksi" bson:"id_transaksi,omitempty"`
	IdKonsumen  uuid.UUID `json:"id_konsumen" bson:"id_konsumen,omitempty"`
}

func NewGetTransaksiQuery(idTransaksi uuid.UUID, idKonsumen uuid.UUID) *GetTransaksiQuery {
	return &GetTransaksiQuery{IdTransaksi: idTransaksi, IdKonsumen: idKonsumen}
}
