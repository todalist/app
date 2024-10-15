package models

type Todo struct {
	BaseModel
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int    `json:"status"`
}

const (
	TODO_STATUS_INIT = iota + 1
	TODO_STATUS_DONE
)
