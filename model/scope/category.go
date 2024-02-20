package scope

import (
	"fmt"
	"fuxiaochen-api-with-go/model/param"
	"gorm.io/gorm"
)

func GetCategoriesScope(params param.ParamsGetCategories) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if params.Name != "" {
			arg := fmt.Sprintf("%%%s%%", params.Name)
			return db.Where("name LIKE ?", arg)
		}

		return db
	}
}
