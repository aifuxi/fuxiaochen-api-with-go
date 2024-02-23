package mysql

import (
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/model"
	"fuxiaochen-api-with-go/model/param"
	"fuxiaochen-api-with-go/model/scope"
	"gorm.io/gorm/clause"
)

func CreateCategory(params param.ParamsCreateCategory) (category model.Category, err error) {
	category = model.Category{
		Name: params.Name,
		Slug: params.Slug,
	}

	result := global.DB.Create(&category)

	return category, result.Error
}

func GetCategories(params param.ParamsGetCategories) (categories []model.Category, total int64, err error) {
	result := global.DB.Scopes(
		scope.PaginationScope(params.Page, params.Limit),
		scope.GetCategoriesScope(params),
	).Find(&categories).Scopes(scope.CountScope).Count(&total)

	return categories, total, result.Error
}

func GetCategoryByID(id int64) (category model.Category, err error) {
	result := global.DB.Preload(model.PostsRetrieveKey).First(&category, id)

	return category, result.Error
}

func UpdateCategoryByID(id int64, params param.ParamsUpdateCategory) (category model.Category, err error) {

	if category, err = GetCategoryByID(id); err != nil {
		return model.Category{}, err
	}

	result := global.DB.Model(&category).Updates(model.Category{
		Name: params.Name,
		Slug: params.Slug,
	})

	return category, result.Error
}

func DeleteCategoryByID(id int64) (category model.Category, err error) {

	if category, err = GetCategoryByID(id); err != nil {
		return model.Category{}, err
	}

	result := global.DB.Clauses(clause.Returning{}).Delete(&category)

	return category, result.Error
}
