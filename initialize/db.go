package initialize

import (
	"fmt"
	"fuxiaochen-api-with-go/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() error {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.Conf.MySQLConfig.Username,
		global.Conf.MySQLConfig.Password,
		global.Conf.MySQLConfig.Host,
		global.Conf.MySQLConfig.Port,
		global.Conf.MySQLConfig.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	global.DB = db

	return err
}
