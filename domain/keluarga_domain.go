package domain

import (
	"context"

	"github.com/kekasicoid/restapi-socketio/enum"
)

type KeluargaRepository interface {
	AddKeluarga(ctx context.Context, req *ReqAddKeluarga) (err error)
	UpdateKeluarga(ctx context.Context, req *ReqUpdateKeluarga) (err error)
	DeleteKeluarga(ctx context.Context, req *ReqDeleteKeluarga) (err error)
}

type KeluargaUsecase interface {
	AddKeluarga(ctx context.Context, req *ReqAddKeluarga) (err error)
	UpdateKeluarga(ctx context.Context, req *ReqUpdateKeluarga) (err error)
	DeleteKeluarga(ctx context.Context, req *ReqDeleteKeluarga) (err error)
}

type ReqAddKeluarga struct {
	Nama         string             `json:"nama" validate:"required,alphanum-space"`
	JenisKelamin *enum.JenisKelamin `json:"jenis_kelamin" validate:"req-numeric"`
	OrangTua     int                `json:"orang_tua" validate:"null-numeric"`
}

type ReqUpdateKeluarga struct {
	Nama         string             `json:"nama" validate:"required,alphanum-space"`
	JenisKelamin *enum.JenisKelamin `json:"jenis_kelamin" validate:"req-numeric"`
	OrangTua     int                `json:"orang_tua" validate:"null-numeric"`
	IdKeluarga   int                `json:"id_keluarga" validate:"req-numeric"`
}

type ReqDeleteKeluarga struct {
	OrangTua   int `json:"orang_tua" validate:"null-numeric"`
	IdKeluarga int `json:"id_keluarga" validate:"req-numeric"`
}
