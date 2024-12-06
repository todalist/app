package dto

import (
	"github.com/todalist/app/internal/common"
)

type UserTodaQuerier struct {
	common.Pager
	Id     *uint `json:"id"`
	UserId *uint `json:"userId"`
	TodaId *uint `json:"todaId"`
}

type ListUserTodaQuerier struct {
	TodaQuerier
	TodaTagId *uint `json:"todaTagId" cond:"-"`
	UserId    *uint `json:"userId" cond:"-"`
}
