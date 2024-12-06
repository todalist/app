package entity

import (
	"github.com/todalist/app/internal/common"
)

type UserToda struct {
	common.BaseModel
	UserId uint `json:"userId" gorm:"index"`
	TodaId uint `json:"todaId" gorm:"index"`
}
