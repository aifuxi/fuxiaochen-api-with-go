package controller

import (
	"fuxiaochen-api-with-go/logic"
	"fuxiaochen-api-with-go/model"
	"fuxiaochen-api-with-go/model/param"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var params param.ParamsGetUsers
	var users []model.User
	var total int64
	var err error

	if err = c.ShouldBindQuery(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams, err)
		return
	}

	if users, total, err = logic.GetUsers(params); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccessWithTotal(c, users, total)
}

func CreatUser(c *gin.Context) {
	var params param.ParamsCreateUser
	var user model.User
	var err error

	if err = c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams, err)
		return
	}

	if user, err = logic.CreateUser(params); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, user)
}

func UpdateUser(c *gin.Context) {
	var params param.ParamsUpdateUser
	var id int64
	var user model.User
	var err error

	if err = GetIDFromParam(c, &id); err != nil {
		ResponseError(c, CodeIncorrectIDParams, err)
		return
	}

	if err = c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams, err)
		return
	}

	if user, err = logic.UpdateUserByID(id, params); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, user)
}

func GetUser(c *gin.Context) {
	var id int64
	var user model.User
	var err error

	if err = GetIDFromParam(c, &id); err != nil {
		ResponseError(c, CodeIncorrectIDParams, err)
		return
	}

	if user, err = logic.GetUserByID(id); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, user)
}

func DeleteUser(c *gin.Context) {
	var id int64
	var user model.User
	var err error

	if err = GetIDFromParam(c, &id); err != nil {
		ResponseError(c, CodeIncorrectIDParams, err)
		return
	}

	if user, err = logic.DeleteUserByID(id); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, user)
}
