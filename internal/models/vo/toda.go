package vo

import (
	"github.com/todalist/app/internal/models/entity"
)

type TodaVO struct {
	*entity.Toda `gorm:"embedded"`
	Tags []*TodaTagRefVO `json:"tags" gorm:"-"`
}
