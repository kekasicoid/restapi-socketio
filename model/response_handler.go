package model

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
