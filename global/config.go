package global

import "time"

type AppConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
}

type JWTConfig struct {
	Secret         string        `mapstructure:"secret"`
	ExpireDuration time.Duration `mapstructure:"expire_duration"`
}

type Config struct {
	AppConfig   `mapstructure:"app"`
	MySQLConfig `mapstructure:"mysql"`
	JWTConfig   `mapstructure:"jwt"`
}
