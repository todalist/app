package dto

import (
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/models/entity"
)

type TodaTagQuerier struct {
	common.Pager
	Id          *uint   `json:"id"`
	Name        *string `json:"name"`
	AccentColor *string `json:"accentColor"`
	OwnerUserId *uint   `json:"userId"`
}

type TodaTagSaveDTO struct {
	entity.TodaTag `json:"todaTag"`
	PinTop bool `json:"pinTop"`
}
