package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code int) {
	response := &ResponseData{
		Code: code,
		Msg:  getMsg(code),
		Data: nil,
	}
	c.JSON(http.StatusOK, response)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	response := &ResponseData{
		Code: CODE_SUCCESS,
		Msg:  getMsg(CODE_SUCCESS),
		Data: data,
	}
	c.JSON(http.StatusOK, response)
}
