package vo

type UserTodaVO struct {
	TodaVO     TodaVO `json:"toda"`
	UserTodaId uint         `json:"id"`
	UserId     uint         `json:"userId"`
}
