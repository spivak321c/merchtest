package routes

import (
	"api-customer-merchant/internal/api/handlers"
	"api-customer-merchant/internal/db/repositories"
	"api-customer-merchant/internal/services/product"

	//"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
	// "github.com/shopspring/decimal"
	 "go.uber.org/zap"
	//"gorm.io/gorm"
)

func RegisterProductRoutes(r *gin.Engine) {
	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logger.Sync() // Ensure logger flushes logs

	// Initialize validator
	//validator := validator.New()
	



	// Initialize repository
	repo := repositories.NewProductRepository()

	// Initialize service
	service := product.NewProductService(repo, logger)

	// Create product route group
	productGroup := r.Group("/product")
	{
		// Initialize handlers
		// Adjust NewProductHandlers if validator is needed
		ProductHandler := handlers.NewProductHandlers(service, logger)

		// Example route (replace with actual routes as needed)
		productGroup.POST("/create", ProductHandler.CreateProduct)
		productGroup.DELETE("/delete/:id", ProductHandler.DeleteProduct)
		productGroup.GET("/products", ProductHandler.GetAllProducts)
		productGroup.GET("/:id", ProductHandler.GetProductByID)
		productGroup.GET("/merchant/:id", ProductHandler.ListProductsByMerchant)
		productGroup.PUT("/products/inventory", ProductHandler.UpdateInventory)


	}
}