package toda

import (
	"time"

	"github.com/todalist/app/internal/common"
)

type Toda struct {
	common.BaseModel
	Title       string     `json:"title" `
	Description string     `json:"description" `
	UserId      uint       `json:"userId" gorm:"index"`
	Priority    int        `json:"priority" `
	Deadline    *time.Time `json:"deadline" `
	Status      int        `json:"status" `
	Estimate    int        `json:"estimate" `
	Elapsed     int        `json:"elapsed" `
}

type TodaQuerier struct {
	common.Pager
	Id          *uint      `json:"id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	UserId      *uint      `json:"userId"`
	Priority    *int       `json:"priority"`
	Deadline    *time.Time `json:"deadline"`
	Status      *int       `json:"status"`
	Estimate    *int       `json:"estimate"`
	Elapsed     *int       `json:"elapsed"`
}

const (
	TodaStatusTodo = iota + 1
	TodaStatusFinished
	TodaStatusArchived
)

const (
	TodaPriorityLow = iota + 1
	TodaPriorityMedium
	TodaPriorityHigh
)
