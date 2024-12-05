package userToda

import (
	"github.com/todalist/app/internal/common"
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
