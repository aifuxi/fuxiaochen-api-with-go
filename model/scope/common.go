package scope

import "gorm.io/gorm"

func PaginationScope(page int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * limit

		return db.Offset(offset).Limit(limit)
	}
}

func CountScope(db *gorm.DB) *gorm.DB {
	// Offset(-1)和Limit(-1) 很重要，也是一个小技巧，不加的话会在统计条数后也加上offset和limit，导致查不到条数
	return db.Offset(-1).Limit(-1)
}
