package globals

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
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

type ContextDBKey struct{}

func ContextDB(ctx context.Context, tx *gorm.DB) context.Context {
	v := ctx.Value(ContextDBKey{})
	if v == nil {
		if tx == nil {
			tx = DB
		}
		return context.WithValue(ctx, ContextDBKey{}, tx)
	}
	return ctx
}

func GetFromContext(ctx context.Context) *gorm.DB {
	v := ctx.Value(ContextDBKey{})
	if v == nil {
		return DB
	}
	db, ok := v.(*gorm.DB)
	if ok {
		return db
	}
	LOG.Fatal("failed to get database from context", zap.Any("value", v))
	return nil
}
