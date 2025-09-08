package config

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