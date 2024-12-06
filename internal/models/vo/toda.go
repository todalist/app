package vo

import (
	"github.com/todalist/app/internal/models/entity"
)

type TodaVO struct {
	entity.Toda
	Tags []*TodaTagVO `json:"tags"`
}
