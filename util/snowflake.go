package util

import "fuxiaochen-api-with-go/initialization"

func NewSnowflakeID() int64 {
	return initialization.SnowflakeNode.Generate().Int64()
}
