package common

import (
	"time"

	"gorm.io/gorm"
)

type TimeRange struct {
	Begin *time.Time
	End   *time.Time
}

func (tr *TimeRange) RangeSql(sql *gorm.DB, field string) *gorm.DB {
	if tr.Begin != nil {
		sql = sql.Where(field+" >= ?", tr.Begin)
	}
	if tr.End != nil {
		sql = sql.Where(field+" <= ?", tr.End)
	}
	return sql
}
