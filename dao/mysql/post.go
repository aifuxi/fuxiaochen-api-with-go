package mysql

import (
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/model"
	"fuxiaochen-api-with-go/model/param"
	"fuxiaochen-api-with-go/model/scope"
	"github.com/samber/lo"
	"gorm.io/gorm/clause"
	"strconv"
)

func CreatePost(params param.ParamsCreatePost) (post model.Post, err error) {
	var categoryID int64

	if params.CategoryID != "" {
		categoryID, err = strconv.ParseInt(params.CategoryID, 10, 64)

		if err != nil {
			return model.Post{}, err
		}
	}

	post = model.Post{
		Title:      params.Title,
		Body:       params.Body,
		Desc:       params.Desc,
		Slug:       params.Slug,
		Cover:      params.Cover,
		Author:     params.Author,
		Type:       params.Type,
		Status:     params.Status,
		Secret:     params.Secret,
		CategoryID: categoryID,
	}

	ids := lo.FilterMap(params.TagIDs, func(v string, _ int) (int64, bool) {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, false
		}

		return i, true
	})

	var tags []model.Tag
	if len(ids) > 0 {
		tags, err = GetTagsByIDs(ids)
		global.L.Debugf("关联的tags %#v\n", tags)
		if err != nil {
			return model.Post{}, err
		}
	}

	result := global.DB.Scopes(
		model.PostUpdateCategoryScope(categoryID),
	).Create(&post)

	if err = global.DB.Model(&post).Association(model.TagsRetrieveKey).Append(tags); err != nil {
		return model.Post{}, err
	}

	return post, result.Error
}

func GetPosts(params param.ParamsGetPosts) (posts []model.Post, total int64, err error) {
	result := global.DB.Scopes(
		scope.PaginationScope(params.Page, params.Limit),
		scope.GetPostsScope(params),
	).Preload(model.TagsRetrieveKey).Find(&posts).Scopes(scope.CountScope).Count(&total)

	return posts, total, result.Error
}

func GetPublishedPostsForSite() (posts []model.Post, total int64, err error) {
	result := global.DB.Scopes(scope.GetPublishedPostsScope).Preload(model.TagsRetrieveKey).Find(&posts).Scopes(scope.CountScope).Count(&total)

	return posts, total, result.Error
}

func GetPostByID(id int64) (post model.Post, err error) {
	result := global.DB.Preload(model.TagsRetrieveKey).First(&post, id)

	return post, result.Error
}

func GetPublishedPostBySlugForSite(slug string) (post model.Post, err error) {
	result := global.DB.Scopes(scope.GetPublishedPostsScope).Preload(model.TagsRetrieveKey).First(&post, "slug = ?", slug)

	return post, result.Error
}

func UpdatePostByID(id int64, params param.ParamsUpdatePost) (post model.Post, err error) {

	if post, err = GetPostByID(id); err != nil {
		return model.Post{}, err
	}

	var categoryID int64

	if params.CategoryID != "" {
		categoryID, err = strconv.ParseInt(params.CategoryID, 10, 64)
		if err != nil {
			return model.Post{}, err
		}
	}

	result := global.DB.Scopes(model.PostUpdateCategoryScope(categoryID)).Model(&post).Updates(model.Post{
		Title:      params.Title,
		Body:       params.Body,
		Desc:       params.Desc,
		Slug:       params.Slug,
		Cover:      params.Cover,
		Author:     params.Author,
		Type:       params.Type,
		Status:     params.Status,
		Secret:     params.Secret,
		CategoryID: categoryID,
	})

	ids := lo.FilterMap(params.TagIDs, func(v string, _ int) (int64, bool) {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, false
		}

		return i, true
	})

	var tags []model.Tag

	if tags, err = GetTagsByIDs(ids); err != nil {
		return model.Post{}, err
	}

	if err = global.DB.Model(&post).Association(model.TagsRetrieveKey).Replace(tags); err != nil {
		return model.Post{}, err
	}

	return post, result.Error
}

func DeletePostByID(id int64) (post model.Post, err error) {

	if post, err = GetPostByID(id); err != nil {
		return model.Post{}, err
	}

	result := global.DB.Clauses(clause.Returning{}).Delete(&post)

	return post, result.Error
}
