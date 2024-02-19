package util

import "fuxiaochen-api-with-go/global"

func NewSnowflakeID() int64 {
	return global.SnowflakeNode.Generate().Int64()
}
