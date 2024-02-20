package model

import (
	"fuxiaochen-api-with-go/util"
	"gorm.io/gorm"
)

type Category struct {
	Name string `gorm:"size:256;uniqueIndex;comment:分类名称" json:"name"`
	Slug string `gorm:"size:256;uniqueIndex;comment:分类slug" json:"slug"`

	Posts []Post `json:"posts"`

	Model
}

func (c *Category) BeforeCreate(_ *gorm.DB) (err error) {
	if c.ID == 0 {
		c.ID = util.NewSnowflakeID()
	}

	return
}
