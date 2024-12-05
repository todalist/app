package user

import (
	"github.com/todalist/app/internal/common"
)

type User struct {
	common.BaseModel
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password" `
	Avatar   string `json:"avatar" `
}

type UserQuerier struct {
	common.Pager
	Id       *uint   `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Avatar   *string `json:"avatar"`
}
