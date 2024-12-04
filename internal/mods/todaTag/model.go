package todaTag

import (
	"dailydo.fe1.xyz/internal/common"
)

type TodaTag struct {
	common.BaseModel  
	Name string `json:"name" ` 
	AccentColor string `json:"accentColor" ` 
	UserId uint `json:"userId" gorm:"index"`
}

type TodaTagQuerier struct {
	common.Pager 
	Id *uint `json:"id"` 
	Name *string `json:"name"` 
	AccentColor *string `json:"accentColor"` 
	UserId *uint `json:"userId"` 
}