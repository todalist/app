package userTodaTag

import (
	"github.com/todalist/app/internal/common"
)

type UserTodaTag struct {
	common.BaseModel
	UserId    uint `json:"userId" gorm:"index"`
	TodaTagId uint `json:"todaTagId" gorm:"index"`
}

type UserTodaTagQuerier struct {
	common.Pager
	Id        *uint `json:"id"`
	UserId    *uint `json:"userId"`
	TodaTagId *uint `json:"todaTagId"`
}
