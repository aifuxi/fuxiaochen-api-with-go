package model

import (
	"fmt"
	"fuxiaochen-api-with-go/util"
	"gorm.io/gorm"
)

type User struct {
	Name     string `gorm:"size:256;uniqueIndex;comment:用户名称" json:"name,omitempty"`
	Password string `gorm:"size:256;comment:用户密码（已加密）" json:"-"`

	Model
}

func (u *User) BeforeCreate(_ *gorm.DB) error {
	u.ID = util.NewSnowflakeID()
	hashedPwd, err := util.HashPassword([]byte(u.Password))
	u.Password = fmt.Sprintf("%s", hashedPwd)
	return err
}
