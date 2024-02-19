package global

import (
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var DB *gorm.DB

var SnowflakeNode *snowflake.Node

var Conf = new(Config)

var Logger *zap.Logger
var L *zap.SugaredLogger
