package model

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/kekasicoid/kekasigohelper"
)

func SocketHandleNotifikasi(s *socketio.Server, namespace string, room string, event string, data interface{}) {
	responseData := Response{
		ResponseModel: ResponseModel{
			Status:  "200",
			Message: "Notifikasi"},
		Data: data,
	}
	s.BroadcastToRoom(namespace, room, event, responseData)
	kekasigohelper.LoggerWarning("Notifikasi dikirim")
}

func SocketHandleSuccess(s socketio.Conn, event string, data interface{}) {
	responseData := Response{
		ResponseModel: ResponseModel{
			Status:  "200",
			Message: "Success"},
		Data: data,
	}
	s.Emit(event, responseData)
}

func SocketHandleError(s socketio.Conn, event string, status int, message string) {
	responseData := Response{
		ResponseModel: ResponseModel{
			Status:  strconv.Itoa(status),
			Message: message},
		Data: nil,
	}
	s.Emit(event, responseData)
}

func HandleSuccess(c *gin.Context, data interface{}) {
	responseData := Response{
		ResponseModel: ResponseModel{
			Status:  "200",
			Message: "Success",
		},
		Data: data,
	}
	c.JSON(http.StatusOK, responseData)
}

func HandleError(c *gin.Context, status int, message string) {
	responseData := Response{
		ResponseModel: ResponseModel{
			Status:  strconv.Itoa(status),
			Message: message},
		Data: nil,
	}
	c.JSON(status, responseData)
}

type ResponseModel struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type Response struct {
	ResponseModel
	Data interface{} `json:"data"`
}
