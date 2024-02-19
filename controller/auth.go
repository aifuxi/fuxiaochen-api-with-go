package controller

import (
	"errors"
	"fuxiaochen-api-with-go/logic"
	"fuxiaochen-api-with-go/model"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var params model.ParamsLogin

	if err := c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams)
		return
	}

	token, err := logic.Login(params)
	if err != nil {
		if errors.Is(err, logic.ErrorUsernameOrPassword) {
			ResponseError(c, CodeIncorrectUsernameOrPassword)
			return
		}
		ResponseErrorWithErr(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, map[string]any{
		"token": token,
	})
}
