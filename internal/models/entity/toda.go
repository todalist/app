package entity

import (
	"github.com/todalist/app/internal/common"
	"time"
)

type Toda struct {
	common.BaseModel
	Title       string     `json:"title" `
	Description string     `json:"description" `
	OwnerUserId uint       `json:"userId" gorm:"index"`
	Priority    int        `json:"priority" `
	Deadline    *time.Time `json:"deadline" `
	Status      int        `json:"status" `
	Estimate    int        `json:"estimate" `
	Elapsed     int        `json:"elapsed" `
	CompletedAt *time.Time `json:"completedAt" `
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
