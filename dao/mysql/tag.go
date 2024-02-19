package mysql

import (
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/model"
	"gorm.io/gorm/clause"
)

func CreateTag(params model.ParamsCreateTag) (tag model.Tag, err error) {
	tag = model.Tag{
		Name: params.Name,
		Icon: params.Icon,
		Slug: params.Slug,
	}

	result := global.DB.Create(&tag)

	return tag, result.Error
}

func GetTags() (tags []model.Tag, err error) {
	result := global.DB.Preload(model.PostsAssociationKey).Find(&tags)

	return tags, result.Error
}

func GetTagsByIDs(ids []int64) (tags []model.Tag, err error) {
	result := global.DB.Find(&tags, ids)

	return tags, result.Error
}

func GetTagByID(id int64) (tag model.Tag, err error) {
	result := global.DB.Preload(model.PostsAssociationKey).First(&tag, id)

	return tag, result.Error
}

func UpdateTagByID(id int64, params model.ParamsUpdateTag) (tag model.Tag, err error) {
	tag, err = GetTagByID(id)
	if err != nil {
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
	tag, err = GetTagByID(id)
	if err != nil {
		return model.Tag{}, err
	}

	result := global.DB.Clauses(clause.Returning{}).Delete(&tag)

	return tag, result.Error
}
