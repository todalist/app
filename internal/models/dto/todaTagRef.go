package dto

import (
	"github.com/todalist/app/internal/common"
)

type TodaTagRefQuerier struct {
	common.Pager
	Id        *uint `json:"id"`
	TodaId    *uint `json:"todaId"`
	TodaTagId *uint `json:"todaTagId"`
}
