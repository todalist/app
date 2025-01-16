package dto

import (
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/models/entity"
)

type TodaTagQuerier struct {
	common.Pager
	Id          *uint   `json:"id"`
	Ids         *[]uint `json:"ids" gorm:"-"`
	Name        *string `json:"name" gorm:"-"`
	AccentColor *string `json:"accentColor"`
	OwnerUserId *uint   `json:"userId"`
}

type TodaTagSaveDTO struct {
	entity.TodaTag `json:"todaTag"`
	PinTop         bool `json:"pinTop"`
}
