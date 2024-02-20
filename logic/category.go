package logic

import (
	"fuxiaochen-api-with-go/dao/mysql"
	"fuxiaochen-api-with-go/model"
	"fuxiaochen-api-with-go/model/param"
)

func GetCategories(params param.ParamsGetCategories) (categories []model.Category, total int64, err error) {
	return mysql.GetCategories(params)
}

func CreateCategory(params param.ParamsCreateCategory) (category model.Category, err error) {
	return mysql.CreateCategory(params)
}

func GetCategoryByID(id int64) (category model.Category, err error) {
	return mysql.GetCategoryByID(id)
}

func DeleteCategoryByID(id int64) (category model.Category, err error) {
	return mysql.DeleteCategoryByID(id)
}

func UpdateCategoryByID(id int64, params param.ParamsUpdateCategory) (category model.Category, err error) {
	return mysql.UpdateCategoryByID(id, params)
}
