package todaTagRef

import (
	"github.com/todalist/app/internal/common"
)

type TodaTagRef struct {
	common.BaseModel  
	TodaId uint `json:"todaId" gorm:"index"` 
	TodaTagId uint `json:"todaTagId" gorm:"index"`
}

type TodaTagRefQuerier struct {
	common.Pager 
	Id *uint `json:"id"` 
	TodaId *uint `json:"todaId"` 
	TodaTagId *uint `json:"todaTagId"` 
}