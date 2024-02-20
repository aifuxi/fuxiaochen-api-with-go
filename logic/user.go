package logic

import (
	"fuxiaochen-api-with-go/dao/mysql"
	"fuxiaochen-api-with-go/model"
	"fuxiaochen-api-with-go/model/param"
)

func GetUsers(params param.ParamsGetUsers) (users []model.User, total int64, err error) {
	return mysql.GetUsers(params)
}

func CreateUser(params param.ParamsCreateUser) (users model.User, err error) {
	return mysql.CreateUser(params)
}

func GetUserByID(id int64) (users model.User, err error) {
	return mysql.GetUserByID(id)
}

func DeleteUserByID(id int64) (users model.User, err error) {
	return mysql.DeleteUserByID(id)
}

func UpdateUserByID(id int64, params param.ParamsUpdateUser) (users model.User, err error) {
	return mysql.UpdateUserByID(id, params)
}
