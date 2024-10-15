package migration

import (
	"log"
	"dailydo.fe1.xyz/internal/globals"
	"dailydo.fe1.xyz/internal/models"
)

func MustMigration() {
	if err := globals.DB.AutoMigrate(
		&models.User{},
	); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}
