package entity

import (
	"github.com/todalist/app/internal/common"
)

type TodaTag struct {
	common.BaseModel
	Name        string `json:"name" `
	AccentColor string `json:"accentColor" `
	OwnerUserId uint   `json:"userId" gorm:"index"`
}
