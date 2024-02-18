package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        int64          `gorm:"primaryKey;comment=雪花算法生成的ID" json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
