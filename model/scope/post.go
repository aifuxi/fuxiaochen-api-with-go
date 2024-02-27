package scope

import (
	"fmt"
	"fuxiaochen-api-with-go/model/param"
	"gorm.io/gorm"
)

func GetPostsScope(params param.ParamsGetPosts) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if params.Title != "" {
			arg := fmt.Sprintf("%%%s%%", params.Title)
			db = db.Where("title LIKE ?", arg)
		}

		if params.Author != "" {
			arg := fmt.Sprintf("%%%s%%", params.Author)
			db = db.Where("author LIKE ?", arg)
		}

		if params.CategoryID != "" {
			db = db.Where("category_id = ?", params.CategoryID)
		}

		if len(params.Types) > 0 {
			db = db.Where("type IN ?", params.Types)
		}

		if len(params.Statuses) > 0 {
			db = db.Where("status IN ?", params.Statuses)
		}

		return db
	}
}

func GetPublishedPostsScope(db *gorm.DB) *gorm.DB {

	db = db.Where("type = ?", 1)

	return db
}
