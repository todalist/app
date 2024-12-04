package userTodaTag

import (
	"dailydo.fe1.xyz/internal/common"
)

type UserTodaTag struct {
	common.BaseModel  
	UserId uint `json:"userId" gorm:"index"` 
	TodaTagId uint `json:"todaTagId" gorm:"index"`
}

type UserTodaTagQuerier struct {
	common.Pager 
	UserId uint `json:"userId"` 
	TodaTagId uint `json:"todaTagId"` 
}