package migration

import (
	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/models/entity"
	"log"
)

func MustMigration() {
	if err := globals.DB.AutoMigrate(
		&entity.User{},
		&entity.Toda{},
		&entity.TodaFlow{},
		&entity.TodaTag{},
		&entity.TodaTagRef{},
		&entity.UserToda{},
		&entity.UserTodaTag{},
	); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}
