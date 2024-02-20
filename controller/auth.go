package controller

import (
	"errors"
	"fuxiaochen-api-with-go/logic"
	"fuxiaochen-api-with-go/model/param"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var params param.ParamsLogin
	var err error
	var token string

	if err = c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams, err)
		return
	}

	if token, err = logic.Login(params); err != nil {
		if errors.Is(err, logic.ErrorUsernameOrPassword) {
			ResponseError(c, CodeIncorrectUsernameOrPassword, err)
			return
		}
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, map[string]any{
		"token": token,
	})
}
