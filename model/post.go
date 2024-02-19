package model

import (
	"fuxiaochen-api-with-go/util"
	"gorm.io/gorm"
)

const TagsAssociationKey = "Tags"

type Post struct {
	Title  string `gorm:"size:256;uniqueIndex;comment:文章标题" json:"title,omitempty"`
	Body   string `gorm:"comment:文章内容" json:"body,omitempty"`
	Desc   string `gorm:"size:256;comment:文章描述" json:"desc,omitempty"`
	Slug   string `gorm:"size:256;comment:文章slug" json:"slug,omitempty"`
	Cover  string `gorm:"size:256;comment:文章封面" json:"cover,omitempty"`
	Author string `gorm:"size:256;comment:文章作者" json:"author,omitempty"`
	Type   int    `gorm:"comment:文章类型" json:"type,omitempty"`
	Status int    `gorm:"comment:文章状态" json:"status,omitempty"`
	Secret string `gorm:"size:256;comment:文章加密的密码" json:"secret,omitempty"`

	Tags       []*Tag `gorm:"many2many:post_tags;" json:"tags,omitempty"`
	CategoryID int64  `gorm:"comment:文章分类id" json:"categoryID,omitempty"`

	Model
}

func (p *Post) BeforeCreate(_ *gorm.DB) (err error) {
	p.ID = util.NewSnowflakeID()

	return
}
