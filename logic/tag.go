package logic

import (
	"fuxiaochen-api-with-go/dao/mysql"
	"fuxiaochen-api-with-go/model"
	"fuxiaochen-api-with-go/model/param"
)

func GetTags(params param.ParamsGetTags) (tags []model.Tag, total int64, err error) {
	return mysql.GetTags(params)
}

func CreateTag(params param.ParamsCreateTag) (tags model.Tag, err error) {
	return mysql.CreateTag(params)
}

func GetTagByID(id int64) (tags model.Tag, err error) {
	return mysql.GetTagByID(id)
}

func DeleteTagByID(id int64) (tags model.Tag, err error) {
	return mysql.DeleteTagByID(id)
}

func UpdateTagByID(id int64, params param.ParamsUpdateTag) (tags model.Tag, err error) {
	return mysql.UpdateTagByID(id, params)
}
