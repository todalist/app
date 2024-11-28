package models

type TodoFolder struct {
	BaseModel
	TodoID             uint   `gorm:"unique" json:"todoId"`
	AccentColor        string `json:"accentColor"`
	ItemsCount         int    `json:"itemsCount"`
	ItemsFinishedCount int    `json:"itemsFinishedCount"`
}
