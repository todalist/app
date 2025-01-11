package vo

import "github.com/todalist/app/internal/models/entity"

type UserTodaTagVO struct {
	Tag           *entity.TodaTag `json:"tag" gorm:"embedded"`
	UserTodaTagId uint            `json:"userTodaTagId" gorm:"embedded;embeddedPrefix:utt"`
	UserId        uint            `json:"userId" gorm:"embedded;embeddedPrefix:utt"`
	PinTop        bool            `json:"pinTop" gorm:"embedded;embeddedPrefix:utt"`
}
