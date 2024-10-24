package models

import "dailydo.fe1.xyz/internal/common"

type Collection struct {
	BaseModel
	UserID         uint   `json:"user_id"`
	ParentID       uint   `json:"parent_id"`
	TodoCount      int    `json:"todo_count"`
	CompletedCount int    `json:"completed_count"`
	Title          string `json:"title"`
	Status         string `json:"status"`
}

type CollectionQuerier struct {
	ID       *uint   `json:"id"`
	UserID   *uint   `json:"user_id"`
	ParentID *uint   `json:"parent_id"`
	Title    *string `json:"title"`
	Status   *string `json:"status"`
	common.Pager
}

const (
	COLLECTION_STATUS_NORMAL = iota + 1
	COLLECTION_STATUS_DISABLED
)
