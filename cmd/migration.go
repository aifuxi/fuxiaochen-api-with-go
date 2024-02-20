package main

import (
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/initialize"
	"fuxiaochen-api-with-go/model"
	"log"
)

func main() {
	var err error

	err = initialize.SetupConfig()
	if err != nil {
		global.L.Fatalf("初始化配置失败: %v\n", err)
	}

	initialize.SetupLogger()

	err = initialize.SetupSnowflakeNode()
	if err != nil {
		global.L.Fatalf("初始化雪花算法节点失败: %v\n", err)
	}

	err = initialize.SetupDB()
	if err != nil {
		global.L.Fatalf("数据库连接失败: %v\n", err)
	}

	err = global.DB.AutoMigrate(&model.User{}, &model.Tag{}, &model.Category{}, &model.Post{})
	if err != nil {
		global.L.Fatalf("迁移表结构失败: %v\n", err)
	}

	global.DB.Create(&model.User{
		Name:     global.Conf.AppConfig.AdminUsername,
		Password: global.Conf.AppConfig.AdminPassword,
		IsAdmin:  true,
	})

	log.Println("迁移表结构成功")
}
