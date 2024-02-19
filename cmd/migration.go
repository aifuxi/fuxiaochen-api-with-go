package main

import (
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/model"
	"log"
)

func main() {

	err := global.SetupDB()
	if err != nil {
		log.Fatalf("数据库连接失败: %v\n", err)
	}

	err = global.DB.AutoMigrate(&model.User{}, &model.Tag{}, &model.Category{}, &model.Post{})
	if err != nil {
		log.Fatalf("迁移表结构失败: %v\n", err)
	}

	log.Println("迁移表结构成功")
}
