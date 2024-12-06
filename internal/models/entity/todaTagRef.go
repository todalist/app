package entity

import (
	"github.com/todalist/app/internal/common"
)

type TodaTagRef struct {
	common.BaseModel
	TodaId    uint `json:"todaId" gorm:"index"`
	TodaTagId uint `json:"todaTagId" gorm:"index"`
}
