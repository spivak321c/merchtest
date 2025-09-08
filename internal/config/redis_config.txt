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