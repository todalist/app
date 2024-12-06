package vo

import "github.com/todalist/app/internal/models/entity"


type UserTodaTagVO struct {
	Tag entity.TodaTag `json:"tag"`
	UserTodaTagId uint `json:"userTodaTagId"`
	UserId uint `json:"userId"`
}
