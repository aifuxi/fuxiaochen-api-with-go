package controller

import (
	"fuxiaochen-api-with-go/logic"
	"fuxiaochen-api-with-go/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUsers(c *gin.Context) {
	users, err := logic.GetUsers()
	if err != nil {
		ResponseErrorWithErr(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, users)
}

func CreatUser(c *gin.Context) {
	var params model.ParamsCreateUser

	if err := c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams)
		return
	}

	user, err := logic.CreateUser(params)
	if err != nil {
		ResponseErrorWithErr(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, user)
}

func UpdateUser(c *gin.Context) {
	var params model.ParamsUpdateUser
	tmp := c.Param("id")

	id, parseErr := strconv.ParseInt(tmp, 10, 64)
	if parseErr != nil {
		ResponseError(c, CodeIncorrectIDParams)
		return
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams)
		return
	}

	user, err := logic.UpdateUserByID(id, params)
	if err != nil {
		ResponseErrorWithErr(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, user)
}

func GetUser(c *gin.Context) {
	tmp := c.Param("id")

	id, parseErr := strconv.ParseInt(tmp, 10, 64)
	if parseErr != nil {
		ResponseError(c, CodeIncorrectIDParams)
		return
	}

	user, err := logic.GetUserByID(id)
	if err != nil {
		ResponseErrorWithErr(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, user)
}

func DeleteUser(c *gin.Context) {
	tmp := c.Param("id")

	id, parseErr := strconv.ParseInt(tmp, 10, 64)
	if parseErr != nil {
		ResponseError(c, CodeIncorrectIDParams)
		return
	}

	user, err := logic.DeleteUserByID(id)
	if err != nil {
		ResponseErrorWithErr(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, user)
}
