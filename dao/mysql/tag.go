package mysql

import (
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/model"
	"fuxiaochen-api-with-go/model/param"
	"fuxiaochen-api-with-go/model/scope"
	"gorm.io/gorm/clause"
)

func CreateTag(params param.ParamsCreateTag) (tag model.Tag, err error) {
	tag = model.Tag{
		Name: params.Name,
		Icon: params.Icon,
		Slug: params.Slug,
	}

	result := global.DB.Create(&tag)

	return tag, result.Error
}

func GetTags(params param.ParamsGetTags) (tags []model.Tag, total int64, err error) {
	// Offset(-1)和Limit(-1) 很重要，也是一个小技巧，不加的话会在统计条数后也加上offset和limit，导致查不到条数
	result := global.DB.Debug().Scopes(
		scope.PaginationScope(params.Page, params.Limit),
		scope.GetTagsScope(params),
	).Find(&tags).Offset(-1).Limit(-1).Count(&total)

	return tags, total, result.Error
}

func GetTagsByIDs(ids []int64) (tags []model.Tag, err error) {
	result := global.DB.Find(&tags, ids)

	return tags, result.Error
}

func GetTagByID(id int64) (tag model.Tag, err error) {
	result := global.DB.First(&tag, id)

	return tag, result.Error
}

func UpdateTagByID(id int64, params param.ParamsUpdateTag) (tag model.Tag, err error) {

	if tag, err = GetTagByID(id); err != nil {
		return model.Tag{}, err
	}

	result := global.DB.Model(&tag).Updates(model.Tag{
		Name: params.Name,
		Icon: params.Icon,
		Slug: params.Slug,
	})

	return tag, result.Error
}

func DeleteTagByID(id int64) (tag model.Tag, err error) {

	if tag, err = GetTagByID(id); err != nil {
		return model.Tag{}, err
	}

	result := global.DB.Clauses(clause.Returning{}).Delete(&tag)

	return tag, result.Error
}
