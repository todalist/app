package toda

import (
	"time"
	"dailydo.fe1.xyz/internal/common"
)

type Toda struct {
	common.BaseModel  
	Title string `json:"title" ` 
	Description string `json:"description" ` 
	UserId uint `json:"userId" gorm:"index"` 
	Priority int `json:"priority" ` 
	Deadline time.Time `json:"deadline" ` 
	Status int `json:"status" ` 
	Estimate int `json:"estimate" ` 
	Elapsed int `json:"elapsed" `
}

type TodaQuerier struct {
	common.Pager 
	Title string `json:"title"` 
	Description string `json:"description"` 
	UserId uint `json:"userId"` 
	Priority int `json:"priority"` 
	Deadline time.Time `json:"deadline"` 
	Status int `json:"status"` 
	Estimate int `json:"estimate"` 
	Elapsed int `json:"elapsed"` 
}