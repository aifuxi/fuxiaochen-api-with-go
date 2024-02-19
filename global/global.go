package global

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

var DB *gorm.DB

var SnowflakeNode *snowflake.Node

var Conf = new(Config)
