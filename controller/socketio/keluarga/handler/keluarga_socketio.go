package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	socketio "github.com/googollee/go-socket.io"
	"github.com/kekasicoid/kekasigohelper"
	"github.com/kekasicoid/restapi-socketio/domain"
	"github.com/kekasicoid/restapi-socketio/helper"
	"github.com/kekasicoid/restapi-socketio/model"
	"github.com/spf13/viper"
)

type KeluargaSioHandler struct {
	keluargaUC    domain.KeluargaUsecase
	keluargaSioUC *domain.KeluargaSioUsecase
	validate      *validator.Validate
	socket        *socketio.Server
}

func NewKeluargaSioHandler(g *gin.Engine, uc domain.KeluargaUsecase, sioUc domain.KeluargaSioUsecase, pg *validator.Validate, srvr *socketio.Server, Middle *helper.GoMiddleware) {
	handler := &KeluargaSioHandler{
		keluargaUC:    uc,
		keluargaSioUC: &sioUc,
		validate:      pg,
		socket:        srvr,
	}
	handler.socket.OnError("/", func(s socketio.Conn, e error) {
		kekasigohelper.LoggerWarning("socketio error:" + e.Error())
	})

	handler.socket.OnDisconnect("/", func(s socketio.Conn, msg string) {
		kekasigohelper.LoggerInfo("OnDisconnect ID : " + s.ID() + " - " + s.RemoteAddr().String() + " - " + s.LocalAddr().String())
		kekasigohelper.LoggerInfo("Message : " + msg)
	})

	handler.socket.OnConnect("/", func(s socketio.Conn) error {
		kekasigohelper.LoggerInfo("Connected ID : " + s.ID() + " - " + s.RemoteAddr().String() + " - " + s.LocalAddr().String())
		s.SetContext("")
		return nil
	})
	handler.socket.OnEvent(viper.Get("SOCK_NS_NOTIFIKASI").(string), "join", func(s socketio.Conn, msg interface{}) (err string) {
		err = handler.JoinNotifikasi(s, "join", msg)
		return err
	})

	g.GET("/kekasigen/*any", Middle.GinMiddlewareSocketIo(viper.Get("ALLOW_ORIGIN").(string)), gin.WrapH(handler.socket))
	g.POST("/kekasigen/*any", Middle.GinMiddlewareSocketIo(viper.Get("ALLOW_ORIGIN").(string)), gin.WrapH(handler.socket))
}

func (k *KeluargaSioHandler) JoinNotifikasi(s socketio.Conn, event string, msg interface{}) string {
	req := new(domain.ReqGetKeluarga)
	ctx := context.Background()
	if err := json.Unmarshal([]byte(kekasigohelper.ObjectToString(msg)), &req); err != nil {
		kekasigohelper.LoggerWarning("keluaraga_handler_socket.JoinNotifikasi.json.Unmarshal " + err.Error())
		model.SocketHandleError(s, event, http.StatusBadRequest, model.ErrJSONFormat)
		return err.Error()
	}

	if err := k.validate.Struct(req); err != nil {
		kekasigohelper.LoggerWarning("keluaraga_handler_socket.JoinNotifikasi.validate.Struct " + err.Error())
		model.SocketHandleError(s, event, http.StatusBadRequest, model.ErrInvalidParameter)
		return err.Error()
	}
	dOrang, err := k.keluargaUC.GetKeluarga(ctx, req)
	if err != nil {
		kekasigohelper.LoggerWarning("keluaraga_handler_socket.JoinNotifikasi.validate.Struct " + err.Error())
		model.SocketHandleError(s, event, http.StatusBadRequest, model.ErrRecordNotFound)
		return err.Error()
	}

	if dOrang[0].OrangTua != req.OrangTua {
		kekasigohelper.LoggerWarning("keluaraga_handler_socket.JoinNotifikasi.Check " + strconv.Itoa(req.OrangTua))
		model.SocketHandleError(s, event, http.StatusBadRequest, model.ErrRecordNotFound)
		return model.ErrRecordNotFound
	}

	info := ""
	if dOrang[0].OrangTua != 0 {
		s.Join(strconv.Itoa(dOrang[0].OrangTua))
		info = model.ANAK_KELUARGA
	} else {
		s.Join(strconv.Itoa(dOrang[0].Id))
		info = model.KEPALA_KELUARGA
	}

	model.SocketHandleSuccess(s, event, info)
	return "kekasi.co.id"
}
