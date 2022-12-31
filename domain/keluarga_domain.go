package domain

import (
	"context"

	"github.com/kekasicoid/restapi-socketio/enum"
	"github.com/kekasicoid/restapi-socketio/table"
)

type KeluargaRepository interface {
	AddKeluarga(ctx context.Context, req *ReqAddKeluarga) (err error)
	UpdateKeluarga(ctx context.Context, req *ReqUpdateKeluarga) (err error)
	DeleteKeluarga(ctx context.Context, req *ReqDeleteKeluarga) (err error)
	CheckOrangById(ctx context.Context, req int) (err error)
	SwitchKeluarga(ctx context.Context, req *ReqSwitchKeluarga) (err error)
	GetKeluarga(ctx context.Context, req *ReqGetKeluarga) (res []table.Orang, err error)
	AddAssetKeluarga(ctx context.Context, req *table.Asset) (err error)
	UpdateAssetKeluarga(ctx context.Context, req *table.Asset) (err error)
	DeleteAssetKeluarga(ctx context.Context, req *table.Asset) (err error)
	GetKeluargaAssets(ctx context.Context, req *ReqGetKeluargaAssets) (res []table.Asset, err error)
}

type KeluargaUsecase interface {
	AddKeluarga(ctx context.Context, req *ReqAddKeluarga) (err error)
	UpdateKeluarga(ctx context.Context, req *ReqUpdateKeluarga) (err error)
	DeleteKeluarga(ctx context.Context, req *ReqDeleteKeluarga) (err error)
	CheckOrangById(ctx context.Context, req int) (err error)
	SwitchKeluarga(ctx context.Context, req *ReqSwitchKeluarga) (err error)
	GetKeluarga(ctx context.Context, req *ReqGetKeluarga) (res []table.Orang, err error)
	GetAllProduct(ctx context.Context) (res interface{}, err error)
	GetProductById(ctx context.Context, req string) (res interface{}, err error)
	AddAssetKeluarga(ctx context.Context, req *ReqAddAssetKeluarga, data interface{}) (err error)
	UpdateAssetKeluarga(ctx context.Context, req *ReqUpdatessetKeluarga, data interface{}) (err error)
	DeleteAssetKeluarga(ctx context.Context, req *ReqDeletessetKeluarga) (err error)
	GetKeluargaAsset(ctx context.Context, req *ReqGetKeluargaAssets) (res []ResRichesKeluarga, err error)
}

type KeluargaSioUsecase interface{}

type ReqGetKeluargaAssets struct {
	OrangTua   int `json:"orang_tua" validate:"null-numeric"`
	IdKeluarga int `json:"id_keluarga" validate:"req-numeric"`
}
type ReqDeletessetKeluarga struct {
	OrangTua   int `json:"orang_tua" validate:"null-numeric"`
	IdKeluarga int `json:"id_keluarga" validate:"req-numeric"`
	IdProduct  int `json:"id_product" validate:"req-numeric"`
	Id         int `json:"id" validate:"req-numeric"`
}
type ReqUpdatessetKeluarga struct {
	OrangTua    int `json:"orang_tua" validate:"null-numeric"`
	IdKeluarga  int `json:"id_keluarga" validate:"req-numeric"`
	IdProduct   int `json:"id_product" validate:"req-numeric"`
	ProductBaru int `json:"product_baru" validate:"req-numeric"`
}
type ReqAddAssetKeluarga struct {
	OrangTua   int `json:"orang_tua" validate:"null-numeric"`
	IdProduct  int `json:"id_product" validate:"req-numeric"`
	IdKeluarga int `json:"id_keluarga" validate:"req-numeric"`
}
type ReqGetKeluarga struct {
	IdKeluarga int `json:"id_keluarga" validate:"req-numeric"`
	OrangTua   int `json:"orang_tua,omitempty" validate:"null-numeric"`
}

type ReqSwitchKeluarga struct {
	OrangTua     int `json:"orang_tua" validate:"null-numeric"`
	OrangTuaBaru int `json:"orang_tua_baru" validate:"req-numeric"`
	IdKeluarga   int `json:"id_keluarga" validate:"req-numeric"`
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

type ResRichesKeluarga struct {
	Id           int               `json:"id"`
	Nama         string            `json:"nama"`
	JenisKelamin enum.JenisKelamin `json:"jenis_kelamin"`
	OrangTua     int               `json:"orang_tua"`
	Anak         []interface{}     `json:"anak"`
	AssetsLocal  []table.Asset
	AssetsOnline []interface{}
	TPLocal      int `json:"total_price_local"`
	TPOnline     int `json:"total_price_online"`
}
