package controller

import (
	"fuxiaochen-api-with-go/logic"
	"fuxiaochen-api-with-go/model"
	"fuxiaochen-api-with-go/model/param"
	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
	var params param.ParamsGetTags
	var tags []model.Tag
	var total int64
	var err error

	if err = c.ShouldBindQuery(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams, err)
		return
	}

	if tags, total, err = logic.GetTags(params); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccessWithTotal(c, tags, total)
}

func CreatTag(c *gin.Context) {
	var params param.ParamsCreateTag
	var tag model.Tag
	var err error

	if err = c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams, err)
		return
	}

	if tag, err = logic.CreateTag(params); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, tag)
}

func UpdateTag(c *gin.Context) {
	var params param.ParamsUpdateTag
	var id int64
	var tag model.Tag
	var err error

	if id, err = GetIDFromParam(c); err != nil {
		ResponseError(c, CodeIncorrectIDParams, err)
		return
	}

	if err = c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams, err)
		return
	}

	if tag, err = logic.UpdateTagByID(id, params); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, tag)
}

func GetTag(c *gin.Context) {
	var id int64
	var tag model.Tag
	var err error

	if id, err = GetIDFromParam(c); err != nil {
		ResponseError(c, CodeIncorrectIDParams, err)
		return
	}

	if tag, err = logic.GetTagByID(id); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, tag)
}

func DeleteTag(c *gin.Context) {
	var id int64
	var tag model.Tag
	var err error

	if id, err = GetIDFromParam(c); err != nil {
		ResponseError(c, CodeIncorrectIDParams, err)
		return
	}

	if tag, err = logic.DeleteTagByID(id); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, tag)
}
