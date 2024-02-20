package logic

import (
	"fuxiaochen-api-with-go/dao/mysql"
	"fuxiaochen-api-with-go/model"
)

func GetPosts() (posts []model.Post, err error) {
	return mysql.GetPosts()
}

func CreatePost(params model.ParamsCreatePost) (post model.Post, err error) {
	return mysql.CreatePost(params)
}

func GetPostByID(id int64) (post model.Post, err error) {
	return mysql.GetPostByID(id)
}

func DeletePostByID(id int64) (post model.Post, err error) {
	return mysql.DeletePostByID(id)
}

func UpdatePostByID(id int64, params model.ParamsUpdatePost) (post model.Post, err error) {
	return mysql.UpdatePostByID(id, params)
}
