package scope

import "gorm.io/gorm"

func PaginationScope(page int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * limit

		return db.Offset(offset).Limit(limit)
	}
}
