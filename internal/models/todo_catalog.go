package models

import "dailydo.fe1.xyz/internal/common"

type TodoCatalog struct {
	BaseModel
	UserID             uint   `json:"userId" gorm:"index"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	ParentID           *uint  `json:"parentId" gorm:"index"`
	PinToSideBar       bool   `json:"pinToSidebar"`
	AccentColor        string `json:"accentColor"`
	ItemsCount         int    `json:"itemsCount"`
	ItemsFinishedCount int    `json:"itemsFinishedCount"`
}

type TodoCatalogQuerier struct {
	ID           *uint   `json:"id"`
	UserID       uint    `json:"userId"`
	Title        *string `json:"title"`
	ParentID     *uint   `json:"parentId"`
	PinToSideBar *bool    `json:"pinToSidebar"`
	common.Pager
}
