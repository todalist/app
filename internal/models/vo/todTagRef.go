package vo

import (
	"github.com/todalist/app/internal/models/entity"
)

type TodaTagRefVO struct {
	entity.TodaTag `gorm:"embedded"`
	TodaId       uint `json:"todaId"`
	TodaTagRefId uint `json:"todaTagId"`
}
