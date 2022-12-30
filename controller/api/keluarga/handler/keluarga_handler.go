package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	socketio "github.com/googollee/go-socket.io"
	"github.com/kekasicoid/kekasigohelper"
	"github.com/kekasicoid/restapi-socketio/domain"
	"github.com/kekasicoid/restapi-socketio/model"
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
}

// AddKeluarga godoc
// @tags Keluarga
// @Accept  json
// @Produce  json
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

	model.HandleSuccess(g, nil)
}
