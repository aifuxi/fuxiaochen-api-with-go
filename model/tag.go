package model

import (
	"fuxiaochen-api-with-go/util"
	"gorm.io/gorm"
)

const PostsAssociationKey = "Posts"

type Tag struct {
	Name string `gorm:"size:256;uniqueIndex;comment:标签名称" json:"name"`
	Icon string `gorm:"comment:标签图标" json:"icon"`
	Slug string `gorm:"size:256;comment:标签slug" json:"slug"`

	Posts *[]Post `gorm:"many2many:post_tags;" json:"posts"`

	Model
}

func (t *Tag) BeforeCreate(_ *gorm.DB) (err error) {
	if t.ID == 0 {
		t.ID = util.NewSnowflakeID()
	}
	return
}
