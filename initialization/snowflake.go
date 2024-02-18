package initialization

import "github.com/bwmarrin/snowflake"

var SnowflakeNode *snowflake.Node

func SetupSnowflakeNode() error {
	var err error

	SnowflakeNode, err = snowflake.NewNode(1)

	return err
}
