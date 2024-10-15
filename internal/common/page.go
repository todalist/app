package common

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

const (
	MAXIMUM_PAGE_SIZE = 64
	DEFAULT_PAGE_SIZE = 10
)

type Pager struct {
	PageNum  int `json:"pageNum" validate:"required,min=1"`
	PageSize int `json:"pageSize" validate:"required,min=1,max=64"`
}

func CalcPageOffset(pager *Pager) (offset int) {
	offset = int(pager.PageNum-1) * pager.PageSize
	return
}

// gorm paginate support
func PaginateWithDefaultOrder(c fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pager := new(Pager)
		c.Bind().Query(pager)
		return paginate(db, pager).Order("created_at desc").Order("id desc")
	}
}

func paginate(db *gorm.DB, pager *Pager) *gorm.DB {
	if pager.PageNum <= 0 {
		pager.PageNum = 1
	}
	if pager.PageSize <= 0 || pager.PageSize >= MAXIMUM_PAGE_SIZE {
		pager.PageSize = DEFAULT_PAGE_SIZE
	}
	return db.Offset(CalcPageOffset(pager)).Limit(pager.PageSize)
}

func Paginate1(sql *gorm.DB, pager *Pager) *gorm.DB {
	return sql.Offset(CalcPageOffset(pager)).Limit(pager.PageSize)
}

// gorm paginate support
func Paginate(c fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pager := new(Pager)
		c.Bind().Query(pager)
		return paginate(db, pager)
	}
}
