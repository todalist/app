package models

import (
	"dailydo.fe1.xyz/internal/common"
	"time"
)

type Todo struct {
	BaseModel
	Title    string     `json:"title" validate:"required"`
	UserID   uint       `json:"userId"`
	Status   int        `json:"status"`
	Type     int        `json:"type"`
	CatalogID *uint      `json:"catalogId"`
	Deadline *time.Time `json:"deadline"`
}

type TodoQuerier struct {
	ID        *uint             `json:"id" uri:"id"`
	UserID    uint             `json:"userId"`
	Type      *int              `json:"type"`
	TimeRange *common.TimeRange `json:"timeRange"`
	Status    *int              `json:"status"`
	CatalogID *uint      `json:"catalogId"`
	common.Pager
}

const (
	T_STATUS_INIT = iota + 1
	T_STATUS_DONE
)
