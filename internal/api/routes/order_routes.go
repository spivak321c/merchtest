package routes

import (
	"api-customer-merchant/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(r *gin.RouterGroup, orderHandler *handlers.OrderHandler) {
	//protected := middleware.AuthMiddleware("user") // Consider adding auth middleware
	r.POST("/orders", orderHandler.CreateOrder)
	r.POST("/orders", orderHandler.CreateOrder)       //create order
	r.GET("/orders/:id", orderHandler.GetOrder) //get order

	//r.GET("/orders/:id", protected, orderHandler.GetOrder)
}
