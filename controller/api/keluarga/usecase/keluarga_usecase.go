package usecase

import (
	"context"
	"time"

	"github.com/kekasicoid/restapi-socketio/domain"
)

type KeluargaUsecase struct {
	keluargaRepo   domain.KeluargaRepository
	contextTimeout time.Duration
}

func NewKeluargaUsecase(keluargaRepo domain.KeluargaRepository, timeout time.Duration) domain.KeluargaUsecase {
	return &KeluargaUsecase{
		keluargaRepo:   keluargaRepo,
		contextTimeout: timeout,
	}
}

// AddKeluarga implements domain.KeluargaUsecase
func (u *KeluargaUsecase) AddKeluarga(ctx context.Context, req *domain.ReqAddKeluarga) (err error) {
	return u.keluargaRepo.AddKeluarga(ctx, req)
}
