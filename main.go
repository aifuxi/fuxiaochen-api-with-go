package main

import (
	"fmt"
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/initialize"
	"fuxiaochen-api-with-go/router"
	"log"
)

func main() {
	var err error

	err = initialize.SetupConfig()
	if err != nil {
		log.Fatalf("初始化配置失败: %v\n", err)
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

	r := router.NewRouter()
	err = r.Run(fmt.Sprintf("%s:%d", global.Conf.AppConfig.Host, global.Conf.AppConfig.Port))
	if err != nil {
		global.L.Fatalf("启动服务器失败：%v\n", err)
	}
}
