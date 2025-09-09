package db

import (
	"log"
	"os"
	"time"

	//"api-customer-merchant/internal/config"
	"api-customer-merchant/internal/db/models"

	//"api-customer-merchant/internal/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"gorm.io/gorm/logger"
)

// Connect initializes GORM DB for dev environment. It will only AutoMigrate if cfg.DevAutoMigrate==true
var DB *gorm.DB

func Connect() {
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        log.Fatal("DB_DSN environment variable not set")
    }

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Configure connection pool
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to get SQL DB: %v", err)
    }
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetConnMaxLifetime(30 * time.Minute)

    log.Println("Database connected successfully")
}

func AutoMigrate() {
    // Run AutoMigrate with all models
	if err := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		log.Fatal("Failed to enable uuid-ossp extension:", err)
	}

    err := DB.AutoMigrate(
        //&models.User{},
        &models.MerchantApplication{},
        // &models.Product{},
        // &models.Cart{},
        // &models.Order{},
        // &models.OrderItem{},
        // &models.CartItem{},
        // &models.Category{},
        // &models.Inventory{},
        // &models.Promotion{},
        // &models.ReturnRequest{},
        // &models.Payout{},
    )
    if err != nil {
        log.Fatalf("Failed to auto-migrate: %v", err)
    }

    // Get the underlying SQL database connection
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to get SQL DB: %v", err)
    }

    // Close all connections to clear cached plans
    if err := sqlDB.Close(); err != nil {
        log.Printf("Failed to close connections: %v", err)
    }

    // Reconnect to ensure fresh connections
    dsn := os.Getenv("DB_DSN")
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to reconnect to database: %v", err)
    }

    // Reconfigure connection pool
    sqlDB, err = DB.DB()
    if err != nil {
        log.Fatalf("Failed to get SQL DB after reconnect: %v", err)
    }
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetConnMaxLifetime(30 * time.Minute)

    log.Println("Database migrated and reconnected successfully")
}