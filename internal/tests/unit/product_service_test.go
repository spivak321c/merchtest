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