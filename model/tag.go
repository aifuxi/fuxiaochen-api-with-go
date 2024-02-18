package model

import (
	"fuxiaochen-api-with-go/util"
	"gorm.io/gorm"
)

type Tag struct {
	Name string `gorm:"size:256;uniqueIndex;comment:标签名称" json:"name,omitempty"`
	Icon string `gorm:"comment:标签图标" json:"icon,omitempty"`

	Model
}

func (t *Tag) BeforeCreate(_ *gorm.DB) (err error) {
	t.ID = util.NewSnowflakeID()

	return
}
