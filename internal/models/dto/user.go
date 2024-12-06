package dto

import (
	"github.com/todalist/app/internal/common"
)

type UserQuerier struct {
	common.Pager
	Id       *uint   `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Avatar   *string `json:"avatar"`
}
