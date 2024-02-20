package initialize

import (
	"fmt"
	"fuxiaochen-api-with-go/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func SetupConfig() (err error) {
	viper.SetConfigName("dev")
	viper.AddConfigPath("./config")

	if err = viper.ReadInConfig(); err != nil {
		return err
	}

	if err = viper.Unmarshal(global.Conf); err != nil {
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("%v: 配置文件发生改变\n", e.Name)

		if err := viper.Unmarshal(global.Conf); err != nil {
			panic(fmt.Errorf("unmarshal 配置文件失败: %v", err))
		}
	})

	return nil
}
