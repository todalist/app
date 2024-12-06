package entity

import (
	"github.com/todalist/app/internal/common"
)

type UserTodaTag struct {
	common.BaseModel
	UserId    uint `json:"userId" gorm:"index"`
	TodaTagId uint `json:"todaTagId" gorm:"index"`
}
