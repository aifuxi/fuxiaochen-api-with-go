package controller

import (
	"fuxiaochen-api-with-go/logic"
	"fuxiaochen-api-with-go/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPosts(c *gin.Context) {
	posts, err := logic.GetPosts()
	if err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, posts)
}

func CreatPost(c *gin.Context) {
	var params model.ParamsCreatePost

	if err := c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams, err)
		return
	}

	post, err := logic.CreatePost(params)
	if err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, post)
}

func UpdatePost(c *gin.Context) {
	var params model.ParamsUpdatePost
	tmp := c.Param("id")

	id, parseErr := strconv.ParseInt(tmp, 10, 64)
	if parseErr != nil {
		ResponseError(c, CodeIncorrectIDParams, parseErr)
		return
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams, parseErr)
		return
	}

	post, err := logic.UpdatePostByID(id, params)
	if err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, post)
}

func GetPost(c *gin.Context) {
	tmp := c.Param("id")

	id, parseErr := strconv.ParseInt(tmp, 10, 64)
	if parseErr != nil {
		ResponseError(c, CodeIncorrectIDParams, parseErr)
		return
	}

	post, err := logic.GetPostByID(id)
	if err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, post)
}

func DeletePost(c *gin.Context) {
	tmp := c.Param("id")

	id, parseErr := strconv.ParseInt(tmp, 10, 64)
	if parseErr != nil {
		ResponseError(c, CodeIncorrectIDParams, parseErr)
		return
	}

	post, err := logic.DeletePostByID(id)
	if err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, post)
}
