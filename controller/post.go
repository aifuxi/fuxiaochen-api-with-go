package controller

import (
	"fuxiaochen-api-with-go/logic"
	"fuxiaochen-api-with-go/model"
	"fuxiaochen-api-with-go/model/param"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var params param.ParamsGetPosts
	var posts []model.Post
	var total int64
	var err error

	if err = c.ShouldBindQuery(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams, err)
		return
	}

	// 从 query中获取 array 类型的数据，给params对应字段重新赋值
	params.Types = c.QueryArray("types")
	params.Statuses = c.QueryArray("statuses")

	if posts, total, err = logic.GetPosts(params); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccessWithTotal(c, posts, total)
}

func CreatPost(c *gin.Context) {
	var params param.ParamsCreatePost
	var post model.Post
	var err error

	if err = c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams, err)
		return
	}

	if post, err = logic.CreatePost(params); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, post)
}

func UpdatePost(c *gin.Context) {
	var params param.ParamsUpdatePost
	var id int64
	var post model.Post
	var err error

	if id, err = GetIDFromParam(c); err != nil {
		ResponseError(c, CodeIncorrectIDParams, err)
		return
	}

	if err = c.ShouldBindJSON(&params); err != nil {
		ResponseError(c, CodeIncorrectRequestParams, err)
		return
	}

	if post, err = logic.UpdatePostByID(id, params); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, post)
}

func GetPost(c *gin.Context) {
	var id int64
	var post model.Post
	var err error

	if id, err = GetIDFromParam(c); err != nil {
		ResponseError(c, CodeIncorrectIDParams, err)
		return
	}

	if post, err = logic.GetPostByID(id); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, post)
}

func DeletePost(c *gin.Context) {
	var id int64
	var post model.Post
	var err error

	if id, err = GetIDFromParam(c); err != nil {
		ResponseError(c, CodeIncorrectIDParams, err)
		return
	}

	if post, err = logic.DeletePostByID(id); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, post)
}

func GetPublishedPostsForSite(c *gin.Context) {
	var posts []model.Post
	var total int64
	var err error

	if posts, total, err = logic.GetPublishedPostsForSite(); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccessWithTotal(c, posts, total)
}

func GetPublishedPostBySlugForSite(c *gin.Context) {
	var post model.Post
	var err error

	slug := c.Param("slug")

	if post, err = logic.GetPublishedPostBySlugForSite(slug); err != nil {
		ResponseError(c, CodeInternalServerError, err)
		return
	}

	ResponseSuccess(c, post)
}
