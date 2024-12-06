package vo

import (
	"github.com/todalist/app/internal/models/entity"
)

type TodaTagRefVO struct {
	entity.TodaTag
	TodaId       uint `json:"todaId"`
	TodaTagRefId uint `json:"todaTagId"`
}
