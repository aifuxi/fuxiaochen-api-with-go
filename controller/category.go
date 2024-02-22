package controller

import (
	"fuxiaochen-api-with-go/logic"
	"fuxiaochen-api-with-go/model"
	"fuxiaochen-api-with-go/model/param"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var params param.ParamsGetCategories
	var categories []model.Category
	var total int64
	var err error

	if err = c.ShouldBindQuery(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams, err)
		return
	}

	if categories, total, err = logic.GetCategories(params); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccessWithTotal(c, categories, total)
}

func CreatCategory(c *gin.Context) {
	var params param.ParamsCreateCategory
	var category model.Category
	var err error

	if err = c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams, err)
		return
	}

	if category, err = logic.CreateCategory(params); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, category)
}

func UpdateCategory(c *gin.Context) {
	var params param.ParamsUpdateCategory
	var id int64
	var category model.Category
	var err error

	if id, err = GetIDFromParam(c); err != nil {
		ResponseError(c, CodeIncorrectIDParams, err)
		return
	}

	if err = c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams, err)
		return
	}

	if category, err = logic.UpdateCategoryByID(id, params); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, category)
}

func GetCategory(c *gin.Context) {
	var id int64
	var category model.Category
	var err error

	if id, err = GetIDFromParam(c); err != nil {
		ResponseError(c, CodeIncorrectIDParams, err)
		return
	}

	if category, err = logic.GetCategoryByID(id); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, category)
}

func DeleteCategory(c *gin.Context) {
	var id int64
	var category model.Category
	var err error

	if id, err = GetIDFromParam(c); err != nil {
		ResponseError(c, CodeIncorrectIDParams, err)
		return
	}

	if category, err = logic.DeleteCategoryByID(id); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, category)
}
