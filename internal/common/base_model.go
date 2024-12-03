package common

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        uint           `gorm:"primaryKey" json:"id,omitempty" query:"id" uri:"id"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt *time.Time     `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
