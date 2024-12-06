package vo

type UserTodaVO struct {
	Toda       TodaVO `json:"toda"`
	UserTodaId uint   `json:"id"`
	UserId     uint   `json:"userId"`
}
