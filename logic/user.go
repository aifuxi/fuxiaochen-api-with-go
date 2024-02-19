package logic

import (
	"fuxiaochen-api-with-go/dao/mysql"
	"fuxiaochen-api-with-go/model"
)

func GetUsers() (users []model.User, err error) {
	return mysql.GetUsers()
}

func CreateUser(params model.ParamsCreateUser) (users model.User, err error) {
	return mysql.CreateUser(params)
}

func GetUserByID(id int64) (users model.User, err error) {
	return mysql.GetUserByID(id)
}

func DeleteUserByID(id int64) (users model.User, err error) {
	return mysql.DeleteUserByID(id)
}

func UpdateUserByID(id int64, params model.ParamsUpdateUser) (users model.User, err error) {
	return mysql.UpdateUserByID(id, params)
}
