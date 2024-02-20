package mysql

import (
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/model"
	"github.com/samber/lo"
	"gorm.io/gorm/clause"
	"strconv"
)

func CreatePost(params model.ParamsCreatePost) (post model.Post, err error) {
	//categoryID, parseErr := strconv.ParseInt(params.CategoryID, 10, 64)
	//if parseErr != nil {
	//	return model.Post{}, parseErr
	//}
	post = model.Post{
		Title:  params.Title,
		Body:   params.Body,
		Desc:   params.Desc,
		Slug:   params.Slug,
		Cover:  params.Cover,
		Author: params.Author,
		Type:   params.Type,
		Status: params.Status,
		Secret: params.Secret,
		//CategoryID: categoryID,
	}

	global.L.Debugf("创建文章参数 %#v\n", params)
	ids := lo.FilterMap(params.TagIDs, func(v string, _ int) (int64, bool) {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, false
		}

		return i, true
	})

	var tags []model.Tag
	tags, err = GetTagsByIDs(ids)
	global.L.Debugf("关联的tags %#v\n", tags)
	if err != nil {
		return model.Post{}, err
	}

	result := global.DB.Create(&post)
	err = global.DB.Model(&post).Association(model.TagsAssociationKey).Append(tags)
	if err != nil {
		return model.Post{}, err
	}

	return post, result.Error
}

func GetPosts() (posts []model.Post, err error) {
	result := global.DB.Preload(model.TagsAssociationKey).Find(&posts)

	return posts, result.Error
}

func GetPostsByIDs(ids []int64) (posts []model.Post, err error) {
	result := global.DB.Find(&posts, ids)

	return posts, result.Error
}

func GetPostByID(id int64) (post model.Post, err error) {
	result := global.DB.Preload(model.TagsAssociationKey).First(&post, id)

	return post, result.Error
}

func UpdatePostByID(id int64, params model.ParamsUpdatePost) (post model.Post, err error) {
	post, err = GetPostByID(id)
	if err != nil {
		return model.Post{}, err
	}
	result := global.DB.Model(&post).Updates(model.Post{
		Title:  params.Title,
		Body:   params.Body,
		Desc:   params.Desc,
		Slug:   params.Slug,
		Cover:  params.Cover,
		Author: params.Author,
		Type:   params.Type,
		Status: params.Status,
		Secret: params.Secret,
	})

	ids := lo.FilterMap(params.TagIDs, func(v string, _ int) (int64, bool) {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, false
		}

		return i, true
	})

	var tags []model.Tag
	tags, err = GetTagsByIDs(ids)
	if err != nil {
		return model.Post{}, err
	}

	err = global.DB.Model(&post).Association(model.TagsAssociationKey).Replace(tags)
	if err != nil {
		return model.Post{}, err
	}

	return post, result.Error
}

func DeletePostByID(id int64) (post model.Post, err error) {
	post, err = GetPostByID(id)
	if err != nil {
		return model.Post{}, err
	}

	result := global.DB.Clauses(clause.Returning{}).Delete(&post)

	return post, result.Error
}
