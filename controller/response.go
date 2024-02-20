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

type ResponseStructWithTotal struct {
	ResponseStruct
	Total int64 `json:"total,omitempty"`
}

func ResponseSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, ResponseStruct{
		Code: CodeSuccess,
		Msg:  CodeSuccess.GetMsg(),
		Data: data,
	})
}

func ResponseSuccessWithTotal(c *gin.Context, data any, total int64) {
	c.JSON(http.StatusOK, ResponseStructWithTotal{
		ResponseStruct: ResponseStruct{
			Code: CodeSuccess, Msg: CodeSuccess.GetMsg(), Data: data,
		},
		Total: total,
	})
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {
	msg := code.GetMsg()

	if err != nil {
		global.L.Error(err)

		if code == CodeInternalServerError {
			msg = err.Error()
		}

	} else {
		global.L.Error(msg)
	}

	c.JSON(http.StatusOK, ResponseStruct{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
