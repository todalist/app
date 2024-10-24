package models

import (
	"time"

	"dailydo.fe1.xyz/internal/common"
)

type Todo struct {
	BaseModel
	Title    string    `json:"title" validate:"required"`
	Content  string    `json:"content"`
	UserID   uint      `json:"userId"`
	Status   int       `json:"status"`
	Deadline time.Time `json:"deadline"`
}

type TodoQuerier struct {
	ID        *uint             `json:"id" uri:"id"`
	UserID    *uint             `json:"userId"`
	TimeRange *common.TimeRange `json:"timeRange"`
	Status    *int              `json:"status"`
	common.Pager
}

const (
	T_STATUS_INIT = iota + 1
	T_STATUS_DONE
)
