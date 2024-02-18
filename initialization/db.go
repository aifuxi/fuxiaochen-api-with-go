package initialization

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() error {
	var err error

	dsn := "root:123456@tcp(127.0.0.1:3306)/fuxiaochen?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return err
}
