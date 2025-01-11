package dto

import (
	"github.com/todalist/app/internal/common"
)

type TodaFlowQuerier struct {
	common.Pager
	Id          *uint   `json:"id"`
	TodaId      *uint   `json:"todaId"`
	UserId      *uint   `json:"userId"`
	Prev        *int    `json:"prev"`
	Next        *int    `json:"next"`
	Description *string `json:"description"`
}
