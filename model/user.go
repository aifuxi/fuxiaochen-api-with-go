package model

type User struct {
	Name     string `gorm:"size:256;uniqueIndex;comment:用户名称" json:"name,omitempty"`
	Password string `gorm:"size:256;comment:用户密码（已加密）" json:"password,omitempty"`

	Model
}
