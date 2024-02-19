package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	// id 为 int64，超出了js里面的Number类型的最大值，所以这里转化成字符串
	ID        int64          `gorm:"primaryKey;comment=雪花算法生成的ID" json:"id,string"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
