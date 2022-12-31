package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
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

// GetKeluargaAsset implements domain.KeluargaUsecase
func (u *KeluargaUsecase) GetKeluargaAsset(ctx context.Context, req *domain.ReqGetKeluargaAssets) (res []domain.ResRichesKeluarga, err error) {
	var resData = []domain.ResRichesKeluarga{}

	data, err := u.keluargaRepo.GetKeluarga(ctx, &domain.ReqGetKeluarga{IdKeluarga: req.IdKeluarga})
	if err != nil {
		return nil, err
	}
	if len(data) > 0 {
		ldata := len(data)
		for i := 0; i < ldata; i++ {
			rd := domain.ResRichesKeluarga{}
			rd.Id = data[i].Id
			rd.Nama = data[i].Nama
			rd.OrangTua = data[i].OrangTua
			assets, _ := u.keluargaRepo.GetKeluargaAssets(ctx, &domain.ReqGetKeluargaAssets{
				OrangTua:   data[i].OrangTua,
				IdKeluarga: data[i].Id,
			})
			if len(assets) > 0 {

				for _, v := range assets {
					v.Orang = nil
					rd.TPLocal = rd.TPLocal + v.Price
					getAssetsOL, _ := u.GetProductById(ctx, strconv.Itoa(v.IdProduct))
					if getAssetsOL != nil {
						rd.TPOnline = rd.TPOnline + int(getAssetsOL.(map[string]interface{})["price"].(float64))
						rd.AssetsOnline = append(rd.AssetsOnline, getAssetsOL)
					}
					rd.AssetsLocal = append(rd.AssetsLocal, v)
				}
			}
			//hitung asset anak
			for _, k := range data[i].Anak {
				rda := domain.ResRichesKeluarga{}
				rda.Id = k.Id
				rda.Nama = k.Nama
				rda.OrangTua = k.OrangTua
				assets, _ := u.keluargaRepo.GetKeluargaAssets(ctx, &domain.ReqGetKeluargaAssets{
					OrangTua:   k.OrangTua,
					IdKeluarga: k.Id,
				})
				if len(assets) > 0 {

					for _, v := range assets {
						v.Orang = nil
						rda.TPLocal = rda.TPLocal + v.Price
						getAssetsOL, _ := u.GetProductById(ctx, strconv.Itoa(v.IdProduct))
						if getAssetsOL != nil {
							rda.TPOnline = rda.TPOnline + int(getAssetsOL.(map[string]interface{})["price"].(float64))
							rda.AssetsOnline = append(rda.AssetsOnline, getAssetsOL)
						}
						rda.AssetsLocal = append(rda.AssetsLocal, v)
					}

				}
				rd.TPLocal = rd.TPLocal + rda.TPLocal
				rd.TPOnline = rd.TPOnline + rda.TPOnline
				rd.Anak = append(rd.Anak, rda)
			}

			resData = append(resData, rd)

		}
	}

	return resData, nil
}

// DeleteAssetKeluarga implements domain.KeluargaUsecase
func (u *KeluargaUsecase) DeleteAssetKeluarga(ctx context.Context, req *domain.ReqDeletessetKeluarga) (err error) {
	dAsset := new(table.Asset)
	dAsset.OrangID = req.IdKeluarga
	dAsset.Id = req.Id
	dAsset.IdProduct = req.IdProduct
	return u.keluargaRepo.DeleteAssetKeluarga(ctx, dAsset)
}

// UpdateAssetKeluarga implements domain.KeluargaUsecase
func (u *KeluargaUsecase) UpdateAssetKeluarga(ctx context.Context, req *domain.ReqUpdatessetKeluarga, data interface{}) (err error) {
	dAsset := new(table.Asset)
	dAsset.OrangID = req.IdKeluarga
	dAsset.Id = req.IdProduct
	dAsset.IdProduct = req.ProductBaru
	dAsset.Tittle = data.(map[string]interface{})["title"].(string)
	dAsset.Description = data.(map[string]interface{})["description"].(string)
	dAsset.Price = int(data.(map[string]interface{})["price"].(float64))
	dAsset.Brand = data.(map[string]interface{})["brand"].(string)
	dAsset.Category = data.(map[string]interface{})["category"].(string)
	dAsset.Thumbnail = data.(map[string]interface{})["thumbnail"].(string)
	return u.keluargaRepo.UpdateAssetKeluarga(ctx, dAsset)
}

// AddAssetKeluarga implements domain.KeluargaUsecase
func (u *KeluargaUsecase) AddAssetKeluarga(ctx context.Context, req *domain.ReqAddAssetKeluarga, data interface{}) (err error) {
	dAsset := new(table.Asset)
	dAsset.OrangID = req.IdKeluarga
	dAsset.IdProduct = req.IdProduct
	dAsset.Tittle = data.(map[string]interface{})["title"].(string)
	dAsset.Description = data.(map[string]interface{})["description"].(string)
	dAsset.Price = int(data.(map[string]interface{})["price"].(float64))
	dAsset.Brand = data.(map[string]interface{})["brand"].(string)
	dAsset.Category = data.(map[string]interface{})["category"].(string)
	dAsset.Thumbnail = data.(map[string]interface{})["thumbnail"].(string)
	return u.keluargaRepo.AddAssetKeluarga(ctx, dAsset)
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
