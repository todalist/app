package dto

import (
	"time"

	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/models/entity"
)

type TodaQuerier struct {
	common.Pager
	Id          *uint      `json:"id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	OwnerUserId *uint      `json:"userId"`
	Priority    *int       `json:"priority"`
	Deadline    *time.Time `json:"deadline"`
	Status      *int       `json:"status"`
	Estimate    *int       `json:"estimate"`
	Elapsed     *int       `json:"elapsed"`
}

type TodaSaveDTO struct {
	entity.Toda `json:"toda"`
	Tags        []uint `json:"tags"`
}
