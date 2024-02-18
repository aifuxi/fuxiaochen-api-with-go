package main

import (
	"fuxiaochen-api-with-go/initialization"
	"fuxiaochen-api-with-go/router"
	"log"
)

func main() {
	var err error

	err = initialization.SetupDB()
	if err != nil {
		log.Fatalf("数据库连接失败: %v\n", err)
	}

	r := router.NewRouter()
	err = r.Run("127.0.0.1:6122")
	if err != nil {
		log.Fatalf("启动服务器失败：%v\n", err)
	}
}
