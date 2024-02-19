package logic

import (
	"fuxiaochen-api-with-go/dao/mysql"
	"fuxiaochen-api-with-go/model"
)

func GetCategories() (categories []model.Category, err error) {
	return mysql.GetCategories()
}

func CreateCategory(params model.ParamsCreateCategory) (category model.Category, err error) {
	return mysql.CreateCategory(params)
}

func GetCategoryByID(id int64) (category model.Category, err error) {
	return mysql.GetCategoryByID(id)
}

func DeleteCategoryByID(id int64) (category model.Category, err error) {
	return mysql.DeleteCategoryByID(id)
}

func UpdateCategoryByID(id int64, params model.ParamsUpdateCategory) (category model.Category, err error) {
	return mysql.UpdateCategoryByID(id, params)
}
