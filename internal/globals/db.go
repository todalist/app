package globals

import (
	"fmt"
	"log"
	"time"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)


func InitDatabase() {
	if CONF == nil {
		log.Fatalf("config must be present")
	}
	dbConfig := CONF.DB
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Database,
		dbConfig.Port,
	)
	LOG.Debug("database config", zap.String("dsn", dsn))
	var loggerLevel logger.LogLevel
	if IsProd() {
		loggerLevel = logger.Silent
	} else {
		loggerLevel = logger.Info
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(loggerLevel),
	})
	if err != nil {
		LOG.Fatal("connection database error.", zap.String("dsn", dsn), zap.Error(err))
	}
	DB = db
}
