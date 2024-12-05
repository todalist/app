package userToda

import (
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/mods/toda"
)

type UserToda struct {
	common.BaseModel
	UserId uint `json:"userId" gorm:"index"`
	TodaId uint `json:"todaId" gorm:"index"`
}

type UserTodaQuerier struct {
	common.Pager
	Id     *uint `json:"id"`
	UserId *uint `json:"userId"`
	TodaId *uint `json:"todaId"`
}

type ListUserTodaQuerier struct {
	toda.TodaQuerier
}

type UserTodaVO struct {
	TodaVO     *toda.TodaVO `json:"toda"`
	UserTodaId uint         `json:"id"`
	UserId     uint         `json:"userId"`
}
