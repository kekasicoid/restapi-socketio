package domain

import (
	"context"

	"github.com/kekasicoid/restapi-socketio/enum"
)

type KeluargaRepository interface {
	AddKeluarga(ctx context.Context, req *ReqAddKeluarga) (err error)
}

type KeluargaUsecase interface {
	AddKeluarga(ctx context.Context, req *ReqAddKeluarga) (err error)
}

type ReqAddKeluarga struct {
	Nama         string            `json:"nama" validate:"required,alphanum-space"`
	JenisKelamin enum.JenisKelamin `json:"jenis_kelamin" validate:"req-numeric"`
	OrangTua     int               `json:"orang_tua" validate:"null-numeric"`
}
