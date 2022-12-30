package repository

import (
	"context"
	"encoding/json"

	"github.com/kekasicoid/kekasigohelper"
	"github.com/kekasicoid/restapi-socketio/domain"
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

// AddKeluarga implements domain.KeluargaRepository
func (r *KeluargaRepository) AddKeluarga(ctx context.Context, req *domain.ReqAddKeluarga) (err error) {
	dOrang := new(table.Orang)
	_ = json.Unmarshal([]byte(kekasigohelper.ObjectToString(req)), dOrang)
	if err := r.Conn.WithContext(ctx).Debug().Create(dOrang).Error; err != nil {
		return err
	}
	return nil
}
