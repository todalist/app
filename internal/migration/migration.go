package migration

import (
	"log"
	"dailydo.fe1.xyz/internal/globals"
	"dailydo.fe1.xyz/internal/models"
)

func MustMigration() {
	if err := globals.DB.AutoMigrate(
		&models.User{},
		&models.Todo{},
	); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}
