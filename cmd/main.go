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