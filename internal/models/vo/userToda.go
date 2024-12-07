package vo

type UserTodaVO struct {
	TodaVO `gorm:"embedded"`
	UserTodaId uint   `json:"id"`
	UserId     uint   `json:"userId"`
}
