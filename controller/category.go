package controller

import (
	"fuxiaochen-api-with-go/logic"
	"fuxiaochen-api-with-go/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetCategories(c *gin.Context) {
	categories, err := logic.GetCategories()
	if err != nil {
		ResponseErrorWithErr(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, categories)
}

func CreatCategory(c *gin.Context) {
	var params model.ParamsCreateCategory

	if err := c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams)
		return
	}

	tag, err := logic.CreateCategory(params)
	if err != nil {
		ResponseErrorWithErr(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, tag)
}

func UpdateCategory(c *gin.Context) {
	var params model.ParamsUpdateCategory
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

	tag, err := logic.UpdateCategoryByID(id, params)
	if err != nil {
		ResponseErrorWithErr(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, tag)
}

func GetCategory(c *gin.Context) {
	tmp := c.Param("id")

	id, parseErr := strconv.ParseInt(tmp, 10, 64)
	if parseErr != nil {
		ResponseError(c, CodeIncorrectIDParams)
		return
	}

	tag, err := logic.GetCategoryByID(id)
	if err != nil {
		ResponseErrorWithErr(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, tag)
}

func DeleteCategory(c *gin.Context) {
	tmp := c.Param("id")

	id, parseErr := strconv.ParseInt(tmp, 10, 64)
	if parseErr != nil {
		ResponseError(c, CodeIncorrectIDParams)
		return
	}

	tag, err := logic.DeleteCategoryByID(id)
	if err != nil {
		ResponseErrorWithErr(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, tag)
}
