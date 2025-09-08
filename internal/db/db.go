package db

import (
	"log"
	"time"

	"api-customer-merchant/internal/config"
	"api-customer-merchant/internal/db/models"
	//"api-customer-merchant/internal/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect initializes GORM DB for dev environment. It will only AutoMigrate if cfg.DevAutoMigrate==true
func Connect(cfg *config.Config) error {
	if cfg.DB_DSN == "" {
		return nil // allow running without DB in small dev scenarios
	}
	var gormLogger logger.Interface
	if cfg.Env == "development" {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default.LogMode(logger.Silent)
	}

	gormCfg := &gorm.Config{Logger: gormLogger}
	db, err := gorm.Open(postgres.Open(cfg.DB_DSN), gormCfg)
	if err != nil {
		//logger.Default.Error("failed to open db", zapError(err))
		log.Fatal("failed to open db", zapError(err))
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	DB = db

	if cfg.DevAutoMigrate {
		// Run safe AutoMigrate for dev only. Do not use in prod.
		// List only a small set of models used during local dev to avoid surprises.
		// import your models package in this file when enabling.
		 err = db.AutoMigrate( &models.MerchantApplication{})
		if err != nil {
			log.Println("aut migrate skipped or failed:", err)
		}
	}

	return nil
}

// Helper to wrap errors for zap without importing zap in many files
func zapError(err error) any { return map[string]any{"error": err.Error()} }