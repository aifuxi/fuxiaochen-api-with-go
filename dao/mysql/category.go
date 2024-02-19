package mysql

import (
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/model"
	"gorm.io/gorm/clause"
)

func CreateCategory(params model.ParamsCreateCategory) (category model.Category, err error) {
	category = model.Category{
		Name: params.Name,
		Slug: params.Slug,
	}

	result := global.DB.Create(&category)

	return category, result.Error
}

func GetCategories() (categories []model.Category, err error) {
	result := global.DB.Preload(model.PostsAssociationKey).Find(&categories)

	return categories, result.Error
}

func GetCategoriesByIDs(ids []int64) (categories []model.Category, err error) {
	result := global.DB.Find(&categories, ids)

	return categories, result.Error
}

func GetCategoryByID(id int64) (category model.Category, err error) {
	result := global.DB.Preload(model.PostsAssociationKey).First(&category, id)

	return category, result.Error
}

func UpdateCategoryByID(id int64, params model.ParamsUpdateCategory) (category model.Category, err error) {
	category, err = GetCategoryByID(id)
	if err != nil {
		return model.Category{}, err
	}
	result := global.DB.Model(&category).Updates(model.Category{
		Name: params.Name,
		Slug: params.Slug,
	})

	return category, result.Error
}

func DeleteCategoryByID(id int64) (category model.Category, err error) {
	category, err = GetCategoryByID(id)
	if err != nil {
		return model.Category{}, err
	}

	result := global.DB.Clauses(clause.Returning{}).Delete(&category)

	return category, result.Error
}
