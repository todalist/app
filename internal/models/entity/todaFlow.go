package entity

import (
	"github.com/todalist/app/internal/common"
)

type TodaFlow struct {
	common.BaseModel
	TodaId      uint   `json:"todaId" gorm:"index"`
	UserId      uint   `json:"userId" gorm:"index"`
	Prev        int    `json:"prev" `
	Next        int    `json:"next" `
	Description string `json:"description" `
}
