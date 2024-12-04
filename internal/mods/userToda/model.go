package userToda

import (
	"dailydo.fe1.xyz/internal/common"
)

type UserToda struct {
	common.BaseModel  
	UserId uint `json:"userId" gorm:"index"` 
	TodaId uint `json:"todaId" gorm:"index"`
}

type UserTodaQuerier struct {
	common.Pager 
	UserId uint `json:"userId"` 
	TodaId uint `json:"todaId"` 
}