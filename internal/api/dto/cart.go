package dto

import (
	"api-customer-merchant/internal/db/models" // For enums if needed
	//"time"
)

// AddItemRequest: For POST /cart/add (add new item or increment existing)
type AddItemRequest struct {
	UserID uint               `json:"user_id,omitempty"`
    ProductID uint `json:"product_id" validate:"required"`
    Quantity  int  `json:"quantity" validate:"required,gt=0"`
}

// UpdateItemRequest: For PUT /cart/items/:id (full replace quantity) or PATCH /cart/items/:id (partial update)
type UpdateItemRequest struct {
	UserID uint               `json:"user_id,omitempty"`
    Quantity int `json:"quantity" validate:"omitempty,gt=0"` // Omitempty for PATCH
}

// BulkUpdateRequest: For POST /cart/bulk (add/update multiple—extension for prod)
type BulkUpdateRequest struct {
	UserID uint               `json:"user_id,omitempty"`
    Items []struct {
        ProductID uint `json:"product_id" validate:"required"`
        Quantity  int  `json:"quantity" validate:"required,gt=0"`
    } `json:"items" validate:"dive"`
}

// CartResponse: For all responses (shared output DTO)
type CartResponse struct {
    ID     string             `json:"id"`
    UserID uint               `json:"user_id,omitempty"` // Omit for security
    Status models.CartStatus  `json:"status"`
    Items  []CartItemResponse `json:"items"`
    Total  float64            `json:"total,omitempty"` // Computed in service
 
}

type CartItemResponse struct {
    ID        string  `json:"id"`
    ProductID string  `json:"product_id"`
    Product   *ProductResponse // Embed lightweight product (from dto/product.go)
    Quantity  int     `json:"quantity"`
    Subtotal  float64 `json:"subtotal"` // Quantity * Product.BasePrice
}