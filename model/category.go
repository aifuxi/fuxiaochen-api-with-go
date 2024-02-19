package model

import (
	"fuxiaochen-api-with-go/util"
	"gorm.io/gorm"
)

type Category struct {
	Name string `gorm:"size:256;uniqueIndex;comment:分类名称" json:"name,omitempty"`
	Slug string `gorm:"size:256;comment:分类slug" json:"slug,omitempty"`

	Posts []Post `json:"posts"`

	Model
}

func (c *Category) BeforeCreate(_ *gorm.DB) (err error) {
	c.ID = util.NewSnowflakeID()

	return
}
