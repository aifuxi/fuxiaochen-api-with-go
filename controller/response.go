package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseStruct struct {
	Code ResponseCode `json:"code"`
	Msg  any          `json:"msg"`
	Data any          `json:"data"`
}

func ResponseSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, ResponseStruct{
		Code: CodeSuccess,
		Msg:  CodeSuccess.GetMsg(),
		Data: data,
	})
}

func ResponseError(c *gin.Context, code ResponseCode) {
	c.JSON(http.StatusOK, ResponseStruct{
		Code: code,
		Msg:  code.GetMsg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResponseCode, msg any) {
	c.JSON(http.StatusOK, ResponseStruct{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
