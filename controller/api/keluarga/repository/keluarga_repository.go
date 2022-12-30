package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/kekasicoid/kekasigohelper"
	"github.com/kekasicoid/restapi-socketio/domain"
	"github.com/kekasicoid/restapi-socketio/model"
	"github.com/kekasicoid/restapi-socketio/table"
	"gorm.io/gorm"
)

type KeluargaRepository struct {
	Conn *gorm.DB
}

func NewKeluargaRepository(Conn *gorm.DB) domain.KeluargaRepository {
	return &KeluargaRepository{
		Conn: Conn,
	}
}

// DeleteKeluarga implements domain.KeluargaRepository
func (r *KeluargaRepository) DeleteKeluarga(ctx context.Context, req *domain.ReqDeleteKeluarga) (err error) {
	q := r.Conn.WithContext(ctx).Model(table.Orang{}).Select("orang_tua").Where("id = ? ", req.IdKeluarga)
	if req.OrangTua == 0 {
		q.Where("orang_tua IS NULL")
	} else {
		q.Where("orang_tua = ?", req.OrangTua)
	}
	affected := q.Updates(map[string]interface{}{"orang_tua": nil})
	if err := affected.Error; err != nil {
		return err
	}
	if affected.RowsAffected > 0 {
		return nil
	}
	return errors.New(model.ErrRecordNotFound)
}

// UpdateKeluarga implements domain.KeluargaRepository
func (r *KeluargaRepository) UpdateKeluarga(ctx context.Context, req *domain.ReqUpdateKeluarga) (err error) {
	dOrang := new(table.Orang)
	_ = json.Unmarshal([]byte(kekasigohelper.ObjectToString(req)), dOrang)

	q := r.Conn.WithContext(ctx).Model(&dOrang).Select("nama", "jenis_kelamin").Where("id = ?", req.IdKeluarga)
	if req.OrangTua == 0 {
		q.Where("orang_tua IS NULL")
	} else {
		q.Where("orang_tua = ?", req.OrangTua)
	}
	update := q.Updates(&dOrang)
	if err := update.Error; err != nil {
		return err
	}
	if update.RowsAffected > 0 {
		return nil
	}
	return errors.New(model.ErrRecordNotFound)
}

// AddKeluarga implements domain.KeluargaRepository
func (r *KeluargaRepository) AddKeluarga(ctx context.Context, req *domain.ReqAddKeluarga) (err error) {
	dOrang := new(table.Orang)
	_ = json.Unmarshal([]byte(kekasigohelper.ObjectToString(req)), dOrang)
	if err := r.Conn.WithContext(ctx).Create(dOrang).Error; err != nil {
		return err
	}
	return nil
}
