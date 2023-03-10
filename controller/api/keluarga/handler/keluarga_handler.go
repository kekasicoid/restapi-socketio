package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	socketio "github.com/googollee/go-socket.io"
	"github.com/kekasicoid/kekasigohelper"
	"github.com/kekasicoid/restapi-socketio/domain"
	"github.com/kekasicoid/restapi-socketio/model"
	"github.com/spf13/viper"
)

type KeluargaHandler struct {
	keluargaUC domain.KeluargaUsecase
	validate   *validator.Validate
	socket     *socketio.Server
}

func NewKeluargaHandler(g *gin.Engine, uc domain.KeluargaUsecase, pg *validator.Validate, srvr *socketio.Server) {
	handler := &KeluargaHandler{
		keluargaUC: uc,
		validate:   pg,
		socket:     srvr,
	}

	g.POST("/keluarga/add", handler.AddKeluarga)
	g.POST("/keluarga/update", handler.UpdateKeluarga)
	g.POST("/keluarga/delete", handler.DeleteKeluarga)
	g.POST("/keluarga/switch", handler.SwitchKeluarga)
	g.POST("/keluarga/get", handler.GetKeluarga)
	g.GET("/3rd/product/all", handler.GetAllProduct)
	g.GET("/3rd/product/:id", handler.GetProductById)
	g.POST("/keluarga/asset/add", handler.AddAssetsKeluarga)
	g.POST("/keluarga/asset/update", handler.UpdateAssetsKeluarga)
	g.POST("/keluarga/asset/delete", handler.DeleteAssetsKeluarga)
	g.POST("/keluarga/asset/riches", handler.RichesAssetsKeluarga)
}

// RichesAssetsKeluarga godoc
// @tags Keluarga
// @description Dapat menampilkan total nilai (price) aset yang dimiliki oleh masing-masing orang dalam suatu keluarga berdasarkan data harga product di api produk berikut https://dummyjson.com/docs/products
// @Accept  json
// @Produce  json
// @Param Keluarga body domain.ReqGetKeluargaAssets    true  "Total Asset Keluarga"
// @Success 200 {object} model.Response
// @Router /keluarga/asset/riches [post]
func (k *KeluargaHandler) RichesAssetsKeluarga(g *gin.Context) {
	req := new(domain.ReqGetKeluargaAssets)
	ctx := g.Request.Context()
	if ctx != nil {
		ctx = context.Background()
	}
	if err := g.BindJSON(&req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.BindJSON " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrJSONFormat)
		return
	}
	if err := k.validate.Struct(req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.validate.struct " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrInvalidParameter)
		return
	}

	dOrang, err := k.keluargaUC.GetKeluargaAsset(ctx, req)
	if err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.GetKeluargaAsset " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrRecordNotFound)
		return
	}
	// if dOrang[0].Orang.OrangTua != req.OrangTua {
	// 	kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.OrangTua " + err.Error())
	// 	model.HandleError(g, http.StatusBadRequest, model.ErrRecordNotFound)
	// 	return
	// }
	room := dOrang[0].Id
	if dOrang[0].OrangTua != 0 {
		room = dOrang[0].OrangTua
	}
	model.SocketHandleNotifikasi(k.socket, viper.Get("SOCK_NS_NOTIFIKASI").(string), strconv.Itoa(room), viper.Get("SOCK_EVENT_NOTIFIKASI").(string), model.UPDATE_DATA_ASSETS)
	model.HandleSuccess(g, dOrang)
}

// DeleteAssetsKeluarga godoc
// @tags Keluarga
// @description 2.f Dapat menghapus data aset keluarga && 2.h Dapat mengurangi asset yang dimiliki oleh suatu orang dalam keluarga
// @Accept  json
// @Produce  json
// @Param Keluarga body domain.ReqDeletessetKeluarga   true  "hapus Asset Keluarga"
// @Success 200 {object} model.Response
// @Router /keluarga/asset/delete [post]
func (k *KeluargaHandler) DeleteAssetsKeluarga(g *gin.Context) {
	req := new(domain.ReqDeletessetKeluarga)
	ctx := g.Request.Context()
	if ctx != nil {
		ctx = context.Background()
	}
	if err := g.BindJSON(&req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.BindJSON " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrJSONFormat)
		return
	}
	if err := k.validate.Struct(req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.validate.struct " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrInvalidParameter)
		return
	}

	dOrang, err := k.keluargaUC.GetKeluarga(ctx, &domain.ReqGetKeluarga{IdKeluarga: req.IdKeluarga})
	if err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.GetKeluarga " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrRecordNotFound)
		return
	}
	if (len(dOrang) < 1) || (dOrang[0].OrangTua != req.OrangTua) {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.OrangTua " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrRecordNotFound)
		return
	}

	if err := k.keluargaUC.DeleteAssetKeluarga(ctx, req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.UpdateAssetKeluarga " + err.Error())
		model.HandleError(g, http.StatusBadRequest, err.Error())
		return
	}
	room := dOrang[0].Id
	if dOrang[0].OrangTua != 0 {
		room = dOrang[0].OrangTua
	}
	model.SocketHandleNotifikasi(k.socket, viper.Get("SOCK_NS_NOTIFIKASI").(string), strconv.Itoa(room), viper.Get("SOCK_EVENT_NOTIFIKASI").(string), model.UPDATE_DATA_ASSETS)
	model.HandleSuccess(g, nil)
}

// UpdateAssetsKeluarga godoc
// @tags Keluarga
// @description 2.e Dapat mengedit data aset keluarga
// @Accept  json
// @Produce  json
// @Param Keluarga body domain.ReqUpdatessetKeluarga   true  "Ubah Asset Keluarga"
// @Success 200 {object} model.Response
// @Router /keluarga/asset/update [post]
func (k *KeluargaHandler) UpdateAssetsKeluarga(g *gin.Context) {
	req := new(domain.ReqUpdatessetKeluarga)
	ctx := g.Request.Context()
	if ctx != nil {
		ctx = context.Background()
	}
	if err := g.BindJSON(&req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.BindJSON " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrJSONFormat)
		return
	}
	if err := k.validate.Struct(req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.validate.struct " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrInvalidParameter)
		return
	}

	productBaru, err := k.keluargaUC.GetProductById(ctx, strconv.Itoa(req.ProductBaru))
	if err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.GetProductById " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrProductNotFound)
		return
	}

	dOrang, err := k.keluargaUC.GetKeluarga(ctx, &domain.ReqGetKeluarga{IdKeluarga: req.IdKeluarga})
	if err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.GetKeluarga " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrRecordNotFound)
		return
	}
	if dOrang[0].OrangTua != req.OrangTua {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.OrangTua " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrRecordNotFound)
		return
	}

	if err := k.keluargaUC.UpdateAssetKeluarga(ctx, req, productBaru); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.UpdateAssetKeluarga " + err.Error())
		model.HandleError(g, http.StatusBadRequest, err.Error())
		return
	}
	room := dOrang[0].Id
	if dOrang[0].OrangTua != 0 {
		room = dOrang[0].OrangTua
	}
	model.SocketHandleNotifikasi(k.socket, viper.Get("SOCK_NS_NOTIFIKASI").(string), strconv.Itoa(room), viper.Get("SOCK_EVENT_NOTIFIKASI").(string), model.UPDATE_DATA_ASSETS)
	model.HandleSuccess(g, nil)
}

// AddAssetsKeluarga godoc
// @tags Keluarga
// @description 2.d Dapat menambah data aset keluarga && 2.g Dapat menambahkan aset yang dimiliki oleh suatu orang dalam keluarga
// @Accept  json
// @Produce  json
// @Param Keluarga body domain.ReqAddAssetKeluarga   true  "Tambah Asset Keluarga"
// @Success 200 {object} model.Response
// @Router /keluarga/asset/add [post]
func (k *KeluargaHandler) AddAssetsKeluarga(g *gin.Context) {
	req := new(domain.ReqAddAssetKeluarga)
	ctx := g.Request.Context()
	if ctx != nil {
		ctx = context.Background()
	}
	if err := g.BindJSON(&req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.BindJSON " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrJSONFormat)
		return
	}
	if err := k.validate.Struct(req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.validate.struct " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrInvalidParameter)
		return
	}

	data, err := k.keluargaUC.GetProductById(ctx, strconv.Itoa(req.IdProduct))
	if err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.GetProductById " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrProductNotFound)
		return
	}

	dOrang, err := k.keluargaUC.GetKeluarga(ctx, &domain.ReqGetKeluarga{IdKeluarga: req.IdKeluarga})
	if err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.GetKeluarga " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrRecordNotFound)
		return
	}
	if dOrang[0].OrangTua != req.OrangTua {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.OrangTua " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrRecordNotFound)
		return
	}

	if err := k.keluargaUC.AddAssetKeluarga(ctx, req, data); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.AddAssetKeluarga " + err.Error())
		model.HandleError(g, http.StatusBadRequest, err.Error())
		return
	}
	room := dOrang[0].Id
	if dOrang[0].OrangTua != 0 {
		room = dOrang[0].OrangTua
	}
	model.SocketHandleNotifikasi(k.socket, viper.Get("SOCK_NS_NOTIFIKASI").(string), strconv.Itoa(room), viper.Get("SOCK_EVENT_NOTIFIKASI").(string), model.UPDATE_DATA_ASSETS)

	model.HandleSuccess(g, nil)
}

// GetProductById godoc
// @tags 3rd Party
// @description https://dummyjson.com/docs/products#single
// @Param id path string true "Product ID"
// @Accept  json
// @Produce  json
// @Router /3rd/product/{id} [get]
func (k *KeluargaHandler) GetProductById(g *gin.Context) {
	idStr := g.Param("id")
	if err := k.validate.Var(idStr, "req-numeric"); err != nil {
		kekasigohelper.LoggerWarning(err)
		model.HandleError(g, http.StatusBadRequest, model.ErrInvalidNumber)
		return
	}
	ctx := g.Request.Context()
	if ctx != nil {
		ctx = context.Background()
	}
	data, err := k.keluargaUC.GetProductById(ctx, idStr)
	if err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.GetProductById " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrProductNotFound)
		return
	}
	model.HandleSuccess(g, data)
}

// GetAllProduct godoc
// @tags 3rd Party
// @description https://dummyjson.com/docs/products#all
// @Accept  json
// @Produce  json
// @Router /3rd/product/all [get]
func (k *KeluargaHandler) GetAllProduct(g *gin.Context) {
	ctx := g.Request.Context()
	if ctx != nil {
		ctx = context.Background()
	}
	data, err := k.keluargaUC.GetAllProduct(ctx)
	if err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.GetAllProduct " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrProductNotFound)
		return
	}

	model.HandleSuccess(g, data)
}

// GetKeluarga godoc
// @tags Keluarga
// @description Menampilkan anggota keluarga 1 tingkat di bawah
// @Accept  json
// @Produce  json
// @Param Keluarga body domain.ReqGetKeluarga   true  "Pindah Keluarga"
// @Success 200 {object} model.Response
// @Router /keluarga/get [post]
func (k *KeluargaHandler) GetKeluarga(g *gin.Context) {
	req := new(domain.ReqGetKeluarga)
	ctx := g.Request.Context()
	if ctx != nil {
		ctx = context.Background()
	}
	if err := g.BindJSON(&req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.BindJSON " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrJSONFormat)
		return
	}
	if err := k.validate.Struct(req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.validate.struct " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrInvalidParameter)
		return
	}
	data, err := k.keluargaUC.GetKeluarga(ctx, req)
	if err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.GetKeluarga " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrOrangTuaBaru)
		return
	}

	model.HandleSuccess(g, data)
}

// SwitchKeluarga godoc
// @tags Keluarga
// @description 2.a Dapat menambahkan data orang baru ke keluarga (Pindah KK)
// @Accept  json
// @Produce  json
// @Param Keluarga body domain.ReqSwitchKeluarga   true  "Pindah Keluarga"
// @Success 200 {object} model.Response
// @Router /keluarga/switch [post]
func (k *KeluargaHandler) SwitchKeluarga(g *gin.Context) {
	req := new(domain.ReqSwitchKeluarga)
	ctx := g.Request.Context()
	if ctx != nil {
		ctx = context.Background()
	}
	if err := g.BindJSON(&req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.BindJSON " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrJSONFormat)
		return
	}
	if err := k.validate.Struct(req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.validate.struct " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrInvalidParameter)
		return
	}

	if err := k.keluargaUC.CheckOrangById(ctx, req.OrangTuaBaru); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.CheckOrangById " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrOrangTuaBaru)
		return
	}

	if err := k.keluargaUC.SwitchKeluarga(ctx, req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.SwitchKeluarga " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrOrangTuaBaru)
		return
	}

	model.SocketHandleNotifikasi(k.socket, viper.Get("SOCK_NS_NOTIFIKASI").(string), strconv.Itoa(req.OrangTuaBaru), viper.Get("SOCK_EVENT_NOTIFIKASI").(string), model.UPDATE_DATA_KELUARGA)

	room := req.IdKeluarga
	if req.OrangTua != 0 {
		room = req.OrangTua
	}
	model.SocketHandleNotifikasi(k.socket, viper.Get("SOCK_NS_NOTIFIKASI").(string), strconv.Itoa(room), viper.Get("SOCK_EVENT_NOTIFIKASI").(string), model.UPDATE_DATA_KELUARGA)
	model.HandleSuccess(g, nil)
}

// DeleteKeluarga godoc
// @tags Keluarga
// @description 3.c Dapat menghapus data orang dalam keluarga
// @Accept  json
// @Produce  json
// @Param Keluarga body domain.ReqDeleteKeluarga   true  "Hapus Keluarga"
// @Success 200 {object} model.Response
// @Router /keluarga/delete [post]
func (k *KeluargaHandler) DeleteKeluarga(g *gin.Context) {
	req := new(domain.ReqDeleteKeluarga)
	ctx := g.Request.Context()
	if ctx != nil {
		ctx = context.Background()
	}
	if err := g.BindJSON(&req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.BindJSON " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrJSONFormat)
		return
	}
	if err := k.validate.Struct(req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.validate.struct " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrInvalidParameter)
		return
	}

	if err := k.keluargaUC.DeleteKeluarga(ctx, req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.DeleteKeluarga " + err.Error())
		model.HandleError(g, http.StatusBadRequest, err.Error())
		return
	}

	room := req.IdKeluarga
	if req.OrangTua != 0 {
		room = req.OrangTua
	}
	model.SocketHandleNotifikasi(k.socket, viper.Get("SOCK_NS_NOTIFIKASI").(string), strconv.Itoa(room), viper.Get("SOCK_EVENT_NOTIFIKASI").(string), model.UPDATE_DATA_KELUARGA)
	model.HandleSuccess(g, nil)
}

// UpdateKeluarga godoc
// @tags Keluarga
// @description 2.b Dapat mengedit data orang dalam keluarga
// @Accept  json
// @Produce  json
// @Param Keluarga body domain.ReqUpdateKeluarga  true  "Ubah Keluarga"
// @Success 200 {object} model.Response
// @Router /keluarga/update [post]
func (k *KeluargaHandler) UpdateKeluarga(g *gin.Context) {
	req := new(domain.ReqUpdateKeluarga)
	ctx := g.Request.Context()
	if ctx != nil {
		ctx = context.Background()
	}
	if err := g.BindJSON(&req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.BindJSON " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrJSONFormat)
		return
	}
	if err := k.validate.Struct(req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.validate.struct " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrInvalidParameter)
		return
	}

	if err := k.keluargaUC.UpdateKeluarga(ctx, req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.UpdateKeluarga " + err.Error())
		model.HandleError(g, http.StatusBadRequest, err.Error())
		return
	}
	room := req.IdKeluarga
	if req.OrangTua != 0 {
		room = req.OrangTua
	}
	model.SocketHandleNotifikasi(k.socket, viper.Get("SOCK_NS_NOTIFIKASI").(string), strconv.Itoa(room), viper.Get("SOCK_EVENT_NOTIFIKASI").(string), model.UPDATE_DATA_KELUARGA)
	model.HandleSuccess(g, nil)
}

// AddKeluarga godoc
// @tags Keluarga
// @description 2.a Dapat menambahkan data orang baru ke keluarga (baru)
// @Accept  json
// @Produce  json
// @Param Keluarga body domain.ReqAddKeluarga  true  "Tambah Keluarga"
// @Success 200 {object} model.Response
// @Router /keluarga/add [post]
func (k *KeluargaHandler) AddKeluarga(g *gin.Context) {
	req := new(domain.ReqAddKeluarga)
	ctx := g.Request.Context()
	if ctx != nil {
		ctx = context.Background()
	}
	if err := g.BindJSON(&req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.BindJSON " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrJSONFormat)
		return
	}
	if err := k.validate.Struct(req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.validate.struct " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrInvalidParameter)
		return
	}

	if err := k.keluargaUC.AddKeluarga(ctx, req); err != nil {
		kekasigohelper.LoggerWarning("keluarga_handler.keluargaUC.AddKeluarga " + err.Error())
		model.HandleError(g, http.StatusBadRequest, model.ErrFailedInsertData)
		return
	}

	if req.OrangTua != 0 {
		model.SocketHandleNotifikasi(k.socket, viper.Get("SOCK_NS_NOTIFIKASI").(string), strconv.Itoa(req.OrangTua), viper.Get("SOCK_EVENT_NOTIFIKASI").(string), model.UPDATE_DATA_KELUARGA)
	}
	model.HandleSuccess(g, nil)
}
