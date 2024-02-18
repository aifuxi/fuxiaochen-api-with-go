package main

import (
	"fuxiaochen-api-with-go/initialization"
	"fuxiaochen-api-with-go/model"
	"log"
)

func main() {

	err := initialization.SetupDB()
	if err != nil {
		log.Fatalf("数据库连接失败: %v\n", err)
	}

	err = initialization.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("迁移表结构失败: %v\n", err)
	}

	log.Println("迁移表结构成功")
}
