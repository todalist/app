package dto

import (
	"github.com/todalist/app/internal/common"
)

type UserTodaTagQuerier struct {
	common.Pager
	Id        *uint `json:"id"`
	UserId    *uint `json:"userId"`
	TodaTagId *uint `json:"todaTagId"`
	PinTop    *bool `json:"pinTop"`
}

type ListUserTodaTagQuerier struct {
	TodaTagQuerier
	UserId *uint `json:"userId" cond:"-"`
}
