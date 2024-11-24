package models

import (
	"dailydo.fe1.xyz/internal/common"
	"time"
)

type Todo struct {
	BaseModel
	Title    string     `json:"title" validate:"required"`
	Content  string     `json:"content"`
	UserID   uint       `json:"userId"`
	Status   int        `json:"status"`
	Type     int        `json:"type"`
	ParentID *uint      `json:"parentId"`
	Deadline *time.Time `json:"deadline"`
}

type TodoQuerier struct {
	ID        *uint             `json:"id" uri:"id"`
	UserID    *uint             `json:"userId"`
	Type      *int              `json:"type"`
	TimeRange *common.TimeRange `json:"timeRange"`
	Status    *int              `json:"status"`
	ParentID  *uint             `json:"parentId"`
	common.Pager
}

const (
	T_STATUS_INIT = iota + 1
	T_STATUS_DONE
)

const (
	T_TYPE_TODO = iota + 1
	T_TYPE_FOLDER
)
