package model

import (
	"fmt"
	"fuxiaochen-api-with-go/util"
	"gorm.io/gorm"
)

type User struct {
	Name     string `gorm:"size:256;uniqueIndex;comment:用户名称" json:"name"`
	Password string `gorm:"size:256;comment:用户密码（已加密）" json:"-"`
	IsAdmin  bool   `gorm:"comment:是否为管理员" json:"isAdmin"`

	Model
}

func (u *User) BeforeCreate(_ *gorm.DB) error {
	if u.ID == 0 {
		u.ID = util.NewSnowflakeID()
	}
	hashedPwd, err := util.HashPassword([]byte(u.Password))
	u.Password = fmt.Sprintf("%s", hashedPwd)
	return err
}
