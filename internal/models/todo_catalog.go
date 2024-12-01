package models

import "dailydo.fe1.xyz/internal/common"

type TodoCatalog struct {
	BaseModel
	UserID             uint   `json:"userId" gorm:"index"`
	Title              string `json:"title"`
	ParentID           *uint  `json:"parentId" gorm:"index"`
	AccentColor        string `json:"accentColor"`
	ItemsCount         int    `json:"itemsCount"`
	ItemsFinishedCount int    `json:"itemsFinishedCount"`
}

type TodoCatalogQuerier struct {
	ID       *uint   `json:"id"`
	UserID   uint   `json:"userId"`
	Title    *string `json:"title"`
	ParentID *uint   `json:"parentId"`
	common.Pager
}