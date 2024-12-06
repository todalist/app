package dto

import (
	"github.com/todalist/app/internal/common"
	"time"
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
