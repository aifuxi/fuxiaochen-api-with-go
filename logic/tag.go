package logic

import (
	"fuxiaochen-api-with-go/dao/mysql"
	"fuxiaochen-api-with-go/model"
)

func GetTags() (tags []model.Tag, err error) {
	return mysql.GetTags()
}

func CreateTag(params model.ParamsCreateTag) (tags model.Tag, err error) {
	return mysql.CreateTag(params)
}

func GetTagByID(id int64) (tags model.Tag, err error) {
	return mysql.GetTagByID(id)
}

func DeleteTagByID(id int64) (tags model.Tag, err error) {
	return mysql.DeleteTagByID(id)
}

func UpdateTagByID(id int64, params model.ParamsUpdateTag) (tags model.Tag, err error) {
	return mysql.UpdateTagByID(id, params)
}
