package vo

type UserTodaVO struct {
	*TodaVO `gorm:"embedded" json:"toda"`
	UserTodaId uint   `json:"userTodaId"`
	UserId     uint   `json:"userId"`
}
