package model

import (
	"fuxiaochen-api-with-go/util"
	"gorm.io/gorm"
)

const TagsAssociationKey = "Tags"

type Post struct {
	Title  string `gorm:"size:256;uniqueIndex;comment:文章标题" json:"title"`
	Body   string `gorm:"comment:文章内容" json:"body"`
	Desc   string `gorm:"size:256;comment:文章描述" json:"desc"`
	Slug   string `gorm:"size:256;comment:文章slug" json:"slug"`
	Cover  string `gorm:"size:256;comment:文章封面" json:"cover"`
	Author string `gorm:"size:256;comment:文章作者" json:"author"`
	Type   int    `gorm:"comment:文章类型" json:"type"`
	Status int    `gorm:"comment:文章状态" json:"status"`
	Secret string `gorm:"size:256;comment:文章加密的密码" json:"secret"`

	Tags []*Tag `gorm:"many2many:post_tags;" json:"tags"`
	//CategoryID int64  `gorm:"comment:文章分类id" json:"categoryID"`

	Model
}

func (p *Post) BeforeCreate(_ *gorm.DB) (err error) {

	if p.ID == 0 {
		p.ID = util.NewSnowflakeID()
	}

	return
}
