package routes

import (
	"api-customer-merchant/internal/api/handlers"
	"api-customer-merchant/internal/middleware"
	"api-customer-merchant/internal/services/cart"
	"github.com/gin-gonic/gin"
)

func SetupCartRoutes(r *gin.RouterGroup, cartService *cart.CartService) {
	customerHandlers := handlers.NewCartHandler(cartService)
	protected := middleware.AuthMiddleware("user")
	r.GET("/cart", protected, customerHandlers.GetCart)
	r.POST("/cart/items", protected, customerHandlers.AddToCart)
	r.GET("/cart/items/:id", protected, customerHandlers.GetCartItem)
	r.PUT("/cart/items/:id", protected, customerHandlers.UpdateCartItemQuantity)
	r.DELETE("/cart/items/:id", protected, customerHandlers.RemoveCartItem)
	//r.POST("/cart/clear", protected, customerHandlers.ClearCart)
}