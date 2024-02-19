package initialize

import (
	"fuxiaochen-api-with-go/global"
	"github.com/bwmarrin/snowflake"
)

func SetupSnowflakeNode() error {
	node, err := snowflake.NewNode(1)
	global.SnowflakeNode = node
	return err
}
