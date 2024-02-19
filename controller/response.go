package controller

import (
	"fuxiaochen-api-with-go/global"
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

	global.L.Error(code.GetMsg())
}

func ResponseErrorWithErr(c *gin.Context, code ResponseCode, err error) {
	c.JSON(http.StatusOK, ResponseStruct{
		Code: code,
		Msg:  err.Error(),
		Data: nil,
	})

	global.L.Error(err)
}
