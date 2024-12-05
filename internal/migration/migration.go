package migration

import (
	"log"

	"github.com/todalist/app/internal/globals"
)

func MustMigration() {
	if err := globals.DB.AutoMigrate(
	// &models.User{},
	// &models.Todo{},
	// &models.TodoCatalog{},
	); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}
