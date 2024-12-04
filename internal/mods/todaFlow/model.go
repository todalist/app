package todaFlow

import (
	"dailydo.fe1.xyz/internal/common"
)

type TodaFlow struct {
	common.BaseModel  
	TodaId uint `json:"todaId" gorm:"index"` 
	UserId uint `json:"userId" gorm:"index"` 
	Prev int `json:"prev" ` 
	Next int `json:"next" ` 
	Description string `json:"description" `
}

type TodaFlowQuerier struct {
	common.Pager 
	Id *uint `json:"id"` 
	TodaId *uint `json:"todaId"` 
	UserId *uint `json:"userId"` 
	Prev *int `json:"prev"` 
	Next *int `json:"next"` 
	Description *string `json:"description"` 
}