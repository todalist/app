package todaTagRef

import (
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/mods/todaTag"
)

type TodaTagRef struct {
	common.BaseModel
	TodaId    uint `json:"todaId" gorm:"index"`
	TodaTagId uint `json:"todaTagId" gorm:"index"`
}

type TodaTagRefQuerier struct {
	common.Pager
	Id        *uint `json:"id"`
	TodaId    *uint `json:"todaId"`
	TodaTagId *uint `json:"todaTagId"`
}

type TodaTagVO struct {
	todaTag.TodaTag
	TodaId       uint `json:"todaId"`
	TodaTagRefId uint `json:"todaTagId"`
}
