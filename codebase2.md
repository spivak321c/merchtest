# Codebase Analysis: merchtest
Generated: 2025-09-12 07:13:17
---

## 📂 Project Structure
```tree
📁 merchtest
├── cmd/
│   └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── auth_handler.go
│   │   │   └── merchant_handlers.go
│   │   └── routes.go
│   ├── config/
│   │   ├── config.go
│   │   └── redis_config.go
│   ├── db/
│   │   ├── models/
│   │   │   ├── cart.go
│   │   │   ├── cart_item.go
│   │   │   ├── category.go
│   │   │   ├── inventory.go
│   │   │   ├── merchant.bak.txt
│   │   │   ├── merchants.go
│   │   │   ├── order.go
│   │   │   ├── order_item.go
│   │   │   ├── payment.go
│   │   │   ├── payout.go
│   │   │   ├── product.go
│   │   │   ├── promotion.go
│   │   │   ├── return_request.go
│   │   │   └── user.go
│   │   ├── repositories/
│   │   │   ├── cart_item_repository.go
│   │   │   ├── cart_repository.go
│   │   │   ├── category_repositry.go
│   │   │   ├── inventory_repository.go
│   │   │   ├── merchant_repository.go
│   │   │   ├── order_item_repository.go
│   │   │   ├── order_repository.go
│   │   │   ├── payment_repository.go
│   │   │   ├── payout_repository.go
│   │   │   ├── product_repository.go
│   │   │   ├── promotion_repository.go
│   │   │   └── return_repository.go
│   │   └── db.go
│   ├── domain/
│   │   ├── audit/
│   │   │   └── audit_service.go
│   │   ├── cart/
│   │   │   └── cart_service.go
│   │   ├── merchant/
│   │   │   └── merchant_service.go
│   │   ├── notifications/
│   │   │   └── notifcation_service.go
│   │   ├── order/
│   │   │   └── order_service.go
│   │   ├── payment/
│   │   │   └── payment_service.go
│   │   ├── payout/
│   │   │   └── payout_service.go
│   │   ├── pricing/
│   │   │   └── pricing_service.go
│   │   ├── product/
│   │   │   └── product_service.go
│   │   ├── promotions/
│   │   │   └── promotion_service.go
│   │   └── returns/
│   │       └── return_service.go
│   ├── logger/
│   │   └── logger.go
│   ├── middleware/
│   │   ├── auth.go
│   │   └── rate_limit.go
│   ├── tests/
│   │   └── unit/
│   │       ├── product_service_test.go
│   │       └── test_service.go
│   └── utils/
│       ├── blacklist.go
│       └── redis.go
└── .env
```
---

## 📄 File Contents
### .env
- Size: 0.74 KB
- Lines: 10
- Last Modified: 2025-09-08 22:26:37

<xaiArtifact artifact_id="519711b0-be3f-419b-b74f-cbd2063a8fe4" artifact_version_id="dd953004-5a73-498f-87ac-c6be0e401bf5" title=".env" contentType="text/plain">
```plain
DB_DSN=postgresql://neondb_owner:npg_CcwoeLb6V1XH@ep-wild-haze-adu0bdvq-pooler.c-2.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require
JWT_SECRET=9072a74677e95918103b4993b96ef0455995408610c82c3cb3433f718d4838e0
GOOGLE_CLIENT_ID=269870327937-9qlv0sl9lt374slkcqicpus76tnk9cle.apps.googleusercontent.com
GOOGLE_CLIENT_SECRET=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
REDIS_ADDR=testing-zrth-cbdi-356640.leapcell.cloud:6379
REDIS_PASS=Ae00000HR66rpYGCp6yeR5xEA0hCzBidoSQqoRfbSr78m/kXbVTTMxpwptseLynFm2zhNKt
REDIS_DB=0
REDIS_URL=rediss://default:Ae00000HR66rpYGCp6yeR5xEA0hCzBidoSQqoRfbSr78m/kXbVTTMxpwptseLynFm2zhNKt@testing-zrth-cbdi-356640.leapcell.cloud:6379
BASE_URL=https://8080-spivak321c-testing-j0cbwp1c30e.ws-eu121.gitpod.io
PORT=8080

```
</xaiArtifact>

---
### cmd/main.go
- Size: 0.97 KB
- Lines: 49
- Last Modified: 2025-09-08 22:32:08

<xaiArtifact artifact_id="3de8c5f1-5b94-47b5-bc4c-16b2e988d489" artifact_version_id="8dcb9007-012f-42d9-a38d-615cc0379602" title="cmd/main.go" contentType="text/go">
```go
package main

import (
	"fmt"
	"log"
	"os"

	merchant "api-customer-merchant/internal/api"
	"api-customer-merchant/internal/config"
	"api-customer-merchant/internal/db"

	//"api-customer-merchant/internal/logger"
	"api-customer-merchant/internal/utils"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	 if err := godotenv.Load(); err != nil {
		 	log.Fatal("Error loading .env file")
		 }
		 conf := config.Load()
		 utils.InitRedis(conf)
		 secret := os.Getenv("JWT_SECRET")
		 if secret == "" {
		 	log.Fatal("JWT_SECRET not set")
		 }
	// Connect to database and migrate
	db.Connect()
	db.AutoMigrate()
	r := gin.Default()
	r.Use(gin.Recovery())

	merchant.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run on 0.0.0.0:port for Railway compatibility
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Example app listening on port %s", port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("API failed: %v", err)
	}
}
```
</xaiArtifact>

---
### internal/utils/blacklist.go
- Size: 0.93 KB
- Lines: 31
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="96c411f9-6c79-4633-a4bd-7e38b99080ca" artifact_version_id="c4b07368-43d6-47c4-abe5-b98dc725e2cf" title="internal/utils/blacklist.go" contentType="text/go">
```go
package utils

import (
    "context"
    "errors"
    "log"
    "time"
    //"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// Add adds a token to the Redis blacklist with an expiration
func Add(token string) error {
    if RedisClient == nil {
        log.Println("RedisClient is nil, cannot add token to blacklist")
        return errors.New("redis client not initialized")
    }
    // Set token in Redis with a 24-hour expiration
    return RedisClient.Set(ctx, "blacklist:"+token, "true", 24*time.Hour).Err()
}

// IsBlacklisted checks if a token is in the Redis blacklist
func IsBlacklisted(token string) bool {
    if RedisClient == nil {
        log.Println("RedisClient is nil, skipping blacklist check")
        return false // Fallback to allow operation if Redis is unavailable
    }
    _, err := RedisClient.Get(ctx, "blacklist:"+token).Result() // Line 22
    return err == nil // Token exists in Redis if no error
}
```
</xaiArtifact>

---
### internal/utils/redis.go
- Size: 1.41 KB
- Lines: 56
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="6fb80b5d-3a0b-4eaa-83dd-e9114c4db759" artifact_version_id="2f76d51e-966d-491d-b749-c3866bf91125" title="internal/utils/redis.go" contentType="text/go">
```go
package utils

import (
	"api-customer-merchant/internal/config"
	"context"
	"crypto/tls"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis(conf *config.Config) {
    
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     conf.RedisAddr,
        Password: conf.RedisPass,
        DB:       conf.RedisDB,
        TLSConfig: &tls.Config{
            MinVersion: tls.VersionTLS12, 
            InsecureSkipVerify: true, // Use with caution, only for testing
        },
    })

    ctx := context.Background()
    if err := RedisClient.Ping(ctx).Err(); err != nil {
       log.Printf("Failed to connect to Redis: %v, continuing without caching", err)
        RedisClient = nil // Fallback to avoid crashes
    } else {
        log.Println("Connected to Redis successfully")
    }
}

// Helper to get cached value or fetch and cache
func GetOrSetCache(ctx context.Context, key string, ttl time.Duration, fetch func() (any, error)) (any, error) {
    val, err := RedisClient.Get(ctx, key).Result()
    if err == nil {
        return val, nil // Deserialize if needed (e.g., JSON)
    }
    if err != redis.Nil {
        return nil, err
    }

    data, err := fetch()
    if err != nil {
        return nil, err
    }

    // Serialize if complex (e.g., JSON marshal)
    if err := RedisClient.Set(ctx, key, data, ttl).Err(); err != nil {
        return nil, err
    }
    return data, nil
}
```
</xaiArtifact>

---
### internal/config/config.go
- Size: 1.30 KB
- Lines: 45
- Last Modified: 2025-09-08 22:28:46

<xaiArtifact artifact_id="54494747-e1e6-41c0-a80c-4637f4316648" artifact_version_id="ea09facb-c51b-4ab0-bfcc-7b7e28e88bdb" title="internal/config/config.go" contentType="text/go">
```go
package config
/*
import (
	"time"

	"github.com/spf13/viper"
)

// Config contains all runtime configuration for dev environment
type Config struct {
	Env                string
	DB_DSN             string
	RedisAddr          string
	RedisPass          string
	RedisDB            int
	JWTSecret          string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
	DevAutoMigrate     bool
}

// Load reads environment using viper and returns Config
func Load() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetDefault("ENV", "development")
	viper.SetDefault("REDIS_DB", 0)
	viper.SetDefault("ACCESS_TOKEN_EXPIRY_MIN", 15) // minutes
	viper.SetDefault("REFRESH_TOKEN_EXPIRY_DAYS", 7)
	viper.SetDefault("DEV_AUTO_MIGRATE", true)

	cfg := &Config{
		Env:                viper.GetString("ENV"),
		DB_DSN:             viper.GetString("DB_DSN"),
		RedisAddr:          viper.GetString("REDIS_ADDR"),
		RedisPass:          viper.GetString("REDIS_PASS"),
		RedisDB:            viper.GetInt("REDIS_DB"),
		JWTSecret:          viper.GetString("JWT_SECRET"),
		AccessTokenExpiry:  time.Minute * time.Duration(viper.GetInt("ACCESS_TOKEN_EXPIRY_MIN")),
		RefreshTokenExpiry: 24 * time.Hour * time.Duration(viper.GetInt("REFRESH_TOKEN_EXPIRY_DAYS")),
		DevAutoMigrate:     viper.GetBool("DEV_AUTO_MIGRATE"),
	}

	return cfg, nil
}
*/
```
</xaiArtifact>

---
### internal/config/redis_config.go
- Size: 1.77 KB
- Lines: 70
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="3db5a036-0966-470e-8b55-48c4e61b1f58" artifact_version_id="ef8e3a4c-5106-405a-95be-f255f9825dcc" title="internal/config/redis_config.go" contentType="text/go">
```go
package config

import (
	//"log"
	//"net/url"
	"os"
	"strconv"
	//"time"
)

type Config struct {
	RedisAddr string
	RedisPass string
	RedisDB   int
	// Other fields...
}

 func Load() *Config {
 	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
 	// AccessTokenExp, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXP_MINUTES")) // e.g., 15
 	// RefreshTokenExp, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXP_DAYS"))  // e.g., 7
 	 // AccessTokenExp = time.Duration(AccessTokenExp) * time.Minute
 	 // RefreshTokenExp = time.Duration(RefreshTokenExp) * 24 * time.Hour
 	return &Config{
 		RedisAddr: os.Getenv("REDIS_ADDR"), // e.g., "localhost:6379"
 		RedisPass: os.Getenv("REDIS_PASS"),
 		RedisDB:   redisDB, // Default 0
 		// ...
 	}
 }

// func Load() *Config {
//     redisURL := os.Getenv("REDIS_URL")
//     if redisURL == "" {
//         // Fallback to individual environment variables
//         redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
//         return &Config{
//             RedisAddr: os.Getenv("REDIS_ADDR"),
//             RedisPass: os.Getenv("REDIS_PASS"),
//             RedisDB:   redisDB,
//         }
//     }

//     // Parse Redis URL
//     u, err := url.Parse(redisURL)
//     if err != nil {
//         log.Fatalf("Failed to parse REDIS_URL: %v", err)
//     }

//     redisAddr := u.Host // Includes host:port
//     redisPass := ""
//     redisDB := 0

//     // Extract password if present
//     if u.User != nil {
//         redisPass, _ = u.User.Password()
//     }

//     // Extract database if present
//     if u.Path != "" {
//         dbStr := u.Path[1:] // Remove leading "/"
//         redisDB, _ = strconv.Atoi(dbStr)
//     }

//     return &Config{
//         RedisAddr: redisAddr,
//         RedisPass: redisPass,
//         RedisDB:   redisDB,
//     }
// }
```
</xaiArtifact>

---
### internal/logger/logger.go
- Size: 0.29 KB
- Lines: 20
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="0b5d305a-c6b4-4eb4-bf49-78b05f849e65" artifact_version_id="7a01d824-4d0c-49f8-8d08-79cabe90ba0e" title="internal/logger/logger.go" contentType="text/go">
```go
package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func Init(dev bool) error {
	if dev {
		l, err := zap.NewDevelopment()
		if err != nil { return err }
		Log = l
		return nil
	}
	l, err := zap.NewProduction()
	if err != nil { return err }
	Log = l
	return nil
}
```
</xaiArtifact>

---
### internal/middleware/rate_limit.go
- Size: 0.45 KB
- Lines: 20
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="91892676-efdd-475a-8e65-b40c8e2fd0a0" artifact_version_id="fa763479-e268-4f49-ab6b-1e2d21e3c4d8" title="internal/middleware/rate_limit.go" contentType="text/go">
```go
package middleware

import (
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
)

func RateLimitMiddleware() gin.HandlerFunc {
    limiter := rate.NewLimiter(rate.Every(time.Minute), 100) // 100/min
    return func(c *gin.Context) {
        if !limiter.Allow() {
            c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
            c.Abort()
            return
        }
        c.Next()
    }
}
```
</xaiArtifact>

---
### internal/middleware/auth.go
- Size: 1.66 KB
- Lines: 61
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="a976b9f7-f988-45d6-95b0-443f4a7189ed" artifact_version_id="f82718fe-35a0-49c2-b333-70373ce89209" title="internal/middleware/auth.go" contentType="text/go">
```go
package middleware

import (
	"net/http"
	"os"
	"strings"

	"api-customer-merchant/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(entityType string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if utils.IsBlacklisted(tokenString) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is blacklisted"})
            c.Abort()
            return
        }
        key := os.Getenv("JWT_SECRET")

        secret := []byte(key) // Load from env
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrSignatureInvalid
            }
            return secret, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || claims["entityType"] != entityType {
            c.JSON(http.StatusForbidden, gin.H{"error": "Invalid entity type"})
            c.Abort()
            return
        }

        id := claims["id"].(float64)
        switch entityType {
case "user":
            c.Set("userID", uint(id))
        case "merchant":
            c.Set("merchantID", uint(id))
        }
        c.Next()
    }
}
```
</xaiArtifact>

---
### internal/db/db.go
- Size: 2.47 KB
- Lines: 97
- Last Modified: 2025-09-12 02:25:28

<xaiArtifact artifact_id="62131410-c04a-47c5-a18a-460d60fd0d6b" artifact_version_id="2bc723f0-d598-4997-9281-93c2fc0c3735" title="internal/db/db.go" contentType="text/go">
```go
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
         &models.Category{},
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
```
</xaiArtifact>

---
### internal/api/routes.go
- Size: 2.15 KB
- Lines: 63
- Last Modified: 2025-09-08 22:24:53

<xaiArtifact artifact_id="8acc87a7-fff8-4aab-9b38-6b3d3e378ea0" artifact_version_id="76ef50e4-8cb2-4318-bc2b-20b7f284505e" title="internal/api/routes.go" contentType="text/go">
```go
package api

    import (
       "api-customer-merchant/internal/api/handlers"
       "api-customer-merchant/internal/db/repositories"
       //"api-customer-merchant/internal/middleware"
        "api-customer-merchant/internal/domain/merchant"

        "github.com/gin-gonic/gin"
    )

/*
    func RegisterRoutes(r *gin.Engine) {
        merchant := r.Group("/merchant")
        {
           //userRepo := repositories.NewUserRepository()
            // merchantRepo := repositories.NewMerchantRepository()
            appRepo := repositories.NewMerchantApplicationRepository()
            repo := repositories.NewMerchantRepository()
            service := merchant.NewMerchantService(appRepo, repo)
            h := handlers.NewMerchantHandler(service)

            // authHandler := handlers.NewAuthHandler(merchantRepo)
            // merchant.POST("/submitApplication", authHandler.Register)
            // merchant.POST("/login", authHandler.Login)
            //merchant.GET("/auth/google", authHandler.GoogleAuth)
            //merchant.GET("/auth/google/callback", authHandler.GoogleCallback)
            merchant.POST("/apply",  h.Apply)
            merchant.GET("/application/:id",  h.GetApplication)
            
            
            // Merchant account access (once approved by admin via Express API)
            //protected.POST("/me",  h.GetMyMerchant)
    
           

            protected := merchant.Group("/")
            protected.Use(middleware.AuthMiddleware("merchant"))
            protected.POST("/me",  h.GetMyMerchant)
            //protected.POST("/logout", authHandler.Logout)
            
        }
    }
*/

func RegisterRoutes(r *gin.Engine) {
    appRepo := repositories.NewMerchantApplicationRepository()
    repo := repositories.NewMerchantRepository()
    service := merchant.NewMerchantService(appRepo, repo)
    h := handlers.NewMerchantHandler(service)
    
    
    mg := r.Group("/merchant")
    {
    // Application submission and status
    mg.POST("/apply",  h.Apply)
    mg.GET("/application/:id",  h.GetApplication)
    
    
    // Merchant account access (once approved by admin via Express API)
    mg.GET("/me",  h.GetMyMerchant)
    }
    }
```
</xaiArtifact>

---
### internal/domain/promotions/promotion_service.go
- Size: 0.82 KB
- Lines: 31
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="78d8f58b-db5d-492e-ad06-f313a34c60f5" artifact_version_id="a513b522-bbc6-43b2-8596-c1b6a09acd43" title="internal/domain/promotions/promotion_service.go" contentType="text/go">
```go
package promotions

import (
    "errors"
    "time"
    "api-customer-merchant/internal/db/models"
    "api-customer-merchant/internal/db/repositories"
)

type PromotionService struct {
    repo *repositories.PromotionRepository
}

func NewPromotionService(repo *repositories.PromotionRepository) *PromotionService {
    return &PromotionService{repo: repo}
}

func (s *PromotionService) CreatePromotion(promo *models.Promotion) error {
    if promo.Code == "" {
        return errors.New("code required")
    }
    return s.repo.Create(promo)
}

func (s *PromotionService) ValidateCode(code string) (*models.Promotion, error) {
    promo, err := s.repo.FindByCode(code)
    if err != nil || time.Now().Before(promo.ValidFrom) || time.Now().After(promo.ValidTo) {
        return nil, errors.New("invalid promo")
    }
    return promo, nil
}
```
</xaiArtifact>

---
### internal/domain/pricing/pricing_service.go
- Size: 1.03 KB
- Lines: 40
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="91210056-955d-4db7-8af1-ca20802331f6" artifact_version_id="af704a50-ee53-4706-8965-50aee71f1dea" title="internal/domain/pricing/pricing_service.go" contentType="text/go">
```go
package pricing

import (
    "errors"
    "api-customer-merchant/internal/db/models"
)

type PricingService struct {
    // Repos if needed
}

func NewPricingService() *PricingService {
    return &PricingService{}
}

func (s *PricingService) CalculateShipping(cart *models.Cart, address string) (float64, error) {
    // Integrate with shipping API (e.g., UPS); placeholder
    if address == "" {
        return 0, errors.New("address required")
    }
    return 10.00, nil // Flat rate per vendor count
}

func (s *PricingService) CalculateTax(cart *models.Cart, country string) (float64, error) {
    // Tax API or rules; placeholder
    var subtotal float64
    for _, item := range cart.CartItems {
        subtotal += float64(item.Quantity) * item.Product.Price
    }
    rate := 0.1 // 10% VAT for luxury
    return subtotal * rate, nil
}

func (s *PricingService) ApplyPromotion(cart *models.Cart, code string) (float64, error) {
    // Validate coupon; placeholder
    if code == "" {
        return 0, nil
    }
    return 5.00, nil // Discount
}
```
</xaiArtifact>

---
### internal/domain/product/product_service.go
- Size: 4.59 KB
- Lines: 167
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="18f893ad-a512-4f0a-aa6b-28de4ecd4dd1" artifact_version_id="cbec4cac-c437-4ee6-b892-36f89032f738" title="internal/domain/product/product_service.go" contentType="text/go">
```go
package product

import (
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"
	"api-customer-merchant/internal/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type ProductService struct {
	productRepo  *repositories.ProductRepository
	inventoryRepo *repositories.InventoryRepository
}

func NewProductService(productRepo *repositories.ProductRepository, inventoryRepo *repositories.InventoryRepository) *ProductService {
	return &ProductService{
		productRepo:  productRepo,
		inventoryRepo: inventoryRepo,
	}
}

// GetProductByID retrieves a product by its ID
func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
    if id == 0 {
        return nil, errors.New("invalid product ID")
    }

    ctx := context.Background()
    key := fmt.Sprintf("product:%d", id)
    data, err := utils.GetOrSetCache(ctx, key, 10*time.Minute, func() (any, error) {
        product, err := s.productRepo.FindByID(id)
        if err != nil {
            return nil, err
        }
        return json.Marshal(product) // Serialize
    })
    if err != nil {
        return nil, err
    }

    var product models.Product
    if err := json.Unmarshal([]byte(data.(string)), &product); err != nil {
        return nil, err
    }
    return &product, nil
}

// GetProductBySKU retrieves a product by its SKU
func (s *ProductService) GetProductBySKU(sku string) (*models.Product, error) {
	if strings.TrimSpace(sku) == "" {
		return nil, errors.New("SKU cannot be empty")
	}
	return s.productRepo.FindBySKU(sku)
}

// SearchProductsByName searches for products by name (case-insensitive)
func (s *ProductService) SearchProductsByName(name string) ([]models.Product, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("search name cannot be empty")
	}
	return s.productRepo.SearchByName(name)
}

// GetProductsByCategory retrieves products in a category
// In ProductService
func (s *ProductService) GetProductsByCategory(categoryID uint, limit, offset int) ([]models.Product, error) {
    if categoryID == 0 {
        return nil, errors.New("invalid category ID")
    }
    return s.productRepo.FindByCategoryWithPagination(categoryID, limit, offset)
}

// CreateProduct creates a new product for a merchant
func (s *ProductService) CreateProduct(product *models.Product, merchantID uint) error {
	if product == nil {
		return errors.New("product cannot be nil")
	}
	if merchantID == 0 {
		return errors.New("invalid merchant ID")
	}
	if strings.TrimSpace(product.Name) == "" {
		return errors.New("product name cannot be empty")
	}
	if strings.TrimSpace(product.SKU) == "" {
		return errors.New("SKU cannot be empty")
	}
	if product.Price <= 0 {
		return errors.New("price must be positive")
	}
	if product.CategoryID == 0 {
		return errors.New("category ID must be set")
	}

	// Check if SKU is unique
	if _, err := s.productRepo.FindBySKU(product.SKU); err == nil {
		return errors.New("SKU already exists")
	}

	product.MerchantID = merchantID
	return s.productRepo.Create(product)
}

// UpdateProduct updates an existing product (merchant only)
func (s *ProductService) UpdateProduct(product *models.Product, merchantID uint) error {
	if product == nil || product.ID == 0 {
		return errors.New("invalid product or product ID")
	}
	if merchantID == 0 {
		return errors.New("invalid merchant ID")
	}
	if strings.TrimSpace(product.Name) == "" {
		return errors.New("product name cannot be empty")
	}
	if strings.TrimSpace(product.SKU) == "" {
		return errors.New("SKU cannot be empty")
	}
	if product.Price <= 0 {
		return errors.New("price must be positive")
	}
	if product.CategoryID == 0 {
		return errors.New("category ID must be set")
	}

	// Verify product belongs to merchant
	existing, err := s.productRepo.FindByID(product.ID)
	if err != nil {
		return err
	}
	if existing.MerchantID != merchantID {
		return errors.New("product does not belong to merchant")
	}

	// Check if SKU is unique (excluding current product)
	if p, err := s.productRepo.FindBySKU(product.SKU); err == nil && p.ID != product.ID {
		return errors.New("SKU already exists")
	}

	return s.productRepo.Update(product)
}

// DeleteProduct deletes a product (merchant only)
func (s *ProductService) DeleteProduct(id uint, merchantID uint) error {
	if id == 0 {
		return errors.New("invalid product ID")
	}
	if merchantID == 0 {
		return errors.New("invalid merchant ID")
	}

	// Verify product belongs to merchant
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return err
	}
	if product.MerchantID != merchantID {
		return errors.New("product does not belong to merchant")
	}

	return s.productRepo.Delete(id)
}



```
</xaiArtifact>

---
### internal/domain/cart/cart_service.go
- Size: 4.86 KB
- Lines: 174
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="06da0d96-7b66-4559-9e64-d4eb3ab343b0" artifact_version_id="8121be50-f9e0-48f5-8f4c-3e92ed224ec8" title="internal/domain/cart/cart_service.go" contentType="text/go">
```go
package cart

import (
	"errors"
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"

	"gorm.io/gorm"
)

type CartService struct {
	cartRepo      *repositories.CartRepository
	cartItemRepo  *repositories.CartItemRepository
	productRepo   *repositories.ProductRepository
	inventoryRepo *repositories.InventoryRepository
}

func NewCartService(cartRepo *repositories.CartRepository, cartItemRepo *repositories.CartItemRepository, productRepo *repositories.ProductRepository, inventoryRepo *repositories.InventoryRepository) *CartService {
	return &CartService{
		cartRepo:      cartRepo,
		cartItemRepo:  cartItemRepo,
		productRepo:   productRepo,
		inventoryRepo: inventoryRepo,
	}
}

// GetActiveCart retrieves or creates an active cart for a user
func (s *CartService) GetActiveCart(userID uint) (*models.Cart, error) {
	if userID == 0 {
		return nil, errors.New("invalid user ID")
	}

	// Try to find an active cart
	cart, err := s.cartRepo.FindActiveCart(userID)
	if err == nil {
		return cart, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Create a new active cart if none exists
	cart = &models.Cart{
		UserID: userID,
		Status: models.CartStatusActive,
	}
	if err := s.cartRepo.Create(cart); err != nil {
		return nil, err
	}
	return s.cartRepo.FindByID(cart.ID)
}

// AddItemToCart adds a product to the user's active cart
func (s *CartService) AddItemToCart(userID, productID uint, quantity int) (*models.Cart, error) {
    if userID == 0 {
        return nil, errors.New("invalid user ID")
    }
    if productID == 0 {
        return nil, errors.New("invalid product ID")
    }
    if quantity <= 0 {
        return nil, errors.New("quantity must be positive")
    }

    // Get active cart
    cart, err := s.GetActiveCart(userID)
    if err != nil {
        return nil, err
    }

    // Check if product exists
    product, err := s.productRepo.FindByID(productID)
    if err != nil {
        return nil, errors.New("product not found")
    }
    if product.MerchantID == 0 {
        return nil, errors.New("invalid merchant for product")
    }

    // Check stock availability
    inventory, err := s.inventoryRepo.FindByProductID(productID)
    if err != nil {
        return nil, errors.New("inventory not found")
    }
    if inventory.StockQuantity < quantity {
        return nil, errors.New("insufficient stock")
    }

    // Check if product is already in cart
    cartItem, err := s.cartItemRepo.FindByProductIDAndCartID(productID, cart.ID)
    if err == nil {
        // Update quantity if item exists
        newQuantity := cartItem.Quantity + quantity
        if inventory.StockQuantity < newQuantity {
            return nil, errors.New("insufficient stock for updated quantity")
        }
        if err := s.cartItemRepo.UpdateQuantity(cartItem.ID, newQuantity); err != nil {
            return nil, err
        }
    } else if errors.Is(err, gorm.ErrRecordNotFound) {
        // Create new cart item
        cartItem = &models.CartItem{
            CartID:     cart.ID,
            ProductID:  productID,
            Quantity:   quantity,
            MerchantID: product.MerchantID,
        }
        if err := s.cartItemRepo.Create(cartItem); err != nil {
            return nil, err
        }
    } else {
        return nil, err
    }

    return s.cartRepo.FindByID(cart.ID)
}
// UpdateCartItemQuantity updates the quantity of a cart item
func (s *CartService) UpdateCartItemQuantity(cartItemID uint, quantity int) (*models.Cart, error) {
	if cartItemID == 0 {
		return nil, errors.New("invalid cart item ID")
	}
	if quantity <= 0 {
		return nil, errors.New("quantity must be positive")
	}

	// Get cart item
	cartItem, err := s.cartItemRepo.FindByID(cartItemID)
	if err != nil {
		return nil, errors.New("cart item not found")
	}

	// Check stock availability
	inventory, err := s.inventoryRepo.FindByProductID(cartItem.ProductID)
	if err != nil {
		return nil, errors.New("inventory not found")
	}
	if inventory.StockQuantity < quantity {
		return nil, errors.New("insufficient stock")
	}

	// Update quantity
	if err := s.cartItemRepo.UpdateQuantity(cartItemID, quantity); err != nil {
		return nil, err
	}

	return s.cartRepo.FindByID(cartItem.CartID)
}

// RemoveCartItem removes an item from the cart
func (s *CartService) RemoveCartItem(cartItemID uint) (*models.Cart, error) {
	if cartItemID == 0 {
		return nil, errors.New("invalid cart item ID")
	}

	// Get cart item to find cart ID
	cartItem, err := s.cartItemRepo.FindByID(cartItemID)
	if err != nil {
		return nil, errors.New("cart item not found")
	}

	// Delete cart item
	if err := s.cartItemRepo.Delete(cartItemID); err != nil {
		return nil, err
	}

	return s.cartRepo.FindByID(cartItem.CartID)
}

func (s *CartService) GetCartItemByID(cartItemID uint) (*models.CartItem, error) {
	if cartItemID == 0 {
		return nil, errors.New("invalid cart item ID")
	}
	return s.cartItemRepo.FindByID(cartItemID)
}
```
</xaiArtifact>

---
### internal/domain/returns/return_service.go
- Size: 1.22 KB
- Lines: 39
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="2138dc53-b9d9-4c96-825a-96b19ddc55b0" artifact_version_id="6338111e-a107-4e5f-8a0d-be024a037fb4" title="internal/domain/returns/return_service.go" contentType="text/go">
```go
package returns

import (
    "errors"
    "api-customer-merchant/internal/db/models"
    "api-customer-merchant/internal/db/repositories"
)

type ReturnService struct {
    repo          *repositories.ReturnRepository
    orderItemRepo *repositories.OrderItemRepository
}

func NewReturnService(repo *repositories.ReturnRepository, orderItemRepo *repositories.OrderItemRepository) *ReturnService {
    return &ReturnService{repo: repo, orderItemRepo: orderItemRepo}
}

func (s *ReturnService) CreateReturn(orderItemID uint, reason string, userID uint) error {
    item, err := s.orderItemRepo.FindByID(orderItemID)
    if err != nil || item.Order.UserID != userID {
        return errors.New("invalid item")
    }
    req := &models.ReturnRequest{OrderItemID: orderItemID, Reason: reason, Status: "Pending"}
    return s.repo.Create(req)
}

func (s *ReturnService) ApproveReturn(returnID uint, merchantID uint) error {
    req, err := s.repo.FindByID(returnID)
    if err != nil {
        return err
    }
    item, _ := s.orderItemRepo.FindByID(req.OrderItemID)
    if item.MerchantID != merchantID {
        return errors.New("permission denied")
    }
    req.Status = "Approved"
    // Trigger refund, restock
    return s.repo.Update(req)
}
```
</xaiArtifact>

---
### internal/domain/payout/payout_service.go
- Size: 1.93 KB
- Lines: 73
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="c25c6191-8a58-4362-aedc-1c15ce75ac40" artifact_version_id="9a26aff1-d4be-4651-8091-2ed9f27b9674" title="internal/domain/payout/payout_service.go" contentType="text/go">
```go
package payout


import (
	"errors"
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"
)

type PayoutService struct {
	payoutRepo *repositories.PayoutRepository
}

func NewPayoutService(payoutRepo *repositories.PayoutRepository) *PayoutService {
	return &PayoutService{
		payoutRepo: payoutRepo,
	}
}

// CreatePayout creates a payout for a merchant
func (s *PayoutService) CreatePayout(merchantID uint, amount float64) (*models.Payout, error) {
    if merchantID == 0 {
        return nil, errors.New("invalid merchant ID")
    }
    if amount <= 0 {
        return nil, errors.New("amount must be positive")
    }

    // Placeholder: Assume GetMerchantBalance fetches earnings
    earnings := GetMerchantBalance(merchantID) // Implement this
    if earnings < amount {
        return nil, errors.New("insufficient balance")
    }

    // Simulate payout processing (placeholder for Stripe)
    payout := &models.Payout{
        MerchantID: merchantID,
        Amount:     amount,
        Status:     models.PayoutStatusPending,
    }
    if err := s.payoutRepo.Create(payout); err != nil {
        return nil, err
    }

    // Simulate successful payout
    payout.Status = models.PayoutStatusCompleted
    if err := s.payoutRepo.Update(payout); err != nil {
        return nil, err
    }

    return s.payoutRepo.FindByID(payout.ID)
}

func GetMerchantBalance(merchantID uint) float64 {
	//panic("unimplemented")
    return 1000.00
}

// GetPayoutByID retrieves a payout by ID
func (s *PayoutService) GetPayoutByID(id uint) (*models.Payout, error) {
	if id == 0 {
		return nil, errors.New("invalid payout ID")
	}
	return s.payoutRepo.FindByID(id)
}

// GetPayoutsByMerchantID retrieves all payouts for a merchant
func (s *PayoutService) GetPayoutsByMerchantID(merchantID uint) ([]models.Payout, error) {
	if merchantID == 0 {
		return nil, errors.New("invalid merchant ID")
	}
	return s.payoutRepo.FindByMerchantID(merchantID)
}
```
</xaiArtifact>

---
### internal/domain/payment/payment_service.go
- Size: 3.38 KB
- Lines: 120
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="73c66036-4bec-4125-8ee1-93dc2d9a0af6" artifact_version_id="4a11a05e-4d8e-4a70-b3c2-44405cc53548" title="internal/domain/payment/payment_service.go" contentType="text/go">
```go
package payment

/*
import (
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"
	"errors"
	"os"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type PaymentService struct {
	paymentRepo *repositories.PaymentRepository
	orderRepo   *repositories.OrderRepository
}

func NewPaymentService(paymentRepo *repositories.PaymentRepository, orderRepo *repositories.OrderRepository) *PaymentService {
	return &PaymentService{
		paymentRepo: paymentRepo,
		orderRepo:   orderRepo,
	}
}

// ProcessPayment creates a payment for an order (placeholder for Stripe)
func (s *PaymentService) ProcessPayment(orderID uint, amount float64) (*models.Payment, error) {
    if orderID == 0 {
        return nil, errors.New("invalid order ID")
    }
    if amount <= 0 {
        return nil, errors.New("amount must be positive")
    }

    // Verify order exists
    order, err := s.orderRepo.FindByID(orderID)
    if err != nil {
        return nil, errors.New("order not found")
    }

    // Verify order amount matches
    if order.TotalAmount != amount {
        return nil, errors.New("payment amount does not match order total")
    }

    // Stripe setup (key from env)
    stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

    // Create Payment Intent with commission (application fee)
    pi, err := paymentintent.New(&stripe.PaymentIntentParams{  
        Amount:   stripe.Int64(int64(amount * 100)),
        Currency: stripe.String(string(stripe.CurrencyUSD)),
        // Add ApplicationFeeAmount for platform commission
        ApplicationFeeAmount: stripe.Int64(int64(amount * 100 * 0.1)), // 10% fee
        TransferData: &stripe.PaymentIntentTransferDataParams{
            Destination: stripe.String("merchant_stripe_id"), // From merchant model
        },
    })
    if err != nil {
        return nil, err
    }

    payment := &models.Payment{
        OrderID: orderID,
        Amount:  amount,
        Status:  models.PaymentStatusPending,
        //ExternalID: pi.ID,
    }
    if err := s.paymentRepo.Create(payment); err != nil {
        return nil, err
    }

    // Assume confirmation; in real, handle webhook
    payment.Status = models.PaymentStatusCompleted
    if err := s.paymentRepo.Update(payment); err != nil {
        return nil, err
    }

    return s.paymentRepo.FindByID(payment.ID)
}

// GetPaymentByOrderID retrieves a payment by order ID
func (s *PaymentService) GetPaymentByOrderID(orderID uint) (*models.Payment, error) {
	if orderID == 0 {
		return nil, errors.New("invalid order ID")
	}
	return s.paymentRepo.FindByOrderID(orderID)
}

// GetPaymentsByUserID retrieves all payments for a user
func (s *PaymentService) GetPaymentsByUserID(userID uint) ([]models.Payment, error) {
	if userID == 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.paymentRepo.FindByUserID(userID)
}

// UpdatePaymentStatus updates the status of a payment
func (s *PaymentService) UpdatePaymentStatus(paymentID uint, status string) (*models.Payment, error) {
	if paymentID == 0 {
		return nil, errors.New("invalid payment ID")
	}
	if err := models.PaymentStatus(status).Valid(); err != nil {
		return nil, err
	}

	payment, err := s.paymentRepo.FindByID(paymentID)
	if err != nil {
		return nil, err
	}

	payment.Status = models.PaymentStatus(status)
	if err := s.paymentRepo.Update(payment); err != nil {
		return nil, err
	}

	return s.paymentRepo.FindByID(paymentID)
}

*/
```
</xaiArtifact>

---
### internal/domain/merchant/merchant_service.go
- Size: 2.55 KB
- Lines: 84
- Last Modified: 2025-09-08 23:27:36

<xaiArtifact artifact_id="5f05aaae-cf21-494c-aa9e-55b7618ca899" artifact_version_id="7ab9761e-4b76-4b04-86db-3d0f6e094c00" title="internal/domain/merchant/merchant_service.go" contentType="text/go">
```go
package merchant

import (
	"context"
	"encoding/json"
	"errors"

	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"
	"github.com/go-playground/validator/v10"
	
)
type MerchantService struct {
	appRepo  *repositories.MerchantApplicationRepository
	repo     *repositories.MerchantRepository
	validate *validator.Validate
}

func NewMerchantService(appRepo *repositories.MerchantApplicationRepository, repo *repositories.MerchantRepository) *MerchantService {
	return &MerchantService{
		appRepo:  appRepo,
		repo:     repo,
		validate: validator.New(),
	}
}

// SubmitApplication allows a prospective merchant to submit an application.
func (s *MerchantService) SubmitApplication(ctx context.Context, app *models.MerchantApplication) (*models.MerchantApplication, error) {
	// Validate required fields
	if err := s.validate.Struct(app); err != nil {
		return nil, errors.New("validation failed: " + err.Error())
	}

	// Validate JSONB fields
	if len(app.PersonalAddress) == 0 {
		return nil, errors.New("personal_address cannot be empty")
	}
	if len(app.WorkAddress) == 0 {
		return nil, errors.New("work_address cannot be empty")
	}
	var temp map[string]interface{}
	if err := json.Unmarshal(app.PersonalAddress, &temp); err != nil {
		return nil, errors.New("invalid personal_address JSON")
	}
	if err := json.Unmarshal(app.WorkAddress, &temp); err != nil {
		return nil, errors.New("invalid work_address JSON")
	}

	
	

	// Set Status to pending and ensure ID is not set
	app.Status = "pending"
	app.ID = ""

	if err := s.appRepo.Create(ctx, app); err != nil {
		return nil, err
	}
	return app, nil
}

// GetApplication returns an application by ID (for applicant to check their own status).
func (s *MerchantService) GetApplication(ctx context.Context, id string) (*models.MerchantApplication, error) {
	if id == "" {
		return nil, errors.New("application ID cannot be empty")
	}
	return s.appRepo.GetByID(ctx, id)
}

// GetMerchantByUserID returns an active merchant account for the authenticated user.
func (s *MerchantService) GetMerchantByUserID(ctx context.Context, uid string) (*models.Merchant, error) {
	if uid == "" {
		return nil, errors.New("user ID cannot be empty")
	}
	return s.repo.GetByUserID(ctx, uid)
}

// GetMerchantByID returns an active merchant by ID.
func (s *MerchantService) GetMerchantByID(ctx context.Context, id string) (*models.Merchant, error) {
	if id == "" {
		return nil, errors.New("merchant ID cannot be empty")
	}
	return s.repo.GetByID(ctx, id)
}
```
</xaiArtifact>

---
### internal/domain/audit/audit_service.go
- Size: 0.51 KB
- Lines: 28
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="66c6e466-0b9d-44d9-ac30-1d5ed88e257a" artifact_version_id="6e50c0ef-4f36-41fa-8691-60f96f195a3b" title="internal/domain/audit/audit_service.go" contentType="text/go">
```go
package audit
/*
import (
    "context"
    "log"
    "api-customer-merchant/internal/db/models"
)

type AuditLog struct {
    models.Model
    UserID    uint
    Action    string
    Details   string
}

type AuditService struct {
    // Repo if DB logging
}

func NewAuditService() *AuditService {
    return &AuditService{}
}

func (s *AuditService) LogAction(ctx context.Context, userID uint, action, details string) {
    log.Printf("Audit: User %d - %s: %s", userID, action, details)
    // Save to DB if needed
}
*/	
```
</xaiArtifact>

---
### internal/domain/notifications/notifcation_service.go
- Size: 0.74 KB
- Lines: 31
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="5d3a4571-ce65-422f-a347-f154e26ac8d4" artifact_version_id="becde4a0-826d-443c-a3dd-4fc1f9da95fa" title="internal/domain/notifications/notifcation_service.go" contentType="text/go">
```go
package notifications

import (
    "context"
    "fmt"
    // Assume SMTP or Twilio lib
)

type NotificationService struct {
    // Email/SMS clients
}

func NewNotificationService() *NotificationService {
    return &NotificationService{}
}

func (s *NotificationService) SendOrderConfirmation(ctx context.Context, orderID uint, email string) error {
    // Use template: "Your order {orderID} is confirmed"
    fmt.Println("Sending email to", email) // Replace with real send
    return nil
}

func (s *NotificationService) SendMerchantNewOrder(merchantID uint, orderID uint) error {
    // Fetch merchant email, send
    return nil
}

func (s *NotificationService) SendStockAlert(merchantID uint, productID uint) error {
    // On low stock
    return nil
}
```
</xaiArtifact>

---
### internal/domain/order/order_service.go
- Size: 6.88 KB
- Lines: 222
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="d5409fb1-2248-4b52-8617-01a7e529af5d" artifact_version_id="c053e62c-ab33-432c-948a-fb6cd1f6ac4a" title="internal/domain/order/order_service.go" contentType="text/go">
```go
package order

import (
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"
	"api-customer-merchant/internal/domain/notifications"
	"api-customer-merchant/internal/domain/pricing"
	"context"
	"fmt"

	"errors"

	"gorm.io/gorm"
)

type OrderService struct {
    orderRepo          *repositories.OrderRepository
    cartRepo           *repositories.CartRepository
    cartItemRepo       *repositories.CartItemRepository
    orderItemRepo      *repositories.OrderItemRepository
    inventoryRepo      *repositories.InventoryRepository
    //paymentRepo        *repositories.PaymentRepository
    pricingService     *pricing.PricingService
    notificationService *notifications.NotificationService
}

// NewOrderService creates a new OrderService with dependencies
func NewOrderService(
    orderRepo *repositories.OrderRepository,
    cartRepo *repositories.CartRepository,
    cartItemRepo *repositories.CartItemRepository,
    orderItemRepo *repositories.OrderItemRepository,
    inventoryRepo *repositories.InventoryRepository,
    //paymentRepo *repositories.PaymentRepository,
    pricingService *pricing.PricingService,
    notificationService *notifications.NotificationService,
) *OrderService {
    return &OrderService{
        orderRepo:          orderRepo,
        cartRepo:           cartRepo,
        cartItemRepo:       cartItemRepo,
        orderItemRepo:      orderItemRepo,
        inventoryRepo:      inventoryRepo,
        //paymentRepo:        paymentRepo,
        pricingService:     pricingService,        
        notificationService: notificationService,  
    }
}

// CreateOrderFromCart creates an order from the user's active cart
func (s *OrderService) CreateOrderFromCart(userID uint) (*models.Order, error) {
    if userID == 0 {
        return nil, errors.New("invalid user ID")
    }

    // Get active cart
    cart, err := s.cartRepo.FindActiveCart(userID)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("no active cart found")
        }
        return nil, err
    }

    // Check if cart has items
    cartItems, err := s.cartItemRepo.FindByCartID(cart.ID)
    if err != nil {
        return nil, err
    }
    if len(cartItems) == 0 {
        return nil, errors.New("cart is empty")
    }

    // Validate stock for all items
    for _, item := range cartItems {
        inventory, err := s.inventoryRepo.FindByProductID(item.ProductID)
        if err != nil {
            return nil, errors.New("inventory not found for product")
        }
        if inventory.StockQuantity < item.Quantity {
            return nil, errors.New("insufficient stock for product")
        }
    }
	// Assume pricingService injected in NewOrderService
		
    // Calculate total amount
    var totalAmount float64
    for _, item := range cartItems {
        totalAmount += float64(item.Quantity) * item.Product.Price
    }

order := &models.Order{
		UserID:      userID,
		TotalAmount: totalAmount,
		Status:      models.OrderStatusPending,
	}
	uniqueMerchants := make(map[uint]bool)
	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}
 // Assume pricingService injected in NewOrderService
	shipping, err := s.pricingService.CalculateShipping(cart, "user_address") // Fetch from user
	if err != nil {
		return nil, err
	}
	tax, err := s.pricingService.CalculateTax(cart, "user_country")
	if err != nil {
		return nil, err
	}
	discount, err := s.pricingService.ApplyPromotion(cart, "promo_code") // From request
	if err != nil {
		return nil, err
	}
	order.TotalAmount = totalAmount + shipping + tax - discount
	order.ShippingCost = shipping
	order.TaxAmount = tax

	// Create order items and update inventory
	for _, item := range cartItems {
		orderItem := &models.OrderItem{
			OrderID:    order.ID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			Price:      item.Product.Price,
			MerchantID: item.MerchantID,
		}
		if err := s.orderItemRepo.Create(orderItem); err != nil {
			return nil, err
		}
		// Update inventory
		if err := s.inventoryRepo.UpdateStock(item.ProductID, -item.Quantity); err != nil {
			return nil, err
		}
	}

	// Mark cart as converted
	cart.Status = models.CartStatusConverted
	if err := s.cartRepo.Update(cart); err != nil {
		return nil, err
	}

	 go func() {
        ctx := context.Background()
        // Assume user_email is fetched from user data; replace "user_email" with actual email
        if err := s.notificationService.SendOrderConfirmation(ctx, order.ID, "user_email"); err != nil {
            // Log, don't fail
            fmt.Printf("Failed to send confirmation for order %d: %v\n", order.ID, err)
        }
        // For each unique merchant in items
        for merchantID := range uniqueMerchants {
            if err := s.notificationService.SendMerchantNewOrder(merchantID, order.ID); err != nil {
                fmt.Printf("Failed to notify merchant %d for order %d: %v\n", merchantID, order.ID, err)
            }
        }
    }()

	return s.orderRepo.FindByID(order.ID)
}
// GetOrderByID retrieves an order by ID
func (s *OrderService) GetOrderByID(id uint) (*models.Order, error) {
	if id == 0 {
		return nil, errors.New("invalid order ID")
	}
	return s.orderRepo.FindByID(id)
}

// GetOrdersByUserID retrieves all orders for a user
func (s *OrderService) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	if userID == 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.orderRepo.FindByUserID(userID)
}

// GetOrdersByMerchantID retrieves orders containing a merchant's products
func (s *OrderService) GetOrdersByMerchantID(merchantID uint) ([]models.Order, error) {
	if merchantID == 0 {
		return nil, errors.New("invalid merchant ID")
	}
	return s.orderRepo.FindByMerchantID(merchantID)
}

// UpdateOrderStatus updates the status of an order
func (s *OrderService) UpdateOrderStatus(orderID uint, status string) (*models.Order, error) {
	if orderID == 0 {
		return nil, errors.New("invalid order ID")
	}
	if err := models.OrderStatus(status).Valid(); err != nil {
		return nil, err
	}

	order, err := s.orderRepo.FindByID(orderID)
	if err != nil {
		return nil, err
	}

	order.Status = models.OrderStatus(status)
	if err := s.orderRepo.Update(order); err != nil {
		return nil, err
	}

	return s.orderRepo.FindByID(orderID)
}


func (s *OrderService) UpdateOrderItemStatus(orderItemID uint, status models.FulfillmentStatus, merchantID uint) error {
    item, err := s.orderItemRepo.FindByID(orderItemID)
    if err != nil || item.MerchantID != merchantID {
        return errors.New("invalid item or permission")
    }
    item.FulfillmentStatus = status
    return s.orderItemRepo.Update(item)
}

func (s *OrderService) ProcessRefund(orderID uint, amount float64) error {
    // payment, err := s.paymentRepo.FindByOrderID(orderID)
    // if err != nil || payment.Status != models.PaymentStatusCompleted {
    //     return errors.New("cannot refund")
    // }
    // Placeholder for Stripe refund
    return nil
}
```
</xaiArtifact>

---
### internal/db/repositories/payment_repository.go
- Size: 1.51 KB
- Lines: 52
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="8a783785-7581-4284-b2eb-68e25fc52c22" artifact_version_id="28fc89c5-c012-4011-9be3-165f24bb2395" title="internal/db/repositories/payment_repository.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{db: db.DB}
}

// Create adds a new payment
func (r *PaymentRepository) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

// FindByID retrieves a payment by ID with associated Order and User
func (r *PaymentRepository) FindByID(id uint) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.Preload("Order.User").First(&payment, id).Error
	return &payment, err
}

// FindByOrderID retrieves a payment by order ID
func (r *PaymentRepository) FindByOrderID(orderID uint) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.Preload("Order.User").Where("order_id = ?", orderID).First(&payment).Error
	return &payment, err
}

// FindByUserID retrieves all payments for a user
func (r *PaymentRepository) FindByUserID(userID uint) ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Preload("Order.User").Joins("JOIN orders ON orders.id = payments.order_id").Where("orders.user_id = ?", userID).Find(&payments).Error
	return payments, err
}

// Update modifies an existing payment
func (r *PaymentRepository) Update(payment *models.Payment) error {
	return r.db.Save(payment).Error
}

// Delete removes a payment by ID
func (r *PaymentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Payment{}, id).Error
}
```
</xaiArtifact>

---
### internal/db/repositories/merchant_repository.go
- Size: 2.22 KB
- Lines: 70
- Last Modified: 2025-09-08 23:05:28

<xaiArtifact artifact_id="47fb0b6c-1904-4bb1-baca-e35e4d3c7dd2" artifact_version_id="ee03d8ac-c6dc-48f2-aa63-c03c1f772e12" title="internal/db/repositories/merchant_repository.go" contentType="text/go">
```go
package repositories

import (
	"context"
	"log"

	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"
	//"gorm.io/gorm"
)

// MerchantApplicationRepository handles CRUD for merchant applications
// Note: Admin service (in Express) will be responsible for updating status/approval.
type MerchantApplicationRepository struct{}

func NewMerchantApplicationRepository() *MerchantApplicationRepository {
	return &MerchantApplicationRepository{}
}

func (r *MerchantApplicationRepository) Create(ctx context.Context, m *models.MerchantApplication) error {
	err := db.DB.WithContext(ctx).Create(m).Error
	if err != nil {
		log.Printf("Failed to create merchant application: %v", err)
		return err
	}
	return nil
}

func (r *MerchantApplicationRepository) GetByID(ctx context.Context, id string) (*models.MerchantApplication, error) {
	var m models.MerchantApplication
	if err := db.DB.WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		log.Printf("Failed to get merchant application by ID %s: %v", id, err)
		return nil, err
	}
	return &m, nil
}

func (r *MerchantApplicationRepository) GetByUserEmail(ctx context.Context, email string) (*models.MerchantApplication, error) {
	var m models.MerchantApplication
	if err := db.DB.WithContext(ctx).Where("personal_email = ? OR work_email = ?", email, email).First(&m).Error; err != nil {
		log.Printf("Failed to get merchant application by email %s: %v", email, err)
		return nil, err
	}
	return &m, nil
}

// MerchantRepository handles active merchants
type MerchantRepository struct{}

func NewMerchantRepository() *MerchantRepository {
	return &MerchantRepository{}
}

func (r *MerchantRepository) GetByID(ctx context.Context, id string) (*models.Merchant, error) {
	var m models.Merchant
	if err := db.DB.WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		log.Printf("Failed to get merchant by ID %s: %v", id, err)
		return nil, err
	}
	return &m, nil
}

func (r *MerchantRepository) GetByUserID(ctx context.Context, uid string) (*models.Merchant, error) {
	var m models.Merchant
	if err := db.DB.WithContext(ctx).Where("user_id = ?", uid).First(&m).Error; err != nil {
		log.Printf("Failed to get merchant by user ID %s: %v", uid, err)
		return nil, err
	}
	return &m, nil
}
```
</xaiArtifact>

---
### internal/db/repositories/payout_repository.go
- Size: 1.18 KB
- Lines: 45
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="d248915f-6927-483a-8c6d-59b1e9e2b234" artifact_version_id="5c04ce9d-a86b-4901-9c71-1ab510617547" title="internal/db/repositories/payout_repository.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"

	"gorm.io/gorm"
)

type PayoutRepository struct {
	db *gorm.DB
}

func NewPayoutRepository() *PayoutRepository {
	return &PayoutRepository{db: db.DB}
}

// Create adds a new payout record
func (r *PayoutRepository) Create(payout *models.Payout) error {
	return r.db.Create(payout).Error
}

// FindByID retrieves a payout by ID with associated Merchant
func (r *PayoutRepository) FindByID(id uint) (*models.Payout, error) {
	var payout models.Payout
	err := r.db.Preload("Merchant").First(&payout, id).Error
	return &payout, err
}

// FindByMerchantID retrieves all payouts for a merchant
func (r *PayoutRepository) FindByMerchantID(merchantID uint) ([]models.Payout, error) {
	var payouts []models.Payout
	err := r.db.Preload("Merchant").Where("merchant_id = ?", merchantID).Find(&payouts).Error
	return payouts, err
}

// Update modifies an existing payout
func (r *PayoutRepository) Update(payout *models.Payout) error {
	return r.db.Save(payout).Error
}

// Delete removes a payout by ID
func (r *PayoutRepository) Delete(id uint) error {
	return r.db.Delete(&models.Payout{}, id).Error
}
```
</xaiArtifact>

---
### internal/db/repositories/cart_item_repository.go
- Size: 1.90 KB
- Lines: 57
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="dd38a138-92b7-49b8-957a-ec349c8f7f6c" artifact_version_id="bebd6ea3-c196-48e9-9749-f84514361809" title="internal/db/repositories/cart_item_repository.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"

	"gorm.io/gorm"
)

type CartItemRepository struct {
	db *gorm.DB
}

func NewCartItemRepository() *CartItemRepository {
	return &CartItemRepository{db: db.DB}
}

// Create adds a new cart item
func (r *CartItemRepository) Create(cartItem *models.CartItem) error {
	return r.db.Create(cartItem).Error
}

// FindByID retrieves a cart item by ID with associated Cart, Product, and Merchant
func (r *CartItemRepository) FindByID(id uint) (*models.CartItem, error) {
	var cartItem models.CartItem
	err := r.db.Preload("Cart.User").Preload("Product.Merchant").Preload("Merchant").First(&cartItem, id).Error
	return &cartItem, err
}

// FindByCartID retrieves all cart items for a cart
func (r *CartItemRepository) FindByCartID(cartID uint) ([]models.CartItem, error) {
	var cartItems []models.CartItem
	err := r.db.Preload("Product.Merchant").Preload("Merchant").Where("cart_id = ?", cartID).Find(&cartItems).Error
	return cartItems, err
}

// FindByProductIDAndCartID retrieves a cart item by product ID and cart ID
func (r *CartItemRepository) FindByProductIDAndCartID(productID, cartID uint) (*models.CartItem, error) {
	var cartItem models.CartItem
	err := r.db.Preload("Product.Merchant").Preload("Merchant").Where("product_id = ? AND cart_id = ?", productID, cartID).First(&cartItem).Error
	return &cartItem, err
}

// UpdateQuantity updates the quantity of a cart item
func (r *CartItemRepository) UpdateQuantity(id uint, quantity int) error {
	return r.db.Model(&models.CartItem{}).Where("id = ?", id).Update("quantity", quantity).Error
}

// Update modifies an existing cart item
func (r *CartItemRepository) Update(cartItem *models.CartItem) error {
	return r.db.Save(cartItem).Error
}

// Delete removes a cart item by ID
func (r *CartItemRepository) Delete(id uint) error {
	return r.db.Delete(&models.CartItem{}, id).Error
}
```
</xaiArtifact>

---
### internal/db/repositories/return_repository.go
- Size: 0.65 KB
- Lines: 29
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="5a145adf-64c5-4e72-8c49-682a4a45cd72" artifact_version_id="bdd1d54e-e80e-40a6-b7a3-e37daa557a0c" title="internal/db/repositories/return_repository.go" contentType="text/go">
```go
package repositories

import (
    "api-customer-merchant/internal/db"
    "api-customer-merchant/internal/db/models"
    "gorm.io/gorm"
)

type ReturnRepository struct {
    db *gorm.DB
}

func NewReturnRepository() *ReturnRepository {
    return &ReturnRepository{db: db.DB}
}

func (r *ReturnRepository) Create(req *models.ReturnRequest) error {
    return r.db.Create(req).Error
}

func (r *ReturnRepository) FindByID(id uint) (*models.ReturnRequest, error) {
    var req models.ReturnRequest
    err := r.db.First(&req, id).Error
    return &req, err
}

func (r *ReturnRepository) Update(req *models.ReturnRequest) error {
    return r.db.Save(req).Error
}
```
</xaiArtifact>

---
### internal/db/repositories/category_repositry.go
- Size: 1.14 KB
- Lines: 45
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="957233ca-edf5-40ad-a87e-63141e18f536" artifact_version_id="ac6ab862-2289-4bb9-aced-4dd174dcf75e" title="internal/db/repositories/category_repositry.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{db: db.DB}
}

// Create adds a new category
func (r *CategoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

// FindByID retrieves a category by ID with parent category
func (r *CategoryRepository) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.Preload("Parent").First(&category, id).Error
	return &category, err
}

// FindAll retrieves all categories
func (r *CategoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Preload("Parent").Find(&categories).Error
	return categories, err
}

// Update modifies an existing category
func (r *CategoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

// Delete removes a category by ID
func (r *CategoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}
```
</xaiArtifact>

---
### internal/db/repositories/promotion_repository.go
- Size: 0.58 KB
- Lines: 25
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="c4cbca1a-d950-4b8c-b2b4-162bf8a4c561" artifact_version_id="80dc11e1-ecbe-429f-b627-aa899809d986" title="internal/db/repositories/promotion_repository.go" contentType="text/go">
```go
package repositories

import (
    "api-customer-merchant/internal/db"
    "api-customer-merchant/internal/db/models"
    "gorm.io/gorm"
)

type PromotionRepository struct {
    db *gorm.DB
}

func NewPromotionRepository() *PromotionRepository {
    return &PromotionRepository{db: db.DB}
}

func (r *PromotionRepository) Create(promo *models.Promotion) error {
    return r.db.Create(promo).Error
}

func (r *PromotionRepository) FindByCode(code string) (*models.Promotion, error) {
    var promo models.Promotion
    err := r.db.Where("code = ?", code).First(&promo).Error
    return &promo, err
}
```
</xaiArtifact>

---
### internal/db/repositories/inventory_repository.go
- Size: 1.29 KB
- Lines: 44
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="5329e4c2-84b5-494b-979c-f143f321b437" artifact_version_id="651d7dc5-2c44-4a01-ab11-1a148bbf5ea0" title="internal/db/repositories/inventory_repository.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"

	"gorm.io/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository() *InventoryRepository {
	return &InventoryRepository{db: db.DB}
}

// Create adds a new inventory record
func (r *InventoryRepository) Create(inventory *models.Inventory) error {
	return r.db.Create(inventory).Error
}

// FindByProductID retrieves inventory by product ID
func (r *InventoryRepository) FindByProductID(productID uint) (*models.Inventory, error) {
	var inventory models.Inventory
	err := r.db.Where("product_id = ?", productID).First(&inventory).Error
	return &inventory, err
}

// UpdateStock updates the stock quantity for a product
func (r *InventoryRepository) UpdateStock(productID uint, quantityChange int) error {
	return r.db.Model(&models.Inventory{}).Where("product_id = ?", productID).
		Update("stock_quantity", gorm.Expr("stock_quantity + ?", quantityChange)).Error
}

// Update modifies an existing inventory record
func (r *InventoryRepository) Update(inventory *models.Inventory) error {
	return r.db.Save(inventory).Error
}

// Delete removes an inventory record by ID
func (r *InventoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Inventory{}, id).Error
}
```
</xaiArtifact>

---
### internal/db/repositories/order_repository.go
- Size: 1.54 KB
- Lines: 53
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="9ccfedfa-e54d-4b33-8b6d-ff1758845e3d" artifact_version_id="8f202fd4-23bd-4fcc-b01f-7864e16f63d0" title="internal/db/repositories/order_repository.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{db: db.DB}
}

// Create adds a new order
func (r *OrderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

// FindByID retrieves an order by ID with associated User and OrderItems
func (r *OrderRepository) FindByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("User").Preload("OrderItems.Product.Merchant").First(&order, id).Error
	return &order, err
}

// FindByUserID retrieves all orders for a user
func (r *OrderRepository) FindByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("OrderItems.Product.Merchant").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

// FindByMerchantID retrieves all orders containing items from a merchant
func (r *OrderRepository) FindByMerchantID(merchantID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("OrderItems.Product").Joins("JOIN order_items oi ON oi.order_id = orders.id").
		Where("oi.merchant_id = ?", merchantID).Find(&orders).Error
	return orders, err
}

// Update modifies an existing order
func (r *OrderRepository) Update(order *models.Order) error {
	return r.db.Save(order).Error
}

// Delete removes an order by ID
func (r *OrderRepository) Delete(id uint) error {
	return r.db.Delete(&models.Order{}, id).Error
}
```
</xaiArtifact>

---
### internal/db/repositories/cart_repository.go
- Size: 1.82 KB
- Lines: 59
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="26425c17-e4e6-4e39-baec-7728ef66d4c5" artifact_version_id="6c3ce220-77fb-48f4-870a-6983fa2893c3" title="internal/db/repositories/cart_repository.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository() *CartRepository {
	return &CartRepository{db: db.DB}
}

// Create adds a new cart
func (r *CartRepository) Create(cart *models.Cart) error {
	return r.db.Create(cart).Error
}

// FindByID retrieves a cart by ID with associated User and CartItems
func (r *CartRepository) FindByID(id uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("User").Preload("CartItems.Product.Merchant").First(&cart, id).Error
	return &cart, err
}

// FindActiveCart retrieves the user's most recent active cart
func (r *CartRepository) FindActiveCart(userID uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("CartItems.Product.Merchant").Where("user_id = ? AND status = ?", userID, models.CartStatusActive).Order("created_at DESC").First(&cart).Error
	return &cart, err
}

// FindByUserIDAndStatus retrieves carts for a user by status
func (r *CartRepository) FindByUserIDAndStatus(userID uint, status models.CartStatus) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("CartItems.Product.Merchant").Where("user_id = ? AND status = ?", userID, status).Find(&carts).Error
	return carts, err
}

// FindByUserID retrieves all carts for a user
func (r *CartRepository) FindByUserID(userID uint) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("CartItems.Product.Merchant").Where("user_id = ?", userID).Find(&carts).Error
	return carts, err
}

// Update modifies an existing cart
func (r *CartRepository) Update(cart *models.Cart) error {
	return r.db.Save(cart).Error
}

// Delete removes a cart by ID
func (r *CartRepository) Delete(id uint) error {
	return r.db.Delete(&models.Cart{}, id).Error
}
```
</xaiArtifact>

---
### internal/db/repositories/product_repository.go
- Size: 2.67 KB
- Lines: 80
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="95eb363d-7f38-4fb9-8143-e3efbe6c27c9" artifact_version_id="18d2f8de-b05c-4eac-b1fa-83256bb10b49" title="internal/db/repositories/product_repository.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"
	"errors"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{db: db.DB}
}

// Create adds a new product to the database
func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// FindByID retrieves a product by ID with associated Merchant and Category
func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
    var product models.Product
    err := r.db.Preload("Merchant").Preload("Category").First(&product, id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("product not found")
        }
        return nil, err
    }
    return &product, nil
}

// FindBySKU retrieves a product by SKU with associated Merchant and Category
func (r *ProductRepository) FindBySKU(sku string) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Merchant").Preload("Category").Where("sku = ?", sku).First(&product).Error
	return &product, err
}

// SearchByName retrieves products matching a name (partial match)
func (r *ProductRepository) SearchByName(name string) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Merchant").Preload("Category").Where("name ILIKE ?", "%"+name+"%").Find(&products).Error
	return products, err
}

// FindByMerchantID retrieves all products for a merchant
func (r *ProductRepository) FindByMerchantID(merchantID uint) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Category").Where("merchant_id = ?", merchantID).Find(&products).Error
	return products, err
}

// FindByCategoryID retrieves all products in a category
func (r *ProductRepository) FindByCategoryID(categoryID uint) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Merchant").Where("category_id = ?", categoryID).Find(&products).Error
	return products, err
}

// Update modifies an existing product
func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// Delete removes a product by ID
func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
// In ProductRepository
func (r *ProductRepository) FindByCategoryWithPagination(categoryID uint, limit, offset int) ([]models.Product, error) {
    var products []models.Product
    err := r.db.Preload("Merchant").Preload("Category").Where("category_id = ?", categoryID).Limit(limit).Offset(offset).Find(&products).Error
    return products, err
}


```
</xaiArtifact>

---
### internal/db/repositories/order_item_repository.go
- Size: 1.35 KB
- Lines: 45
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="4955b7e7-fcb2-43cd-bbf1-b1ad39410cae" artifact_version_id="0b2c480c-40c5-4868-ad97-7e954ecd33b0" title="internal/db/repositories/order_item_repository.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"

	"gorm.io/gorm"
)

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository() *OrderItemRepository {
	return &OrderItemRepository{db: db.DB}
}

// Create adds a new order item
func (r *OrderItemRepository) Create(orderItem *models.OrderItem) error {
	return r.db.Create(orderItem).Error
}

// FindByID retrieves an order item by ID with associated Order, Product, and Merchant
func (r *OrderItemRepository) FindByID(id uint) (*models.OrderItem, error) {
	var orderItem models.OrderItem
	err := r.db.Preload("Order.User").Preload("Product.Merchant").Preload("Merchant").First(&orderItem, id).Error
	return &orderItem, err
}

// FindByOrderID retrieves all order items for an order
func (r *OrderItemRepository) FindByOrderID(orderID uint) ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	err := r.db.Preload("Product.Merchant").Preload("Merchant").Where("order_id = ?", orderID).Find(&orderItems).Error
	return orderItems, err
}

// Update modifies an existing order item
func (r *OrderItemRepository) Update(orderItem *models.OrderItem) error {
	return r.db.Save(orderItem).Error
}

// Delete removes an order item by ID
func (r *OrderItemRepository) Delete(id uint) error {
	return r.db.Delete(&models.OrderItem{}, id).Error
}
```
</xaiArtifact>

---
### internal/db/models/payout.go
- Size: 1.21 KB
- Lines: 49
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="b35d4031-427d-4d4d-b7a9-d9c37fbbc0d8" artifact_version_id="e608c06f-9606-4406-9285-22eabac19060" title="internal/db/models/payout.go" contentType="text/go">
```go
package models

import (
	"fmt"
	"gorm.io/gorm"
)

// PayoutStatus defines possible payout status values
type PayoutStatus string

const (
	PayoutStatusPending   PayoutStatus = "Pending"
	PayoutStatusCompleted PayoutStatus = "Completed"
)

// Valid checks if the status is one of the allowed values
func (s PayoutStatus) Valid() error {
	switch s {
	case PayoutStatusPending, PayoutStatusCompleted:
		return nil
	default:
		return fmt.Errorf("invalid payout status: %s", s)
	}
}

type Payout struct {
	gorm.Model
	MerchantID       uint         `gorm:"not null" json:"merchant_id"`
	Amount           float64      `gorm:"type:decimal(10,2);not null" json:"amount"`
	Status           PayoutStatus `gorm:"type:varchar(20);not null;default:'Pending'" json:"status"`
	PayoutAccountID  string       `gorm:"size:255;not null" json:"payout_account_id"`
	Merchant         Merchant     `gorm:"foreignKey:MerchantID"`
}

// BeforeCreate validates the Status field
func (p *Payout) BeforeCreate(tx *gorm.DB) error {
	if err := p.Status.Valid(); err != nil {
		return err
	}
	return nil
}

// BeforeUpdate validates the Status field
func (p *Payout) BeforeUpdate(tx *gorm.DB) error {
	if err := p.Status.Valid(); err != nil {
		return err
	}
	return nil
}
```
</xaiArtifact>

---
### internal/db/models/cart_item.go
- Size: 0.44 KB
- Lines: 16
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="06d73c93-78aa-4e35-b8e2-854d4305d95b" artifact_version_id="bcc96403-7213-4710-aaea-244d6cf6a1e4" title="internal/db/models/cart_item.go" contentType="text/go">
```go
package models

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	CartID     uint    `gorm:"not null" json:"cart_id"`
	ProductID  uint    `gorm:"not null" json:"product_id"`
	Quantity   int     `gorm:"not null" json:"quantity"`
	MerchantID uint    `gorm:"not null" json:"merchant_id"`
	Cart       Cart    `gorm:"foreignKey:CartID"`
	Product    Product `gorm:"foreignKey:ProductID"`
	Merchant   Merchant `gorm:"foreignKey:MerchantID"`
}
```
</xaiArtifact>

---
### internal/db/models/promotion.go
- Size: 0.23 KB
- Lines: 15
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="2a79bde7-38e2-4b14-8e7a-6dc727e1b585" artifact_version_id="a9785e68-1ea6-438b-888a-e76487c1f1c1" title="internal/db/models/promotion.go" contentType="text/go">
```go
package models

import (
    "time"
    "gorm.io/gorm"
)

type Promotion struct {
    gorm.Model
    Code       string    `gorm:"unique"`
    Discount   float64
    ValidFrom  time.Time
    ValidTo    time.Time
    MerchantID uint
}
```
</xaiArtifact>

---
### internal/db/models/merchants.go
- Size: 10.88 KB
- Lines: 180
- Last Modified: 2025-09-10 09:18:48

<xaiArtifact artifact_id="d5690827-9207-4bc1-8725-e643d8c97b87" artifact_version_id="e9adb3ea-8557-40ba-9bf9-f56f74447dbf" title="internal/db/models/merchants.go" contentType="text/go">
```go
package models

import (
	"time"
	"gorm.io/datatypes"
)

// MerchantBasicInfo holds basic merchant information
/*
type MerchantBasicInfo struct {
	StoreName     string `gorm:"column:store_name;size:255;not null;index" json:"store_name" validate:"required"`
	Name          string `gorm:"column:name;size:255;not null" json:"name" validate:"required"`
	PersonalEmail string `gorm:"column:personal_email;size:255;not null;unique;index" json:"personal_email" validate:"required,email"`
	WorkEmail     string `gorm:"column:work_email;size:255;not null;unique;index" json:"work_email" validate:"required,email"`
	PhoneNumber   string `gorm:"column:phone_number;size:50" json:"phone_number"`
	
}

// MerchantAddress holds address information as JSONB
type MerchantAddress struct {
	PersonalAddress datatypes.JSON `gorm:"column:personal_address;type:jsonb;not null" json:"personal_address" validate:"required"`
	WorkAddress     datatypes.JSON `gorm:"column:work_address;type:jsonb;not null" json:"work_address" validate:"required"`
}

// MerchantBusinessInfo holds business-related information
type MerchantBusinessInfo struct {
	BusinessType               string `gorm:"column:business_type;size:100" json:"business_type"`
	Website                    string `gorm:"column:website;size:255" json:"website"`
	BusinessDescription        string `gorm:"column:business_description;type:text" json:"business_description"`
	BusinessRegistrationNumber string `gorm:"column:business_registration_number;size:255;not null;unique" json:"business_registration_number" validate:"required"`
}

// MerchantDocuments holds document-related information
type MerchantDocuments struct {
	StoreLogoURL                   string `gorm:"column:store_logo_url;size:255" json:"store_logo_url"`
	BusinessRegistrationCertificate string `gorm:"column:business_registration_certificate;size:255" json:"business_registration_certificate"`
}

// MerchantApplication holds the information required for a merchant onboarding application
type MerchantApplication struct {
	ID              string            `gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4();index" json:"id,omitempty"`
	MerchantBasicInfo                 `gorm:"embedded"`
	MerchantAddress                   `gorm:"embedded"`
	MerchantBusinessInfo              `gorm:"embedded"`
	MerchantDocuments                 `gorm:"embedded"`
	Status          string            `gorm:"column:status;type:varchar(20);default:pending;not null" json:"-"` // Fixed to pending
	CreatedAt       time.Time         `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time         `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (MerchantApplication) TableName() string {
	return "merchant_application"
}

// MerchantStatus defines the possible statuses for a merchant
type MerchantStatus string

const (
	MerchantStatusActive    MerchantStatus = "active"
	MerchantStatusSuspended MerchantStatus = "suspended"
)

// Merchant holds the active merchant account details after approval
// This model is ONLY QUERIED by the merchant API (Go/Gin); schema migrations are handled by the admin API (Express).
type Merchant struct {
	ID                string            `gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4();index" json:"id,omitempty"`
	ApplicationID     string            `gorm:"column:application_id;type:uuid;not null;unique;index" json:"application_id"`
	UserID            string            `gorm:"column:user_id;type:uuid;not null;unique;index" json:"user_id"`
	MerchantBasicInfo                   `gorm:"embedded"`
	MerchantAddress                     `gorm:"embedded"`
	MerchantBusinessInfo                `gorm:"embedded"`
	MerchantDocuments                   `gorm:"embedded"`
	Password      string `gorm:"column:password;size:255;not null" json:"password" validate:"required"`
	Status            MerchantStatus    `gorm:"column:status;type:varchar(20);default:active;index" json:"status"`
	CommissionTier    string            `gorm:"column:commission_tier;default:standard" json:"commission_tier"`
	CommissionRate    float64           `gorm:"column:commission_rate;default:5.00" json:"commission_rate"`
	AccountBalance    float64           `gorm:"column:account_balance;default:0.00" json:"account_balance"`
	TotalSales        float64           `gorm:"column:total_sales;default:0.00" json:"total_sales"`
	TotalPayouts      float64           `gorm:"column:total_payouts;default:0.00" json:"total_payouts"`
	StripeAccountID   string            `gorm:"column:stripe_account_id" json:"stripe_account_id"`
	PayoutSchedule    string            `gorm:"column:payout_schedule;default:weekly" json:"payout_schedule"`
	LastPayoutDate    *time.Time        `gorm:"column:last_payout_date" json:"last_payout_date"`
	Banner            string            `gorm:"column:banner;size:255" json:"banner"`
	Policies          datatypes.JSON    `gorm:"column:policies;type:jsonb" json:"policies"`
	CreatedAt         time.Time         `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time         `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	Products          []Product         `gorm:"foreignKey:MerchantID" json:"products,omitempty"`
	CartItems         []CartItem        `gorm:"foreignKey:MerchantID" json:"cart_items,omitempty"`
	OrderItems        []OrderItem       `gorm:"foreignKey:MerchantID" json:"order_items,omitempty"`
	Payouts           []Payout          `gorm:"foreignKey:MerchantID" json:"payouts,omitempty"`
}

func (Merchant) TableName() string {
	return "merchant"
}
*/	
type MerchantBasicInfo struct {
	StoreName     string `gorm:"column:store_name;size:255;not null" json:"store_name" validate:"required"`
	Name          string `gorm:"column:name;size:255;not null" json:"name" validate:"required"`
	PersonalEmail string `gorm:"column:personal_email;size:255;not null;unique" json:"personal_email" validate:"required,email"`
	WorkEmail     string `gorm:"column:work_email;size:255;not null;unique" json:"work_email" validate:"required,email"`
	PhoneNumber   string `gorm:"column:phone_number;size:50" json:"phone_number"`
}

// MerchantAddress holds address information as JSONB
type MerchantAddress struct {
	PersonalAddress datatypes.JSON `gorm:"column:personal_address;type:jsonb;not null" json:"personal_address" validate:"required"`
	WorkAddress     datatypes.JSON `gorm:"column:work_address;type:jsonb;not null" json:"work_address" validate:"required"`
}

// MerchantBusinessInfo holds business-related information
type MerchantBusinessInfo struct {
	BusinessType               string `gorm:"column:business_type;size:100" json:"business_type"`
	Website                    string `gorm:"column:website;size:255" json:"website"`
	BusinessDescription        string `gorm:"column:business_description;type:text" json:"business_description"`
	BusinessRegistrationNumber string `gorm:"column:business_registration_number;size:255;not null;unique" json:"business_registration_number" validate:"required"`
}

// MerchantDocuments holds document-related information
type MerchantDocuments struct {
	StoreLogoURL                   string `gorm:"column:store_logo_url;size:255" json:"store_logo_url"`
	BusinessRegistrationCertificate string `gorm:"column:business_registration_certificate;size:255" json:"business_registration_certificate"`
}

// MerchantApplication holds the information required for a merchant onboarding application
type MerchantApplication struct {
	ID                string            `gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4()" json:"id,omitempty"`
	MerchantBasicInfo                   `gorm:"embedded"`
	MerchantAddress                     `gorm:"embedded"`
	MerchantBusinessInfo                `gorm:"embedded"`
	MerchantDocuments                   `gorm:"embedded"`
	Status            string            `gorm:"column:status;type:varchar(20);default:pending;not null" json:"status"`
	CreatedAt         time.Time         `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time         `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (MerchantApplication) TableName() string {
	return "merchant_application"
}

// MerchantStatus defines the possible statuses for a merchant
type MerchantStatus string

const (
	MerchantStatusActive    MerchantStatus = "active"
	MerchantStatusSuspended MerchantStatus = "suspended"
)

// Merchant holds the active merchant account details after approval
type Merchant struct {
	ID                string            `gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4()" json:"id,omitempty"`
	ApplicationID     string            `gorm:"column:application_id;type:uuid;not null;unique" json:"application_id"`
	MerchantID            string            `gorm:"column:merchant_id;type:uuid;not null;unique" json:"user_id"`
	MerchantBasicInfo                   `gorm:"embedded"`
	MerchantAddress                     `gorm:"embedded"`
	MerchantBusinessInfo                `gorm:"embedded"`
	MerchantDocuments                   `gorm:"embedded"`
	Password          string            `gorm:"column:password;size:255;not null" json:"password" validate:"required"`
	Status            MerchantStatus    `gorm:"column:status;type:varchar(20);default:active;index" json:"status"`
	CommissionTier    string            `gorm:"column:commission_tier;default:standard" json:"commission_tier"`
	CommissionRate    float64           `gorm:"column:commission_rate;default:5.00" json:"commission_rate"`
	AccountBalance    float64           `gorm:"column:account_balance;default:0.00" json:"account_balance"`
	TotalSales        float64           `gorm:"column:total_sales;default:0.00" json:"total_sales"`
	TotalPayouts      float64           `gorm:"column:total_payouts;default:0.00" json:"total_payouts"`
	StripeAccountID   string            `gorm:"column:stripe_account_id" json:"stripe_account_id"`
	PayoutSchedule    string            `gorm:"column:payout_schedule;default:weekly" json:"payout_schedule"`
	LastPayoutDate    *time.Time        `gorm:"column:last_payout_date" json:"last_payout_date"`
	Banner            string            `gorm:"column:banner;size:255" json:"banner"`
	Policies          datatypes.JSON    `gorm:"column:policies;type:jsonb" json:"policies"`
	CreatedAt         time.Time         `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time         `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	Products          []Product         `gorm:"foreignKey:MerchantID" json:"products,omitempty"`
	CartItems         []CartItem        `gorm:"foreignKey:MerchantID" json:"cart_items,omitempty"`
	OrderItems        []OrderItem       `gorm:"foreignKey:MerchantID" json:"order_items,omitempty"`
	Payouts           []Payout          `gorm:"foreignKey:MerchantID" json:"payouts,omitempty"`
}

func (Merchant) TableName() string {
	return "merchant"
}
```
</xaiArtifact>

---
### internal/db/models/payment.go
- Size: 1.16 KB
- Lines: 49
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="944afb6b-eddc-4d07-be1a-e8de17381865" artifact_version_id="d9b87c7b-fec4-400e-b15b-31cfbea1ae5d" title="internal/db/models/payment.go" contentType="text/go">
```go
package models

import (
	"fmt"
	"gorm.io/gorm"
)

// PaymentStatus defines possible payment status values
type PaymentStatus string

const (
	PaymentStatusPending  PaymentStatus = "Pending"
	PaymentStatusCompleted PaymentStatus = "Completed"
	PaymentStatusFailed   PaymentStatus = "Failed"
)

// Valid checks if the status is one of the allowed values
func (s PaymentStatus) Valid() error {
	switch s {
	case PaymentStatusPending, PaymentStatusCompleted, PaymentStatusFailed:
		return nil
	default:
		return fmt.Errorf("invalid payment status: %s", s)
	}
}

type Payment struct {
	gorm.Model
	OrderID uint          `gorm:"not null" json:"order_id"`
	Amount  float64       `gorm:"not null" json:"amount"`
	Status  PaymentStatus `gorm:"type:varchar(20);not null;default:'Pending'" json:"status"`
	Order   Order         `gorm:"foreignKey:OrderID"`
}

// BeforeCreate validates the Status field
func (p *Payment) BeforeCreate(tx *gorm.DB) error {
	if err := p.Status.Valid(); err != nil {
		return err
	}
	return nil
}

// BeforeUpdate validates the Status field
func (p *Payment) BeforeUpdate(tx *gorm.DB) error {
	if err := p.Status.Valid(); err != nil {
		return err
	}
	return nil
}
```
</xaiArtifact>

---
### internal/db/models/return_request.go
- Size: 0.19 KB
- Lines: 10
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="81b01ce0-63db-4842-a260-2b3b930de228" artifact_version_id="80d3180f-5b33-45a3-bf40-90205094c6a2" title="internal/db/models/return_request.go" contentType="text/go">
```go
package models

import "gorm.io/gorm"

type ReturnRequest struct {
    gorm.Model
    OrderItemID uint   `gorm:"not null"`
    Reason      string
    Status      string `gorm:"default:'Pending'"`
}
```
</xaiArtifact>

---
### internal/db/models/cart.go
- Size: 1.09 KB
- Lines: 49
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="368d4705-a62d-484b-839f-f7e6683452f5" artifact_version_id="9ffe4230-0cfc-4d12-9e4f-92b1834b9437" title="internal/db/models/cart.go" contentType="text/go">
```go
package models

import (
	"fmt"
	"gorm.io/gorm"
)

// CartStatus defines possible cart status values
type CartStatus string

const (
	CartStatusActive    CartStatus = "Active"
	CartStatusAbandoned CartStatus = "Abandoned"
	CartStatusConverted CartStatus = "Converted"
)

// Valid checks if the status is one of the allowed values
func (s CartStatus) Valid() error {
	switch s {
	case CartStatusActive, CartStatusAbandoned, CartStatusConverted:
		return nil
	default:
		return fmt.Errorf("invalid cart status: %s", s)
	}
}

type Cart struct {
	gorm.Model
	UserID uint       `gorm:"not null" json:"user_id"`
	Status CartStatus `gorm:"type:varchar(20);not null;default:'Active'" json:"status"`
	User   User       `gorm:"foreignKey:UserID"`
	CartItems []CartItem `gorm:"foreignKey:CartID"`
}

// BeforeCreate validates the Status field
func (c *Cart) BeforeCreate(tx *gorm.DB) error {
	if err := c.Status.Valid(); err != nil {
		return err
	}
	return nil
}

// BeforeUpdate validates the Status field
func (c *Cart) BeforeUpdate(tx *gorm.DB) error {
	if err := c.Status.Valid(); err != nil {
		return err
	}
	return nil
}
```
</xaiArtifact>

---
### internal/db/models/inventory.go
- Size: 0.34 KB
- Lines: 13
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="1e4a72bd-b9b1-4256-a39b-d9d1b3995f3a" artifact_version_id="7e40f2e7-91a7-47aa-a8e4-bb53781206dd" title="internal/db/models/inventory.go" contentType="text/go">
```go
package models

import (
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	ProductID         uint   `gorm:"not null" json:"product_id"`
	StockQuantity     int    `gorm:"not null" json:"stock_quantity"`
	LowStockThreshold int    `gorm:"not null;default:10" json:"low_stock_threshold"`
	Product           Product `gorm:"foreignKey:ProductID"`
}
```
</xaiArtifact>

---
### internal/db/models/product.go
- Size: 0.94 KB
- Lines: 38
- Last Modified: 2025-09-10 21:16:59

<xaiArtifact artifact_id="03698a8b-a5d2-453e-ad36-c977d55a2340" artifact_version_id="36efd6a8-e49d-4801-bb6e-ebfcfc88a045" title="internal/db/models/product.go" contentType="text/go">
```go
package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	MerchantID  uint    `gorm:"not null" json:"merchant_id"`
	Name        string  `gorm:"size:255;not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	SKU         string  `gorm:"size:100;unique;not null" json:"sku"`
	Price       float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	CategoryID  uint    `gorm:"not null" json:"category_id"`
	Merchant    Merchant `gorm:"foreignKey:MerchantID"`
	Category    Category `gorm:"foreignKey:CategoryID"`
	Variants    []Variant `gorm:"foreignKey:ProductID"`
    Media       []Media   `gorm:"foreignKey:ProductID"`
}


type Variant struct {
    gorm.Model
    ProductID  uint
    Attributes map[string]string `gorm:"type:jsonb"`
    Price      float64
    SKU        string
}


//define merchant variant

type Media struct {
    gorm.Model
    ProductID uint
    URL       string
    Type      string // image/video
}
```
</xaiArtifact>

---
### internal/db/models/order_item.go
- Size: 1.59 KB
- Lines: 53
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="27ebd439-d9fc-4587-8d72-e667821f6f8c" artifact_version_id="02a8bce8-aef9-42d2-b424-796a29273994" title="internal/db/models/order_item.go" contentType="text/go">
```go
package models

import (
	"fmt"
	"gorm.io/gorm"
)

// FulfillmentStatus defines possible fulfillment status values
type FulfillmentStatus string

const (
	FulfillmentStatusNew     FulfillmentStatus = "New"
	FulfillmentStatusShipped FulfillmentStatus = "Shipped"
)

// Valid checks if the status is one of the allowed values
func (s FulfillmentStatus) Valid() error {
	switch s {
	case FulfillmentStatusNew, FulfillmentStatusShipped:
		return nil
	default:
		return fmt.Errorf("invalid fulfillment status: %s", s)
	}
}

type OrderItem struct {
	gorm.Model
	OrderID           uint              `gorm:"not null" json:"order_id"`
	ProductID         uint              `gorm:"not null" json:"product_id"`
	MerchantID        uint              `gorm:"not null" json:"merchant_id"`
	Quantity          int               `gorm:"not null" json:"quantity"`
	Price             float64           `gorm:"type:decimal(10,2);not null" json:"price"`
	FulfillmentStatus FulfillmentStatus `gorm:"type:varchar(20);not null;default:'New'" json:"fulfillment_status"`
	Order             Order             `gorm:"foreignKey:OrderID"`
	Product           Product           `gorm:"foreignKey:ProductID"`
	Merchant          Merchant          `gorm:"foreignKey:MerchantID"`
}

// BeforeCreate validates the FulfillmentStatus field
func (oi *OrderItem) BeforeCreate(tx *gorm.DB) error {
	if err := oi.FulfillmentStatus.Valid(); err != nil {
		return err
	}
	return nil
}

// BeforeUpdate validates the FulfillmentStatus field
func (oi *OrderItem) BeforeUpdate(tx *gorm.DB) error {
	if err := oi.FulfillmentStatus.Valid(); err != nil {
		return err
	}
	return nil
}
```
</xaiArtifact>

---
### internal/db/models/merchant.bak.txt
- Size: 2.99 KB
- Lines: 90
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="cc540a18-4e36-40eb-8ca7-717aa89fd1a0" artifact_version_id="54c17de0-d116-4d45-a321-1aa2252204b9" title="internal/db/models/merchant.bak.txt" contentType="text/plain">
```plain
package models

import (
	"fmt"
	"gorm.io/gorm"
)

// MerchantStatus defines the possible status values for a merchant
type MerchantStatus string

const (
	MerchantStatusApproved  MerchantStatus = "Approved"
	MerchantStatusPending   MerchantStatus = "Pending"
	MerchantStatusInReview  MerchantStatus = "InReview"
	MerchantStatusRejected  MerchantStatus = "Rejected"
)

// Valid checks if the status is one of the allowed values
func (s MerchantStatus) Valid() error {
	switch s {
	case MerchantStatusApproved, MerchantStatusPending, MerchantStatusInReview, MerchantStatusRejected:
		return nil
	default:
		return fmt.Errorf("invalid merchant status: %s", s)
	}
}

// Basic merchant info
type MerchantBasicInfo struct {
	Name          string `gorm:"size:255;not null" json:"name"`
	StoreName     string `gorm:"size:255;not null" json:"store_name"`
	PersonalEmail string `gorm:"size:255;not null;unique" json:"personal_email"`
	WorkEmail     string `gorm:"size:255;not null;unique" json:"work_email"`
	PhoneNumber   string `gorm:"size:50" json:"phone_number"`
	Password      string `gorm:"size:255;not null" json:"password"`
}

// Address info
type MerchantAddress struct {
	StreetAddress string `gorm:"size:255" json:"street_address"`
	City          string `gorm:"size:100" json:"city"`
	Country       string `gorm:"size:100" json:"country"`
	ZipCode       string `gorm:"size:20" json:"zip_code"`
	WorkAddress   string `gorm:"size:255" json:"work_address"`
}

// Business info
type MerchantBusinessInfo struct {
	BusinessType        string `gorm:"size:100" json:"business_type"`
	Website             string `gorm:"size:255" json:"website"`
	BusinessDescription string `gorm:"type:text" json:"business_description"`
}

// Document info (only store logo and business registration certificate)
type MerchantDocuments struct {
	StoreLogoURL                   string `gorm:"size:255" json:"store_logo_url"`                    // image file (jpg, png)
	BusinessRegistrationCertificate string `gorm:"size:255" json:"business_registration_certificate"` // file (pdf, doc, docx, jpg, png)
}

// Main merchant model
type Merchant struct {
	gorm.Model
	MerchantBasicInfo
	MerchantAddress
	MerchantBusinessInfo
	MerchantDocuments
	Status      MerchantStatus `gorm:"type:varchar(20);not null;default:'Pending'" json:"status"`
	Products    []Product      `gorm:"foreignKey:MerchantID" json:"products,omitempty"`
	CartItems   []CartItem     `gorm:"foreignKey:MerchantID" json:"cart_items,omitempty"`
	OrderItems  []OrderItem    `gorm:"foreignKey:MerchantID" json:"order_items,omitempty"`
	Payouts     []Payout       `gorm:"foreignKey:MerchantID" json:"payouts,omitempty"`
}

// BeforeCreate validates the Status field before saving to the database
func (m *Merchant) BeforeCreate(tx *gorm.DB) error {
	if err := m.Status.Valid(); err != nil {
		return err
	}
	return nil
}

// BeforeUpdate validates the Status field before updating the database
func (m *Merchant) BeforeUpdate(tx *gorm.DB) error {
	if err := m.Status.Valid(); err != nil {
		return err
	}
	return nil
}



```
</xaiArtifact>

---
### internal/db/models/user.go
- Size: 0.54 KB
- Lines: 15
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="392cc212-df6b-4889-a278-19b8ebdd34cd" artifact_version_id="c2421f7a-8f4e-4daa-b3e1-bc74d73ee4f7" title="internal/db/models/user.go" contentType="text/go">
```go
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Name     string `gorm:"type:varchar(100);not null"`
	Password string // Empty for OAuth users
	//Role     string `gorm:"not null"` // "customer" (default) or "merchant" (upgraded by admin)
	GoogleID string    // Google ID for OAuth
	Country  string `gorm:"type:varchar(100)"` // Optional country field
	Carts    []Cart  `gorm:"foreignKey:UserID" json:"carts,omitempty"`
	Orders   []Order `gorm:"foreignKey:UserID" json:"orders,omitempty"`
}
```
</xaiArtifact>

---
### internal/db/models/category.go
- Size: 0.33 KB
- Lines: 13
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="1648c4cb-3567-4824-8a19-85e58965a6a8" artifact_version_id="88bfb9fe-7f22-4a10-8eab-7a8144bf0728" title="internal/db/models/category.go" contentType="text/go">
```go
package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name       string                 `gorm:"size:255;not null" json:"name"`
	ParentID   *uint                  `json:"parent_id"`
	Attributes map[string]any `gorm:"type:jsonb" json:"attributes"`
	Parent     *Category              `gorm:"foreignKey:ParentID"`
}
```
</xaiArtifact>

---
### internal/db/models/order.go
- Size: 1.33 KB
- Lines: 50
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="ce457b88-a38c-417b-83dd-e99dd6c5027f" artifact_version_id="8877fdc8-d912-4536-8b5b-4e159153aeb4" title="internal/db/models/order.go" contentType="text/go">
```go
package models

import (
	"fmt"
	"gorm.io/gorm"
)

// OrderStatus defines possible order status values
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "Pending"
	OrderStatusCompleted OrderStatus = "Completed"
	OrderStatusCancelled OrderStatus = "Cancelled"
)

// Valid checks if the status is one of the allowed values
func (s OrderStatus) Valid() error {
	switch s {
	case OrderStatusPending, OrderStatusCompleted, OrderStatusCancelled:
		return nil
	default:
		return fmt.Errorf("invalid order status: %s", s)
	}
}

type Order struct {
    gorm.Model
    UserID       uint        `gorm:"not null" json:"user_id"`
    TotalAmount  float64     `gorm:"type:decimal(10,2);not null" json:"total_amount"`
    ShippingCost float64     `gorm:"type:decimal(10,2)" json:"shipping_cost"`
    TaxAmount    float64     `gorm:"type:decimal(10,2)" json:"tax_amount"`
    Status       OrderStatus `gorm:"type:varchar(20);not null;default:'Pending'" json:"status"`
    User         User        `gorm:"foreignKey:UserID"`
}
// BeforeCreate validates the Status field
func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if err := o.Status.Valid(); err != nil {
		return err
	}
	return nil
}

// BeforeUpdate validates the Status field
func (o *Order) BeforeUpdate(tx *gorm.DB) error {
	if err := o.Status.Valid(); err != nil {
		return err
	}
	return nil
}
```
</xaiArtifact>

---
### internal/tests/unit/product_service_test.go
- Size: 0.64 KB
- Lines: 19
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="8cc83c35-cdb1-4743-88b0-89190d07c9cc" artifact_version_id="1a44e8fd-8f40-4f69-bc68-def37f34be67" title="internal/tests/unit/product_service_test.go" contentType="text/go">
```go
package unit

import (
    "testing"
    "api-customer-merchant/internal/db/models"
    "api-customer-merchant/internal/domain/product"
    "api-customer-merchant/internal/db/repositories" // Mock if needed
)

func TestCreateProduct(t *testing.T) {
    productRepo := repositories.NewProductRepository() // Or mock
    inventoryRepo := repositories.NewInventoryRepository()
    service := product.NewProductService(productRepo, inventoryRepo)
    product := &models.Product{Name: "Test", SKU: "TEST123", Price: 100, CategoryID: 1}
    err := service.CreateProduct(product, 1)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
}
```
</xaiArtifact>

---
### internal/tests/unit/test_service.go
- Size: 5.61 KB
- Lines: 172
- Last Modified: 2025-09-08 22:19:09

<xaiArtifact artifact_id="0aa6075b-d6b8-40c2-84b2-3c511c9b2e89" artifact_version_id="8348c047-5e21-45cd-bab9-cccd17722843" title="internal/tests/unit/test_service.go" contentType="text/go">
```go
package unit
/*
 import (
 	"fmt"
 	"api-customer-merchant/internal/db"
 	"api-customer-merchant/internal/db/models"
 	"api-customer-merchant/internal/db/repositories"
 	"api-customer-merchant/internal/domain/cart"
 	"api-customer-merchant/internal/domain/order"
 	"api-customer-merchant/internal/domain/payment"
 	"api-customer-merchant/internal/domain/payout"
 	"api-customer-merchant/internal/domain/product"
 )

 func TestServices() {
 	db.Connect()
 	db.AutoMigrate()

 	 Setup repositories
 	productRepo := repositories.NewProductRepository()
 	inventoryRepo := repositories.NewInventoryRepository()
 	cartRepo := repositories.NewCartRepository()
 	cartItemRepo := repositories.NewCartItemRepository()
 	orderRepo := repositories.NewOrderRepository()
 	orderItemRepo := repositories.NewOrderItemRepository()
 	paymentRepo := repositories.NewPaymentRepository()
 	payoutRepo := repositories.NewPayoutRepository()

 	 Setup services
 	productService := product.NewProductService(productRepo, inventoryRepo)
 	cartService := cart.NewCartService(cartRepo, cartItemRepo, productRepo, inventoryRepo)
 	orderService := order.NewOrderService(orderRepo, orderItemRepo, cartRepo, cartItemRepo, inventoryRepo)
 	paymentService := payment.NewPaymentService(paymentRepo, orderRepo)
 	payoutService := payout.NewPayoutService(payoutRepo)

 	 Insert test data
 	user := &models.User{Email: "test@example.com", Name: "Test User", Password: "$2a$10$examplehashedpassword", Country: "Nigeria"}
 	if err := db.DB.Create(user).Error; err != nil {
 		fmt.Println("Error creating user:", err)
 	}

 	merchant := &models.Merchant{MerchantBasicInfo: models.MerchantBasicInfo{Name: "Test Merchant", StoreName: "Test Store", PersonalEmail: "personal@example.com", WorkEmail: "work@example.com", Password: "$2a$10$examplehashedpassword"}, Status: models.MerchantStatusApproved}
 	if err := db.DB.Create(merchant).Error; err != nil {
 		fmt.Println("Error creating merchant:", err)
 	}

 	category := &models.Category{Name: "Electronics", Attributes: map[string]interface{}{"color": []string{"black"}}}
 	if err := db.DB.Create(category).Error; err != nil {
 		fmt.Println("Error creating category:", err)
 	}

 	 Test ProductService
 	product := &models.Product{Name: "Smartphone", SKU: "SM123", Price: 599.99, CategoryID: category.ID, MerchantID: merchant.ID}
 	if err := productService.CreateProduct(product, merchant.ID); err != nil {
 		fmt.Println("Error creating product:", err)
 	}
 	p, err := productService.GetProductBySKU("SM123")
 	if err != nil {
 		fmt.Println("Error finding product by SKU:", err)
 	} else {
 		fmt.Printf("Found product by SKU: %+v\n", p)
 	}
 	products, err := productService.SearchProductsByName("Smart")
 	if err != nil {
 		fmt.Println("Error searching products:", err)
 	} else {
 		fmt.Printf("Found products by name: %+v\n", products)
 	}

 	 Add inventory
 	inventory := &models.Inventory{ProductID: p.ID, StockQuantity: 100, LowStockThreshold: 10}
 	if err := inventoryRepo.Create(inventory); err != nil {
 		fmt.Println("Error creating inventory:", err)
 	}

 	 Test inventory stock update
 	if err := inventoryRepo.UpdateStock(p.ID, -10); err != nil {
 		fmt.Println("Error updating stock:", err)
 	} else {
 		inv, _ := inventoryRepo.FindByProductID(p.ID)
 		fmt.Printf("Updated stock: %+v\n", inv)
 	}

 	 Test CartService
 	cart, err := cartService.GetActiveCart(user.ID)
 	if err != nil {
 		fmt.Println("Error getting active cart:", err)
 	} else {
 		fmt.Printf("Active cart: %+v\n", cart)
 	}

 	 Add item to cart
 	cart, err = cartService.AddItemToCart(user.ID, p.ID, 2)
 	if err != nil {
 		fmt.Println("Error adding item to cart:", err)
 	} else {
 		fmt.Printf("Cart after adding item: %+v\n", cart)
 	}

 	 Test OrderService
 	order, err := orderService.CreateOrderFromCart(user.ID)
 	if err != nil {
 		fmt.Println("Error creating order:", err)
 	} else {
 		fmt.Printf("Created order: %+v\n", order)
 	}

 	 Verify stock after order
 	inv, err := inventoryRepo.FindByProductID(p.ID)
 	if err != nil {
 		fmt.Println("Error checking stock after order:", err)
 	} else {
 		fmt.Printf("Stock after order: %+v\n", inv)
 	}

 	 Retrieve order
 	o, err := orderService.GetOrderByID(order.ID)
 	if err != nil {
 		fmt.Println("Error finding order by ID:", err)
 	} else {
 		fmt.Printf("Found order by ID: %+v\n", o)
 	}

 	 Update order status
 	o, err = orderService.UpdateOrderStatus(order.ID, string(models.OrderStatusShipped))
 	if err != nil {
 		fmt.Println("Error updating order status:", err)
 	} else {
 		fmt.Printf("Updated order status: %+v\n", o)
 	}

 	 Test PaymentService
 	payment, err := paymentService.ProcessPayment(order.ID, order.TotalAmount)
 	if err != nil {
 		fmt.Println("Error processing payment:", err)
 	} else {
 		fmt.Printf("Processed payment: %+v\n", payment)
 	}

 	 Retrieve payment
 	pay, err := paymentService.GetPaymentByOrderID(order.ID)
 	if err != nil {
 		fmt.Println("Error finding payment by order ID:", err)
 	} else {
 		fmt.Printf("Found payment by order ID: %+v\n", pay)
 	}

 	 Update payment status
 	pay, err = paymentService.UpdatePaymentStatus(payment.ID, string(models.PaymentStatusCompleted))
 	if err != nil {
 		fmt.Println("Error updating payment status:", err)
 	} else {
 		fmt.Printf("Updated payment status: %+v\n", pay)
 	}

 	 Test PayoutService
 	payout, err := payoutService.CreatePayout(merchant.ID, 500.00)
 	if err != nil {
 		fmt.Println("Error creating payout:", err)
 	} else {
 		fmt.Printf("Created payout: %+v\n", payout)
 	}

 	 Retrieve payout
 	po, err := payoutService.GetPayoutByID(payout.ID)
 	if err != nil {
 		fmt.Println("Error finding payout by ID:", err)
 	} else {
 		fmt.Printf("Found payout by ID: %+v\n", po)
 	}
 }
*/
```
</xaiArtifact>

---
### internal/api/handlers/merchant_handlers.go
- Size: 6.61 KB
- Lines: 233
- Last Modified: 2025-09-12 03:30:50

<xaiArtifact artifact_id="1bc59277-c9e8-424b-89c2-5919192e1e7b" artifact_version_id="80e5ed7b-357a-42af-805d-e67bddd1a53d" title="internal/api/handlers/merchant_handlers.go" contentType="text/go">
```go
package handlers

import (
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/domain/order"
	"api-customer-merchant/internal/domain/payout"
	"api-customer-merchant/internal/domain/product"
	"api-customer-merchant/internal/domain/promotions"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type MerchantHandlers struct {
	productService *product.ProductService
	orderService   *order.OrderService
	payoutService  *payout.PayoutService
	promotionService *promotions.PromotionService
}

func NewMerchantHandlers(productService *product.ProductService, orderService *order.OrderService, payoutService *payout.PayoutService,promotionService *promotions.PromotionService) *MerchantHandlers {
	return &MerchantHandlers{
		productService: productService,
		orderService:   orderService,
		payoutService:  payoutService,
		promotionService:promotionService,
	}
}

// CreateProduct handles POST /merchant/products
func (h *MerchantHandlers) CreateProduct(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.productService.CreateProduct(&product, merchantID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// UpdateProduct handles PUT /merchant/products/:id
func (h *MerchantHandlers) UpdateProduct(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = uint(id)

	if err := h.productService.UpdateProduct(&product, merchantID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct handles DELETE /merchant/products/:id
func (h *MerchantHandlers) DeleteProduct(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	if err := h.productService.DeleteProduct(uint(id), merchantID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
/*
// GetOrders handles GET /merchant/orders
func (h *MerchantHandlers) GetOrders(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	orders, err := h.orderService.GetOrdersByMerchantID(merchantID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GetPayouts handles GET /merchant/payouts
func (h *MerchantHandlers) GetPayouts(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	payouts, err := h.payoutService.GetPayoutsByMerchantID(merchantID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payouts)


}


func (h *MerchantHandlers) CreatePromotion(c *gin.Context) {
    merchantID, _ := c.Get("merchantID")
    var promo models.Promotion
    if err := c.ShouldBindJSON(&promo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    promo.MerchantID = merchantID.(uint)
    if err := h.promotionService.CreatePromotion(&promo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, promo)
}

func (h *MerchantHandlers) UpdateOrderItemStatus(c *gin.Context) {
    merchantID, _ := c.Get("merchantID")
    itemIDStr := c.Param("itemID")
    itemID, _ := strconv.ParseUint(itemIDStr, 10, 32)
    var req struct { Status string `json:"status"` }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.orderService.UpdateOrderItemStatus(uint(itemID), models.FulfillmentStatus(req.Status), merchantID.(uint)); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "updated"})
}
*/

func (h *MerchantHandlers) BulkUploadProducts(c *gin.Context) {
    merchantID, _ := c.Get("merchantID")
    file, err := c.FormFile("csv")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "csv file required"})
        return
    }
    f, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer f.Close()

    reader := csv.NewReader(f)
    // Skip header
    _, err = reader.Read()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid csv header"})
        return
    }

    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid csv format"})
            return
        }

        // Parse price with error handling
        price, err := strconv.ParseFloat(record[2], 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid price format in CSV: %s", record[2])})
            return
        }

        // Parse record: name, sku, price, etc.
        product := &models.Product{
            Name:       record[0],
            SKU:        record[1],
            Price:      price,
            // Add other fields as needed (e.g., CategoryID if required in CSV)
            MerchantID: merchantID.(uint),
        }

        if err := h.productService.CreateProduct(product, merchantID.(uint)); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to create product: %s", err.Error())})
            return
        }
    }
    c.JSON(http.StatusOK, gin.H{"message": "products uploaded"})
}

```
</xaiArtifact>

---
### internal/api/handlers/auth_handler.go
- Size: 2.76 KB
- Lines: 89
- Last Modified: 2025-09-08 23:28:28

<xaiArtifact artifact_id="cf701535-8319-4c4e-aa20-db6fb8d921ed" artifact_version_id="76c53782-2fd0-4908-8b16-3a55f01f3c31" title="internal/api/handlers/auth_handler.go" contentType="text/go">
```go
package handlers

import (
	"encoding/json"
	"net/http"

	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/domain/merchant"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	service *merchant.MerchantService
}

func NewMerchantHandler(s *merchant.MerchantService) *MerchantHandler {
	return &MerchantHandler{service: s}
}

// POST /merchant/apply
// Used by a prospective merchant to submit an application.
func (h *MerchantHandler) Apply(c *gin.Context) {
	var req struct {
		models.MerchantBasicInfo
		PersonalAddress            map[string]any `json:"personal_address" validate:"required"`
		WorkAddress                map[string]any `json:"work_address" validate:"required"`
		models.MerchantBusinessInfo
		models.MerchantDocuments
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	// Convert personal_address and work_address to JSONB
	personalAddressJSON, err := json.Marshal(req.PersonalAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid personal_address JSON: " + err.Error()})
		return
	}
	workAddressJSON, err := json.Marshal(req.WorkAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid work_address JSON: " + err.Error()})
		return
	}

	app := &models.MerchantApplication{
		MerchantBasicInfo:    req.MerchantBasicInfo,
		MerchantAddress:      models.MerchantAddress{PersonalAddress: personalAddressJSON, WorkAddress: workAddressJSON},
		MerchantBusinessInfo: req.MerchantBusinessInfo,
		MerchantDocuments:    req.MerchantDocuments,
	}

	app, err = h.service.SubmitApplication(c.Request.Context(), app)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to submit application: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, app)
}

// GET /merchant/application/:id
// Allows an applicant to view their application status.
func (h *MerchantHandler) GetApplication(c *gin.Context) {
	id := c.Param("id")
	app, err := h.service.GetApplication(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "application not found: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, app)
}

// GET /merchant/me
// Fetches the current user's merchant account (if approved).
func (h *MerchantHandler) GetMyMerchant(c *gin.Context) {
	userID, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	m, err := h.service.GetMerchantByUserID(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "merchant not found: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, m)
}
```
</xaiArtifact>

---

---
## 📊 Summary
- Total files: 52
- Total size: 99.38 KB
- File types: .go, .txt, unknown
