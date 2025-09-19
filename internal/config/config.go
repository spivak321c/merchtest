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
