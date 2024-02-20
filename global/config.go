package global

import "time"

type AppConfig struct {
	Host          string `mapstructure:"host"`
	Port          int    `mapstructure:"port"`
	Mode          string `mapstructure:"mode"`
	AdminUsername string `mapstructure:"admin_username"`
	AdminPassword string `mapstructure:"admin_password"`
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

type LogConfig struct {
	Filepath   string `mapstructure:"filepath"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type Config struct {
	AppConfig   `mapstructure:"app"`
	MySQLConfig `mapstructure:"mysql"`
	JWTConfig   `mapstructure:"jwt"`
	LogConfig   `mapstructure:"log"`
}
