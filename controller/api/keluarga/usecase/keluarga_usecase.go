package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/kekasicoid/kekasigohelper"
	"github.com/kekasicoid/kekasigohelper/httpclient"
	"github.com/kekasicoid/restapi-socketio/domain"
	"github.com/kekasicoid/restapi-socketio/model"
	"github.com/kekasicoid/restapi-socketio/table"
	"github.com/spf13/viper"
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

// GetProductById implements domain.KeluargaUsecase
func (*KeluargaUsecase) GetProductById(ctx context.Context, req string) (res interface{}, err error) {
	reqOption := httpclient.RequestOptions{
		Payload:       nil,
		URL:           viper.Get("VENDOR_PRODUCT_URL").(string) + "/" + req,
		TimeoutSecond: 60,
		Method:        http.MethodGet,
		Context:       ctx,
		Header: map[string]string{
			httpclient.ContentType: httpclient.MediaTypeJSON,
		},
	}

	resp, err := httpclient.Request(reqOption)
	if err != nil {
		kekasigohelper.LoggerWarning(err.Error())
		return nil, errors.New(model.GENERAL_MSG_COMM)
	}
	if resp.Status() != http.StatusOK {
		kekasigohelper.LoggerWarning("http.Do() error: " + string(resp.RawByte()))
		return nil, errors.New(resp.String())
	}
	_ = json.Unmarshal(resp.RawByte(), &res)
	return res, nil
}

// GetAllProduct implements domain.KeluargaUsecase
func (*KeluargaUsecase) GetAllProduct(ctx context.Context) (res interface{}, err error) {
	reqOption := httpclient.RequestOptions{
		Payload:       nil,
		URL:           viper.Get("VENDOR_PRODUCT_URL").(string),
		TimeoutSecond: 60,
		Method:        http.MethodGet,
		Context:       ctx,
		Header: map[string]string{
			httpclient.ContentType: httpclient.MediaTypeJSON,
		},
	}

	resp, err := httpclient.Request(reqOption)
	if err != nil {
		kekasigohelper.LoggerWarning(err.Error())
		return nil, errors.New(model.GENERAL_MSG_COMM)
	}
	if resp.Status() != http.StatusOK {
		kekasigohelper.LoggerWarning("http.Do() error: " + string(resp.RawByte()))
		return nil, errors.New(resp.String())
	}
	_ = json.Unmarshal(resp.RawByte(), &res)
	return res, nil
}

// GetKeluarga implements domain.KeluargaUsecase
func (u *KeluargaUsecase) GetKeluarga(ctx context.Context, req *domain.ReqGetKeluarga) (res []table.Orang, err error) {
	return u.keluargaRepo.GetKeluarga(ctx, req)
}

// SwitchKeluarga implements domain.KeluargaUsecase
func (u *KeluargaUsecase) SwitchKeluarga(ctx context.Context, req *domain.ReqSwitchKeluarga) (err error) {
	return u.keluargaRepo.SwitchKeluarga(ctx, req)
}

// SwitchKeluarga implements domain.KeluargaUsecase
func (u *KeluargaUsecase) CheckOrangById(ctx context.Context, req int) (err error) {
	return u.keluargaRepo.CheckOrangById(ctx, req)
}

// DeleteKeluarga implements domain.KeluargaUsecase
func (u *KeluargaUsecase) DeleteKeluarga(ctx context.Context, req *domain.ReqDeleteKeluarga) (err error) {
	return u.keluargaRepo.DeleteKeluarga(ctx, req)
}

// UpdateKeluarga implements domain.KeluargaUsecase
func (u *KeluargaUsecase) UpdateKeluarga(ctx context.Context, req *domain.ReqUpdateKeluarga) (err error) {
	return u.keluargaRepo.UpdateKeluarga(ctx, req)
}

// AddKeluarga implements domain.KeluargaUsecase
func (u *KeluargaUsecase) AddKeluarga(ctx context.Context, req *domain.ReqAddKeluarga) (err error) {
	return u.keluargaRepo.AddKeluarga(ctx, req)
}
