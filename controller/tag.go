package controller

import (
	"fuxiaochen-api-with-go/logic"
	"fuxiaochen-api-with-go/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetTags(c *gin.Context) {
	tags, err := logic.GetTags()
	if err != nil {
		ResponseErrorWithMsg(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, tags)
}

func CreatTag(c *gin.Context) {
	var params model.ParamsCreateTag

	if err := c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams)
		return
	}

	tag, err := logic.CreateTag(params)
	if err != nil {
		ResponseErrorWithMsg(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, tag)
}

func UpdateTag(c *gin.Context) {
	var params model.ParamsUpdateTag
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

	tag, err := logic.UpdateTagByID(id, params)
	if err != nil {
		ResponseErrorWithMsg(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, tag)
}

func GetTag(c *gin.Context) {
	tmp := c.Param("id")

	id, parseErr := strconv.ParseInt(tmp, 10, 64)
	if parseErr != nil {
		ResponseError(c, CodeIncorrectIDParams)
		return
	}

	tag, err := logic.GetTagByID(id)
	if err != nil {
		ResponseErrorWithMsg(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, tag)
}

func DeleteTag(c *gin.Context) {
	tmp := c.Param("id")

	id, parseErr := strconv.ParseInt(tmp, 10, 64)
	if parseErr != nil {
		ResponseError(c, CodeIncorrectIDParams)
		return
	}

	tag, err := logic.DeleteTagByID(id)
	if err != nil {
		ResponseErrorWithMsg(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, tag)
}
