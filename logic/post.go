package logic

import (
	"fuxiaochen-api-with-go/dao/mysql"
	"fuxiaochen-api-with-go/model"
	"fuxiaochen-api-with-go/model/param"
)

func GetPosts(params param.ParamsGetPosts) (posts []model.Post, total int64, err error) {
	return mysql.GetPosts(params)
}

func CreatePost(params param.ParamsCreatePost) (post model.Post, err error) {
	return mysql.CreatePost(params)
}

func GetPostByID(id int64) (post model.Post, err error) {
	return mysql.GetPostByID(id)
}

func DeletePostByID(id int64) (post model.Post, err error) {
	return mysql.DeletePostByID(id)
}

func UpdatePostByID(id int64, params param.ParamsUpdatePost) (post model.Post, err error) {
	return mysql.UpdatePostByID(id, params)
}

func GetPublishedPostsForSite() (posts []model.Post, total int64, err error) {
	return mysql.GetPublishedPostsForSite()
}

func GetPublishedPostByIDForSite(id int64) (post model.Post, err error) {
	return mysql.GetPostByID(id)
}
