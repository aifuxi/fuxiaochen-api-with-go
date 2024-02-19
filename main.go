package main

import (
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/router"
	"log"
)

func main() {
	var err error

	err = global.SetupDB()
	if err != nil {
		log.Fatalf("数据库连接失败: %v\n", err)
	}

	err = global.SetupSnowflakeNode()
	if err != nil {
		log.Fatalf("初始化雪花算法节点失败: %v\n", err)
	}

	r := router.NewRouter()
	err = r.Run("127.0.0.1:6122")
	if err != nil {
		log.Fatalf("启动服务器失败：%v\n", err)
	}
}
