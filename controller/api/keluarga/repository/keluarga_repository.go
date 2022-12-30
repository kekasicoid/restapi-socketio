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

// DeleteAssetKeluarga implements domain.KeluargaRepository
func (r *KeluargaRepository) DeleteAssetKeluarga(ctx context.Context, req *table.Asset) (err error) {
	if err := r.Conn.WithContext(ctx).Debug().Where("orang_id = ? and id_product = ? ", req.OrangID, req.IdProduct).Delete(req).Error; err != nil {
		return err
	}
	return nil
}

// UpdateAssetKeluarga implements domain.KeluargaRepository
func (r *KeluargaRepository) UpdateAssetKeluarga(ctx context.Context, req *table.Asset) (err error) {
	// db Transaction
	trx := r.Conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// menghapus product
		rm := new(table.Asset)
		if err := tx.WithContext(ctx).Where("orang_id = ? and id = ? ", req.OrangID, req.Id).Delete(&rm).Error; err != nil {
			kekasigohelper.LoggerWarning("keluarga_repository.UpdateAssetKeluarga.delete " + err.Error())
			return err
		}

		// menambah product
		req.Id = 0
		if err := tx.WithContext(ctx).Create(req).Error; err != nil {
			kekasigohelper.LoggerWarning("keluarga_repository.UpdateAssetKeluarga.create " + err.Error())
			return err
		}
		return nil
	})
	return trx
}

// AddAssetKeluarga implements domain.KeluargaRepository
func (r *KeluargaRepository) AddAssetKeluarga(ctx context.Context, req *table.Asset) (err error) {
	if err := r.Conn.WithContext(ctx).Create(req).Error; err != nil {
		return err
	}
	return nil
}

// GetKeluarga implements domain.KeluargaRepository
func (r *KeluargaRepository) GetKeluarga(ctx context.Context, req *domain.ReqGetKeluarga) (res []table.Orang, err error) {
	dOrang := []table.Orang{}
	r.Conn.WithContext(ctx).Where("id = ?", req.IdKeluarga).Preload("Anak").Find(&dOrang)
	if len(dOrang) > 0 {
		return dOrang, nil
	}
	return nil, errors.New(model.ErrRecordNotFound)
}

// SwitchKeluarga implements domain.KeluargaRepository
func (r *KeluargaRepository) SwitchKeluarga(ctx context.Context, req *domain.ReqSwitchKeluarga) (err error) {
	q := r.Conn.WithContext(ctx).Model(table.Orang{}).Select("orang_tua").Where("id = ? ", req.IdKeluarga)
	if req.OrangTua == 0 {
		q.Where("orang_tua IS NULL")
	} else {
		q.Where("orang_tua = ?", req.OrangTua)
	}
	affected := q.Updates(map[string]interface{}{"orang_tua": req.OrangTuaBaru})
	if err := affected.Error; err != nil {
		return err
	}
	if affected.RowsAffected > 0 {
		return nil
	}
	return errors.New(model.ErrRecordNotFound)
}

// CheckOrangById implements domain.KeluargaRepository
func (r *KeluargaRepository) CheckOrangById(ctx context.Context, req int) (err error) {
	dOrang := new(table.Orang)
	dOrang.Id = req
	affected := r.Conn.WithContext(ctx).First(&dOrang)
	if err := affected.Error; err != nil {
		return err
	}
	if affected.RowsAffected > 0 {
		return nil
	}
	return errors.New(model.ErrRecordNotFound)
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
