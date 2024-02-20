package main

import (
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/initialize"
	"fuxiaochen-api-with-go/model"
)

func main() {
	var err error

	if err = initialize.SetupConfig(); err != nil {
		global.L.Fatalf("初始化配置失败: %v\n", err)
	}

	initialize.SetupLogger()

	if err = initialize.SetupSnowflakeNode(); err != nil {
		global.L.Fatalf("初始化雪花算法节点失败: %v\n", err)
	}

	if err = initialize.SetupDB(); err != nil {
		global.L.Fatalf("数据库连接失败: %v\n", err)
	}

	if err = global.DB.AutoMigrate(&model.User{}, &model.Tag{}, &model.Category{}, &model.Post{}); err != nil {
		global.L.Fatalf("迁移表结构失败: %v\n", err)
	}

	global.DB.Create(&model.User{
		Name:     global.Conf.AppConfig.AdminUsername,
		Password: global.Conf.AppConfig.AdminPassword,
		IsAdmin:  true,
	})

	global.L.Infoln("迁移表结构成功")
}
