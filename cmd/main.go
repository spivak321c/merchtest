package main

import (
	"os"

	merchant"api-customer-merchant/internal/api"
	"api-customer-merchant/internal/config"
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/logger"
	"api-customer-merchant/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, _ := config.Load()
	logger.Init(cfg.Env == "development")
	utils.InitRedis(cfg)
	db.Connect(cfg)

	r := gin.Default()
	r.Use(gin.Recovery())

	merchant.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" { port = "8000" }
	r.Run(":" + port)
}