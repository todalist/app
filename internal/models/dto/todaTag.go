package dto

import (
	"github.com/todalist/app/internal/common"
)

type TodaTagQuerier struct {
	common.Pager 
	Id           *uint   `json:"id"`
	Name         *string `json:"name"`
	AccentColor  *string `json:"accentColor"`
	OwnerUserId  *uint   `json:"userId"`
}
