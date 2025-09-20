# Codebase Analysis: internal
Generated: 2025-09-19 22:01:26
---

## 📂 Project Structure
```tree
📁 internal
├── api/
│   ├── dto/
│   │   ├── cart.go
│   │   ├── order.go
│   │   └── product.go
│   ├── handlers/
│   │   ├── cart_handler.go
│   │   ├── customer_auth_handler.go
│   │   ├── customer_handlers.go
│   │   ├── merchant_auth_handler.go
│   │   ├── merchant_handlers.go
│   │   ├── order_handler.go
│   │   └── product_handler.go
│   └── routes/
│       ├── carts.go
│       ├── customer.go
│       ├── merchant.go
│       ├── order_routes.go
│       └── product_route.go
├── config/
│   └── config.go
├── db/
│   ├── models/
│   │   ├── cart.go
│   │   ├── cart_item.go
│   │   ├── category.go
│   │   ├── inventory.go
│   │   ├── merchant.go
│   │   ├── order.go
│   │   ├── order_item.go
│   │   ├── payment.go
│   │   ├── payout.go
│   │   ├── product.go
│   │   ├── product2.go
│   │   └── user.go
│   ├── repositories/
│   │   ├── cart_item_repositry.go
│   │   ├── cart_repositry.go
│   │   ├── category_repositry.go
│   │   ├── inventory_repository.go
│   │   ├── merchant_repositry.go
│   │   ├── order_item_repository.go
│   │   ├── order_repository.go
│   │   ├── payment_repository.go
│   │   ├── payout_repository.go
│   │   ├── product_repo.go
│   │   ├── product_repositry.go
│   │   └── user_repository.go
│   └── db.go
├── middleware/
│   ├── auth.go
│   └── rate_limit.go
├── services/
│   ├── cart/
│   │   └── cart_service.go
│   ├── merchant/
│   │   └── merchant_service.go
│   ├── notifications/
│   │   └── notifcation_service.go
│   ├── order/
│   │   └── order_service.go
│   ├── payment/
│   │   └── payment_service.go
│   ├── payout/
│   │   └── payout_service.go
│   ├── pricing/
│   │   └── pricing_service.go
│   ├── product/
│   │   ├── product_service.go
│   │   └── service_product.go
│   └── user/
│       └── user_service.go
├── tests/
│   └── unit/
│       ├── cart_service_test.go
│       ├── test_handlers.go
│       └── test_service.go
└── utils/
    ├── blacklist.go
    └── redis.go
```
---

## 📄 File Contents
### api/dto/cart.go
- Size: 1.76 KB
- Lines: 47
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="070984c3-63cc-4d18-b894-dd310296bd6a" artifact_version_id="ba518519-62cc-46db-975c-559c8fd511f3" title="api/dto/cart.go" contentType="text/go">
```go
package dto

import (
	"api-customer-merchant/internal/db/models" // For enums if needed
	"time"
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
	ID        string             `json:"id"`
	UserID    uint               `json:"user_id,omitempty"`
	Status    models.CartStatus  `json:"status"`
	Items     []CartItemResponse `json:"items"`
	Total     float64            `json:"total,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"` // Added
	UpdatedAt time.Time          `json:"updated_at,omitempty"` // Added
}

type CartItemResponse struct {
    ID        string  `json:"id"`
    ProductID string  `json:"product_id"`
    Product   *ProductResponse // Embed lightweight product (from dto/product.go)
    Quantity  int     `json:"quantity"`
    Subtotal  float64 `json:"subtotal"` // Quantity * Product.BasePrice
}
```
</xaiArtifact>

---
### api/dto/order.go
- Size: 0.89 KB
- Lines: 25
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="5cf210fd-971b-44dc-a50d-a096c7a055ea" artifact_version_id="08ddfee9-b9f1-418c-9536-96da7cc217e8" title="api/dto/order.go" contentType="text/go">
```go
package dto

//import "api-customer-merchant/internal/db/models"

// CreateOrderRequest defines the request body for creating an order.
type CreateOrderRequest struct {
	UserID uint `json:"user_id"` // The ID of the user placing the order.
	//UserID uint `json:"user_id"` // The ID of the user placing the order
}

// OrderResponse defines the structure for order-related responses.
type OrderResponse struct {
	ID            uint            `json:"id"`
	UserID        uint            `json:"user_id"`
	Status        string          `json:"status"` // Consider using the models.OrderStatus type directly
	OrderItems    []OrderItemResponse `json:"order_items"`
}

// OrderItemResponse defines the structure for individual items in an order.
type OrderItemResponse struct {
	ProductID         string  `json:"product_id"`
	Quantity          int     `json:"quantity"`
	Price             float64 `json:"price"`
}


```
</xaiArtifact>

---
### api/dto/product.go
- Size: 3.05 KB
- Lines: 78
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="17146f70-e913-400e-99c6-77256eed4894" artifact_version_id="cc587f7b-68ff-4428-8904-0f52fb7cd395" title="api/dto/product.go" contentType="text/go">
```go
package dto

import (
	"time"

	//"github.com/shopspring/decimal"
)

// ProductInput represents the request body for creating a product
type ProductInput struct {
	Name         string          `json:"name" validate:"required,max=255"`
	MerchantID    string          `json:"merchant_id" validate:"required,max=255"`
	Description  string          `json:"description" validate:"max=1000"`
	SKU          string          `json:"sku" validate:"required,max=100"`
	BasePrice    float64 `json:"base_price" validate:"required,gt=0"`
	CategoryID   uint            `json:"category_id" validate:"required"`
	InitialStock *int            `json:"initial_stock" validate:"omitempty,gte=0"` // For simple products
	Variants     []VariantInput  `json:"variants,omitempty" validate:"dive,omitempty"`
	Media        []MediaInput    `json:"media,omitempty" validate:"dive,omitempty"`
}

type VariantInput struct {
	SKU             string         `json:"sku" validate:"required,max=100"`
	PriceAdjustment float64 `json:"price_adjustment" validate:"gte=0"`
	Attributes      map[string]string `json:"attributes" validate:"required,dive,required"`
	InitialStock    int            `json:"initial_stock" validate:"gte=0"`
}

type MediaInput struct {
	URL  string `json:"url" validate:"required,url,max=500"`
	Type string `json:"type" validate:"required,oneof=image video"`
}

// ProductResponse for API output
type ProductResponse struct {
	ID           string            `json:"id"`
	MerchantID   string            `json:"merchant_id"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	SKU          string            `json:"sku"`
	BasePrice    float64   `json:"base_price"`
	CategoryID   uint              `json:"category_id"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	Variants     []VariantResponse  `json:"variants,omitempty"`
	Media        []MediaResponse    `json:"media,omitempty"`
	SimpleInventory *InventoryResponse `json:"simple_inventory,omitempty"` // For simple products
}

type VariantResponse struct {
	ID             string         `json:"id"`
	ProductID      string         `json:"product_id"`
	SKU            string         `json:"sku"`
	PriceAdjustment float64 `json:"price_adjustment"`
	TotalPrice     float64 `json:"total_price"`
	Attributes     map[string]string `json:"attributes"`
	IsActive       bool           `json:"is_active"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	Inventory      InventoryResponse `json:"inventory"`
}

type MediaResponse struct {
	ID        string    `json:"id"`
	ProductID string    `json:"product_id"`
	URL       string    `json:"url"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InventoryResponse struct {
	ID               string `json:"id"`
	Quantity         int    `json:"quantity"`
	ReservedQuantity int    `json:"reserved_quantity"`
	LowStockThreshold int   `json:"low_stock_threshold"`
	BackorderAllowed bool   `json:"backorder_allowed"`
}
```
</xaiArtifact>

---
### api/handlers/cart_handler.go
- Size: 6.24 KB
- Lines: 231
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="1d5b9ba3-085d-40e0-8e57-a0778f903494" artifact_version_id="784d282c-a909-4bd8-b1be-909aef503469" title="api/handlers/cart_handler.go" contentType="text/go">
```go
//package handlers
/*
import (
	//"net/http"
	//"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/api/dto"
	"api-customer-merchant/internal/services/cart"
	"net/http"
	"strconv"

	//"api-customer-merchant/internal/domain/order"
	//"api-customer-merchant/internal/domain/payment"
	"api-customer-merchant/internal/services/product"
	//"strconv"
	"github.com/gin-gonic/gin"
)

type CartHandlers struct {
	productService *product.ProductService
	cartService    *cart.CartService
	// orderService   *order.OrderService
	// paymentService *payment.PaymentService
}

func NewCartHandlers(productService *product.ProductService, cartService *cart.CartService) *CartHandlers {
	return &CartHandlers{
		productService: productService,
		cartService:    cartService,
		// orderService:   orderService,
		// paymentService: paymentService,
	}
}





func (h *CartHandlers) AddItemToCart(c *gin.Context) {
	// userID, exists := c.Get("userID")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	// 	return
	// }

	var req dto.AddItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cart, err := h.cartService.AddItemToCart(req.UserID, req.ProductID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cart)
}


// UpdateCartItemQuantity handles PUT /customer/cart/update/:cartItemID
func (h *CartHandlers) UpdateCartItemQuantity(c *gin.Context) {
	// userID, exists := c.Get("userID")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	// 	return
	// }

	cartItemIDStr := c.Param("cartItemID")
	cartItemID, err := strconv.ParseUint(cartItemIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart item ID"})
		return
	}
	// Verify cart item belongs to user's active cart
	cartItem, err := h.cartService.GetCartItemByID(uint(cartItemID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cart item not found"})
		return
	}

	var req dto.UpdateItemRequest 
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cart, err := h.cartService.GetActiveCart(req.UserID)
	if err != nil || cart.ID != cartItem.CartID {
		c.JSON(http.StatusForbidden, gin.H{"error": "cart item does not belong to user"})
		return
	}

	updatedCart, err := h.cartService.UpdateCartItemQuantity(uint(cartItemID), req.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCart)
}



// RemoveCartItem handles DELETE /customer/cart/remove/:cartItemID
func (h *CartHandlers) RemoveCartItem(c *gin.Context) {
	 userID, exists := c.Get("userID")
	 if !exists {
	 	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	 	return
	 }

	cartItemIDStr := c.Param("cartItemID")
	cartItemID, err := strconv.ParseUint(cartItemIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart item ID"})
		return
	}
	// Verify cart item belongs to user's active cart
	cartItem, err := h.cartService.GetCartItemByID(uint(cartItemID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cart item not found"})
		return
	}
	 cart, err := h.cartService.GetActiveCart(userID.(uint))
	 if err != nil || cart.ID != cartItem.CartID {
	 	c.JSON(http.StatusForbidden, gin.H{"error": "cart item does not belong to user"})
	 	return
	 }
	

	updatedCart, err := h.cartService.RemoveCartItem(uint(cartItemID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCart)
}
*/

package handlers

import (
	"net/http"
	"strconv"

	"api-customer-merchant/internal/services/cart" // Assuming service import
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartService *cart.CartService
}

func NewCartHandler(cartService *cart.CartService) *CartHandler {
	return &CartHandler{cartService: cartService}
}

// AddToCart handles adding an item to the cart
func (h *CartHandler) AddToCart(c *gin.Context) {
	ctx := c.Request.Context()
	userIDStr := c.Query("user_id") // For testing, get from query/body
	userID, _ := strconv.ParseUint(userIDStr, 10, 32)
	productID := c.Query("product_id")
	quantityStr := c.Query("quantity")
	quantity, _ := strconv.ParseUint(quantityStr, 10, 32)

	updatedCart, err := h.cartService.AddItemToCart(ctx, uint(userID), uint(quantity), productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedCart)
}

// GetCartItem handles getting a cart item
func (h *CartHandler) GetCartItem(c *gin.Context) {
	ctx := c.Request.Context()
	itemIDStr := c.Param("id")
	itemID, _ := strconv.ParseUint(itemIDStr, 10, 32)

	item, err := h.cartService.GetCartItemByID(ctx, uint(itemID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// GetCart handles getting the active cart
func (h *CartHandler) GetCart(c *gin.Context) {
	ctx := c.Request.Context()
	userIDStr := c.Query("user_id")
	userID, _ := strconv.ParseUint(userIDStr, 10, 32)

	cart, err := h.cartService.GetActiveCart(ctx, uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}

// UpdateCartItemQuantity handles updating quantity
func (h *CartHandler) UpdateCartItemQuantity(c *gin.Context) {
	ctx := c.Request.Context()
	itemIDStr := c.Param("id")
	itemID, _ := strconv.ParseUint(itemIDStr, 10, 32)
	quantityStr := c.Query("quantity")
	quantity, _ := strconv.Atoi(quantityStr)

	updatedCart, err := h.cartService.UpdateCartItemQuantity(ctx, uint(itemID), quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedCart)
}

// RemoveCartItem handles removing an item
func (h *CartHandler) RemoveCartItem(c *gin.Context) {
	ctx := c.Request.Context()
	itemIDStr := c.Param("id")
	itemID, _ := strconv.ParseUint(itemIDStr, 10, 32)

	updatedCart, err := h.cartService.RemoveCartItem(ctx, uint(itemID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedCart)
}
```
</xaiArtifact>

---
### api/handlers/customer_auth_handler.go
- Size: 6.12 KB
- Lines: 198
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="012a26fb-02a9-4fdd-a31d-61e16e3c9c9d" artifact_version_id="f3bac88f-e6ab-4eac-8295-8cc63bf0dbea" title="api/handlers/customer_auth_handler.go" contentType="text/go">
```go
package handlers

import (
	"log"
	"net/http"
	"os"

	//"os"
	"strings"

	//"api-customer-merchant/internal/db/models"
	//"api-customer-merchant/internal/db/repositories"
	services "api-customer-merchant/internal/services/user"
	"api-customer-merchant/internal/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	service *services.AuthService
}

// In customer/handlers/auth_handler.go AND merchant/handlers/auth_handler.go
func NewAuthHandler(s *services.AuthService) *AuthHandler {
    return &AuthHandler{
        service: s,
    }
}
// Register godoc
// @Summary Register a new customer
// @Description Creates a new customer account with email, name, password, and optional country
// @Tags Customer
// @Accept json
// @Produce json
// @Param body body object{email=string,name=string,password=string,country=string} true "Customer registration details"
// @Success 200 {object} object{token=string} "JWT token"
// @Failure 400 {object} object{error=string} "Invalid request"
// @Failure 500 {object} object{error=string} "Server error"
// @Router /customer/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
		Country  string `json:"country"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.RegisterUser(req.Email, req.Name, req.Password, req.Country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}


// Login godoc
// @Summary Customer login
// @Description Authenticates a customer using email and password
// @Tags Customer
// @Accept json
// @Produce json
// @Param body body object{email=string,password=string} true "Customer login credentials"
// @Success 200 {object} object{token=string} "JWT token"
// @Failure 400 {object} object{error=string} "Invalid request"
// @Failure 401 {object} object{error=string} "Unauthorized"
// @Failure 403 {object} object{error=string} "Invalid role"
// @Failure 500 {object} object{error=string} "Server error"
// @Router /customer/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.LoginUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// if user.Role != "customer" {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "Invalid role for this API"})
	// 	return
	// }

	token, err := h.service.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}


// GoogleAuth godoc
// @Summary Initiate Google OAuth for customer
// @Description Redirects to Google OAuth login page
// @Tags Customer
// @Produce json
// @Success 307 {object} object{} "Redirect to Google OAuth"
// @Router /customer/auth/google [get]
func (h *AuthHandler) GoogleAuth(c *gin.Context) {
	url := h.service.GetOAuthConfig("customer").AuthCodeURL("state-customer", oauth2.AccessTypeOffline)
	//c.Redirect(http.StatusTemporaryRedirect, url)
	 c.JSON(http.StatusOK, gin.H{"url": url})
}



// GoogleCallback godoc
// @Summary Handle Google OAuth callback for customer
// @Description Processes Google OAuth callback and returns JWT token
// @Tags Customer
// @Produce json
// @Param code query string true "OAuth code"
// @Success 200 {object} object{token=string} "JWT token"
// @Failure 400 {object} object{error=string} "Code not provided"
// @Failure 500 {object} object{error=string} "Server error"
// @Router /customer/auth/google/callback [get]
func (h *AuthHandler) GoogleCallback(c *gin.Context) {
    code := c.Query("code")
    state := c.Query("state")
    if code == "" || state == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code or state"})
        return
    }
    // Verify state (in production, check against stored value)
    if state != "state-customer" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state parameter"})
        return
    }

    user, token, err := h.service.GoogleLogin(code, os.Getenv("BASE_URL"), "customer")
    if err != nil {
        log.Printf("Google login failed: %v", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"token": token, "user": user})
}



// Logout godoc
// @Summary Customer logout
// @Description Invalidates the customer's JWT token
// @Tags Customer
// @Security BearerAuth
// @Produce json
// @Success 200 {object} object{message=string} "Logout successful"
// @Failure 400 {object} object{error=string} "Authorization header required"
// @Router /customer/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header required"})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	utils.Add(tokenString)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (h *AuthHandler) UpdateProfile(c *gin.Context) {
    userID, _ := c.Get("userID")
    var req struct { Name string; Country string; Addresses []string }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.service.UpdateProfile(userID.(uint), req.Name, req.Country, req.Addresses); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "updated"})
}
```
</xaiArtifact>

---
### api/handlers/customer_handlers.go
- Size: 9.14 KB
- Lines: 336
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="0c814b6f-2690-44cb-b7e1-695e3c9a77dd" artifact_version_id="11cbcdb2-cbb2-46ba-b366-e09821c3ebf5" title="api/handlers/customer_handlers.go" contentType="text/go">
```go
package handlers
/*
import (
	"net/http"
	//"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/domain/cart"
	"api-customer-merchant/internal/domain/order"
	"api-customer-merchant/internal/domain/payment"
	"api-customer-merchant/internal/domain/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHandlers struct {
	productService *product.ProductService
	cartService    *cart.CartService
	orderService   *order.OrderService
	paymentService *payment.PaymentService
}

func NewCustomerHandlers(productService *product.ProductService, cartService *cart.CartService, orderService *order.OrderService, paymentService *payment.PaymentService) *CustomerHandlers {
	return &CustomerHandlers{
		productService: productService,
		cartService:    cartService,
		orderService:   orderService,
		paymentService: paymentService,
	}
}

// GetProductByID handles GET /customer/products/:id
func (h *CustomerHandlers) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetProductBySKU handles GET /customer/products/sku/:sku
func (h *CustomerHandlers) GetProductBySKU(c *gin.Context) {
	sku := c.Param("sku")
	if sku == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SKU cannot be empty"})
		return
	}

	product, err := h.productService.GetProductBySKU(sku)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// SearchProductsByName handles GET /customer/products/search?name={name}
func (h *CustomerHandlers) SearchProductsByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "search name cannot be empty"})
		return
	}

	products, err := h.productService.SearchProductsByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// AddItemToCart handles POST /customer/cart/add
func (h *CustomerHandlers) AddItemToCart(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  int  `json:"quantity" binding:"required,gt=0"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart, err := h.cartService.AddItemToCart(userID.(uint), req.ProductID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cart)
}

// UpdateCartItemQuantity handles PUT /customer/cart/update/:cartItemID
func (h *CustomerHandlers) UpdateCartItemQuantity(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	cartItemIDStr := c.Param("cartItemID")
	cartItemID, err := strconv.ParseUint(cartItemIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart item ID"})
		return
	}
	// Verify cart item belongs to user's active cart
	cartItem, err := h.cartService.GetCartItemByID(uint(cartItemID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cart item not found"})
		return
	}
	cart, err := h.cartService.GetActiveCart(userID.(uint))
	if err != nil || cart.ID != cartItem.CartID {
		c.JSON(http.StatusForbidden, gin.H{"error": "cart item does not belong to user"})
		return
	}

	var req struct {
		Quantity int `json:"quantity" binding:"required,gt=0"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCart, err := h.cartService.UpdateCartItemQuantity(uint(cartItemID), req.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCart)
}

// RemoveCartItem handles DELETE /customer/cart/remove/:cartItemID
func (h *CustomerHandlers) RemoveCartItem(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	cartItemIDStr := c.Param("cartItemID")
	cartItemID, err := strconv.ParseUint(cartItemIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart item ID"})
		return
	}
	// Verify cart item belongs to user's active cart
	cartItem, err := h.cartService.GetCartItemByID(uint(cartItemID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cart item not found"})
		return
	}
	cart, err := h.cartService.GetActiveCart(userID.(uint))
	if err != nil || cart.ID != cartItem.CartID {
		c.JSON(http.StatusForbidden, gin.H{"error": "cart item does not belong to user"})
		return
	}
	

	updatedCart, err := h.cartService.RemoveCartItem(uint(cartItemID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCart)
}

// CreateOrder handles POST /customer/orders/create
func (h *CustomerHandlers) CreateOrder(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	order, err := h.orderService.CreateOrderFromCart(userID.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// GetOrderByID handles GET /customer/orders/:id
func (h *CustomerHandlers) GetOrderByID(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	order, err := h.orderService.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	if order.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "order does not belong to user"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// GetOrders handles GET /customer/orders
func (h *CustomerHandlers) GetOrders(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	orders, err := h.orderService.GetOrdersByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// ProcessPayment handles POST /customer/payments/process
func (h *CustomerHandlers) ProcessPayment(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req struct {
		OrderID uint    `json:"order_id" binding:"required"`
		Amount  float64 `json:"amount" binding:"required,gt=0"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify order belongs to user
	order, err := h.orderService.GetOrderByID(req.OrderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	if order.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "order does not belong to user"})
		return
	}

	payment, err := h.paymentService.ProcessPayment(req.OrderID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payment)
}

// GetPaymentByOrderID handles GET /customer/payments/:orderID
func (h *CustomerHandlers) GetPaymentByOrderID(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	orderIDStr := c.Param("orderID")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	// Verify order belongs to user
	order, err := h.orderService.GetOrderByID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	if order.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "order does not belong to user"})
		return
	}

	payment, err := h.paymentService.GetPaymentByOrderID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

func (h *CustomerHandlers) GetProductsByCategory(c *gin.Context) {
    categoryIDStr := c.Param("categoryID")
    categoryID, _ := strconv.ParseUint(categoryIDStr, 10, 32)
    limit, _ := strconv.Atoi(c.Query("limit"))
    offset, _ := strconv.Atoi(c.Query("offset"))
    products, err := h.productService.GetProductsByCategory(uint(categoryID), limit, offset)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, products)
}

*/
```
</xaiArtifact>

---
### api/handlers/merchant_auth_handler.go
- Size: 11.58 KB
- Lines: 247
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="babee554-dcf4-480b-9348-53c4a2dff365" artifact_version_id="6e085399-21fe-4bc7-ae9f-5c8c93fba44d" title="api/handlers/merchant_auth_handler.go" contentType="text/go">
```go
package handlers

import (
	"encoding/json"
	"net/http"

	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/services/merchant"

	"github.com/gin-gonic/gin"
)
/*
type MerchantHandler struct {
	service *merchant.MerchantService
}

func NewMerchantAuthHandler(s *merchant.MerchantService) *MerchantHandler {
	return &MerchantHandler{service: s}
}

// Apply godoc
// @Summary Submit a new merchant application
// @Description Allows a prospective merchant to submit an application with personal, business, and address information
// @Tags Merchant
// @Accept json
// @Produce json
// @Param body body object{first_name=string,last_name=string,email=string,phone=string,personal_address=object{street=string,city=string,state=string,postal_code=string,country=string},work_address=object{street=string,city=string,state=string,postal_code=string,country=string},business_name=string,business_type=string,tax_id=string,documents=object{business_license=string,identification=string}} true "Merchant application details"
// @Success 201 {object} object{id=string,first_name=string,last_name=string,email=string,phone=string,personal_address=object{street=string,city=string,state=string,postal_code=string,country=string},work_address=object{street=string,city=string,state=string,postal_code=string,country=string},business_name=string,business_type=string,tax_id=string,documents=object{business_license=string,identification=string},status=string,created_at=string} "Created application"
// @Failure 400 {object} object{error=string} "Invalid request body or malformed JSON"
// @Failure 500 {object} object{error=string} "Failed to submit application"
// @Router /merchant/apply [post]
func (h *MerchantHandler) Apply(c *gin.Context) {
	var req struct {
		models.MerchantBasicInfo
		PersonalAddress            map[string]any `json:"personal_address" validate:"required"`
		WorkAddress                map[string]any `json:"work_address" validate:"required"`
		models.MerchantBusinessInfo
		models.MerchantDocuments
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	// Convert personal_address and work_address to JSONB
	personalAddressJSON, err := json.Marshal(req.PersonalAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid personal_address JSON: " + err.Error()})
		return
	}
	workAddressJSON, err := json.Marshal(req.WorkAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid work_address JSON: " + err.Error()})
		return
	}

	app := &models.MerchantApplication{
		MerchantBasicInfo:    req.MerchantBasicInfo,
		MerchantAddress:      models.MerchantAddress{PersonalAddress: personalAddressJSON, WorkAddress: workAddressJSON},
		MerchantBusinessInfo: req.MerchantBusinessInfo,
		MerchantDocuments:    req.MerchantDocuments,
	}

	app, err = h.service.SubmitApplication(c.Request.Context(), app)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to submit application: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, app)
}

// GetApplication godoc
// @Summary Retrieve a merchant application by ID
// @Description Allows an applicant to view the status and details of their merchant application
// @Tags Merchant
// @Produce json
// @Param id path string true "Application ID" format(uuid)
// @Success 200 {object} object{id=string,first_name=string,last_name=string,email=string,phone=string,personal_address=object{street=string,city=string,state=string,postal_code=string,country=string},work_address=object{street=string,city=string,state=string,postal_code=string,country=string},business_name=string,business_type=string,tax_id=string,documents=object{business_license=string,identification=string},status=string,created_at=string} "Application details"
// @Failure 400 {object} object{error=string} "Invalid application ID format"
// @Failure 404 {object} object{error=string} "Application not found"
// @Router /merchant/application/{id} [get]
func (h *MerchantHandler) GetApplication(c *gin.Context) {
	id := c.Param("id")
	app, err := h.service.GetApplication(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "application not found: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, app)
}

// GetMyMerchant godoc
// @Summary Retrieve current merchant account
// @Description Fetches the merchant account details for the authenticated user, if their application has been approved
// @Tags Merchant
// @Produce json
// @Security BearerAuth
// @Success 200 {object} object{id=string,user_id=string,business_name=string,business_type=string,tax_id=string,personal_address=object{street=string,city=string,state=string,postal_code=string,country=string},work_address=object{street=string,city=string,state=string,postal_code=string,country=string},status=string,created_at=string,updated_at=string} "Merchant account details"
// @Failure 401 {object} object{error=string} "Unauthorized: Missing or invalid authentication"
// @Failure 404 {object} object{error=string} "Merchant account not found"
// @Router /merchant/me [get]
func (h *MerchantHandler) GetMyMerchant(c *gin.Context) {
	userID, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	m, err := h.service.GetMerchantByUserID(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "merchant not found: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, m)
}
*/


type MerchantHandler struct {
	service *merchant.MerchantService
}

func NewMerchantAuthHandler(s *merchant.MerchantService) *MerchantHandler {
	return &MerchantHandler{service: s}
}

// Apply godoc
// @Summary Submit a new merchant application
// @Description Allows a prospective merchant to submit an application with personal, business, and address information
// @Tags Merchant
// @Accept json
// @Produce json
// @Param body body object{first_name=string,last_name=string,email=string,phone=string,personal_address=object{street=string,city=string,state=string,postal_code=string,country=string},work_address=object{street=string,city=string,state=string,postal_code=string,country=string},business_name=string,business_type=string,tax_id=string,documents=object{business_license=string,identification=string}} true "Merchant application details"
// @Success 201 {object} object{id=string,first_name=string,last_name=string,email=string,phone=string,personal_address=object{street=string,city=string,state=string,postal_code=string,country=string},work_address=object{street=string,city=string,state=string,postal_code=string,country=string},business_name=string,business_type=string,tax_id=string,documents=object{business_license=string,identification=string},status=string,created_at=string} "Created application"
// @Failure 400 {object} object{error=string} "Invalid request body or malformed JSON"
// @Failure 500 {object} object{error=string} "Failed to submit application"
// @Router /merchant/apply [post]
func (h *MerchantHandler) Apply(c *gin.Context) {
	var req struct {
		models.MerchantBasicInfo
		PersonalAddress            map[string]any `json:"personal_address" validate:"required"`
		WorkAddress                map[string]any `json:"work_address" validate:"required"`
		models.MerchantBusinessInfo
		models.MerchantDocuments
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	// Convert personal_address and work_address to JSONB
	personalAddressJSON, err := json.Marshal(req.PersonalAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid personal_address JSON: " + err.Error()})
		return
	}
	workAddressJSON, err := json.Marshal(req.WorkAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid work_address JSON: " + err.Error()})
		return
	}

	app := &models.MerchantApplication{
		MerchantBasicInfo:    req.MerchantBasicInfo,
		MerchantAddress:      models.MerchantAddress{PersonalAddress: personalAddressJSON, WorkAddress: workAddressJSON},
		MerchantBusinessInfo: req.MerchantBusinessInfo,
		MerchantDocuments:    req.MerchantDocuments,
	}

	app, err = h.service.SubmitApplication(c.Request.Context(), app)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to submit application: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, app)
}


func (h *MerchantHandler) Login(c *gin.Context) {
	var req struct {
		Work_Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	merchant, err := h.service.LoginMerchant(c.Request.Context(),req.Work_Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.GenerateJWT(merchant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}


// GetApplication godoc
// @Summary Retrieve a merchant application by ID
// @Description Fetches the details and status of a submitted merchant application (e.g., 'pending', 'approved', 'rejected')
// @Tags Merchant
// @Produce json
// @Param id path string true "Application ID (UUID)"
// @Success 200 {object} object{id=string,status=string,merchant_basic_info=object{name=string,store_name=string,personal_email=string,work_email=string,password=string},personal_address=object{street=string,city=string,state=string,zip=string,country=string},work_address=object{street=string,city=string,state=string,zip=string,country=string},merchant_business_info=object{business_type=string,years_in_business=integer,annual_revenue=number,tax_id=string},merchant_documents=object{id_proof=string,business_license=string,bank_statement=string},created_at=string,updated_at=string} "Application details retrieved"
// @Failure 404 {object} object{error=string} "Application not found"
// @Failure 500 {object} object{error=string} "Failed to retrieve application"
// @Router /merchant/application/{id} [get]
func (h *MerchantHandler) GetApplication(c *gin.Context) {
	id := c.Param("id")
	app, err := h.service.GetApplication(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "application not found: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, app)
}

// GetMyMerchant godoc
// @Summary Retrieve current merchant account
// @Description Fetches the merchant account details for the authenticated user, if their application has been approved
// @Tags Merchant
// @Produce json
// @Security BearerAuth
// @Success 200 {object} object{id=string,user_id=string,business_name=string,business_type=string,tax_id=string,personal_address=object{street=string,city=string,state=string,zip=string,country=string},work_address=object{street=string,city=string,state=string,zip=string,country=string},status=string,created_at=string,updated_at=string} "Merchant account details"
// @Failure 401 {object} object{error=string} "Unauthorized: Missing or invalid authentication"
// @Failure 404 {object} object{error=string} "Merchant account not found"
// @Router /merchant/me [get]
func (h *MerchantHandler) GetMyMerchant(c *gin.Context) {
	userID, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	m, err := h.service.GetMerchantByUserID(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "merchant not found: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, m)
}
```
</xaiArtifact>

---
### api/handlers/merchant_handlers.go
- Size: 12.94 KB
- Lines: 519
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="f32d3d28-c3aa-44b5-9a08-5d5b818ed69b" artifact_version_id="dd1164df-2dd3-49f6-8c7e-307a0b01e23e" title="api/handlers/merchant_handlers.go" contentType="text/go">
```go
package handlers

// import (
// 	"api-customer-merchant/internal/db/models"
// 	"encoding/csv"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"time"

// 	//"api-customer-merchant/internal/domain/order"
// 	//"api-customer-merchant/internal/domain/payout"
// 	"api-customer-merchant/internal/services/product"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

/*
type MerchantHandlers struct {
	productService *product.ProductService
	orderService   *order.OrderService
	payoutService  *payout.PayoutService
}

func NewMerchantHandlers(productService *product.ProductService, orderService *order.OrderService, payoutService *payout.PayoutService) *MerchantHandlers {
	return &MerchantHandlers{
		productService: productService,
		orderService:   orderService,
		payoutService:  payoutService,
	}
}

// CreateProduct handles POST /merchant/products
func (h *MerchantHandlers) CreateProduct(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.productService.CreateProduct(&product, merchantID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct handles PUT /merchant/products/:id
func (h *MerchantHandlers) UpdateProduct(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = uint(id)

	if err := h.productService.UpdateProduct(&product, merchantID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct handles DELETE /merchant/products/:id
func (h *MerchantHandlers) DeleteProduct(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	if err := h.productService.DeleteProduct(uint(id), merchantID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}

// GetOrders handles GET /merchant/orders
func (h *MerchantHandlers) GetOrders(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	orders, err := h.orderService.GetOrdersByMerchantID(merchantID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GetPayouts handles GET /merchant/payouts
func (h *MerchantHandlers) GetPayouts(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	payouts, err := h.payoutService.GetPayoutsByMerchantID(merchantID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payouts)
}
*/


/*
type MerchantHandlers struct {
	productService *product.ProductService
	//orderService   *order.OrderService
	//payoutService  *payout.PayoutService
	//promotionService *promotions.PromotionService
}
//, orderService *order.OrderService, payoutService *payout.PayoutService,promotionService *promotions.PromotionService) *MerchantHandlers {
//orderService:   orderService,
		//payoutService:  payoutService,
		//promotionService:promotionService,
func NewMerchantHandlers(productService *product.ProductService)*MerchantHandlers{
	return &MerchantHandlers{
		productService: productService,
	}
}

func (h *MerchantHandlers) CreateProduct(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	merchantIDStr := merchantID.(string)

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.productService.CreateProduct(&product, merchantIDStr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}


 func (h *MerchantHandlers) UpdateProduct(c *gin.Context) {
 	merchantID, exists := c.Get("merchantID")
 	if !exists {
 		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
 		return
 	}
	merchantIDStr := merchantID.(string)

 	idStr := c.Param("id")
 	//id, err := strconv.ParseUint(idStr, 10, 32)
 	// if err != nil {
 	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
 	// 	return
 	// }

	var product models.Product
 	if err := c.ShouldBindJSON(&product); err != nil {
 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
 		return
 	}
 	product.ID = idStr

 	if err := h.productService.UpdateProduct(&product, merchantIDStr); err != nil {
 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
 		return
 	}

 	c.JSON(http.StatusOK, product)
 }







// func (h *MerchantHandler) UpdateProduct(c *gin.Context) {
// 	merchantID, exists := c.Get("merchantID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
// 		return
// 	}
// 	merchantIDStr := merchantID.(string)

// 	id := c.Param("id")

// 	var product models.Product
// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	product.ID = id
// 	product.UpdatedAt = time.Now()

// 	if err := h.productService.UpdateProduct(&product, merchantIDStr); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, product)
// }



func (h *MerchantHandlers) GetMyProducts(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	merchantIDStr := merchantID.(string)

	//products, err := h.productService.GetProductsByMerchantID(merchantIDStr)
	products, err := h.productService.GetProductsByMerchantID(merchantIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusOK, products)
}

























func (h *MerchantHandlers) DeleteProduct(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	merchantIDStr := merchantID.(string)


	idStr := c.Param("id")
	///id, err := strconv.ParseUint(idStr, 10, 32)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
	// 	return
	// }
	id:=idStr

	if err := h.productService.DeleteProduct(id,merchantIDStr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
/*
// GetOrders handles GET /merchant/orders
func (h *MerchantHandlers) GetOrders(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	orders, err := h.orderService.GetOrdersByMerchantID(merchantID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GetPayouts handles GET /merchant/payouts
func (h *MerchantHandlers) GetPayouts(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	payouts, err := h.payoutService.GetPayoutsByMerchantID(merchantID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payouts)


}


func (h *MerchantHandlers) CreatePromotion(c *gin.Context) {
    merchantID, _ := c.Get("merchantID")
    var promo models.Promotion
    if err := c.ShouldBindJSON(&promo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    promo.MerchantID = merchantID.(uint)
    if err := h.promotionService.CreatePromotion(&promo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, promo)
}

func (h *MerchantHandlers) UpdateOrderItemStatus(c *gin.Context) {
    merchantID, _ := c.Get("merchantID")
    itemIDStr := c.Param("itemID")
    itemID, _ := strconv.ParseUint(itemIDStr, 10, 32)
    var req struct { Status string `json:"status"` }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.orderService.UpdateOrderItemStatus(uint(itemID), models.FulfillmentStatus(req.Status), merchantID.(uint)); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "updated"})
}







 func (h *MerchantHandlers) BulkUploadProducts(c *gin.Context) {
     merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	merchantIDStr := merchantID.(string)
     file, err := c.FormFile("csv")
     if err != nil {
         c.JSON(http.StatusBadRequest, gin.H{"error": "csv file required"})
         return
     }
     f, err := file.Open()
     if err != nil {
         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
         return
     }
     defer f.Close()

     reader := csv.NewReader(f)
     // Skip header
     _, err = reader.Read()
     if err != nil {
         c.JSON(http.StatusBadRequest, gin.H{"error": "invalid csv header"})
         return
     }

     for {
         record, err := reader.Read()
        if err == io.EOF {
             break
         }
         if err != nil {
             c.JSON(http.StatusBadRequest, gin.H{"error": "invalid csv format"})
             return
         }

//         // Parse price with error handling
         price, err := strconv.ParseFloat(record[2], 64)
         if err != nil {
             c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid price format in CSV: %s", record[2])})
             return
         }

//         // Parse record: name, sku, price, etc.
         product := &models.Product{
             Name:       record[0],
             SKU:        record[1],
             Price:      price,
             // Add other fields as needed (e.g., CategoryID if required in CSV)
			 ID:        uuid.New().String(),
			CreatedAt: time.Now(),
             MerchantID:merchantIDStr,
         }

         if err := h.productService.CreateProduct(product, merchantIDStr); err != nil {
             c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to create product: %s", err.Error())})
             return
         }
     }
     c.JSON(http.StatusOK, gin.H{"message": "products uploaded"})
 }
*/


/*
func (h *MerchantHandler) BulkUploadProducts(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	merchantIDStr := merchantID.(string)

	file, err := c.FormFile("csv")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "csv file required"})
		return
	}
	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	_, err = reader.Read() // Skip header
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid csv header"})
		return
	}

	successCount := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid csv format"})
			return
		}

		price, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid price: %s", record[2])})
			return
		}

		product := &models.Product{
			Name:      record[0],
			SKU:       record[1],
			Price:     price,
			MerchantID: merchantIDStr,
			ID:        uuid.New().String(),
			CreatedAt: time.Now(),
		}

		if err := h.productService.CreateProduct(&product, merchantIDStr); err != nil {
			log.Printf("Failed to upload product %s: %v", record[0], err)
			continue // Continue on error (or abort if strict)
		}
		successCount++
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("uploaded %d products", successCount)})
}
*/
```
</xaiArtifact>

---
### api/handlers/order_handler.go
- Size: 1.47 KB
- Lines: 55
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="c89afef6-4ccd-41f4-9078-eefaf38a2a76" artifact_version_id="646a73ec-02d6-4bc9-9fae-bec0582b198e" title="api/handlers/order_handler.go" contentType="text/go">
```go
package handlers

import (
	"net/http"
	"strconv"

	"api-customer-merchant/internal/services/order"
	"github.com/gin-gonic/gin"
)

// OrderHandler manages order-related API requests.
type OrderHandler struct {
	orderService *order.OrderService
}

// NewOrderHandler creates a new OrderHandler instance.
func NewOrderHandler(orderService *order.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

// CreateOrder handles the creation of a new order.
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	userIDStr := c.Query("user_id") // For testing, get from query/body
	userID, _ := strconv.ParseUint(userIDStr, 10, 32)
	ctx := c.Request.Context()
	newOrder, err := h.orderService.CreateOrder(ctx, uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newOrder)
}

// GetOrder handles the request to retrieve a specific order by ID.
func (h *OrderHandler) GetOrder(c *gin.Context) {
	orderIDStr := c.Param("id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := h.orderService.GetOrder(c.Request.Context(), uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Adjust status code as needed
		return
	}

	if order == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}
```
</xaiArtifact>

---
### api/handlers/product_handler.go
- Size: 9.46 KB
- Lines: 291
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="254b4e98-33c9-4e79-a5c9-4892e91f88c2" artifact_version_id="81ebda49-f568-44a0-94b0-23b8784dde95" title="api/handlers/product_handler.go" contentType="text/go">
```go
package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"api-customer-merchant/internal/api/dto" // Assuming this exists for VariantInput
	
	"api-customer-merchant/internal/services/product"
)

type ProductHandler struct {
	productService *product.ProductService
	logger         *zap.Logger
	validator      *validator.Validate
}

func NewProductHandlers(productService *product.ProductService, logger *zap.Logger) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		logger:         logger,
		validator:      validator.New(),
	}
}

// CreateProduct handles product creation for a merchant
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	logger := h.logger.With(zap.String("operation", "CreateProduct"))

	// Check merchant authorization
	// merchantID, exists := c.Get("merchantID")
	// if !exists {
	// 	logger.Warn("Unauthorized access attempt")
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	// 	return
	// }
	// merchantIDStr, ok := merchantID.(string)
	// if !ok || merchantIDStr == "" {
	// 	logger.Warn("Invalid merchant ID in context")
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid merchant ID"})
	// 	return
	// }

	// Bind and validate input
	var input dto.ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}
	if err := h.validator.Struct(&input); err != nil {
		logger.Error("Input validation failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set merchant ID from context
	//input.MerchantID := merchantIDStr
	//merchantIDStr := input.MerchantID

	// Call service
	response, err := h.productService.CreateProductWithVariants(c.Request.Context(), &input)
	if err != nil {
		logger.Error("Failed to create product", zap.Error(err))
		if errors.Is(err, product.ErrInvalidProduct) || errors.Is(err, product.ErrInvalidSKU) ||
			errors.Is(err, product.ErrInvalidMediaURL) || errors.Is(err, product.ErrInvalidAttributes) ||
			errors.Is(err, product.ErrInvalidProduct) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}

	logger.Info("Product created successfully", zap.String("product_id", response.ID))
	c.JSON(http.StatusCreated, response)
}

// GetAllProducts handles fetching paginated products for the landing page
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	logger := h.logger.With(zap.String("operation", "GetAllProducts"))

	// Parse query parameters
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil || limit <= 0 || limit > 100 {
		limit = 20
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		offset = 0
	}
	var categoryID *uint
	if catIDStr := c.Query("category_id"); catIDStr != "" {
		catID, err := strconv.ParseUint(catIDStr, 10, 32)
		if err == nil {
			id := uint(catID)
			categoryID = &id
		} else {
			logger.Error("Invalid category ID", zap.String("category_id", catIDStr))
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
			return
		}
	}

	// Call service
	products, total, err := h.productService.GetAllProducts(c.Request.Context(), limit, offset, categoryID)
	if err != nil {
		logger.Error("Failed to fetch products", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}

	logger.Info("Products fetched successfully", zap.Int("count", len(products)), zap.Int64("total", total))
	c.JSON(http.StatusOK, gin.H{
		"products": products,
		"total":    total,
		"limit":    limit,
		"offset":   offset,
	})
}

// GetProductByID fetches a single product by ID
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	logger := h.logger.With(zap.String("operation", "GetProductByID"))
	productID := c.Param("id")
	if productID == "" {
		logger.Error("Missing product ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "product ID required"})
		return
	}

	// Call service with preloads
	response, err := h.productService.GetProductByID(c.Request.Context(), productID, "Media", "Variants", "Variants.Inventory", "SimpleInventory")
	if err != nil {
		logger.Error("Failed to fetch product", zap.Error(err), zap.String("product_id", productID))
		if errors.Is(err, product.ErrInvalidProduct) {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch product"})
		return
	}

	logger.Info("Product fetched successfully", zap.String("product_id", productID))
	c.JSON(http.StatusOK, response)
}

// ListProductsByMerchant lists a merchant's products with pagination
func (h *ProductHandler) ListProductsByMerchant(c *gin.Context) {
	logger := h.logger.With(zap.String("operation", "ListProductsByMerchant"))

	// Check merchant authorization
	merchantID:= c.Param("id")
	if merchantID=="" {
		logger.Warn("Unauthorized access attempt")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// merchantIDStr, ok := merchantID.(string)
	// if !ok || merchantIDStr == "" {
	// 	logger.Warn("Invalid merchant ID in context")
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid merchant ID"})
	// 	return
	// }

	// Parse query parameters
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil || limit <= 0 || limit > 100 {
		limit = 20
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		offset = 0
	}
	activeOnly := c.Query("active_only") == "true"

	// Call service
	products, err := h.productService.ListProductsByMerchant(c.Request.Context(), merchantID, limit, offset, activeOnly)
	if err != nil {
		logger.Error("Failed to list products", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list products"})
		return
	}

	logger.Info("Products listed successfully", zap.Int("count", len(products)))
	c.JSON(http.StatusOK, gin.H{
		"products": products,
		"total":    len(products), // Note: Repository doesn't return total; add if needed
		"limit":    limit,
		"offset":   offset,
	})
}

// UpdateInventory adjusts stock for a given inventory ID
func (h *ProductHandler) UpdateInventory(c *gin.Context) {
	logger := h.logger.With(zap.String("operation", "UpdateInventory"))
	inventoryID := c.Param("id")
	if inventoryID == "" {
		logger.Error("Missing inventory ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "inventory ID required"})
		return
	}

	// Check merchant authorization
	// merchantID, exists := c.Get("merchantID")
	// if !exists {
	// 	logger.Warn("Unauthorized access attempt")
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	// 	return
	// }
	// merchantIDStr, ok := merchantID.(string)
	// if !ok || merchantIDStr == "" {
	// 	logger.Warn("Invalid merchant ID in context")
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid merchant ID"})
	// 	return
	// }

	// Bind input
	var input struct {
		Delta int `json:"delta" validate:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}
	if err := h.validator.Struct(&input); err != nil {
		logger.Error("Input validation failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service
	err := h.productService.UpdateInventory(c.Request.Context(), inventoryID, input.Delta)
	if err != nil {
		logger.Error("Failed to update inventory", zap.Error(err), zap.String("inventory_id", inventoryID))
		if errors.Is(err, product.ErrInvalidProduct) {
			c.JSON(http.StatusNotFound, gin.H{"error": "inventory not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update inventory"})
		return
	}

	logger.Info("Inventory updated successfully", zap.String("inventory_id", inventoryID), zap.Int("delta", input.Delta))
	c.JSON(http.StatusOK, gin.H{"message": "inventory updated"})
}

// DeleteProduct soft-deletes a product
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	logger := h.logger.With(zap.String("operation", "DeleteProduct"))
	productID := c.Param("id")
	if productID == "" {
		logger.Error("Missing product ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "product ID required"})
		return
	}

	// Check merchant authorization
	// merchantID, exists := c.Get("merchantID")
	// if !exists {
	// 	logger.Warn("Unauthorized access attempt")
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	// 	return
	// }
	// merchantIDStr, ok := merchantID.(string)
	// if !ok || merchantIDStr == "" {
	// 	logger.Warn("Invalid merchant ID in context")
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid merchant ID"})
	// 	return
	// }

	// Call service
	err := h.productService.DeleteProduct(c.Request.Context(), productID)
	if err != nil {
		logger.Error("Failed to delete product", zap.Error(err), zap.String("product_id", productID))
		if errors.Is(err, product.ErrInvalidProduct) {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete product"})
		return
	}

	logger.Info("Product deleted successfully", zap.String("product_id", productID))
	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
```
</xaiArtifact>

---
### api/routes/carts.go
- Size: 0.69 KB
- Lines: 19
- Last Modified: 2025-09-19 21:59:09

<xaiArtifact artifact_id="29879f5a-7ecb-4d1a-b437-e7319a72827b" artifact_version_id="9b6135d8-84af-41c3-bda9-6df87130f374" title="api/routes/carts.go" contentType="text/go">
```go
package routes

import (
	"api-customer-merchant/internal/api/handlers"
	//"api-customer-merchant/internal/middleware"
	"api-customer-merchant/internal/services/cart"
	"github.com/gin-gonic/gin"
)

func SetupCartRoutes(r *gin.RouterGroup, cartService *cart.CartService) {
	cartHandlers := handlers.NewCartHandler(cartService)
	//protected := middleware.AuthMiddleware("user")
	r.GET("/cart", cartHandlers.GetCart)
	r.POST("/cart/items",  cartHandlers.AddToCart)
	r.GET("/cart/items/:id",  cartHandlers.GetCartItem)
	r.PUT("/cart/items/:id",  cartHandlers.UpdateCartItemQuantity)
	r.DELETE("/cart/items/:id",  cartHandlers.RemoveCartItem)
	//r.POST("/cart/clear", protected, customerHandlers.ClearCart)
}
```
</xaiArtifact>

---
### api/routes/customer.go
- Size: 0.94 KB
- Lines: 28
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="7371f80a-ace8-47fa-a032-d3eff7062a33" artifact_version_id="3ef9e71d-e21c-428a-844a-5e54b0f5ee49" title="api/routes/customer.go" contentType="text/go">
```go
package routes

   import (
       "api-customer-merchant/internal/api/handlers"
       "api-customer-merchant/internal/middleware"
       "api-customer-merchant/internal/db/repositories"
       "api-customer-merchant/internal/services/user"


       "github.com/gin-gonic/gin"
   )

   func RegisterCustomerRoutes(r *gin.Engine) {
    repo := repositories.NewUserRepository()
    service := user.NewAuthService( repo)
       customer := r.Group("/customer")
       {
           authHandler := handlers.NewAuthHandler(service)
           customer.POST("/register", authHandler.Register)
           customer.POST("/login", authHandler.Login)
           customer.GET("/auth/google", authHandler.GoogleAuth)
           customer.GET("/auth/google/callback", authHandler.GoogleCallback)

           protected := customer.Group("/")
           protected.Use(middleware.AuthMiddleware("customer"))
           protected.POST("/logout", authHandler.Logout)
       }
   }
```
</xaiArtifact>

---
### api/routes/merchant.go
- Size: 5.03 KB
- Lines: 136
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="0867aa18-e49b-4f34-9444-0c584dd0041c" artifact_version_id="8644c4f8-2766-4748-b203-e797d94f4835" title="api/routes/merchant.go" contentType="text/go">
```go
package routes

/*
   import (
      "api-customer-merchant/internal/api/merchant/handlers"
      "api-customer-merchant/internal/middleware"
      "api-customer-merchant/internal/db/repositories"

       "github.com/gin-gonic/gin"
   )

   func RegisterRoutes(r *gin.Engine) {
       merchant := r.Group("/merchant")
       {
           authHandler := handlers.NewAuthHandler()
           merchant.POST("/submitApplication", authHandler.Register)
           merchant.POST("/login", authHandler.Login)
           //merchant.GET("/auth/google", authHandler.GoogleAuth)
           //merchant.GET("/auth/google/callback", authHandler.GoogleCallback)

           protected := merchant.Group("/")
           protected.Use(middleware.AuthMiddleware("merchant"))
           protected.POST("/logout", authHandler.Logout)
       }
   }
*/

import (
	// "api-customer-merchant/internal/api/handlers"
	// "api-customer-merchant/internal/middleware"
	//"api-customer-merchant/internal/db"
	//"api-customer-merchant/internal/db/models"
	//"api-customer-merchant/internal/db/repositories"

	//"api-customer-merchant/internal/middleware"
	// "api-customer-merchant/internal/services/merchant"
	// "api-customer-merchant/internal/services/product"

	// "api-customer-merchant/internal/domain/order"
	// "api-customer-merchant/internal/domain/payout"
	//"api-customer-merchant/internal/domain/product"
	//"api-customer-merchant/internal/domain/invent"

	//"github.com/gin-gonic/gin"
)

/*
   func RegisterRoutes(r *gin.Engine) {
       merchant := r.Group("/merchant")
       {
          //userRepo := repositories.NewUserRepository()
           // merchantRepo := repositories.NewMerchantRepository()
           appRepo := repositories.NewMerchantApplicationRepository()
           repo := repositories.NewMerchantRepository()
           service := merchant.NewMerchantService(appRepo, repo)
           h := handlers.NewMerchantHandler(service)

           // authHandler := handlers.NewAuthHandler(merchantRepo)
           // merchant.POST("/submitApplication", authHandler.Register)
           // merchant.POST("/login", authHandler.Login)
           //merchant.GET("/auth/google", authHandler.GoogleAuth)
           //merchant.GET("/auth/google/callback", authHandler.GoogleCallback)
           merchant.POST("/apply",  h.Apply)
           merchant.GET("/application/:id",  h.GetApplication)


           // Merchant account access (once approved by admin via Express API)
           //merchant.POST("/me",  h.GetMyMerchant)



           merchant := merchant.Group("/")
           merchant.Use(middleware.AuthMiddleware("merchant"))
           merchant.POST("/me",  h.GetMyMerchant)
           //merchant.POST("/logout", authHandler.Logout)

       }
   }
*/




/*
func RegisterMerchantRoutes(r *gin.Engine) {
    appRepo := repositories.NewMerchantApplicationRepository()
    repo := repositories.NewMerchantRepository()
    service := merchant.NewMerchantService(appRepo, repo)
    productRepo := repositories.NewProductRepository()
    inventoryRepo:= repositories.NewInventoryRepository()
	productService := product.NewProductService(productRepo,inventoryRepo)
    //productRepo := repositories.NewProductRepository()
    //inventoryRepo:= repositories.NewInventoryRepository()
	//productService := product.NewProductService(productRepo,inventoryRepo)

	// Other services (stubs; instantiate as needed)
	// orderService := order.NewOrderService(nil) // Adjust
	// payoutService := payout.NewPayoutService(nil)
	// promotionService := promotions.NewPromotionService(nil)
    //h := handlers.NewMerchantAuthHandler(service)
    
    authHandler := handlers.NewMerchantAuthHandler(service)
    merchhandler:=handlers.NewMerchantHandlers(productService)
    merchant := r.Group("/merchant")
    {
    // Application submission and status
    // merchant.POST("/apply",  authHandler.Apply)
    // merchant.GET("/application/:id",  authHandler.GetApplication)
    // // Merchant account access (once approved by admin via Express API)
    // merchant.GET("/me",  authHandler.GetMyMerchant)

    //merchant.POST("/create/product", merchhandler.CreateProduct)
	//merchant.GET("/products", merchhandler.GetMyProducts)
	//merchant.PUT("/:id", merchhandler.UpdateProduct)
	//merchant.DELETE("/:id", merchhandler.DeleteProduct)
	//merchant.POST("/bulk-upload", merchhandler.BulkUploadProducts)
    merchant.POST("/apply",  authHandler.Apply)
    merchant.POST("/login",  authHandler.Login)
    merchant.GET("/application/:id",  authHandler.GetApplication)
    // Merchant account access (once approved by admin via Express API)
    //merchant.GET("/me",  authHandler.GetMyMerchant)
    protected := merchant.Group("/")
    protected.Use(middleware.AuthMiddleware("merchant"))
    protected.GET("/me",  authHandler.GetMyMerchant)

    protected.POST("/products/create", merchhandler.CreateProduct)
	protected.GET("/products", merchhandler.GetMyProducts)
	protected.PUT("/products/:id", merchhandler.UpdateProduct)
	protected.DELETE("/products/:id", merchhandler.DeleteProduct)
	protected.POST("/products/bulk-upload", merchhandler.BulkUploadProducts)

    }

   
}
*/
```
</xaiArtifact>

---
### api/routes/order_routes.go
- Size: 0.49 KB
- Lines: 15
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="6e3fb464-84ae-4d1b-9771-2a92993e1780" artifact_version_id="b132fafd-c8bb-4f6f-b3db-680705e3616c" title="api/routes/order_routes.go" contentType="text/go">
```go
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

```
</xaiArtifact>

---
### api/routes/product_route.go
- Size: 1.45 KB
- Lines: 53
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="0a625f64-f729-45fe-9cb7-f19b45fa5b84" artifact_version_id="a6df0829-6f06-4f84-9491-8eb01e8ca464" title="api/routes/product_route.go" contentType="text/go">
```go
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
```
</xaiArtifact>

---
### config/config.go
- Size: 0.72 KB
- Lines: 31
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="b64f76e7-9410-4476-b3dc-2b82478b8f00" artifact_version_id="9c091511-625a-4404-be88-f1599b8421c2" title="config/config.go" contentType="text/go">
```go
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

```
</xaiArtifact>

---
### db/db.go
- Size: 6.15 KB
- Lines: 216
- Last Modified: 2025-09-19 21:57:15

<xaiArtifact artifact_id="b7293969-b94a-4233-8198-62f780739985" artifact_version_id="0d4f7ca8-1f1f-4bcb-8a63-b85cf61e0ea2" title="db/db.go" contentType="text/go">
```go
package db

import (
	//"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
/*
func Connect() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable not set")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")
}

func AutoMigrate() {
	//    err := DB.AutoMigrate(
	//        &models.User{},
	// 	   &models.Merchant{},
	//        // Add other models here when implemented
	//    )
	//    if err != nil {
	//        log.Fatalf("Failed to auto-migrate: %v", err)
	//    }

}
*/

func Connect() {
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        log.Fatal("DB_DSN environment variable not set")
    }

    var err error
    //DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true,})
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{PrepareStmt: false})
    //DB = DB.Debug()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Configure connection pool
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to get SQL DB: %v", err)
    }
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetConnMaxLifetime(30 * time.Minute)

    log.Println("Database connected successfully")
}

func AutoMigrate() {
    // Run AutoMigrate with all models
	if err := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		log.Fatal("Failed to enable uuid-ossp extension:", err)

	}

    //DB = DB.Session(&gorm.Session{DisableForeignKeyConstraintWhenMigrating: true})
   log.Println("Starting AutoMigrate...")

    // Independent tables (no incoming FKs or self-referential only)
    // log.Println("Migrating Category...")
    // if err := DB.AutoMigrate(&models.Category{}); err != nil {
    //     log.Printf("Failed to migrate Category: %v", err)
    //     return
    // }

    //  log.Println("Migrating MerchantApplication...")
    //  if err := DB.AutoMigrate(&models.MerchantApplication{}); err != nil {
    //     log.Printf("Failed to migrate MerchantApplication: %v", err)
    //      return
    //  }

    // log.Println("Migrating Merchant...")
    // if err := DB.AutoMigrate(&models.Merchant{}); err != nil {
    //     log.Printf("Failed to migrate Merchant: %v", err)
    //     return
    // }

    //  log.Println("Migrating User...")
    //  if err := DB.AutoMigrate(&models.User{}); err != nil {
    //      log.Printf("Failed to migrate User: %v", err)
    //      return
    //  }

    // Tables depending on Merchant/Category/User
    //    log.Println("Migrating Product ecosystem (Product, Variant, Media)...")
    //  if err := DB.AutoMigrate(&models.Product{}, &models.Variant{}, &models.Media{}, &models.VendorInventory{}); err != nil {
    //      log.Printf("Failed to migrate Product/Variant/Media: %v", err)
    //      return
    //  }

    // log.Println("Migrating Inventory...")
    // if err := DB.AutoMigrate(&models.Inventory{}); err != nil {
    //     log.Printf("Failed to migrate Inventory: %v", err)
    //     return
    // }

    // log.Println("Migrating Promotion...")
    // if err := DB.AutoMigrate(&models.Promotion{}); err != nil {
    //     log.Printf("Failed to migrate Promotion: %v", err)
    //     return
    // }

     log.Println("Migrating Cart...")
     if err := DB.AutoMigrate(&models.Cart{}); err != nil {
         log.Printf("Failed to migrate Cart: %v", err)
         return
     }

     log.Println("Migrating CartItem...")
     if err := DB.AutoMigrate(&models.CartItem{}); err != nil {
         log.Printf("Failed to migrate CartItem: %v", err)
         return
     }

     log.Println("Migrating Order...")
     if err := DB.AutoMigrate(&models.Order{}); err != nil {
         log.Printf("Failed to migrate Order: %v", err)
         return
     }

     log.Println("Migrating OrderItem...")
     if err := DB.AutoMigrate(&models.OrderItem{}); err != nil {
         log.Printf("Failed to migrate OrderItem: %v", err)
         return
     }

    // log.Println("Migrating Payment...")
    // if err := DB.AutoMigrate(&models.Payment{}); err != nil {
    //     log.Printf("Failed to migrate Payment: %v", err)
    //     return
    // }

    // log.Println("Migrating Payout...")
    // if err := DB.AutoMigrate(&models.Payout{}); err != nil {
    //     log.Printf("Failed to migrate Payout: %v", err)
    //     return
    // }

    // log.Println("Migrating ReturnRequest...")
    // if err := DB.AutoMigrate(&models.ReturnRequest{}); err != nil {
    //     log.Printf("Failed to migrate ReturnRequest: %v", err)
    //     return
    // }
 
    err := DB.AutoMigrate(
        //&models.User{},
        //&models.MerchantApplication{},
         //&models.Product{},
         //&models.Variant{},
         //&models.Media{},
        // &models.Cart{},
        // &models.Order{},
        // &models.OrderItem{},
        // &models.CartItem{},
         //&models.Category{},
        // &models.Inventory{},
        // &models.Promotion{},
        // &models.ReturnRequest{},
        // &models.Payout{},
    )
        
    if err != nil {
        log.Fatalf("Failed to auto-migrate: %v", err)
    }


    // Get the underlying SQL database connection
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to get SQL DB: %v", err)
    }

    // Close all connections to clear cached plans
    if err := sqlDB.Close(); err != nil {
        log.Printf("Failed to close connections: %v", err)
    }

    // Reconnect to ensure fresh connections
    dsn := os.Getenv("DB_DSN")
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to reconnect to database: %v", err)
    }

    // Reconfigure connection pool
    sqlDB, err = DB.DB()
    if err != nil {
        log.Fatalf("Failed to get SQL DB after reconnect: %v", err)
    }
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetConnMaxLifetime(30 * time.Minute)

    log.Println("Database migrated and reconnected successfully")
}
```
</xaiArtifact>

---
### db/models/cart.go
- Size: 1.75 KB
- Lines: 68
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="d225194a-7a42-40dd-b299-9bcb0cb1f71e" artifact_version_id="42b1705f-4ef2-4e53-a317-bb5875f681e9" title="db/models/cart.go" contentType="text/go">
```go
package models

import (
	"fmt"
	"gorm.io/gorm"
)

// CartStatus defines possible cart status values
type CartStatus string

const (
	CartStatusActive    CartStatus = "Active"
	CartStatusAbandoned CartStatus = "Abandoned"
	CartStatusConverted CartStatus = "Converted"
)

// Valid checks if the status is one of the allowed values
func (s CartStatus) Valid() error {
	switch s {
	case CartStatusActive, CartStatusAbandoned, CartStatusConverted:
		return nil
	default:
		return fmt.Errorf("invalid cart status: %s", s)
	}
}

type Cart struct {
	gorm.Model
	UserID uint       `gorm:"not null" json:"user_id"`
	Status CartStatus `gorm:"type:varchar(20);not null;default:'Active'" json:"status"`
	SubTotal    float64 `gorm:"-" json:"subtotal"`    // Computed
	TaxTotal    float64 `gorm:"-" json:"tax_total"`
    ShipTotal   float64 `gorm:"-" json:"shipping_total"`
    GrandTotal  float64 `gorm:"-" json:"grand_total"`
	User   User       `gorm:"foreignKey:UserID"`
	CartItems []CartItem `gorm:"foreignKey:CartID"`
}

// BeforeCreate validates the Status field
func (c *Cart) BeforeCreate(tx *gorm.DB) error {
	if err := c.Status.Valid(); err != nil {
		return err
	}
	return nil
}

// BeforeUpdate validates the Status field
func (c *Cart) BeforeUpdate(tx *gorm.DB) error {
	if err := c.Status.Valid(); err != nil {
		return err
	}
	return nil
}

func (c *Cart) AfterFind(tx *gorm.DB) error {
    c.ComputeTotals()
    return nil
}

func (c *Cart) ComputeTotals() {
    c.SubTotal = 0
    for _, item := range c.CartItems {
        c.SubTotal += float64(item.Quantity) * (item.Product.BasePrice).InexactFloat64() // Assume BasePrice in Product
    }
    // Stub: c.TaxTotal = 0.1 * c.SubTotal // Or call pricing
    // c.ShipTotal = 10.00
    c.GrandTotal = c.SubTotal // + c.TaxTotal + c.ShipTotal
}
```
</xaiArtifact>

---
### db/models/cart_item.go
- Size: 0.44 KB
- Lines: 16
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="8dea3499-75ba-43d9-a5a2-41c24f67de8c" artifact_version_id="9053ecb4-4008-4592-9518-6279e281ddb1" title="db/models/cart_item.go" contentType="text/go">
```go
package models

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	CartID     uint    `gorm:"not null" json:"cart_id"`
	ProductID  string    `gorm:"not null" json:"product_id"`
	Quantity   int     `gorm:"not null" json:"quantity"`
	MerchantID string    `gorm:"not null" json:"merchant_id"`
	Cart       Cart    `gorm:"foreignKey:CartID"`
	Product    Product `gorm:"foreignKey:ProductID"`
	Merchant   Merchant `gorm:"foreignKey:MerchantID"`
}
```
</xaiArtifact>

---
### db/models/category.go
- Size: 0.34 KB
- Lines: 13
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="0b491ae5-b90c-425c-9157-c69b910663ec" artifact_version_id="b3703f1d-ecef-4925-95d4-8b9f656ccfee" title="db/models/category.go" contentType="text/go">
```go
package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name       string                 `gorm:"size:255;not null" json:"name"`
	ParentID   *uint                  `json:"parent_id"`
	Attributes map[string]interface{} `gorm:"type:jsonb" json:"attributes"`
	Parent     *Category              `gorm:"foreignKey:ParentID"`
}
```
</xaiArtifact>

---
### db/models/inventory.go
- Size: 0.34 KB
- Lines: 13
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="550829c9-2743-46d4-a7af-8e0f5956be9c" artifact_version_id="1c39d088-bc50-467f-910b-4187d18afe0e" title="db/models/inventory.go" contentType="text/go">
```go
package models

import (
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	ProductID         string   `gorm:"not null" json:"product_id"`
	StockQuantity     int    `gorm:"not null" json:"stock_quantity"`
	LowStockThreshold int    `gorm:"not null;default:10" json:"low_stock_threshold"`
	Product           Product `gorm:"foreignKey:ProductID"`
}
```
</xaiArtifact>

---
### db/models/merchant.go
- Size: 10.60 KB
- Lines: 179
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="431b8a27-5178-4384-bcb7-0262b257ccd5" artifact_version_id="6434c98e-90d4-4e56-bbf3-e022fc5cd8ee" title="db/models/merchant.go" contentType="text/go">
```go
package models


import (
	"time"
	"gorm.io/datatypes"
)
/*
// MerchantStatus defines the possible status values for a merchant
type MerchantBasicInfo struct {
	StoreName     string `gorm:"column:store_name;size:255;not null" json:"store_name" validate:"required"`
	Name          string `gorm:"column:name;size:255;not null" json:"name" validate:"required"`
	PersonalEmail string `gorm:"column:personal_email;size:255;not null;unique" json:"personal_email" validate:"required,email"`
	WorkEmail     string `gorm:"column:work_email;size:255;not null;unique" json:"work_email" validate:"required,email"`
	PhoneNumber   string `gorm:"column:phone_number;size:50" json:"phone_number"`
}

// MerchantAddress holds address information as JSONB
type MerchantAddress struct {
	PersonalAddress datatypes.JSON `gorm:"column:personal_address;type:jsonb;not null" json:"personal_address" validate:"required"`
	WorkAddress     datatypes.JSON `gorm:"column:work_address;type:jsonb;not null" json:"work_address" validate:"required"`
}

// MerchantBusinessInfo holds business-related information
type MerchantBusinessInfo struct {
	BusinessType               string `gorm:"column:business_type;size:100" json:"business_type"`
	Website                    string `gorm:"column:website;size:255" json:"website"`
	BusinessDescription        string `gorm:"column:business_description;type:text" json:"business_description"`
	BusinessRegistrationNumber string `gorm:"column:business_registration_number;size:255;not null;unique" json:"business_registration_number" validate:"required"`
}

// MerchantDocuments holds document-related information
type MerchantDocuments struct {
	StoreLogoURL                   string `gorm:"column:store_logo_url;size:255" json:"store_logo_url"`
	BusinessRegistrationCertificate string `gorm:"column:business_registration_certificate;size:255" json:"business_registration_certificate"`
}

// MerchantApplication holds the information required for a merchant onboarding application
type MerchantApplication struct {
	ID                string            `gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4()" json:"id,omitempty"`
	MerchantBasicInfo                   `gorm:"embedded"`
	MerchantAddress                     `gorm:"embedded"`
	MerchantBusinessInfo                `gorm:"embedded"`
	MerchantDocuments                   `gorm:"embedded"`
	Status            string            `gorm:"column:status;type:varchar(20);default:pending;not null" json:"status"`
	CreatedAt         time.Time         `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time         `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (MerchantApplication) TableName() string {
	return "merchant_application"
}

// MerchantStatus defines the possible statuses for a merchant
type MerchantStatus string

const (
	MerchantStatusActive    MerchantStatus = "active"
	MerchantStatusSuspended MerchantStatus = "suspended"
)

// Merchant holds the active merchant account details after approval
type Merchant struct {
	ID                string            `gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4()" json:"id,omitempty"`
	ApplicationID     string            `gorm:"column:application_id;type:uuid;not null;unique" json:"application_id"`
	MerchantID            string            `gorm:"column:merchant_id;type:uuid;not null;unique" json:"user_id"`
	MerchantBasicInfo                   `gorm:"embedded"`
	MerchantAddress                     `gorm:"embedded"`
	MerchantBusinessInfo                `gorm:"embedded"`
	MerchantDocuments                   `gorm:"embedded"`
	Password          string            `gorm:"column:password;size:255;not null" json:"password" validate:"required"`
	Status            MerchantStatus    `gorm:"column:status;type:varchar(20);default:active;index" json:"status"`
	CommissionTier    string            `gorm:"column:commission_tier;default:standard" json:"commission_tier"`
	CommissionRate    float64           `gorm:"column:commission_rate;default:5.00" json:"commission_rate"`
	AccountBalance    float64           `gorm:"column:account_balance;default:0.00" json:"account_balance"`
	TotalSales        float64           `gorm:"column:total_sales;default:0.00" json:"total_sales"`
	TotalPayouts      float64           `gorm:"column:total_payouts;default:0.00" json:"total_payouts"`
	StripeAccountID   string            `gorm:"column:stripe_account_id" json:"stripe_account_id"`
	PayoutSchedule    string            `gorm:"column:payout_schedule;default:weekly" json:"payout_schedule"`
	LastPayoutDate    *time.Time        `gorm:"column:last_payout_date" json:"last_payout_date"`
	Banner            string            `gorm:"column:banner;size:255" json:"banner"`
	Policies          datatypes.JSON    `gorm:"column:policies;type:jsonb" json:"policies"`
	CreatedAt         time.Time         `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time         `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	//Products          []Product         `gorm:"foreignKey:MerchantID" json:"products,omitempty"`
	//CartItems         []CartItem        `gorm:"foreignKey:MerchantID" json:"cart_items,omitempty"`
	//OrderItems        []OrderItem       `gorm:"foreignKey:MerchantID" json:"order_items,omitempty"`
	//Payouts           []Payout          `gorm:"foreignKey:MerchantID" json:"payouts,omitempty"`
}

func (Merchant) TableName() string {
	return "merchant"
}
*/

type MerchantBasicInfo struct {
	StoreName     string `gorm:"column:store_name;size:255;not null" json:"store_name" validate:"required"`
	Name          string `gorm:"column:name;size:255;not null" json:"name" validate:"required"`
	PersonalEmail string `gorm:"column:personal_email;size:255;not null;unique" json:"personal_email" validate:"required,email"`
	WorkEmail     string `gorm:"column:work_email;size:255;not null;unique" json:"work_email" validate:"required,email"`
	PhoneNumber   string `gorm:"column:phone_number;size:50" json:"phone_number"`
}

// MerchantAddress holds address information as JSONB
type MerchantAddress struct {
	PersonalAddress datatypes.JSON `gorm:"column:personal_address;type:jsonb;not null" json:"personal_address" validate:"required"`
	WorkAddress     datatypes.JSON `gorm:"column:work_address;type:jsonb;not null" json:"work_address" validate:"required"`
}

// MerchantBusinessInfo holds business-related information
type MerchantBusinessInfo struct {
	BusinessType               string `gorm:"column:business_type;size:100" json:"business_type"`
	Website                    string `gorm:"column:website;size:255" json:"website"`
	BusinessDescription        string `gorm:"column:business_description;type:text" json:"business_description"`
	BusinessRegistrationNumber string `gorm:"column:business_registration_number;size:255;not null;unique" json:"business_registration_number" validate:"required"`
}

// MerchantDocuments holds document-related information
type MerchantDocuments struct {
	StoreLogoURL                   string `gorm:"column:store_logo_url;size:255" json:"store_logo_url"`
	BusinessRegistrationCertificate string `gorm:"column:business_registration_certificate;size:255" json:"business_registration_certificate"`
}

// MerchantApplication holds the information required for a merchant onboarding application
type MerchantApplication struct {
	ID                string            `gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4()" json:"id,omitempty"`
	MerchantBasicInfo                   `gorm:"embedded"`
	MerchantAddress                     `gorm:"embedded"`
	MerchantBusinessInfo                `gorm:"embedded"`
	MerchantDocuments                   `gorm:"embedded"`
	Status            string            `gorm:"column:status;type:varchar(20);default:pending;not null" json:"status"`
	CreatedAt         time.Time         `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time         `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (MerchantApplication) TableName() string {
	return "merchant_application"
}

// MerchantStatus defines the possible statuses for a merchant
type MerchantStatus string

const (
	MerchantStatusActive    MerchantStatus = "active"
	MerchantStatusSuspended MerchantStatus = "suspended"
)

// Merchant holds the active merchant account details after approval
type Merchant struct {
	ID                string            `gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4()" json:"id,omitempty"`
	ApplicationID     string            `gorm:"column:application_id;type:uuid;not null;unique" json:"application_id"`
	MerchantID            string            `gorm:"column:merchant_id;type:uuid;not null;unique" json:"user_id"`
	MerchantBasicInfo                   `gorm:"embedded"`
	MerchantAddress                     `gorm:"embedded"`
	MerchantBusinessInfo                `gorm:"embedded"`
	MerchantDocuments                   `gorm:"embedded"`
	Password          string            `gorm:"column:password;size:255;not null" json:"password" validate:"required"`
	Status            MerchantStatus    `gorm:"column:status;type:varchar(20);default:active;index" json:"status"`
	CommissionTier    string            `gorm:"column:commission_tier;default:standard" json:"commission_tier"`
	CommissionRate    float64           `gorm:"column:commission_rate;default:5.00" json:"commission_rate"`
	AccountBalance    float64           `gorm:"column:account_balance;default:0.00" json:"account_balance"`
	TotalSales        float64           `gorm:"column:total_sales;default:0.00" json:"total_sales"`
	TotalPayouts      float64           `gorm:"column:total_payouts;default:0.00" json:"total_payouts"`
	StripeAccountID   string            `gorm:"column:stripe_account_id" json:"stripe_account_id"`
	PayoutSchedule    string            `gorm:"column:payout_schedule;default:weekly" json:"payout_schedule"`
	LastPayoutDate    *time.Time        `gorm:"column:last_payout_date" json:"last_payout_date"`
	Banner            string            `gorm:"column:banner;size:255" json:"banner"`
	Policies          datatypes.JSON    `gorm:"column:policies;type:jsonb" json:"policies"`
	CreatedAt         time.Time         `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time         `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	//Products          []Product         `gorm:"foreignKey:MerchantID" json:"products,omitempty"`
	//CartItems         []CartItem        `gorm:"foreignKey:MerchantID" json:"cart_items,omitempty"`
	//OrderItems        []OrderItem       `gorm:"foreignKey:MerchantID" json:"order_items,omitempty"`
	//Payouts           []Payout          `gorm:"foreignKey:MerchantID" json:"payouts,omitempty"`
}

func (Merchant) TableName() string {
	return "merchant"
}
```
</xaiArtifact>

---
### db/models/order.go
- Size: 1.24 KB
- Lines: 50
- Last Modified: 2025-09-19 21:48:48

<xaiArtifact artifact_id="e9e01809-fba5-4b58-acc2-1b3f090474a4" artifact_version_id="8bf9ca64-787a-45d6-a26e-b388998a308e" title="db/models/order.go" contentType="text/go">
```go
package models

import (
	"fmt"
	"gorm.io/gorm"
)

// OrderStatus defines possible order status values
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "Pending"
	OrderStatusCompleted OrderStatus = "Completed"
	OrderStatusCancelled OrderStatus = "Cancelled"
)

// Valid checks if the status is one of the allowed values
func (s OrderStatus) Valid() error {
	switch s {
	case OrderStatusPending, OrderStatusCompleted, OrderStatusCancelled:
		return nil
	default:
		return fmt.Errorf("invalid order status: %s", s)
	}
}

type Order struct {
	gorm.Model
	UserID      uint        `gorm:"not null" json:"user_id"`
	TotalAmount float64     `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	Status      OrderStatus `gorm:"type:varchar(20);not null;default:'Pending'" json:"status"`
	User        User        `gorm:"foreignKey:UserID"`
	OrderItems    []OrderItem   `gorm:"foreignKey:ProductID" json:"media,omitempty"`
}

// BeforeCreate validates the Status field
func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if err := o.Status.Valid(); err != nil {
		return err
	}
	return nil
}

// BeforeUpdate validates the Status field
func (o *Order) BeforeUpdate(tx *gorm.DB) error {
	if err := o.Status.Valid(); err != nil {
		return err
	}
	return nil
}
```
</xaiArtifact>

---
### db/models/order_item.go
- Size: 2.46 KB
- Lines: 73
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="9519cbae-1780-426b-881a-92f57373cf72" artifact_version_id="9787bd6d-5cab-415c-b5a6-68bdbe635bf8" title="db/models/order_item.go" contentType="text/go">
```go
package models

import (
	"fmt"
	"gorm.io/gorm"
)

// FulfillmentStatus defines possible fulfillment status values
type FulfillmentStatus string

const (
	FulfillmentStatusNew     FulfillmentStatus = "New"
	FulfillmentStatusShipped FulfillmentStatus = "Shipped"
)

// Valid checks if the status is one of the allowed values
func (s FulfillmentStatus) Valid() error {
	switch s {
	case FulfillmentStatusNew, FulfillmentStatusShipped:
		return nil
	default:
		return fmt.Errorf("invalid fulfillment status: %s", s)
	}
}

// type OrderItem struct {
// 	gorm.Model
// 	OrderID           uint              `gorm:"not null" json:"order_id"`
// 	ProductID         string              `gorm:"not null" json:"product_id"`
// 	MerchantID        uint              `gorm:"not null" json:"merchant_id"`
// 	Quantity          int               `gorm:"not null" json:"quantity"`
// 	Price             float64           `gorm:"type:decimal(10,2);not null" json:"price"`
// 	FulfillmentStatus FulfillmentStatus `gorm:"type:varchar(20);not null;default:'New'" json:"fulfillment_status"`
// 	Order             Order             `gorm:"foreignKey:OrderID"`
// 	Product           Product           `gorm:"foreignKey:ProductID"`
// 	Merchant          Merchant          `gorm:"foreignKey:MerchantID"`
// }




type OrderItem struct {
	gorm.Model
	OrderID           uint              `gorm:"not null;index" json:"order_id"`
	ProductID         string              `gorm:"not null;index" json:"product_id"`
	//ProductID         uint              `gorm:"not null;index" json:"product_id"`
	MerchantID        string            `gorm:"not null;index" json:"merchant_id"`
	Quantity          int               `gorm:"not null" json:"quantity"`
	Price             float64           `gorm:"type:decimal(10,2);not null" json:"price"`
	FulfillmentStatus FulfillmentStatus `gorm:"type:varchar(20);not null;default:'New'" json:"fulfillment_status"`
	Order             Order             `gorm:"foreignKey:OrderID"`
	Product           Product           `gorm:"foreignKey:ProductID;references:ID"`
	Merchant          Merchant          `gorm:"foreignKey:MerchantID"`
}




// BeforeCreate validates the FulfillmentStatus field
func (oi *OrderItem) BeforeCreate(tx *gorm.DB) error {
	if err := oi.FulfillmentStatus.Valid(); err != nil {
		return err
	}
	return nil
}

// BeforeUpdate validates the FulfillmentStatus field
func (oi *OrderItem) BeforeUpdate(tx *gorm.DB) error {
	if err := oi.FulfillmentStatus.Valid(); err != nil {
		return err
	}
	return nil
}
```
</xaiArtifact>

---
### db/models/payment.go
- Size: 1.16 KB
- Lines: 49
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="552d9b63-ce01-47c2-84f4-5a93a6fe6708" artifact_version_id="d690e554-c335-467a-b177-f68cf3b01738" title="db/models/payment.go" contentType="text/go">
```go
package models

import (
	"fmt"
	"gorm.io/gorm"
)

// PaymentStatus defines possible payment status values
type PaymentStatus string

const (
	PaymentStatusPending  PaymentStatus = "Pending"
	PaymentStatusCompleted PaymentStatus = "Completed"
	PaymentStatusFailed   PaymentStatus = "Failed"
)

// Valid checks if the status is one of the allowed values
func (s PaymentStatus) Valid() error {
	switch s {
	case PaymentStatusPending, PaymentStatusCompleted, PaymentStatusFailed:
		return nil
	default:
		return fmt.Errorf("invalid payment status: %s", s)
	}
}

type Payment struct {
	gorm.Model
	OrderID uint          `gorm:"not null" json:"order_id"`
	Amount  float64       `gorm:"not null" json:"amount"`
	Status  PaymentStatus `gorm:"type:varchar(20);not null;default:'Pending'" json:"status"`
	Order   Order         `gorm:"foreignKey:OrderID"`
}

// BeforeCreate validates the Status field
func (p *Payment) BeforeCreate(tx *gorm.DB) error {
	if err := p.Status.Valid(); err != nil {
		return err
	}
	return nil
}

// BeforeUpdate validates the Status field
func (p *Payment) BeforeUpdate(tx *gorm.DB) error {
	if err := p.Status.Valid(); err != nil {
		return err
	}
	return nil
}
```
</xaiArtifact>

---
### db/models/payout.go
- Size: 1.21 KB
- Lines: 49
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="12a9f6cb-40ee-4ab9-ad36-eb498c8c1bea" artifact_version_id="07ecc302-edc0-475c-b1e1-e05c694865bc" title="db/models/payout.go" contentType="text/go">
```go
package models

import (
	"fmt"
	"gorm.io/gorm"
)

// PayoutStatus defines possible payout status values
type PayoutStatus string

const (
	PayoutStatusPending   PayoutStatus = "Pending"
	PayoutStatusCompleted PayoutStatus = "Completed"
)

// Valid checks if the status is one of the allowed values
func (s PayoutStatus) Valid() error {
	switch s {
	case PayoutStatusPending, PayoutStatusCompleted:
		return nil
	default:
		return fmt.Errorf("invalid payout status: %s", s)
	}
}

type Payout struct {
	gorm.Model
	MerchantID       uint         `gorm:"not null" json:"merchant_id"`
	Amount           float64      `gorm:"type:decimal(10,2);not null" json:"amount"`
	Status           PayoutStatus `gorm:"type:varchar(20);not null;default:'Pending'" json:"status"`
	PayoutAccountID  string       `gorm:"size:255;not null" json:"payout_account_id"`
	Merchant         Merchant     `gorm:"foreignKey:MerchantID"`
}

// BeforeCreate validates the Status field
func (p *Payout) BeforeCreate(tx *gorm.DB) error {
	if err := p.Status.Valid(); err != nil {
		return err
	}
	return nil
}

// BeforeUpdate validates the Status field
func (p *Payout) BeforeUpdate(tx *gorm.DB) error {
	if err := p.Status.Valid(); err != nil {
		return err
	}
	return nil
}
```
</xaiArtifact>

---
### db/models/product.go
- Size: 3.77 KB
- Lines: 124
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="c5e593d1-8b96-4df0-a0d8-6575b21a663b" artifact_version_id="72b8e95e-6b38-4ac1-bb0c-c5da41d019c7" title="db/models/product.go" contentType="text/go">
```go
package models

/*

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// type Product struct {
// 	gorm.Model
// 	MerchantID  uint    `gorm:"not null" json:"merchant_id"`
// 	Name        string  `gorm:"size:255;not null" json:"name"`
// 	Description string  `gorm:"type:text" json:"description"`
// 	SKU         string  `gorm:"size:100;unique;not null" json:"sku"`
// 	Price       float64 `gorm:"type:decimal(10,2);not null" json:"price"`
// 	CategoryID  uint    `gorm:"not null" json:"category_id"`
// 	Merchant    Merchant `gorm:"foreignKey:MerchantID"`
// 	Category    Category `gorm:"foreignKey:CategoryID"`
// }

type AttributesMap map[string]string

// Value implements driver.Valuer
func (a AttributesMap) Value() (driver.Value, error) {
    return json.Marshal(a)
}


// Scan implements sql.Scanner
func (a *AttributesMap) Scan(value interface{}) error {
    b, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }
    return json.Unmarshal(b, a)
}

type Product struct {
	ID          string                 `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"` // Override gorm.Model ID
	MerchantID  string                 `gorm:"type:uuid;not null;index"` // FK to merchants.id (UUID string)
	Name        string                 `gorm:"size:255;not null" json:"name"`
	Description string                 `gorm:"type:text" json:"description"`
	SKU         string                 `gorm:"size:100;unique;not null;index" json:"sku"`
	Price       float64                `gorm:"type:decimal(10,2);not null" json:"price"`
	CategoryID  uint                 `gorm:"type:int;index" json:"category_id"` // Changed to string for UUID; revert to uint if numeric
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"` 

	Merchant    Merchant    `gorm:"foreignKey:MerchantID;references:id"` // Ensure Merchant.ID is string
	Category    Category    `gorm:"foreignKey:CategoryID"` // Adjust if Category.ID is uint
	Variants    []Variant   `gorm:"foreignKey:ProductID"`  // Keep relational
	Media       []Media     `gorm:"foreignKey:ProductID"`  // Keep relational
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}


// type Variant struct {
//     gorm.Model
//     ProductID  uint
//     Attributes map[string]string `gorm:"type:jsonb"`
//     Price      float64
//     SKU        string
// }


// //define merchant variant

// type Media struct {
//     gorm.Model
//     ProductID uint
//     URL       string
//     Type      string // image/video
//}

type Variant struct {
	gorm.Model
	ID        string             `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"` // Add UUID PK
	ProductID string             `gorm:"type:uuid;not null;index"` // Fixed: string for UUID (references products.id)
	//Attributes map[string]string `gorm:"type:jsonb;default:'{}'"`
	Attributes AttributesMap `gorm:"type:jsonb;default:'{}'"`
	Price     float64            `gorm:"type:decimal(10,2);not null"`
	SKU       string             `gorm:"size:100;unique;not null;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (v *Variant) BeforeCreate(tx *gorm.DB) error {
	if v.ID == "" {
		v.ID = uuid.New().String()
	}
	return nil
}

type Media struct {
	gorm.Model
	ID        string  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"` // Add UUID PK
	ProductID string  `gorm:"type:uuid;not null;index"` // Fixed: string for UUID (references products.id)
	URL       string  `gorm:"size:500;not null"`
	Type      string  `gorm:"size:20;default:'image';not null"` // e.g., "image", "video"
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Media) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	return nil
}

*/
```
</xaiArtifact>

---
### db/models/product2.go
- Size: 5.29 KB
- Lines: 153
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="59380bac-e32e-49b1-8774-f59b1e0128c5" artifact_version_id="19049ff3-d1ce-432a-ac54-76c792b70080" title="db/models/product2.go" contentType="text/go">
```go
package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// AttributesMap for JSONB
type AttributesMap map[string]string

func (a *AttributesMap) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, a)
}

func (a AttributesMap) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// MediaType enum-like type
type MediaType string

const (
	MediaTypeImage MediaType = "image"
	MediaTypeVideo MediaType = "video"
)

func (mt *MediaType) Scan(value interface{}) error {
	*mt = MediaType(value.(string))
	return nil
}

func (mt MediaType) Value() (driver.Value, error) {
	return string(mt), nil
}

type Product struct {
	ID          string         `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	MerchantID  string         `gorm:"type:uuid;not null;index" json:"merchant_id"`
	Name        string         `gorm:"size:255;not null" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	SKU         string         `gorm:"size:100;unique;not null;index" json:"sku"`
	BasePrice   decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"base_price"`
	CategoryID  uint           `gorm:"type:int;index" json:"category_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Merchant Merchant `gorm:"foreignKey:MerchantID;references:id"`
	Category Category `gorm:"foreignKey:CategoryID"`
	Variants []Variant `gorm:"foreignKey:ProductID" json:"variants,omitempty"`
	Media    []Media   `gorm:"foreignKey:ProductID" json:"media,omitempty"`
	SimpleInventory *VendorInventory `gorm:"foreignKey:ProductID"` // Only for simple products (no variants)
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}

type Variant struct {
	ID             string         `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	ProductID      string         `gorm:"type:uuid;not null;index" json:"product_id"`
	SKU            string         `gorm:"size:100;unique;not null;index" json:"sku"`
	PriceAdjustment decimal.Decimal `gorm:"type:decimal(10,2);not null;default:0.00" json:"price_adjustment"`
	TotalPrice     decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"total_price"` // Computed: BasePrice + PriceAdjustment
	Attributes     AttributesMap  `gorm:"type:jsonb;default:'{}'" json:"attributes"`
	IsActive       bool           `gorm:"default:true" json:"is_active"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`

	Product   Product         `gorm:"foreignKey:ProductID"`
	Inventory VendorInventory `gorm:"foreignKey:VariantID"`
}

func (v *Variant) BeforeCreate(tx *gorm.DB) error {
	if v.ID == "" {
		v.ID = uuid.New().String()
	}
	// Fetch Product to compute TotalPrice
	var product Product
	if err := tx.Where("id = ?", v.ProductID).First(&product).Error; err != nil {
		return err
	}
	v.TotalPrice = product.BasePrice.Add(v.PriceAdjustment)
	return nil
}

func (v *Variant) BeforeUpdate(tx *gorm.DB) error {
	// Recompute TotalPrice on update
	var product Product
	if err := tx.Where("id = ?", v.ProductID).First(&product).Error; err != nil {
		return err
	}
	v.TotalPrice = product.BasePrice.Add(v.PriceAdjustment)
	return nil
}

type VendorInventory struct {
	ID               string         `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	VariantID        string         `gorm:"type:uuid;not null;unique;index" json:"variant_id"`
	MerchantID       string         `gorm:"type:uuid;not null;index" json:"merchant_id"`
	ProductID        *string        `gorm:"type:uuid;index"` // Nullable: For simple products
	Quantity         int            `gorm:"default:0;not null;check:quantity >= 0" json:"quantity"`
	ReservedQuantity int            `gorm:"default:0;check:reserved_quantity >= 0" json:"reserved_quantity"`
	LowStockThreshold int           `gorm:"default:10" json:"low_stock_threshold"`
	BackorderAllowed bool           `gorm:"default:false" json:"backorder_allowed"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`

	Variant  *Variant `gorm:"foreignKey:VariantID"`
	Product  *Product `gorm:"foreignKey:ProductID"`
	Merchant Merchant `gorm:"foreignKey:MerchantID"`
}

func (vi *VendorInventory) BeforeCreate(tx *gorm.DB) error {
	if vi.ID == "" {
		vi.ID = uuid.New().String()
	}
	if (vi.VariantID != "" && vi.ProductID != nil) || (vi.VariantID == "" && vi.ProductID == nil) {
		return errors.New("exactly one of VariantID or ProductID must be set")
	}
	return nil
}

type Media struct {
	ID        string      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	ProductID string      `gorm:"type:uuid;not null;index" json:"product_id"`
	URL       string      `gorm:"size:500;not null" json:"url"`
	Type      MediaType   `gorm:"type:varchar(20);default:image;not null" json:"type"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`

	Product Product `gorm:"foreignKey:ProductID"`
}

func (m *Media) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	return nil
}
```
</xaiArtifact>

---
### db/models/user.go
- Size: 0.54 KB
- Lines: 15
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="0641df26-c3ba-4001-a278-f58cf34386b8" artifact_version_id="9b4f7989-ec28-4b0b-bce3-7c4d252bae61" title="db/models/user.go" contentType="text/go">
```go
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Name     string `gorm:"type:varchar(100);not null"`
	Password string // Empty for OAuth users
	//Role     string `gorm:"not null"` // "customer" (default) or "merchant" (upgraded by admin)
	GoogleID string    // Google ID for OAuth
	Country  string `gorm:"type:varchar(100)"` // Optional country field
	//Carts    []Cart  `gorm:"foreignKey:UserID" json:"carts,omitempty"`
	//Orders   []Order `gorm:"foreignKey:UserID" json:"orders,omitempty"`
}
```
</xaiArtifact>

---
### db/repositories/cart_item_repositry.go
- Size: 4.15 KB
- Lines: 121
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="634a4bfc-7c8c-40a7-bf2b-3363f32c0ecd" artifact_version_id="b4b5da5e-148d-4d7d-8f75-6d9dc96edbfb" title="db/repositories/cart_item_repositry.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"
	"fmt"

	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause" // Added for Locking
)

var (
	ErrCartItemNotFound  = errors.New("cart item not found")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrReservationFailed = errors.New("failed to reserve stock")
)

type CartItemRepository struct {
	db *gorm.DB
}

func NewCartItemRepository() *CartItemRepository {
	return &CartItemRepository{db: db.DB}
}

func (r *CartItemRepository) Create(ctx context.Context, cartItem *models.CartItem) error {
	return r.db.WithContext(ctx).Create(cartItem).Error
}

func (r *CartItemRepository) FindByID(ctx context.Context, id uint) (*models.CartItem, error) {
	var cartItem models.CartItem
	err := r.db.WithContext(ctx).
		Preload("Cart.User").
		Preload("Product.Merchant").
		Preload("Merchant").
		First(&cartItem, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrCartItemNotFound
	}
	return &cartItem, err
}

func (r *CartItemRepository) FindByCartID(ctx context.Context, cartID uint) ([]models.CartItem, error) {
	var cartItems []models.CartItem
	err := r.db.WithContext(ctx).
		Preload("Product.Merchant").
		Preload("Merchant").
		Where("cart_id = ?", cartID).Find(&cartItems).Error
	return cartItems, err
}

func (r *CartItemRepository) FindByProductIDAndCartID(ctx context.Context, productID string, cartID uint) (*models.CartItem, error) {
	var cartItem models.CartItem
	err := r.db.WithContext(ctx).
		Preload("Product.Merchant").
		Preload("Merchant").
		Where("product_id = ? AND cart_id = ?", productID, cartID).First(&cartItem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrCartItemNotFound
	}
	return &cartItem, err
}

func (r *CartItemRepository) UpdateQuantityWithReservation(ctx context.Context, itemID uint, newQuantity int, inventoryID uint) error { // Changed to uint
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var item models.CartItem
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&item, itemID).Error; err != nil {
			return fmt.Errorf("failed to lock item: %w", err)
		}

		delta := newQuantity - item.Quantity
		if delta > 0 {
			// Reserve (check first via repo if needed)
			if err := tx.Model(&models.Inventory{}).Where("id = ?", inventoryID).
				Update("reserved_quantity", gorm.Expr("reserved_quantity + ?", delta)).Error; err != nil { // Fixed: Assigned err
				return fmt.Errorf("stock reservation failed: %w", ErrReservationFailed)
			}
		} else if delta < 0 {
			err := tx.Model(&models.Inventory{}).Where("id = ?", inventoryID).
				Update("reserved_quantity", gorm.Expr("reserved_quantity - ?", -delta)).Error // Fixed: Assigned err (unused but for consistency)
			if err != nil {
				return fmt.Errorf("stock unreservation failed: %w", err)
			}
		}

		return tx.Model(&models.CartItem{}).Where("id = ?", itemID).Update("quantity", newQuantity).Error
	})
}

func (r *CartItemRepository) Update(ctx context.Context, cartItem *models.CartItem) error {
	return r.db.WithContext(ctx).Save(cartItem).Error
}

func (r *CartItemRepository) DeleteWithUnreserve(ctx context.Context, id uint, inventoryID uint) error { // uint
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var item models.CartItem
		if err := tx.First(&item, id).Error; err != nil {
			return ErrCartItemNotFound
		}
		// Release
		err := tx.Model(&models.Inventory{}).Where("id = ?", inventoryID).
			Update("reserved_quantity", gorm.Expr("reserved_quantity - ?", item.Quantity)).Error // Fixed: Assigned err
		if err != nil {
			return fmt.Errorf("stock unreservation failed: %w", err)
		}
		return tx.Delete(&models.CartItem{}, id).Error
	})
}

func (r *CartItemRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.CartItem{}, id).Error
}

// Stub for inventory_repo.go if missing:
// func (r *InventoryRepository) UpdateInventoryQuantity(ctx context.Context, inventoryID uint, delta int) error { // Add this method
// 	return r.db.WithContext(ctx).Model(&models.Inventory{}).Where("id = ?", inventoryID).
// 		Update("quantity", gorm.Expr("quantity + ?", delta)).Error
// }
```
</xaiArtifact>

---
### db/repositories/cart_repositry.go
- Size: 2.92 KB
- Lines: 90
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="2e41f28f-751a-4f3b-9c94-e0228c0abbfe" artifact_version_id="a860e480-76de-496b-934d-6473c79224ab" title="db/repositories/cart_repositry.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrCartNotFound = errors.New("cart not found")

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository() *CartRepository {
	return &CartRepository{db: db.DB}
}

// Create adds a new cart
func (r *CartRepository) Create(ctx context.Context, cart *models.Cart) error {
	return r.db.WithContext(ctx).Create(cart).Error
}
// FindByID retrieves a cart by ID with associated User and CartItems
// func (r *CartRepository) FindByID(id uint) (*models.Cart, error) {
// 	var cart models.Cart
// 	err := r.db.Preload("User").Preload("CartItems.Product.Merchant").First(&cart, id).Error
// 	return &cart, err
// }

func (r *CartRepository) FindByID(ctx context.Context, id uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("CartItems.Product.Media"). // Efficient: preload media for UI
		Preload("CartItems.Product.Variants.Inventory").
		First(&cart, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrCartNotFound
	}
	return &cart, err
}

// FindActiveCart retrieves the user's most recent active cart
// func (r *CartRepository) FindActiveCart(userID uint) (*models.Cart, error) {
// 	var cart models.Cart
// 	err := r.db.Preload("CartItems.Product.Merchant").Where("user_id = ? AND status = ?", userID, models.CartStatusActive).Order("created_at DESC").First(&cart).Error
// 	return &cart, err
// }


func (r *CartRepository) FindActiveCart(ctx context.Context, userID uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.WithContext(ctx).
		Preload("CartItems.Product.Merchant"). // As before
		Where("user_id = ? AND status = ?", userID, models.CartStatusActive).
		Order("created_at DESC").First(&cart).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrCartNotFound
	}
	return &cart, err
}


// FindByUserIDAndStatus retrieves carts for a user by status
func (r *CartRepository) FindByUserIDAndStatus(ctx context.Context , userID uint, status models.CartStatus) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.WithContext(ctx).
	Preload("CartItems.Product.Merchant").Where("user_id = ? AND status = ?", userID, status).Find(&carts).Error
	return carts, err
}

// FindByUserID retrieves all carts for a user
func (r *CartRepository) FindByUserID(ctx context.Context, userID uint) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.WithContext(ctx).Preload("CartItems.Product.Merchant").Where("user_id = ?", userID).Find(&carts).Error
	return carts, err
}

// Update modifies an existing cart
func (r *CartRepository) Update(ctx context.Context, cart *models.Cart) error {
	return r.db.WithContext(ctx).Save(cart).Error
}

// Delete removes a cart by ID
func (r *CartRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Cart{}, id).Error
}
```
</xaiArtifact>

---
### db/repositories/category_repositry.go
- Size: 1.14 KB
- Lines: 45
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="ec3c1015-5b03-434d-8fba-4c8c1a4e253b" artifact_version_id="4bd4082e-1cb4-4ee4-abc0-3791340fd91f" title="db/repositories/category_repositry.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{db: db.DB}
}

// Create adds a new category
func (r *CategoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

// FindByID retrieves a category by ID with parent category
func (r *CategoryRepository) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.Preload("Parent").First(&category, id).Error
	return &category, err
}

// FindAll retrieves all categories
func (r *CategoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Preload("Parent").Find(&categories).Error
	return categories, err
}

// Update modifies an existing category
func (r *CategoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

// Delete removes a category by ID
func (r *CategoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}
```
</xaiArtifact>

---
### db/repositories/inventory_repository.go
- Size: 1.58 KB
- Lines: 51
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="091e4f6c-00cc-4627-bdfb-05da7aa7b669" artifact_version_id="8201cbfb-5fea-4083-952f-f5d7a0b653d2" title="db/repositories/inventory_repository.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"

	"context"

	"gorm.io/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository() *InventoryRepository {
	return &InventoryRepository{db: db.DB}
}

// Create adds a new inventory record
func (r *InventoryRepository) Create(inventory *models.Inventory) error {
	return r.db.Create(inventory).Error
}

// FindByProductID retrieves inventory by product ID
func (r *InventoryRepository) FindByProductID(productID string) (*models.Inventory, error) {
	var inventory models.Inventory
	err := r.db.Where("product_id = ?", productID).First(&inventory).Error
	return &inventory, err
}

// UpdateStock updates the stock quantity for a product
func (r *InventoryRepository) UpdateStock(productID string, quantityChange int) error {
	return r.db.Model(&models.Inventory{}).Where("product_id = ?", productID).
		Update("stock_quantity", gorm.Expr("stock_quantity + ?", quantityChange)).Error
}

// Update modifies an existing inventory record
func (r *InventoryRepository) Update(inventory *models.Inventory) error {
	return r.db.Save(inventory).Error
}

// Delete removes an inventory record by ID
func (r *InventoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Inventory{}, id).Error
}

 func (r *InventoryRepository) UpdateInventoryQuantity(ctx context.Context, inventoryID uint, delta int) error { // Add this method
 	return r.db.WithContext(ctx).Model(&models.Inventory{}).Where("id = ?", inventoryID).
 		Update("quantity", gorm.Expr("quantity + ?", delta)).Error
 }
```
</xaiArtifact>

---
### db/repositories/merchant_repositry.go
- Size: 2.57 KB
- Lines: 80
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="b0d3735f-2c49-494b-bf79-6582cc68fcd3" artifact_version_id="b357965d-1799-46bc-bb94-f1305334e226" title="db/repositories/merchant_repositry.go" contentType="text/go">
```go
package repositories

import (
	"context"
	"log"

	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"
	//"gorm.io/gorm"
)

// MerchantApplicationRepository handles CRUD for merchant applications
// Note: Admin service (in Express) will be responsible for updating status/approval.
type MerchantApplicationRepository struct{}

func NewMerchantApplicationRepository() *MerchantApplicationRepository {
	return &MerchantApplicationRepository{}
}

func (r *MerchantApplicationRepository) Create(ctx context.Context, m *models.MerchantApplication) error {
	err := db.DB.WithContext(ctx).Create(m).Error
	if err != nil {
		log.Printf("Failed to create merchant application: %v", err)
		return err
	}
	return nil
}

func (r *MerchantApplicationRepository) GetByID(ctx context.Context, id string) (*models.MerchantApplication, error) {
	var m models.MerchantApplication
	if err := db.DB.WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		log.Printf("Failed to get merchant application by ID %s: %v", id, err)
		return nil, err
	}
	return &m, nil
}

func (r *MerchantApplicationRepository) GetByUserEmail(ctx context.Context, email string) (*models.MerchantApplication, error) {
	var m models.MerchantApplication
	if err := db.DB.WithContext(ctx).Where("personal_email = ? OR work_email = ?", email, email).First(&m).Error; err != nil {
		log.Printf("Failed to get merchant application by email %s: %v", email, err)
		return nil, err
	}
	return &m, nil
}

// MerchantRepository handles active merchants
type MerchantRepository struct{}

func NewMerchantRepository() *MerchantRepository {
	return &MerchantRepository{}
}

func (r *MerchantRepository) GetByID(ctx context.Context, id string) (*models.Merchant, error) {
	var m models.Merchant
	if err := db.DB.WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		log.Printf("Failed to get merchant by ID %s: %v", id, err)
		return nil, err
	}
	return &m, nil
}

func (r *MerchantRepository) GetByUserID(ctx context.Context, uid string) (*models.Merchant, error) {
	var m models.Merchant
	if err := db.DB.WithContext(ctx).Where("user_id = ?", uid).First(&m).Error; err != nil {
		log.Printf("Failed to get merchant by user ID %s: %v", uid, err)
		return nil, err
	}
	return &m, nil
}


func (r *MerchantRepository)GetByWorkEmail(ctx context.Context, email string) (*models.Merchant, error) {
	var m models.Merchant
	if err := db.DB.WithContext(ctx).Where("personal_email = ? OR work_email = ?", email, email).First(&m).Error; err != nil {
		log.Printf("Failed to get merchant  by email %s: %v", email, err)
		return nil, err
	}
	return &m, nil
}
```
</xaiArtifact>

---
### db/repositories/order_item_repository.go
- Size: 1.35 KB
- Lines: 45
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="2be4632d-2546-4e0b-8bb6-5ddf6a1ddeb7" artifact_version_id="3b6df042-8db5-418e-b67b-5b8b9caaae94" title="db/repositories/order_item_repository.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"

	"gorm.io/gorm"
)

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository() *OrderItemRepository {
	return &OrderItemRepository{db: db.DB}
}

// Create adds a new order item
func (r *OrderItemRepository) Create(orderItem *models.OrderItem) error {
	return r.db.Create(orderItem).Error
}

// FindByID retrieves an order item by ID with associated Order, Product, and Merchant
func (r *OrderItemRepository) FindByID(id uint) (*models.OrderItem, error) {
	var orderItem models.OrderItem
	err := r.db.Preload("Order.User").Preload("Product.Merchant").Preload("Merchant").First(&orderItem, id).Error
	return &orderItem, err
}

// FindByOrderID retrieves all order items for an order
func (r *OrderItemRepository) FindByOrderID(orderID uint) ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	err := r.db.Preload("Product.Merchant").Preload("Merchant").Where("order_id = ?", orderID).Find(&orderItems).Error
	return orderItems, err
}

// Update modifies an existing order item
func (r *OrderItemRepository) Update(orderItem *models.OrderItem) error {
	return r.db.Save(orderItem).Error
}

// Delete removes an order item by ID
func (r *OrderItemRepository) Delete(id uint) error {
	return r.db.Delete(&models.OrderItem{}, id).Error
}
```
</xaiArtifact>

---
### db/repositories/order_repository.go
- Size: 1.87 KB
- Lines: 59
- Last Modified: 2025-09-19 21:50:43

<xaiArtifact artifact_id="860c5001-d3e3-4e94-938e-d30e66961b09" artifact_version_id="d4c240ae-89f6-4adb-9f99-a0b14628234b" title="db/repositories/order_repository.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"
	"context"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{db: db.DB}
}

// Create adds a new order
// func (r *OrderRepository) Create(order *models.Order) error {
// 	return r.db.Create(order).Error
// }

func (r *OrderRepository) Create(ctx context.Context, order *models.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

// FindByID retrieves an order by ID with associated User and OrderItems
func (r *OrderRepository) FindByID(ctx context.Context, id uint) (*models.Order, error) {
	var order models.Order
	//err := r.db.Preload("User").Preload("OrderItems.Product.Merchant").First(&order, id).Error
	err := r.db.WithContext(ctx).Preload("User").Preload("OrderItems").Preload("OrderItems.Product").Preload("OrderItems.Merchant").First(&order, id).Error
	return &order, err
}

// FindByUserID retrieves all orders for a user
func (r *OrderRepository) FindByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("OrderItems.Product.Merchant").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

// FindByMerchantID retrieves all orders containing items from a merchant
func (r *OrderRepository) FindByMerchantID(merchantID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("OrderItems.Product").Joins("JOIN order_items oi ON oi.order_id = orders.id").
		Where("oi.merchant_id = ?", merchantID).Find(&orders).Error
	return orders, err
}

// Update modifies an existing order
func (r *OrderRepository) Update(order *models.Order) error {
	return r.db.Save(order).Error
}

// Delete removes an order by ID
func (r *OrderRepository) Delete(id uint) error {
	return r.db.Delete(&models.Order{}, id).Error
}
```
</xaiArtifact>

---
### db/repositories/payment_repository.go
- Size: 1.51 KB
- Lines: 52
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="17eec709-7d12-4a54-bd51-b34d3fbc865c" artifact_version_id="2f6a1ef9-04ce-47cd-bdfe-10cac7b05790" title="db/repositories/payment_repository.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{db: db.DB}
}

// Create adds a new payment
func (r *PaymentRepository) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

// FindByID retrieves a payment by ID with associated Order and User
func (r *PaymentRepository) FindByID(id uint) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.Preload("Order.User").First(&payment, id).Error
	return &payment, err
}

// FindByOrderID retrieves a payment by order ID
func (r *PaymentRepository) FindByOrderID(orderID uint) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.Preload("Order.User").Where("order_id = ?", orderID).First(&payment).Error
	return &payment, err
}

// FindByUserID retrieves all payments for a user
func (r *PaymentRepository) FindByUserID(userID uint) ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Preload("Order.User").Joins("JOIN orders ON orders.id = payments.order_id").Where("orders.user_id = ?", userID).Find(&payments).Error
	return payments, err
}

// Update modifies an existing payment
func (r *PaymentRepository) Update(payment *models.Payment) error {
	return r.db.Save(payment).Error
}

// Delete removes a payment by ID
func (r *PaymentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Payment{}, id).Error
}
```
</xaiArtifact>

---
### db/repositories/payout_repository.go
- Size: 1.18 KB
- Lines: 45
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="98a7f78e-9cb1-46ef-8be4-f72d50ee814e" artifact_version_id="83ca0ba0-6590-4821-ae07-c03d056649e9" title="db/repositories/payout_repository.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"

	"gorm.io/gorm"
)

type PayoutRepository struct {
	db *gorm.DB
}

func NewPayoutRepository() *PayoutRepository {
	return &PayoutRepository{db: db.DB}
}

// Create adds a new payout record
func (r *PayoutRepository) Create(payout *models.Payout) error {
	return r.db.Create(payout).Error
}

// FindByID retrieves a payout by ID with associated Merchant
func (r *PayoutRepository) FindByID(id uint) (*models.Payout, error) {
	var payout models.Payout
	err := r.db.Preload("Merchant").First(&payout, id).Error
	return &payout, err
}

// FindByMerchantID retrieves all payouts for a merchant
func (r *PayoutRepository) FindByMerchantID(merchantID uint) ([]models.Payout, error) {
	var payouts []models.Payout
	err := r.db.Preload("Merchant").Where("merchant_id = ?", merchantID).Find(&payouts).Error
	return payouts, err
}

// Update modifies an existing payout
func (r *PayoutRepository) Update(payout *models.Payout) error {
	return r.db.Save(payout).Error
}

// Delete removes a payout by ID
func (r *PayoutRepository) Delete(id uint) error {
	return r.db.Delete(&models.Payout{}, id).Error
}
```
</xaiArtifact>

---
### db/repositories/product_repo.go
- Size: 5.26 KB
- Lines: 171
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="3c80fe79-adac-413e-8c79-f3d40f2e0e78" artifact_version_id="b0cdfd77-bd61-4402-84d6-20745c4aaa7f" title="db/repositories/product_repo.go" contentType="text/go">
```go
package repositories

import (
	"errors"
	"fmt"
	"context"

	"api-customer-merchant/internal/api/dto"
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"

	"gorm.io/gorm"
)

var (
	ErrProductNotFound   = errors.New("product not found")
	ErrDuplicateSKU      = errors.New("duplicate SKU")
	ErrInvalidInventory  = errors.New("invalid inventory setup")
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{db: db.DB}
}

func (r *ProductRepository) FindBySKU(sku string) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("sku = ?", sku).First(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrProductNotFound
	} else if err != nil {
		return nil, fmt.Errorf("failed to find product by SKU: %w", err)
	}
	return &product, nil
}

func (r *ProductRepository) FindByID(id string, preloads ...string) (*models.Product, error) {
	var product models.Product
	query := r.db.Where("id = ?", id)
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	err := query.First(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrProductNotFound
	} else if err != nil {
		return nil, fmt.Errorf("failed to find product by ID: %w", err)
	}
	return &product, nil
}

func (r *ProductRepository) ListByMerchant(merchantID string, limit, offset int, filterActive bool) ([]models.Product, error) {
	var products []models.Product
	query := r.db.Where("merchant_id = ?", merchantID).Limit(limit).Offset(offset)
	if filterActive {
		query = query.Where("deleted_at IS NULL")
	}
	err := query.Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list products: %w", err)
	}
	return products, nil
}

func (r *ProductRepository) GetAllProducts(limit, offset int, categoryID *uint, preloads ...string) ([]models.Product, int64, error) {
	var products []models.Product
	query := r.db.Model(&models.Product{}).Where("deleted_at IS NULL")
	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	query = query.Limit(limit).Offset(offset).Order("created_at DESC")
	err := query.Find(&products).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to fetch products: %w", err)
	}
	return products, total, nil
}

func (r *ProductRepository) CreateProductWithVariantsAndInventory(ctx context.Context, product *models.Product, variants []models.Variant, variantInputs []dto.VariantInput, media []models.Media, simpleInitialStock *int, isSimple bool) error {
	if isSimple && len(variants) > 0 {
		return ErrInvalidInventory // Cannot have variants for simple products
	}
	if !isSimple && (len(variants) == 0 || len(variants) != len(variantInputs)) {
		return ErrInvalidInventory // Must provide matching variants and inputs
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create product
		if err := tx.Create(product).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return ErrDuplicateSKU
			}
			return fmt.Errorf("failed to create product: %w", err)
		}

		if !isSimple {
			// Variant-based product
			for i := range variants {
				variants[i].ProductID = product.ID
				if err := tx.Create(&variants[i]).Error; err != nil {
					return fmt.Errorf("failed to create variant: %w", err)
				}
				variantIDPtr := variants[i].ID
				inventory := models.VendorInventory{
					VariantID:         variantIDPtr,
					ProductID:         nil,
					MerchantID:        product.MerchantID,
					Quantity:          variantInputs[i].InitialStock,
					ReservedQuantity:  0,
					LowStockThreshold: 10,
					BackorderAllowed:  false,
				}
				if err := tx.Create(&inventory).Error; err != nil {
					return fmt.Errorf("failed to create variant inventory: %w", err)
				}
				variants[i].Inventory = inventory
			}
		}
		// Note: Skip VendorInventory creation for simple products

		// Create media
		for i := range media {
			media[i].ProductID = product.ID
			if err := tx.Create(&media[i]).Error; err != nil {
				return fmt.Errorf("failed to create media: %w", err)
			}
		}

		// Reload with preloads
		preloadQuery := tx.Where("id = ?", product.ID)
		if !isSimple {
			preloadQuery = preloadQuery.Preload("Variants.Inventory")
		}
		preloadQuery = preloadQuery.Preload("Media")
		if err := preloadQuery.First(product).Error; err != nil {
			return fmt.Errorf("failed to preload associations: %w", err)
		}

		return nil
	})
}

func (r *ProductRepository) UpdateInventoryQuantity(inventoryID string, delta int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var inventory models.VendorInventory
		if err := tx.First(&inventory, "id = ?", inventoryID).Error; err != nil {
			return fmt.Errorf("failed to find inventory: %w", err)
		}
		newQuantity := inventory.Quantity + delta
		if newQuantity < 0 && !inventory.BackorderAllowed {
			return errors.New("insufficient stock and backorders not allowed")
		}
		inventory.Quantity = newQuantity
		return tx.Save(&inventory).Error
	})
}


func (r *ProductRepository) SoftDeleteProduct(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.Product{}).Error
}
```
</xaiArtifact>

---
### db/repositories/product_repositry.go
- Size: 5.08 KB
- Lines: 169
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="5c4870ee-0cea-4f8e-ae88-a820b5129f54" artifact_version_id="036d6c29-745a-4083-a144-770a7ee0e68f" title="db/repositories/product_repositry.go" contentType="text/go">
```go
package repositories

// import (
// 	"api-customer-merchant/internal/db"
// 	"api-customer-merchant/internal/db/models"
// 	"errors"

// 	"gorm.io/gorm"
// )

/*
type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{db: db.DB}
}

// Create adds a new product to the database
func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// FindByID retrieves a product by ID with associated Merchant and Category
func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Merchant").Preload("Category").First(&product, id).Error
	return &product, err
}

// FindBySKU retrieves a product by SKU with associated Merchant and Category
func (r *ProductRepository) FindBySKU(sku string) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Merchant").Preload("Category").Where("sku = ?", sku).First(&product).Error
	return &product, err
}

// SearchByName retrieves products matching a name (partial match)
func (r *ProductRepository) SearchByName(name string) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Merchant").Preload("Category").Where("name ILIKE ?", "%"+name+"%").Find(&products).Error
	return products, err
}

// FindByMerchantID retrieves all products for a merchant
func (r *ProductRepository) FindByMerchantID(merchantID uint) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Category").Where("merchant_id = ?", merchantID).Find(&products).Error
	return products, err
}

// FindByCategoryID retrieves all products in a category
func (r *ProductRepository) FindByCategoryID(categoryID uint) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Merchant").Where("category_id = ?", categoryID).Find(&products).Error
	return products, err
}

// Update modifies an existing product
func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// Delete removes a product by ID
func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
// In ProductRepository
func (r *ProductRepository) FindByCategoryWithPagination(categoryID uint, limit, offset int) ([]models.Product, error) {
    var products []models.Product
    err := r.db.Preload("Merchant").Preload("Category").Where("category_id = ?", categoryID).Limit(limit).Offset(offset).Find(&products).Error
    return products, err
}
*/






















/*
type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{db: db.DB}
}

// Create adds a new product to the database
func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// FindByID retrieves a product by ID with associated Merchant and Category
func (r *ProductRepository) FindByID(id string) (*models.Product, error) {
    var product models.Product
    err := r.db.Preload("Merchant").Preload("Category").First(&product, id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("product not found")
        }
        return nil, err
    }
    return &product, nil
}

// FindBySKU retrieves a product by SKU with associated Merchant and Category
func (r *ProductRepository) FindBySKU(sku string) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Merchant").Preload("Category").Where("sku = ?", sku).First(&product).Error
	return &product, err
}

// SearchByName retrieves products matching a name (partial match)
func (r *ProductRepository) SearchByName(name string) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Merchant").Preload("Category").Where("name ILIKE ?", "%"+name+"%").Find(&products).Error
	return products, err
}

// FindByMerchantID retrieves all products for a merchant
func (r *ProductRepository) FindByMerchantID(merchantID string) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Category").Where("merchant_id = ?", merchantID).Find(&products).Error
	return products, err
}

// FindByCategoryID retrieves all products in a category
func (r *ProductRepository) FindByCategoryID(categoryID string) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Merchant").Where("category_id = ?", categoryID).Find(&products).Error
	return products, err
}

// Update modifies an existing product
func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// Delete removes a product by ID
func (r *ProductRepository) Delete(id string) error {
	return r.db.Delete(&models.Product{}, id).Error
}
// In ProductRepository
func (r *ProductRepository) FindByCategoryWithPagination(categoryID string, limit, offset int) ([]models.Product, error) {
    var products []models.Product
    err := r.db.Preload("Merchant").Preload("Category").Where("category_id = ?", categoryID).Limit(limit).Offset(offset).Find(&products).Error
    return products, err
}

*/
```
</xaiArtifact>

---
### db/repositories/user_repository.go
- Size: 1.11 KB
- Lines: 53
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="25426edc-655a-4b36-93e0-c0e563878946" artifact_version_id="da25ae70-38ff-4d38-bb48-d8b53e1c024b" title="db/repositories/user_repository.go" contentType="text/go">
```go
package repositories

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: db.DB}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("Failed to find user by email %s: %v", email, err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByGoogleID(googleID string) (*models.User, error) {
	var user models.User
	err := r.db.Where("google_id = ?", googleID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}
```
</xaiArtifact>

---
### middleware/auth.go
- Size: 1.45 KB
- Lines: 62
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="bcad33a3-0e37-4c76-86c5-6a88888edb54" artifact_version_id="e5aecba5-75c9-4928-8f3f-04acfb0a6d12" title="middleware/auth.go" contentType="text/go">
```go
package middleware

import (
	"net/http"
	"os"
	"strings"

	"api-customer-merchant/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(entityType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if utils.IsBlacklisted(tokenString) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is blacklisted"})
			c.Abort()
			return
		}
		key := os.Getenv("JWT_SECRET")

		secret := []byte(key) // Load from env
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return secret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || claims["entityType"] != entityType {
				c.JSON(http.StatusForbidden, gin.H{"error": "Invalid entity type"})
				c.Abort()
				return
			}

			//c.Set("entityId", claims["id"])
			id := claims["id"].(string)
        switch entityType {
case "user":
            c.Set("userID", id)
        case "merchant":
            c.Set("merchantID", id)
        }
        c.Next()
    }
	}

```
</xaiArtifact>

---
### middleware/rate_limit.go
- Size: 0.45 KB
- Lines: 20
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="97dce7f7-9974-47d7-ab7a-512dc1882580" artifact_version_id="8d5e4d7d-be84-45a6-a6f2-1b20feee7b18" title="middleware/rate_limit.go" contentType="text/go">
```go
package middleware

import (
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
)

func RateLimitMiddleware() gin.HandlerFunc {
    limiter := rate.NewLimiter(rate.Every(time.Minute), 100) // 100/min
    return func(c *gin.Context) {
        if !limiter.Allow() {
            c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
            c.Abort()
            return
        }
        c.Next()
    }
}
```
</xaiArtifact>

---
### services/cart/cart_service.go
- Size: 7.47 KB
- Lines: 258
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="522cc82e-0749-463a-a99c-adadf69868be" artifact_version_id="84702283-dbf0-4905-aee6-fa84572bc7d1" title="services/cart/cart_service.go" contentType="text/go">
```go
package cart

import (
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"
	"api-customer-merchant/internal/api/dto" // Assuming dto.BulkUpdateRequest is defined here with ProductID string, Quantity int
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ErrInvalidUserID     = errors.New("invalid user ID")
	ErrInvalidQuantity   = errors.New("quantity must be positive")
	ErrProductNotFound   = errors.New("product not found")
	ErrInventoryNotFound = errors.New("inventory not found")
	ErrInsufficientStock = errors.New("insufficient stock")
)

type CartService struct {
	cartRepo      *repositories.CartRepository
	cartItemRepo  *repositories.CartItemRepository
	productRepo   *repositories.ProductRepository
	inventoryRepo *repositories.InventoryRepository
	logger        *zap.Logger
	validator     *validator.Validate
}

func NewCartService(cartRepo *repositories.CartRepository, cartItemRepo *repositories.CartItemRepository, productRepo *repositories.ProductRepository, inventoryRepo *repositories.InventoryRepository, logger *zap.Logger) *CartService {
	return &CartService{
		cartRepo:      cartRepo,
		cartItemRepo:  cartItemRepo,
		productRepo:   productRepo,
		inventoryRepo: inventoryRepo,
		logger:        logger,
		validator:     validator.New(),
	}
}

// GetActiveCart retrieves or creates an active cart for a user
func (s *CartService) GetActiveCart(ctx context.Context, userID uint) (*models.Cart, error) {
	if userID == 0 {
		return nil, ErrInvalidUserID
	}
	cart, err := s.cartRepo.FindActiveCart(ctx, userID)
	if err == nil {
		return cart, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.Error("Failed to find active cart", zap.Uint("user_id", userID), zap.Error(err))
		return nil, fmt.Errorf("db error: %w", err)
	}

	cart = &models.Cart{UserID: userID, Status: models.CartStatusActive}
	if err := s.cartRepo.Create(ctx, cart); err != nil {
		s.logger.Error("Failed to create cart", zap.Error(err))
		return nil, fmt.Errorf("create failed: %w", err)
	}
	return s.cartRepo.FindByID(ctx, cart.ID)
}

func (s *CartService) GetCart(ctx context.Context, userID uint) (*models.Cart, error) {
	cart, err := s.GetActiveCart(ctx, userID)
	if err != nil {
		return nil, err
	}
	// Fixed: Use model method
	cart.ComputeTotals()
	s.logger.Info("Cart fetched", zap.Uint("user_id", userID), zap.Float64("total", cart.GrandTotal))
	return cart, nil
}

// AddItemToCart adds a product to the user's active cart
func (s *CartService) AddItemToCart(ctx context.Context, userID uint, quantity uint, productID string) (*models.Cart, error) {
	if userID == 0 {
		return nil, ErrInvalidUserID
	}
	if productID == "" {
		return nil, errors.New("invalid product ID")
	}
	if quantity == 0 {
		return nil, ErrInvalidQuantity
	}

	cart, err := s.GetActiveCart(ctx, userID)
	if err != nil {
		return nil, err
	}

	product, err := s.productRepo.FindByID(productID)
	if err != nil {
		return nil, ErrProductNotFound
	}

	inventory, err := s.inventoryRepo.FindByProductID(productID)
	if err != nil {
		return nil, ErrInventoryNotFound
	}
	if inventory.StockQuantity < int(quantity) {
		return nil, ErrInsufficientStock
	}

	err = db.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		cartItem, err := s.cartItemRepo.FindByProductIDAndCartID(ctx, productID, cart.ID)
		newQty := quantity
		if err == nil {
			newQty += uint(cartItem.Quantity)
			if inventory.StockQuantity < int(newQty) {
				return ErrInsufficientStock
			}
			if err := s.cartItemRepo.UpdateQuantityWithReservation(ctx, cartItem.ID, int(newQty), inventory.ID); err != nil {
				return err
			}
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			cartItem = &models.CartItem{
				CartID:     cart.ID,
				ProductID:  productID,
				Quantity:   int(quantity),
				MerchantID: product.MerchantID,
			}
			if err := s.cartItemRepo.Create(ctx, cartItem); err != nil {
				return err
			}
			if err := s.inventoryRepo.UpdateInventoryQuantity(ctx, inventory.ID, -int(quantity)); err != nil {
				return err
			}
		} else {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	// Reload
	return s.cartRepo.FindByID(ctx, cart.ID)
}

// UpdateCartItemQuantity updates the quantity of a cart item
func (s *CartService) UpdateCartItemQuantity(ctx context.Context, cartItemID uint, quantity int) (*models.Cart, error) {
	if cartItemID == 0 {
		return nil, errors.New("invalid cart item ID")
	}
	if quantity <= 0 {
		return nil, ErrInvalidQuantity
	}

	cartItem, err := s.cartItemRepo.FindByID(ctx, cartItemID)
	if err != nil {
		return nil, repositories.ErrCartItemNotFound
	}

	inventory, err := s.inventoryRepo.FindByProductID(cartItem.ProductID)
	if err != nil {
		return nil, ErrInventoryNotFound
	}
	if inventory.StockQuantity < quantity {
		return nil, ErrInsufficientStock
	}

	if err := s.cartItemRepo.UpdateQuantityWithReservation(ctx, cartItemID, quantity, inventory.ID); err != nil {
		return nil, err
	}

	return s.cartRepo.FindByID(ctx, cartItem.CartID)
}

func (s *CartService) RemoveCartItem(ctx context.Context, cartItemID uint) (*models.Cart, error) {
	if cartItemID == 0 {
		return nil, errors.New("invalid cart item ID")
	}

	cartItem, err := s.cartItemRepo.FindByID(ctx, cartItemID)
	if err != nil {
		return nil, repositories.ErrCartItemNotFound
	}

	inventory, err := s.inventoryRepo.FindByProductID(cartItem.ProductID)
	if err != nil {
		return nil, ErrInventoryNotFound
	}

	if err := s.cartItemRepo.DeleteWithUnreserve(ctx, cartItemID, inventory.ID); err != nil {
		return nil, err
	}

	return s.cartRepo.FindByID(ctx, cartItem.CartID)
}

func (s *CartService) GetCartItemByID(ctx context.Context, cartItemID uint) (*models.CartItem, error) {
	if cartItemID == 0 {
		return nil, errors.New("invalid cart item ID")
	}
	return s.cartItemRepo.FindByID(ctx, cartItemID)
}

// ClearCart, BulkAddItems ... (add ctx to all calls; stub Bulk if not used)
func (s *CartService) ClearCart(ctx context.Context, userID uint) error {
	cart, err := s.GetActiveCart(ctx, userID)
	if err != nil {
		return err
	}
	items, err := s.cartItemRepo.FindByCartID(ctx, cart.ID)
	if err != nil {
		return err
	}
	for _, item := range items {
		s.RemoveCartItem(ctx, item.ID)
	}
	cart.Status = models.CartStatusAbandoned
	return s.cartRepo.Update(ctx, cart)
}

// BulkAddItems stub (implement as needed; fixed DTO)
// func (s *CartService) BulkAddItems(ctx context.Context, userID uint, items []dto.BulkUpdateRequest) (*models.Cart, error) {
// 	// Validation loop...
// 	for _, item := range items {
// 		if err := s.validator.Struct(&item); err != nil {
// 			return nil, err
// 		}
// 		// Add each (loop AddItemToCart)
// 		_, err := s.AddItemToCart(ctx, userID, uint(item.Quantity), item.ProductID)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return s.GetCart(ctx, userID)
// }



func (s *CartService) BulkAddItems(ctx context.Context, userID uint, items dto.BulkUpdateRequest) (*models.Cart, error) {
	if userID == 0 {
		return nil, ErrInvalidUserID
	}
	if err := s.validator.Struct(&items); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// cart, err := s.GetActiveCart(ctx, userID)
	// if err != nil {
	// 	return nil, err
	// }

	for _, item := range items.Items {
		// Convert uint ProductID to string for consistency
		productID := fmt.Sprint(item.ProductID)
		if _, err := s.AddItemToCart(ctx, userID, uint(item.Quantity), productID); err != nil {
			return nil, fmt.Errorf("failed to add item %s: %w", productID, err)
		}
	}
	return s.GetCart(ctx, userID)
}
```
</xaiArtifact>

---
### services/merchant/merchant_service.go
- Size: 5.85 KB
- Lines: 207
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="e7392bbf-dfb3-4da2-a0d8-63441f8c5b82" artifact_version_id="22937357-4e7b-438a-b693-2c53eb889d49" title="services/merchant/merchant_service.go" contentType="text/go">
```go
package merchant

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
)

/*
type MerchantService struct {
	appRepo  *repositories.MerchantApplicationRepository
	repo     *repositories.MerchantRepository
	validate *validator.Validate
}

func NewMerchantService(appRepo *repositories.MerchantApplicationRepository, repo *repositories.MerchantRepository) *MerchantService {
	return &MerchantService{
		appRepo:  appRepo,
		repo:     repo,
		validate: validator.New(),
	}
}

// SubmitApplication allows a prospective merchant to submit an application.
func (s *MerchantService) SubmitApplication(ctx context.Context, app *models.MerchantApplication) (*models.MerchantApplication, error) {
	// Validate required fields
	if err := s.validate.Struct(app); err != nil {
		return nil, errors.New("validation failed: " + err.Error())
	}

	// Validate JSONB fields
	if len(app.PersonalAddress) == 0 {
		return nil, errors.New("personal_address cannot be empty")
	}
	if len(app.WorkAddress) == 0 {
		return nil, errors.New("work_address cannot be empty")
	}
	var temp map[string]interface{}
	if err := json.Unmarshal(app.PersonalAddress, &temp); err != nil {
		return nil, errors.New("invalid personal_address JSON")
	}
	if err := json.Unmarshal(app.WorkAddress, &temp); err != nil {
		return nil, errors.New("invalid work_address JSON")
	}




	// Set Status to pending and ensure ID is not set
	app.Status = "pending"
	app.ID = ""

	if err := s.appRepo.Create(ctx, app); err != nil {
		return nil, err
	}
	return app, nil
}

// GetApplication returns an application by ID (for applicant to check their own status).
func (s *MerchantService) GetApplication(ctx context.Context, id string) (*models.MerchantApplication, error) {
	if id == "" {
		return nil, errors.New("application ID cannot be empty")
	}
	return s.appRepo.GetByID(ctx, id)
}

// GetMerchantByUserID returns an active merchant account for the authenticated user.
func (s *MerchantService) GetMerchantByUserID(ctx context.Context, uid string) (*models.Merchant, error) {
	if uid == "" {
		return nil, errors.New("user ID cannot be empty")
	}
	return s.repo.GetByUserID(ctx, uid)
}

// GetMerchantByID returns an active merchant by ID.
func (s *MerchantService) GetMerchantByID(ctx context.Context, id string) (*models.Merchant, error) {
	if id == "" {
		return nil, errors.New("merchant ID cannot be empty")
	}
	return s.repo.GetByID(ctx, id)
}
*/
type MerchantService struct {
	appRepo  *repositories.MerchantApplicationRepository
	repo     *repositories.MerchantRepository
	validate *validator.Validate
}

func NewMerchantService(appRepo *repositories.MerchantApplicationRepository, repo *repositories.MerchantRepository) *MerchantService {
	return &MerchantService{
		appRepo:  appRepo,
		repo:     repo,
		validate: validator.New(),
	}
}

// SubmitApplication allows a prospective merchant to submit an application.
func (s *MerchantService) SubmitApplication(ctx context.Context, app *models.MerchantApplication) (*models.MerchantApplication, error) {
	// Validate required fields
	if err := s.validate.Struct(app); err != nil {
		return nil, errors.New("validation failed: " + err.Error())
	}

	// Validate JSONB fields
	if len(app.PersonalAddress) == 0 {
		return nil, errors.New("personal_address cannot be empty")
	}
	if len(app.WorkAddress) == 0 {
		return nil, errors.New("work_address cannot be empty")
	}
	var temp map[string]interface{}
	if err := json.Unmarshal(app.PersonalAddress, &temp); err != nil {
		return nil, errors.New("invalid personal_address JSON")
	}
	if err := json.Unmarshal(app.WorkAddress, &temp); err != nil {
		return nil, errors.New("invalid work_address JSON")
	}

	
	

	// Set Status to pending and ensure ID is not set
	app.Status = "pending"
	app.ID = ""

	if err := s.appRepo.Create(ctx, app); err != nil {
		return nil, err
	}
	return app, nil
}

// GetApplication returns an application by ID (for applicant to check their own status).
func (s *MerchantService) GetApplication(ctx context.Context, id string) (*models.MerchantApplication, error) {
	if id == "" {
		return nil, errors.New("application ID cannot be empty")
	}
	return s.appRepo.GetByID(ctx, id)
}

// GetMerchantByUserID returns an active merchant account for the authenticated user.
func (s *MerchantService) GetMerchantByUserID(ctx context.Context, uid string) (*models.Merchant, error) {
	if uid == "" {
		return nil, errors.New("user ID cannot be empty")
	}
	return s.repo.GetByUserID(ctx, uid)
}

// GetMerchantByID returns an active merchant by ID.
func (s *MerchantService) GetMerchantByID(ctx context.Context, id string) (*models.Merchant, error) {
	if id == "" {
		return nil, errors.New("merchant ID cannot be empty")
	}
	return s.repo.GetByID(ctx, id)
}




func (s *MerchantService) LoginMerchant(ctx context.Context,work_email, password string) (*models.Merchant, error) {
    merchant, err := s.repo.GetByWorkEmail(ctx,work_email)
    if err != nil {
        return nil, errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(merchant.Password), []byte(password)); err != nil {
        return nil, errors.New("invalid credentials")
    }

    return merchant, nil
}

func (s *MerchantService) GenerateJWT(entity interface{}) (string, error) {
	var id string
	var entityType string

	switch e := entity.(type) {
	case *models.Merchant:
		id = e.ID
		entityType = "merchant"

	default:
		return "", errors.New("invalid entity type")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        id,
		"entityType": entityType,
        "exp":        time.Now().Add(24 * time.Hour).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET not set")
	}

	return token.SignedString([]byte(secret))
}
```
</xaiArtifact>

---
### services/notifications/notifcation_service.go
- Size: 0.74 KB
- Lines: 31
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="49ae36fc-6bf9-4efb-a47a-f01ae9f8d65d" artifact_version_id="f2b5d743-5eaf-4de6-83b9-615478aea6bb" title="services/notifications/notifcation_service.go" contentType="text/go">
```go
package notifications

import (
    "context"
    "fmt"
    // Assume SMTP or Twilio lib
)

type NotificationService struct {
    // Email/SMS clients
}

func NewNotificationService() *NotificationService {
    return &NotificationService{}
}

func (s *NotificationService) SendOrderConfirmation(ctx context.Context, orderID uint, email string) error {
    // Use template: "Your order {orderID} is confirmed"
    fmt.Println("Sending email to", email) // Replace with real send
    return nil
}

func (s *NotificationService) SendMerchantNewOrder(merchantID uint, orderID uint) error {
    // Fetch merchant email, send
    return nil
}

func (s *NotificationService) SendStockAlert(merchantID uint, productID uint) error {
    // On low stock
    return nil
}
```
</xaiArtifact>

---
### services/order/order_service.go
- Size: 8.90 KB
- Lines: 324
- Last Modified: 2025-09-19 21:51:30

<xaiArtifact artifact_id="bd72086e-ea21-44af-bcbf-e74a85a967f5" artifact_version_id="3b01a637-83e7-4489-9d28-e99bf965c877" title="services/order/order_service.go" contentType="text/go">
```go
//package order

// import (
// 	"errors"
// 	"api-customer-merchant/internal/db/models"
// 	"api-customer-merchant/internal/db/repositories"

// 	"gorm.io/gorm"
// )

/*

type OrderService struct {
	orderRepo     *repositories.OrderRepository
	orderItemRepo *repositories.OrderItemRepository
	cartRepo      *repositories.CartRepository
	cartItemRepo  *repositories.CartItemRepository
	inventoryRepo *repositories.InventoryRepository
}

func NewOrderService(orderRepo *repositories.OrderRepository, orderItemRepo *repositories.OrderItemRepository, cartRepo *repositories.CartRepository, cartItemRepo *repositories.CartItemRepository, inventoryRepo *repositories.InventoryRepository) *OrderService {
	return &OrderService{
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
		cartRepo:      cartRepo,
		cartItemRepo:  cartItemRepo,
		inventoryRepo: inventoryRepo,
	}
}

// CreateOrderFromCart creates an order from the user's active cart
func (s *OrderService) CreateOrderFromCart(userID uint) (*models.Order, error) {
	if userID == 0 {
		return nil, errors.New("invalid user ID")
	}

	// Get active cart
	cart, err := s.cartRepo.FindActiveCart(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no active cart found")
		}
		return nil, err
	}

	// Check if cart has items
	cartItems, err := s.cartItemRepo.FindByCartID(cart.ID)
	if err != nil {
		return nil, err
	}
	if len(cartItems) == 0 {
		return nil, errors.New("cart is empty")
	}

	// Validate stock for all items
	for _, item := range cartItems {
		inventory, err := s.inventoryRepo.FindByProductID(item.ProductID)
		if err != nil {
			return nil, errors.New("inventory not found for product")
		}
		if inventory.StockQuantity < item.Quantity {
			return nil, errors.New("insufficient stock for product")
		}
	}

	// Calculate total amount
	var totalAmount float64
	for _, item := range cartItems {
		totalAmount += float64(item.Quantity) * item.Product.Price
	}

	// Create order
	order := &models.Order{
		UserID:      userID,
		TotalAmount: totalAmount,
		Status:      models.OrderStatusPending,
	}
	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}

	// Create order items and update inventory
	for _, item := range cartItems {
		orderItem := &models.OrderItem{
			OrderID:    order.ID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			Price:      item.Product.Price,
			MerchantID: item.MerchantID,
		}
		if err := s.orderItemRepo.Create(orderItem); err != nil {
			return nil, err
		}
		// Update inventory
		if err := s.inventoryRepo.UpdateStock(item.ProductID, -item.Quantity); err != nil {
			return nil, err
		}
	}

	// Mark cart as converted
	cart.Status = models.CartStatusConverted
	if err := s.cartRepo.Update(cart); err != nil {
		return nil, err
	}

	return s.orderRepo.FindByID(order.ID)
}

// GetOrderByID retrieves an order by ID
func (s *OrderService) GetOrderByID(id uint) (*models.Order, error) {
	if id == 0 {
		return nil, errors.New("invalid order ID")
	}
	return s.orderRepo.FindByID(id)
}

// GetOrdersByUserID retrieves all orders for a user
func (s *OrderService) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	if userID == 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.orderRepo.FindByUserID(userID)
}

// GetOrdersByMerchantID retrieves orders containing a merchant's products
func (s *OrderService) GetOrdersByMerchantID(merchantID uint) ([]models.Order, error) {
	if merchantID == 0 {
		return nil, errors.New("invalid merchant ID")
	}
	return s.orderRepo.FindByMerchantID(merchantID)
}

// UpdateOrderStatus updates the status of an order
func (s *OrderService) UpdateOrderStatus(orderID uint, status string) (*models.Order, error) {
	if orderID == 0 {
		return nil, errors.New("invalid order ID")
	}
	if err := models.OrderStatus(status).Valid(); err != nil {
		return nil, err
	}

	order, err := s.orderRepo.FindByID(orderID)
	if err != nil {
		return nil, err
	}

	order.Status = models.OrderStatus(status)
	if err := s.orderRepo.Update(order); err != nil {
		return nil, err
	}

	return s.orderRepo.FindByID(orderID)
}
*/



package order

import (
	"context"
	"errors"
	"fmt"

	"api-customer-merchant/internal/api/dto"
	"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"
	"gorm.io/gorm"
)

// OrderService provides business logic for handling orders.
type OrderService struct {
	orderRepo     *repositories.OrderRepository
	orderItemRepo *repositories.OrderItemRepository
	cartRepo      *repositories.CartRepository
	cartItemRepo  *repositories.CartItemRepository
	productRepo   *repositories.ProductRepository
	db            *gorm.DB
}

// NewOrderService creates a new instance of OrderService.
func NewOrderService(
	orderRepo *repositories.OrderRepository,
	orderItemRepo *repositories.OrderItemRepository,
	cartRepo *repositories.CartRepository,
	cartItemRepo *repositories.CartItemRepository,
	productRepo *repositories.ProductRepository,
) *OrderService {
	return &OrderService{
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
		cartRepo:      cartRepo,
		cartItemRepo:  cartItemRepo,
		productRepo:   productRepo,
		db:            db.DB,
	}
}

// CreateOrder converts a user's active cart into an order.
// It performs several actions within a single database transaction:
// 1. Finds the user's active cart.
// 2. Validates that the cart is not empty.
// 3. For each item in the cart, it moves the reserved stock to committed stock.
// 4. Creates an Order record.
// 5. Creates OrderItem records corresponding to the CartItems.
// 6. Deletes the cart items.
// 7. Updates the cart status to 'Converted'.
// 8. Returns a DTO representing the newly created order.
func (s *OrderService) CreateOrder(ctx context.Context, userID uint) (*dto.OrderResponse, error) {
	if userID == 0 {
		return nil, errors.New("invalid user ID")
	}

	cart, err := s.cartRepo.FindActiveCart(ctx, userID)
	if err != nil {
		if errors.Is(err, repositories.ErrCartNotFound) {
			return nil, errors.New("no active cart found")
		}
		return nil, err
	}

	if len(cart.CartItems) == 0 {
		return nil, errors.New("cart is empty")
	}

	var newOrder *models.Order
	var totalAmount float64

	// Use a transaction to ensure atomicity
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Calculate total and create order items
		var orderItems []models.OrderItem
		for _, item := range cart.CartItems {
			price := item.Product.BasePrice.InexactFloat64()
			totalAmount += float64(item.Quantity) * price
			orderItems = append(orderItems, models.OrderItem{
				ProductID:  item.ProductID,
				MerchantID: item.Product.MerchantID,
				Quantity:   item.Quantity,
				Price:      price,
			})

			// Here you would typically move reserved stock to committed stock.
			// For now, we assume cart reservation handled this.
			// We'll just update the main inventory.
			// This logic might need to be more robust depending on inventory strategy.
			if err := tx.Model(&models.VendorInventory{}).
				Where("product_id = ?", item.ProductID).
				Updates(map[string]interface{}{
					"quantity":          gorm.Expr("quantity - ?", item.Quantity),
					"reserved_quantity": gorm.Expr("reserved_quantity - ?", item.Quantity),
				}).Error; err != nil {
				return fmt.Errorf("failed to commit stock for product %s: %w", item.ProductID, err)
			}
		}

		// Create the order
		newOrder = &models.Order{
			UserID:      userID,
			TotalAmount: totalAmount,
			Status:      models.OrderStatusPending,
			OrderItems:  orderItems,
		}
		if err := tx.Create(newOrder).Error; err != nil {
			return err
		}

		// Associate order items with the new order ID and create them
		for i := range orderItems {
			orderItems[i].OrderID = newOrder.ID
		}
		if err := tx.Create(&orderItems).Error; err != nil {
			return err
		}

		// Manually associate for the response DTO
		newOrder.OrderItems = orderItems

		// Clear cart items
		if err := tx.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error; err != nil {
			return err
		}

		// Mark cart as converted
		cart.Status = models.CartStatusConverted
		if err := tx.Save(cart).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Convert to DTO for response
	orderResponse := &dto.OrderResponse{
		ID:         newOrder.ID,
		UserID:     newOrder.UserID,
		Status:     string(newOrder.Status),
		OrderItems: make([]dto.OrderItemResponse, len(newOrder.OrderItems)),
	}
	for i, item := range newOrder.OrderItems {
		orderResponse.OrderItems[i] = dto.OrderItemResponse{
			ProductID: fmt.Sprint(item.ProductID),
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
	}

	return orderResponse, nil
}

// GetOrder retrieves a single order by its ID.
func (s *OrderService) GetOrder(ctx context.Context, id uint) (*models.Order, error) {
	if id == 0 {
		return nil, errors.New("invalid order ID")
	}
	// The repository already preloads necessary associations.
	return s.orderRepo.FindByID(ctx, id)
	//return s.orderRepo.FindByID(id)
}

```
</xaiArtifact>

---
### services/payment/payment_service.go
- Size: 2.54 KB
- Lines: 97
- Last Modified: 2025-09-19 21:51:57

<xaiArtifact artifact_id="bdb0557b-0e3c-4dc3-b5b2-11d361816b90" artifact_version_id="e910e97f-7c9f-49fa-ba91-37fc20489c7a" title="services/payment/payment_service.go" contentType="text/go">
```go
package payment

import (
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"
	"context"
	"errors"
)

type PaymentService struct {
	paymentRepo *repositories.PaymentRepository
	orderRepo   *repositories.OrderRepository
}

func NewPaymentService(paymentRepo *repositories.PaymentRepository, orderRepo *repositories.OrderRepository) *PaymentService {
	return &PaymentService{
		paymentRepo: paymentRepo,
		orderRepo:   orderRepo,
	}
}

// ProcessPayment creates a payment for an order (placeholder for Stripe)
func (s *PaymentService) ProcessPayment(ctx context.Context, orderID uint, amount float64) (*models.Payment, error) {
	if orderID == 0 {
		return nil, errors.New("invalid order ID")
	}
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}

	// Verify order exists
	order, err := s.orderRepo.FindByID(ctx,orderID)
	if err != nil {
		return nil, errors.New("order not found")
	}

	// Verify order amount matches
	if order.TotalAmount != amount {
		return nil, errors.New("payment amount does not match order total")
	}

	// Simulate Stripe payment processing
	payment := &models.Payment{
		OrderID: orderID,
		Amount:  amount,
		Status:  models.PaymentStatusPending,
	}
	if err := s.paymentRepo.Create(payment); err != nil {
		return nil, err
	}

	// Simulate successful payment
	payment.Status = models.PaymentStatusCompleted
	if err := s.paymentRepo.Update(payment); err != nil {
		return nil, err
	}

	return s.paymentRepo.FindByID(payment.ID)
}

// GetPaymentByOrderID retrieves a payment by order ID
func (s *PaymentService) GetPaymentByOrderID(orderID uint) (*models.Payment, error) {
	if orderID == 0 {
		return nil, errors.New("invalid order ID")
	}
	return s.paymentRepo.FindByOrderID(orderID)
}

// GetPaymentsByUserID retrieves all payments for a user
func (s *PaymentService) GetPaymentsByUserID(userID uint) ([]models.Payment, error) {
	if userID == 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.paymentRepo.FindByUserID(userID)
}

// UpdatePaymentStatus updates the status of a payment
func (s *PaymentService) UpdatePaymentStatus(paymentID uint, status string) (*models.Payment, error) {
	if paymentID == 0 {
		return nil, errors.New("invalid payment ID")
	}
	if err := models.PaymentStatus(status).Valid(); err != nil {
		return nil, err
	}

	payment, err := s.paymentRepo.FindByID(paymentID)
	if err != nil {
		return nil, err
	}

	payment.Status = models.PaymentStatus(status)
	if err := s.paymentRepo.Update(payment); err != nil {
		return nil, err
	}

	return s.paymentRepo.FindByID(paymentID)
}
```
</xaiArtifact>

---
### services/payout/payout_service.go
- Size: 1.54 KB
- Lines: 62
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="6755782f-fd82-4d12-b70b-b4480885a082" artifact_version_id="98277fb4-eb58-4534-b01c-ffbd82733576" title="services/payout/payout_service.go" contentType="text/go">
```go
package payout


import (
	"errors"
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"
)

type PayoutService struct {
	payoutRepo *repositories.PayoutRepository
}

func NewPayoutService(payoutRepo *repositories.PayoutRepository) *PayoutService {
	return &PayoutService{
		payoutRepo: payoutRepo,
	}
}

// CreatePayout creates a payout for a merchant
func (s *PayoutService) CreatePayout(merchantID uint, amount float64) (*models.Payout, error) {
	if merchantID == 0 {
		return nil, errors.New("invalid merchant ID")
	}
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}

	// Simulate payout processing (placeholder for Stripe)
	payout := &models.Payout{
		MerchantID: merchantID,
		Amount:     amount,
		Status:     models.PayoutStatusPending,
	}
	if err := s.payoutRepo.Create(payout); err != nil {
		return nil, err
	}

	// Simulate successful payout
	payout.Status = models.PayoutStatusCompleted
	if err := s.payoutRepo.Update(payout); err != nil {
		return nil, err
	}

	return s.payoutRepo.FindByID(payout.ID)
}

// GetPayoutByID retrieves a payout by ID
func (s *PayoutService) GetPayoutByID(id uint) (*models.Payout, error) {
	if id == 0 {
		return nil, errors.New("invalid payout ID")
	}
	return s.payoutRepo.FindByID(id)
}

// GetPayoutsByMerchantID retrieves all payouts for a merchant
func (s *PayoutService) GetPayoutsByMerchantID(merchantID uint) ([]models.Payout, error) {
	if merchantID == 0 {
		return nil, errors.New("invalid merchant ID")
	}
	return s.payoutRepo.FindByMerchantID(merchantID)
}
```
</xaiArtifact>

---
### services/pricing/pricing_service.go
- Size: 1.04 KB
- Lines: 41
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="e586a0c2-758c-484b-8218-60ed1f29bfb7" artifact_version_id="6b26e360-2795-4fdf-a2a7-a61b8f55f069" title="services/pricing/pricing_service.go" contentType="text/go">
```go
package pricing
/*
import (
    "errors"
    "api-customer-merchant/internal/db/models"
)

type PricingService struct {
    // Repos if needed
}

func NewPricingService() *PricingService {
    return &PricingService{}
}

func (s *PricingService) CalculateShipping(cart *models.Cart, address string) (float64, error) {
    // Integrate with shipping API (e.g., UPS); placeholder
    if address == "" {
        return 0, errors.New("address required")
    }
    return 10.00, nil // Flat rate per vendor count
}

func (s *PricingService) CalculateTax(cart *models.Cart, country string) (float64, error) {
    // Tax API or rules; placeholder
    var subtotal float64
    for _, item := range cart.CartItems {
        subtotal += float64(item.Quantity) * item.Product.Price
    }
    rate := 0.1 // 10% VAT for luxury
    return subtotal * rate, nil
}

func (s *PricingService) ApplyPromotion(cart *models.Cart, code string) (float64, error) {
    // Validate coupon; placeholder
    if code == "" {
        return 0, nil
    }
    return 5.00, nil // Discount
}
*/
```
</xaiArtifact>

---
### services/product/product_service.go
- Size: 8.63 KB
- Lines: 317
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="08a56441-f27c-4594-b267-5bed97899116" artifact_version_id="c99f3392-87d4-4335-bce2-bb7803f72b76" title="services/product/product_service.go" contentType="text/go">
```go
package product

/*
type ProductService struct {
	productRepo  *repositories.ProductRepository
	inventoryRepo *repositories.InventoryRepository
}

func NewProductService(productRepo *repositories.ProductRepository, inventoryRepo *repositories.InventoryRepository) *ProductService {
	return &ProductService{
		productRepo:  productRepo,
		inventoryRepo: inventoryRepo,
	}
}

// GetProductByID retrieves a product by its ID
func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	if id == 0 {
		return nil, errors.New("invalid product ID")
	}
	return s.productRepo.FindByID(id)
}

// GetProductBySKU retrieves a product by its SKU
func (s *ProductService) GetProductBySKU(sku string) (*models.Product, error) {
	if strings.TrimSpace(sku) == "" {
		return nil, errors.New("SKU cannot be empty")
	}
	return s.productRepo.FindBySKU(sku)
}

// SearchProductsByName searches for products by name (case-insensitive)
func (s *ProductService) SearchProductsByName(name string) ([]models.Product, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("search name cannot be empty")
	}
	return s.productRepo.SearchByName(name)
}

// GetProductsByCategory retrieves products in a category
// In ProductService
func (s *ProductService) GetProductsByCategory(categoryID uint, limit, offset int) ([]models.Product, error) {
    if categoryID == 0 {
        return nil, errors.New("invalid category ID")
    }
    return s.productRepo.FindByCategoryWithPagination(categoryID, limit, offset)
}

// CreateProduct creates a new product for a merchant
func (s *ProductService) CreateProduct(product *models.Product, merchantID uint) error {
	if product == nil {
		return errors.New("product cannot be nil")
	}
	if merchantID == 0 {
		return errors.New("invalid merchant ID")
	}
	if strings.TrimSpace(product.Name) == "" {
		return errors.New("product name cannot be empty")
	}
	if strings.TrimSpace(product.SKU) == "" {
		return errors.New("SKU cannot be empty")
	}
	if product.Price <= 0 {
		return errors.New("price must be positive")
	}
	if product.CategoryID == 0 {
		return errors.New("category ID must be set")
	}

	// Check if SKU is unique
	if _, err := s.productRepo.FindBySKU(product.SKU); err == nil {
		return errors.New("SKU already exists")
	}

	product.MerchantID = merchantID
	return s.productRepo.Create(product)
}

// UpdateProduct updates an existing product (merchant only)
func (s *ProductService) UpdateProduct(product *models.Product, merchantID uint) error {
	if product == nil || product.ID == 0 {
		return errors.New("invalid product or product ID")
	}
	if merchantID == 0 {
		return errors.New("invalid merchant ID")
	}
	if strings.TrimSpace(product.Name) == "" {
		return errors.New("product name cannot be empty")
	}
	if strings.TrimSpace(product.SKU) == "" {
		return errors.New("SKU cannot be empty")
	}
	if product.Price <= 0 {
		return errors.New("price must be positive")
	}
	if product.CategoryID == 0 {
		return errors.New("category ID must be set")
	}

	// Verify product belongs to merchant
	existing, err := s.productRepo.FindByID(product.ID)
	if err != nil {
		return err
	}
	if existing.MerchantID != merchantID {
		return errors.New("product does not belong to merchant")
	}

	// Check if SKU is unique (excluding current product)
	if p, err := s.productRepo.FindBySKU(product.SKU); err == nil && p.ID != product.ID {
		return errors.New("SKU already exists")
	}

	return s.productRepo.Update(product)
}

// DeleteProduct deletes a product (merchant only)
func (s *ProductService) DeleteProduct(id uint, merchantID uint) error {
	if id == 0 {
		return errors.New("invalid product ID")
	}
	if merchantID == 0 {
		return errors.New("invalid merchant ID")
	}

	// Verify product belongs to merchant
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return err
	}
	if product.MerchantID != merchantID {
		return errors.New("product does not belong to merchant")
	}

	return s.productRepo.Delete(id)
}
*/








/*

import (
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"
	"api-customer-merchant/internal/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type ProductService struct {
	productRepo  *repositories.ProductRepository
	inventoryRepo *repositories.InventoryRepository
}

func NewProductService(productRepo *repositories.ProductRepository, inventoryRepo *repositories.InventoryRepository) *ProductService {
	return &ProductService{
		productRepo:  productRepo,
		inventoryRepo: inventoryRepo,
	}
}

// GetProductByID retrieves a product by its ID
func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
    if id == "" {
        return nil, errors.New("invalid product ID")
    }

    ctx := context.Background()
    key := fmt.Sprintf("product:%s", id)
    data, err := utils.GetOrSetCache(ctx, key, 10*time.Minute, func() (any, error) {
        product, err := s.productRepo.FindByID(id)
        if err != nil {
            return nil, err
        }
        return json.Marshal(product) // Serialize
    })
    if err != nil {
        return nil, err
    }

    var product models.Product
    if err := json.Unmarshal([]byte(data.(string)), &product); err != nil {
        return nil, err
    }
    return &product, nil
}

// GetProductBySKU retrieves a product by its SKU
func (s *ProductService) GetProductBySKU(sku string) (*models.Product, error) {
	if strings.TrimSpace(sku) == "" {
		return nil, errors.New("SKU cannot be empty")
	}
	return s.productRepo.FindBySKU(sku)
}

// SearchProductsByName searches for products by name (case-insensitive)
func (s *ProductService) SearchProductsByName(name string) ([]models.Product, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("search name cannot be empty")
	}
	return s.productRepo.SearchByName(name)
}

// GetProductsByCategory retrieves products in a category
// In ProductService
func (s *ProductService) GetProductsByCategory(categoryID string, limit, offset int) ([]models.Product, error) {
    if categoryID == "" {
        return nil, errors.New("invalid category ID")
    }
    return s.productRepo.FindByCategoryWithPagination(categoryID, limit, offset)
}

// CreateProduct creates a new product for a merchant
func (s *ProductService) CreateProduct(product *models.Product, merchantID string) error {
	if product == nil {
		return errors.New("product cannot be nil")
	}
	if merchantID == "" {
		return errors.New("invalid merchant ID")
	}
	if strings.TrimSpace(product.Name) == "" {
		return errors.New("product name cannot be empty")
	}
	if strings.TrimSpace(product.SKU) == "" {
		return errors.New("SKU cannot be empty")
	}
	if product.Price <= 0 {
		return errors.New("price must be positive")
	}
	if product.CategoryID == 0 {
		return errors.New("category ID must be set")
	}

	// Check if SKU is unique
	if _, err := s.productRepo.FindBySKU(product.SKU); err == nil {
		return errors.New("SKU already exists")
	}

	product.MerchantID = merchantID
	return s.productRepo.Create(product)
}

// UpdateProduct updates an existing product (merchant only)
func (s *ProductService) UpdateProduct(product *models.Product, merchantID string) error {
	if product == nil || product.ID == "" {
		return errors.New("invalid product or product ID")
	}
	if merchantID == "" {
		return errors.New("invalid merchant ID")
	}
	if strings.TrimSpace(product.Name) == "" {
		return errors.New("product name cannot be empty")
	}
	if strings.TrimSpace(product.SKU) == "" {
		return errors.New("SKU cannot be empty")
	}
	if product.Price <= 0 {
		return errors.New("price must be positive")
	}
	if product.CategoryID == 0 {
		return errors.New("category ID must be set")
	}

	// Verify product belongs to merchant
	existing, err := s.productRepo.FindByID(product.ID)
	if err != nil {
		return err
	}
	if existing.MerchantID != merchantID {
		return errors.New("product does not belong to merchant")
	}

	// Check if SKU is unique (excluding current product)
	if p, err := s.productRepo.FindBySKU(product.SKU); err == nil && p.ID != product.ID {
		return errors.New("SKU already exists")
	}

	return s.productRepo.Update(product)
}

// DeleteProduct deletes a product (merchant only)
func (s *ProductService) DeleteProduct(id string, merchantID string) error {
	if id == "" {
		return errors.New("invalid product ID")
	}
	if merchantID == "" {
		return errors.New("invalid merchant ID")
	}

	// Verify product belongs to merchant
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return err
	}
	if product.MerchantID != merchantID {
		return errors.New("product does not belong to merchant")
	}

	return s.productRepo.Delete(id)
}

func (s *ProductService) GetProductsByMerchantID(merchantID string) ([]models.Product, error) {
	return s.productRepo.FindByMerchantID(merchantID)
}


*/
```
</xaiArtifact>

---
### services/product/service_product.go
- Size: 14.25 KB
- Lines: 421
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="6ce617d7-2ca4-43b5-9418-f677fbfec3be" artifact_version_id="44e40229-4480-4f42-9d1e-456ad44006bd" title="services/product/service_product.go" contentType="text/go">
```go
package product

import (
	"context"
	"errors"
	"fmt"

	//"net/url"
	"regexp"
	"strings"

	//"github.com/shopspring/decimal"
	"api-customer-merchant/internal/api/dto" // Assuming this exists for VariantInput
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

var (
	ErrInvalidProduct     = errors.New("invalid product data")
	ErrInvalidSKU         = errors.New("invalid SKU format")
	ErrInvalidMediaURL    = errors.New("invalid media URL")
	ErrInvalidAttributes  = errors.New("invalid variant attributes")
	ErrUnauthorized       = errors.New("unauthorized operation")
)

// SKU validation regex: alphanumeric, hyphens, underscores, max 100 chars
var skuRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]{1,100}$`)

type ProductService struct {
	productRepo *repositories.ProductRepository
	logger      *zap.Logger
	validator   *validator.Validate
}

func NewProductService(productRepo *repositories.ProductRepository, logger *zap.Logger) *ProductService {
	return &ProductService{
		productRepo: productRepo,
		logger:      logger,
		validator:   validator.New(),
	}
}

// CreateProductWithVariants creates a product from input DTO
func (s *ProductService) CreateProductWithVariants(ctx context.Context, input *dto.ProductInput) (*dto.ProductResponse, error) {
	logger := s.logger.With(zap.String("operation", "CreateProductWithVariants"), )

	// Validate input
	if err := s.validator.Struct(input); err != nil {
		logger.Error("Input validation failed", zap.Error(err))
		return nil, fmt.Errorf("invalid input: %w", err)
	}

	// Additional validation
	if !skuRegex.MatchString(input.SKU) {
		logger.Error("Invalid SKU format", zap.String("sku", input.SKU))
		return nil, ErrInvalidSKU
	}
	isSimple := len(input.Variants) == 0
	if isSimple && input.InitialStock == nil {
		logger.Error("Initial stock required for simple product")
		return nil, ErrInvalidProduct
	}
	for _, v := range input.Variants {
		if !skuRegex.MatchString(v.SKU) {
			logger.Error("Invalid variant SKU format", zap.String("sku", v.SKU))
			return nil, ErrInvalidSKU
		}
	}

	// Check SKU uniqueness
	if _, err := s.productRepo.FindBySKU(input.SKU); err == nil {
		logger.Warn("Duplicate product SKU", zap.String("sku", input.SKU))
		return nil, fmt.Errorf("product SKU %s already exists", input.SKU)
	} else if !errors.Is(err, repositories.ErrProductNotFound) {
		logger.Error("Failed to check product SKU", zap.Error(err))
		return nil, fmt.Errorf("failed to check SKU: %w", err)
	}
	for _, v := range input.Variants {
		if _, err := s.productRepo.FindBySKU(v.SKU); err == nil {
			logger.Warn("Duplicate variant SKU", zap.String("sku", v.SKU))
			return nil, fmt.Errorf("variant SKU %s already exists", v.SKU)
		} else if !errors.Is(err, repositories.ErrProductNotFound) {
			logger.Error("Failed to check variant SKU", zap.Error(err))
			return nil, fmt.Errorf("failed to check variant SKU: %w", err)
		}
	}

	// Map DTO to models
	product := &models.Product{
		Name:        strings.TrimSpace(input.Name),
		MerchantID: strings.TrimSpace(input.MerchantID),
		Description: strings.TrimSpace(input.Description),
		SKU:         strings.TrimSpace(input.SKU),
		BasePrice:   decimal.NewFromFloat(input.BasePrice),
		CategoryID:  input.CategoryID,
	}
	variants := make([]models.Variant, len(input.Variants))
	for i, v := range input.Variants {
		variants[i] = models.Variant{
			SKU:            strings.TrimSpace(v.SKU),
			PriceAdjustment: decimal.NewFromFloat(v.PriceAdjustment),
			Attributes:     v.Attributes,
			IsActive:       true,
		}
	}
	media := make([]models.Media, len(input.Media))
	for i, m := range input.Media {
		media[i] = models.Media{
			URL:  strings.TrimSpace(m.URL),
			Type: models.MediaType(m.Type),
		}
	}

	// Delegate to repo
	err := s.productRepo.CreateProductWithVariantsAndInventory(ctx,product, variants, input.Variants, media, nil, isSimple)
	if err != nil {
		if errors.Is(err, repositories.ErrDuplicateSKU) {
			return nil, fmt.Errorf("duplicate SKU: %w", err)
		}
		if errors.Is(err, repositories.ErrInvalidInventory) {
			return nil, fmt.Errorf("invalid inventory setup: %w", err)
		}
		logger.Error("Failed to create product", zap.Error(err))
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	// Map to response DTO
	response := &dto.ProductResponse{
		ID:          product.ID,
		MerchantID:  product.MerchantID,
		Name:        product.Name,
		Description: product.Description,
		SKU:         product.SKU,
		BasePrice:   (product.BasePrice).InexactFloat64(),
		CategoryID:  product.CategoryID,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		Media:       make([]dto.MediaResponse, len(product.Media)),
		Variants:    make([]dto.VariantResponse, len(product.Variants)),
	}
	for i, v := range product.Variants {
		response.Variants[i] = dto.VariantResponse{
			ID:             v.ID,
			ProductID:      v.ProductID,
			SKU:            v.SKU,
			PriceAdjustment: v.PriceAdjustment.InexactFloat64(),
			TotalPrice:     v.TotalPrice.InexactFloat64(),
			Attributes:     v.Attributes,
			IsActive:       v.IsActive,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
			Inventory: dto.InventoryResponse{
				ID:               v.Inventory.ID,
				Quantity:         v.Inventory.Quantity,
				ReservedQuantity: v.Inventory.ReservedQuantity,
				LowStockThreshold: v.Inventory.LowStockThreshold,
				BackorderAllowed: v.Inventory.BackorderAllowed,
			},
		}
	}

	// Map media
	for i, m := range product.Media {
		response.Media[i] = dto.MediaResponse{
			ID:        m.ID,
			ProductID: m.ProductID,
			URL:       m.URL,
			Type:      string(m.Type),
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		}
	}

	// SimpleInventory is always nil for simple products
	response.SimpleInventory = nil

	logger.Info("Product created successfully", zap.String("product_id", product.ID))
	return response, nil
}

// GetProductByID fetches a product with optional preloads
func (s *ProductService) GetProductByID(ctx context.Context, id string, preloads ...string) (*dto.ProductResponse, error) {
	logger := s.logger.With(zap.String("operation", "GetProductByID"), zap.String("product_id", id))
	product, err := s.productRepo.FindByID(id, preloads...)
	if err != nil {
		if errors.Is(err, repositories.ErrProductNotFound) {
			return nil, err
		}
		logger.Error("Failed to fetch product", zap.Error(err))
		return nil, fmt.Errorf("failed to fetch product: %w", err)
	}

	response := &dto.ProductResponse{
		ID:          product.ID,
		MerchantID:  product.MerchantID,
		Name:        product.Name,
		Description: product.Description,
		SKU:         product.SKU,
		BasePrice:   (product.BasePrice).InexactFloat64(),
		CategoryID:  product.CategoryID,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		Variants:    make([]dto.VariantResponse, len(product.Variants)),
		Media:       make([]dto.MediaResponse, len(product.Media)),
	}
	for i, v := range product.Variants {
		response.Variants[i] = dto.VariantResponse{
			ID:             v.ID,
			ProductID:      v.ProductID,
			SKU:            v.SKU,
			PriceAdjustment: (v.PriceAdjustment).InexactFloat64(),
			TotalPrice:     (v.TotalPrice).InexactFloat64(),
			Attributes:     v.Attributes,
			IsActive:       v.IsActive,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
			Inventory: dto.InventoryResponse{
				ID:               v.Inventory.ID,
				Quantity:         v.Inventory.Quantity,
				ReservedQuantity: v.Inventory.ReservedQuantity,
				LowStockThreshold: v.Inventory.LowStockThreshold,
				BackorderAllowed: v.Inventory.BackorderAllowed,
			},
		}
	}
	for i, m := range product.Media {
		response.Media[i] = dto.MediaResponse{
			ID:        m.ID,
			ProductID: m.ProductID,
			URL:       m.URL,
			Type:      string(m.Type),
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		}
	}
	if product.SimpleInventory != nil {
		response.SimpleInventory = &dto.InventoryResponse{
			ID:               product.SimpleInventory.ID,
			Quantity:         product.SimpleInventory.Quantity,
			ReservedQuantity: product.SimpleInventory.ReservedQuantity,
			LowStockThreshold: product.SimpleInventory.LowStockThreshold,
			BackorderAllowed: product.SimpleInventory.BackorderAllowed,
		}
	}

	logger.Info("Product fetched successfully")
	return response, nil
}

// ListProductsByMerchant lists products for a merchant
func (s *ProductService) ListProductsByMerchant(ctx context.Context, merchantID string, limit, offset int, activeOnly bool) ([]dto.ProductResponse, error) {
	logger := s.logger.With(zap.String("operation", "ListProductsByMerchant"), zap.String("merchant_id", merchantID))
	products, err := s.productRepo.ListByMerchant(merchantID, limit, offset, activeOnly)
	if err != nil {
		logger.Error("Failed to list products", zap.Error(err))
		return nil, fmt.Errorf("failed to list products: %w", err)
	}

	responses := make([]dto.ProductResponse, len(products))
	for i, p := range products {
		responses[i] = dto.ProductResponse{
			ID:          p.ID,
			MerchantID:  p.MerchantID,
			Name:        p.Name,
			Description: p.Description,
			SKU:         p.SKU,
			BasePrice:   (p.BasePrice).InexactFloat64(),
			CategoryID:  p.CategoryID,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
			Variants:    make([]dto.VariantResponse, len(p.Variants)),
			Media:       make([]dto.MediaResponse, len(p.Media)),
		}
		for j, v := range p.Variants {
			responses[i].Variants[j] = dto.VariantResponse{
				ID:             v.ID,
				ProductID:      v.ProductID,
				SKU:            v.SKU,
				PriceAdjustment: (v.PriceAdjustment).InexactFloat64(),
				TotalPrice:     (v.TotalPrice).InexactFloat64(),
				Attributes:     v.Attributes,
				IsActive:       v.IsActive,
				CreatedAt:      v.CreatedAt,
				UpdatedAt:      v.UpdatedAt,
				Inventory: dto.InventoryResponse{
					ID:               v.Inventory.ID,
					Quantity:         v.Inventory.Quantity,
					ReservedQuantity: v.Inventory.ReservedQuantity,
					LowStockThreshold: v.Inventory.LowStockThreshold,
					BackorderAllowed: v.Inventory.BackorderAllowed,
				},
			}
		}
		for j, m := range p.Media {
			responses[i].Media[j] = dto.MediaResponse{
				ID:        m.ID,
				ProductID: m.ProductID,
				URL:       m.URL,
				Type:      string(m.Type),
				CreatedAt: m.CreatedAt,
				UpdatedAt: m.UpdatedAt,
			}
		}
		if p.SimpleInventory != nil {
			responses[i].SimpleInventory = &dto.InventoryResponse{
				ID:               p.SimpleInventory.ID,
				Quantity:         p.SimpleInventory.Quantity,
				ReservedQuantity: p.SimpleInventory.ReservedQuantity,
				LowStockThreshold: p.SimpleInventory.LowStockThreshold,
				BackorderAllowed: p.SimpleInventory.BackorderAllowed,
			}
		}
	}

	logger.Info("Products listed successfully", zap.Int("count", len(responses)))
	return responses, nil
}

// GetAllProducts fetches all active products for the landing page
func (s *ProductService) GetAllProducts(ctx context.Context, limit, offset int, categoryID *uint) ([]dto.ProductResponse, int64, error) {
	logger := s.logger.With(zap.String("operation", "GetAllProducts"))
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	products, total, err := s.productRepo.GetAllProducts(limit, offset, categoryID, "Media", "Variants", "Variants.Inventory", "SimpleInventory")
	if err != nil {
		logger.Error("Failed to fetch all products", zap.Error(err))
		return nil, 0, fmt.Errorf("failed to fetch products: %w", err)
	}

	responses := make([]dto.ProductResponse, len(products))
	for i, p := range products {
		responses[i] = dto.ProductResponse{
			ID:          p.ID,
			MerchantID:  "", // Exclude for customer-facing API
			Name:        p.Name,
			Description: p.Description,
			SKU:         p.SKU,
			BasePrice:   (p.BasePrice).InexactFloat64(),
			CategoryID:  p.CategoryID,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
			Variants:    make([]dto.VariantResponse, len(p.Variants)),
			Media:       make([]dto.MediaResponse, len(p.Media)),
		}
		for j, v := range p.Variants {
			responses[i].Variants[j] = dto.VariantResponse{
				ID:             v.ID,
				ProductID:      v.ProductID,
				SKU:            v.SKU,
				PriceAdjustment: (v.PriceAdjustment).InexactFloat64(),
				TotalPrice:     (v.TotalPrice).InexactFloat64(),
				Attributes:     v.Attributes,
				IsActive:       v.IsActive,
				CreatedAt:      v.CreatedAt,
				UpdatedAt:      v.UpdatedAt,
				Inventory: dto.InventoryResponse{
					ID:               v.Inventory.ID,
					Quantity:         v.Inventory.Quantity,
					ReservedQuantity: v.Inventory.ReservedQuantity,
					LowStockThreshold: v.Inventory.LowStockThreshold,
					BackorderAllowed: v.Inventory.BackorderAllowed,
				},
			}
		}
		for j, m := range p.Media {
			responses[i].Media[j] = dto.MediaResponse{
				ID:        m.ID,
				ProductID: m.ProductID,
				URL:       m.URL,
				Type:      string(m.Type),
				CreatedAt: m.CreatedAt,
				UpdatedAt: m.UpdatedAt,
			}
		}
		if p.SimpleInventory != nil {
			responses[i].SimpleInventory = &dto.InventoryResponse{
				ID:               p.SimpleInventory.ID,
				Quantity:         p.SimpleInventory.Quantity,
				ReservedQuantity: p.SimpleInventory.ReservedQuantity,
				LowStockThreshold: p.SimpleInventory.LowStockThreshold,
				BackorderAllowed: p.SimpleInventory.BackorderAllowed,
			}
		}
	}

	logger.Info("Products fetched for landing page", zap.Int("count", len(responses)), zap.Int64("total", total))
	return responses, total, nil
}

// UpdateInventory adjusts stock for a given inventory ID
func (s *ProductService) UpdateInventory(ctx context.Context, inventoryID string, delta int) error {
	logger := s.logger.With(zap.String("operation", "UpdateInventory"), zap.String("inventory_id", inventoryID))
	err := s.productRepo.UpdateInventoryQuantity(inventoryID, delta)
	if err != nil {
		logger.Error("Failed to update inventory", zap.Error(err))
		return fmt.Errorf("failed to update inventory: %w", err)
	}
	logger.Info("Inventory updated successfully", zap.Int("delta", delta))
	return nil
}

// DeleteProduct soft-deletes a product
func (s *ProductService) DeleteProduct(ctx context.Context, id string) error {
	logger := s.logger.With(zap.String("operation", "DeleteProduct"), zap.String("product_id", id))
	err := s.productRepo.SoftDeleteProduct(id)
	if err != nil {
		logger.Error("Failed to delete product", zap.Error(err))
		return fmt.Errorf("failed to delete product: %w", err)
	}
	logger.Info("Product deleted successfully")
	return nil
}
```
</xaiArtifact>

---
### services/user/user_service.go
- Size: 6.56 KB
- Lines: 242
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="b8b7955e-9242-428b-9fb0-91d804d03f85" artifact_version_id="e54012fd-ddf0-4213-921c-1a242ccbac30" title="services/user/user_service.go" contentType="text/go">
```go
package user

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	//"api-customer-merchant/internal/db"
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"

	//"google.golang.org/api/oauth2/v2"

	//"api-customer-merchant/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo     *repositories.UserRepository
	
}


func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepo:     userRepo,
		
	}
}

type googleUserInfo struct {
    Email string `json:"email"`
    Name  string `json:"name"`
}


func (s *AuthService) RegisterUser(email, name, password, country string) (*models.User, error) {
    _, err := s.userRepo.FindByEmail(email)
    if err == nil {
        return nil, errors.New("email already exists")
    } else if !errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, err
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &models.User{
        Email:    email,
        Name:     name,
        Password: string(hashedPassword),
        Country:  country,
    }

    if err := s.userRepo.Create(user); err != nil {
        return nil, err
    }

    return user, nil
}



func (s *AuthService) LoginUser(email, password string) (*models.User, error) {
    user, err := s.userRepo.FindByEmail(email)
    if err != nil {
        return nil, errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, errors.New("invalid credentials")
    }

    return user, nil
}

func (s *AuthService) GenerateJWT(entity interface{}) (string, error) {
	var id uint
	var entityType string

	switch e := entity.(type) {
	case *models.User:
		id = e.ID
		entityType = "user"

	default:
		return "", errors.New("invalid entity type")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        float64(id),
		"entityType": entityType,
        "exp":        time.Now().Add(24 * time.Hour).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET not set")
	}

	return token.SignedString([]byte(secret))
}

func (s *AuthService) GetOAuthConfig(entityType string) *oauth2.Config {
	baseURL := os.Getenv("BASE_URL")
    if baseURL == "" {
        log.Println("BASE_URL environment variable is not set")
        return nil
    }
    redirectURL := baseURL + "/customer/auth/google/callback"
    log.Printf("OAuth redirect URL: %s", redirectURL)
	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  redirectURL,
		 Scopes:       []string{
            "https://www.googleapis.com/auth/userinfo.email",
            "https://www.googleapis.com/auth/userinfo.profile",
            "openid"},
		Endpoint:     google.Endpoint,
	}
}

func (s *AuthService) GoogleLogin(code, baseURL, entityType string) (*models.User, string, error) {
    if entityType != "customer" {
        log.Printf("Invalid entityType for OAuth: %s", entityType)
        return nil, "", errors.New("OAuth only supported for customers")
    }

    // Get OAuth config
    oauthConfig := s.GetOAuthConfig(entityType)
    if oauthConfig == nil || oauthConfig.ClientID == "" || oauthConfig.ClientSecret == "" {
        log.Println("Google OAuth credentials not set")
        return nil, "", errors.New("OAuth configuration error")
    }

    // Exchange code for access token
    ctx := context.Background()
    token, err := oauthConfig.Exchange(ctx, code)
    if err != nil {
        log.Printf("Failed to exchange code: %v", err)
        return nil, "", errors.New("failed to exchange code")
    }

    // Fetch user info
    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", nil)
    if err != nil {
        log.Printf("Failed to create userinfo request: %v", err)
        return nil, "", errors.New("failed to create userinfo request")
    }
    req.Header.Set("Authorization", "Bearer "+token.AccessToken)
    resp, err := client.Do(req)
    if err != nil {
        log.Printf("Failed to get user info: %v", err)
        return nil, "", errors.New("failed to get user info")
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Printf("Userinfo endpoint returned status: %d", resp.StatusCode)
        return nil, "", errors.New("failed to get user info")
    }

    var userInfo googleUserInfo
    if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
        log.Printf("Failed to decode user info: %v", err)
        return nil, "", errors.New("failed to decode user info")
    }

    // Validate email
    if userInfo.Email == "" {
        log.Println("No email provided by Google")
        return nil, "", errors.New("no email provided")
    }

    // Check if user exists
    user, err := s.userRepo.FindByEmail(userInfo.Email)
    if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        log.Printf("Failed to find user: %v", err)
        return nil, "", err
    }
    if user == nil {
        // Register new user
        user = &models.User{
            Email:   userInfo.Email,
            Name:    userInfo.Name,
            Country: "", // Set default or prompt later
        }
        if err := s.userRepo.Create(user); err != nil {
            log.Printf("Failed to create user: %v", err)
            return nil, "", err
        }
    }

    // Generate JWT
    jwtToken, err := s.GenerateJWT(user)
    if err != nil {
        log.Printf("Failed to generate JWT: %v", err)
        return nil, "", errors.New("failed to generate JWT")
    }

    return user, jwtToken, nil
}



func (s *AuthService) UpdateProfile(userID uint, name, country string, addresses []string) error {

    user, err := s.userRepo.FindByID(userID)
    if err != nil {
        return err
    }
    user.Name = name
    user.Country = country
    // Addresses as JSON; add field to User model if needed: Addresses jsonb
    return s.userRepo.Update(user)
}

func (s *AuthService) ResetPassword(email, newPassword string) error {
    user, err := s.userRepo.FindByEmail(email)
    if err != nil {
        return err
    }
    hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashed)
    return s.userRepo.Update(user)
}
```
</xaiArtifact>

---
### tests/unit/test_handlers.go
- Size: 0.01 KB
- Lines: 1
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="2e958e76-5345-4cc1-8632-60fc4dd3d603" artifact_version_id="b0a62d2b-2fd4-43cb-98d5-ee625a86f093" title="tests/unit/test_handlers.go" contentType="text/go">
```go
package unit
```
</xaiArtifact>

---
### tests/unit/test_service.go
- Size: 0.01 KB
- Lines: 1
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="c81f954e-de9e-40af-9cd2-5e4f9aa92cad" artifact_version_id="f82d04d0-813d-41be-b019-eacd04534ad7" title="tests/unit/test_service.go" contentType="text/go">
```go
package unit

```
</xaiArtifact>

---
### tests/unit/cart_service_test.go
- Size: 17.74 KB
- Lines: 560
- Last Modified: 2025-09-19 21:41:15

<xaiArtifact artifact_id="11c03662-3642-47d2-8ac8-b8948420ad93" artifact_version_id="21fa9760-ecce-47c2-87bc-8486058e748d" title="tests/unit/cart_service_test.go" contentType="text/go">
```go
package unit
/*
import (
	"api-customer-merchant/internal/api/dto"
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// --- Mocks ---

type MockCartRepository struct {
	mock.Mock
}

func (m *MockCartRepository) FindActiveCart(ctx context.Context, userID uint) (*models.Cart, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Cart), args.Error(1)
}

func (m *MockCartRepository) Create(ctx context.Context, cart *models.Cart) error {
	args := m.Called(ctx, cart)
	return args.Error(0)
}

func (m *MockCartRepository) FindByID(ctx context.Context, id uint) (*models.Cart, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Cart), args.Error(1)
}

func (m *MockCartRepository) Update(ctx context.Context, cart *models.Cart) error {
	args := m.Called(ctx, cart)
	return args.Error(0)
}

type MockCartItemRepository struct {
	mock.Mock
}

func (m *MockCartItemRepository) FindByProductIDAndCartID(ctx context.Context, productID string, cartID uint) (*models.CartItem, error) {
	args := m.Called(ctx, productID, cartID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.CartItem), args.Error(1)
}

func (m *MockCartItemRepository) UpdateQuantityWithReservation(ctx context.Context, itemID uint, newQuantity int, inventoryID uint) error {
	args := m.Called(ctx, itemID, newQuantity, inventoryID)
	return args.Error(0)
}

func (m *MockCartItemRepository) Create(ctx context.Context, cartItem *models.CartItem) error {
	args := m.Called(ctx, cartItem)
	return args.Error(0)
}

func (m *MockCartItemRepository) FindByID(ctx context.Context, id uint) (*models.CartItem, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.CartItem), args.Error(1)
}

func (m *MockCartItemRepository) DeleteWithUnreserve(ctx context.Context, id uint, inventoryID uint) error {
	args := m.Called(ctx, id, inventoryID)
	return args.Error(0)
}

func (m *MockCartItemRepository) FindByCartID(ctx context.Context, cartID uint) ([]models.CartItem, error) {
	args := m.Called(ctx, cartID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.CartItem), args.Error(1)
}

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) FindByID(id string, preloads ...string) (*models.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

type MockInventoryRepository struct {
	mock.Mock
}

func (m *MockInventoryRepository) FindByProductID(productID string) (*models.Inventory, error) {
	args := m.Called(productID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Inventory), args.Error(1)
}

func (m *MockInventoryRepository) UpdateInventoryQuantity(ctx context.Context, inventoryID uint, delta int) error {
	args := m.Called(ctx, inventoryID, delta)
	return args.Error(0)
}

// --- Test Suite ---

type CartServiceTestSuite struct {
	service       *CartService
	cartRepo      *MockCartRepository
	cartItemRepo  *MockCartItemRepository
	productRepo   *MockProductRepository
	inventoryRepo *MockInventoryRepository
	logger        *zap.Logger
}

func (suite *CartServiceTestSuite) SetupTest() {
	suite.cartRepo = new(MockCartRepository)
	suite.cartItemRepo = new(MockCartItemRepository)
	suite.productRepo = new(MockProductRepository)
	suite.inventoryRepo = new(MockInventoryRepository)
	suite.logger, _ = zap.NewDevelopment()

	// We need to cast the mocks to the interface type for the constructor
	var cartRepoInterface repositories.CartRepository = suite.cartRepo
	var cartItemRepoInterface repositories.CartItemRepository = suite.cartItemRepo
	var productRepoInterface repositories.ProductRepository = suite.productRepo
	var inventoryRepoInterface repositories.InventoryRepository = suite.inventoryRepo

	suite.service = NewCartService(
		&cartRepoInterface,
		&cartItemRepoInterface,
		&productRepoInterface,
		&inventoryRepoInterface,
		suite.logger,
	)
}

func TestCartService(t *testing.T) {
	// This is a placeholder function that testify's suite runner will use
}

func TestGetActiveCart_Existing(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()
	userID := uint(1)
	expectedCart := &models.Cart{Model: gorm.Model{ID: 1}, UserID: userID}

	suite.cartRepo.On("FindActiveCart", ctx, userID).Return(expectedCart, nil)

	// Execute
	cart, err := suite.service.GetActiveCart(ctx, userID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedCart, cart)
	suite.cartRepo.AssertExpectations(t)
}

func TestGetActiveCart_New(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()
	userID := uint(1)
	newCart := &models.Cart{UserID: userID, Status: models.CartStatusActive}
	createdCart := &models.Cart{Model: gorm.Model{ID: 2}, UserID: userID, Status: models.CartStatusActive}

	suite.cartRepo.On("FindActiveCart", ctx, userID).Return(nil, gorm.ErrRecordNotFound)
	suite.cartRepo.On("Create", ctx, newCart).Return(nil)
	suite.cartRepo.On("FindByID", ctx, newCart.ID).Return(createdCart, nil)

	// Execute
	cart, err := suite.service.GetActiveCart(ctx, userID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, createdCart, cart)
	suite.cartRepo.AssertExpectations(t)
}

func TestGetActiveCart_InvalidUser(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()

	// Execute
	_, err := suite.service.GetActiveCart(ctx, 0)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidUserID, err)
}

func TestAddItemToCart_NewItem(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()
	userID := uint(1)
	productID := "prod-123"
	quantity := uint(2)

	cart := &models.Cart{Model: gorm.Model{ID: 1}, UserID: userID}
	product := &models.Product{ID: productID, MerchantID: "merch-1"}
	inventory := &models.Inventory{Model: gorm.Model{ID: 10}, ProductID: productID, StockQuantity: 10}

	suite.cartRepo.On("FindActiveCart", ctx, userID).Return(cart, nil)
	suite.productRepo.On("FindByID", productID).Return(product, nil)
	suite.inventoryRepo.On("FindByProductID", productID).Return(inventory, nil)
	suite.cartItemRepo.On("FindByProductIDAndCartID", ctx, productID, cart.ID).Return(nil, gorm.ErrRecordNotFound)

	// Mocking the transaction is complex. We'll assume the inner calls are what matters.
	// In a real scenario with db.DB, this would be an integration test.
	// For unit tests, we assume the transaction block works if its components are mocked correctly.
	suite.cartItemRepo.On("Create", ctx, mock.AnythingOfType("*models.CartItem")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*models.CartItem)
		assert.Equal(t, cart.ID, arg.CartID)
		assert.Equal(t, productID, arg.ProductID)
		assert.Equal(t, int(quantity), arg.Quantity)
	})
	suite.inventoryRepo.On("UpdateInventoryQuantity", ctx, inventory.ID, -int(quantity)).Return(nil)
	suite.cartRepo.On("FindByID", ctx, cart.ID).Return(cart, nil)

	// Execute
	_, err := suite.service.AddItemToCart(ctx, userID, quantity, productID)

	// Assert
	assert.NoError(t, err)
	suite.cartRepo.AssertExpectations(t)
	suite.productRepo.AssertExpectations(t)
	suite.inventoryRepo.AssertExpectations(t)
	suite.cartItemRepo.AssertExpectations(t)
}

func TestAddItemToCart_ExistingItem(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()
	userID := uint(1)
	productID := "prod-123"
	quantity := uint(2)

	cart := &models.Cart{Model: gorm.Model{ID: 1}, UserID: userID}
	product := &models.Product{ID: productID, MerchantID: "merch-1"}
	inventory := &models.Inventory{Model: gorm.Model{ID: 10}, ProductID: productID, StockQuantity: 10}
	existingItem := &models.CartItem{Model: gorm.Model{ID: 100}, ProductID: productID, CartID: cart.ID, Quantity: 3}

	suite.cartRepo.On("FindActiveCart", ctx, userID).Return(cart, nil)
	suite.productRepo.On("FindByID", productID).Return(product, nil)
	suite.inventoryRepo.On("FindByProductID", productID).Return(inventory, nil)
	suite.cartItemRepo.On("FindByProductIDAndCartID", ctx, productID, cart.ID).Return(existingItem, nil)

	newQty := int(quantity) + existingItem.Quantity
	suite.cartItemRepo.On("UpdateQuantityWithReservation", ctx, existingItem.ID, newQty, inventory.ID).Return(nil)
	suite.cartRepo.On("FindByID", ctx, cart.ID).Return(cart, nil)

	// Execute
	_, err := suite.service.AddItemToCart(ctx, userID, quantity, productID)

	// Assert
	assert.NoError(t, err)
	suite.cartRepo.AssertExpectations(t)
	suite.productRepo.AssertExpectations(t)
	suite.inventoryRepo.AssertExpectations(t)
	suite.cartItemRepo.AssertExpectations(t)
}

func TestAddItemToCart_InsufficientStock(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()
	userID := uint(1)
	productID := "prod-123"
	quantity := uint(5)

	cart := &models.Cart{Model: gorm.Model{ID: 1}, UserID: userID}
	product := &models.Product{ID: productID, MerchantID: "merch-1"}
	inventory := &models.Inventory{Model: gorm.Model{ID: 10}, ProductID: productID, StockQuantity: 3}

	suite.cartRepo.On("FindActiveCart", ctx, userID).Return(cart, nil)
	suite.productRepo.On("FindByID", productID).Return(product, nil)
	suite.inventoryRepo.On("FindByProductID", productID).Return(inventory, nil)

	// Execute
	_, err := suite.service.AddItemToCart(ctx, userID, quantity, productID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrInsufficientStock, err)
	suite.inventoryRepo.AssertExpectations(t)
}

func TestUpdateCartItemQuantity_Success(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()
	cartItemID := uint(100)
	newQuantity := 5

	cartItem := &models.CartItem{Model: gorm.Model{ID: cartItemID}, ProductID: "prod-123", CartID: 1, Quantity: 2}
	inventory := &models.Inventory{Model: gorm.Model{ID: 10}, ProductID: "prod-123", StockQuantity: 10}
	cart := &models.Cart{Model: gorm.Model{ID: 1}}

	suite.cartItemRepo.On("FindByID", ctx, cartItemID).Return(cartItem, nil)
	suite.inventoryRepo.On("FindByProductID", cartItem.ProductID).Return(inventory, nil)
	suite.cartItemRepo.On("UpdateQuantityWithReservation", ctx, cartItemID, newQuantity, inventory.ID).Return(nil)
	suite.cartRepo.On("FindByID", ctx, cartItem.CartID).Return(cart, nil)

	// Execute
	_, err := suite.service.UpdateCartItemQuantity(ctx, cartItemID, newQuantity)

	// Assert
	assert.NoError(t, err)
	suite.cartItemRepo.AssertExpectations(t)
	suite.inventoryRepo.AssertExpectations(t)
	suite.cartRepo.AssertExpectations(t)
}

func TestUpdateCartItemQuantity_InvalidQuantity(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()

	// Execute
	_, err := suite.service.UpdateCartItemQuantity(ctx, 1, 0)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidQuantity, err)
}

func TestUpdateCartItemQuantity_ItemNotFound(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()
	cartItemID := uint(999)

	suite.cartItemRepo.On("FindByID", ctx, cartItemID).Return(nil, repositories.ErrCartItemNotFound)

	// Execute
	_, err := suite.service.UpdateCartItemQuantity(ctx, cartItemID, 5)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, repositories.ErrCartItemNotFound, err)
	suite.cartItemRepo.AssertExpectations(t)
}

func TestRemoveCartItem_Success(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()
	cartItemID := uint(100)

	cartItem := &models.CartItem{Model: gorm.Model{ID: cartItemID}, ProductID: "prod-123", CartID: 1, Quantity: 2}
	inventory := &models.Inventory{Model: gorm.Model{ID: 10}, ProductID: "prod-123", StockQuantity: 10}
	cart := &models.Cart{Model: gorm.Model{ID: 1}}

	suite.cartItemRepo.On("FindByID", ctx, cartItemID).Return(cartItem, nil)
	suite.inventoryRepo.On("FindByProductID", cartItem.ProductID).Return(inventory, nil)
	suite.cartItemRepo.On("DeleteWithUnreserve", ctx, cartItemID, inventory.ID).Return(nil)
	suite.cartRepo.On("FindByID", ctx, cartItem.CartID).Return(cart, nil)

	// Execute
	_, err := suite.service.RemoveCartItem(ctx, cartItemID)

	// Assert
	assert.NoError(t, err)
	suite.cartItemRepo.AssertExpectations(t)
	suite.inventoryRepo.AssertExpectations(t)
	suite.cartRepo.AssertExpectations(t)
}

func TestRemoveCartItem_ItemNotFound(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()
	cartItemID := uint(999)

	suite.cartItemRepo.On("FindByID", ctx, cartItemID).Return(nil, repositories.ErrCartItemNotFound)

	// Execute
	_, err := suite.service.RemoveCartItem(ctx, cartItemID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, repositories.ErrCartItemNotFound, err)
	suite.cartItemRepo.AssertExpectations(t)
}

func TestGetCartItemByID_Success(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()
	cartItemID := uint(100)
	expectedItem := &models.CartItem{Model: gorm.Model{ID: cartItemID}}

	suite.cartItemRepo.On("FindByID", ctx, cartItemID).Return(expectedItem, nil)

	// Execute
	item, err := suite.service.GetCartItemByID(ctx, cartItemID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedItem, item)
	suite.cartItemRepo.AssertExpectations(t)
}

func TestClearCart_Success(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()
	userID := uint(1)

	cart := &models.Cart{Model: gorm.Model{ID: 1}, UserID: userID, Status: models.CartStatusActive}
	items := []models.CartItem{
		{Model: gorm.Model{ID: 101}, ProductID: "prod-A", CartID: 1, Quantity: 1},
		{Model: gorm.Model{ID: 102}, ProductID: "prod-B", CartID: 1, Quantity: 2},
	}
	inventoryA := &models.Inventory{Model: gorm.Model{ID: 10}, ProductID: "prod-A"}
	inventoryB := &models.Inventory{Model: gorm.Model{ID: 11}, ProductID: "prod-B"}

	suite.cartRepo.On("FindActiveCart", ctx, userID).Return(cart, nil)
	suite.cartItemRepo.On("FindByCartID", ctx, cart.ID).Return(items, nil)

	// Mock the loop of RemoveCartItem calls
	suite.cartItemRepo.On("FindByID", ctx, items[0].ID).Return(&items[0], nil)
	suite.inventoryRepo.On("FindByProductID", items[0].ProductID).Return(inventoryA, nil)
	suite.cartItemRepo.On("DeleteWithUnreserve", ctx, items[0].ID, inventoryA.ID).Return(nil)
	suite.cartRepo.On("FindByID", ctx, cart.ID).Return(cart, nil)

	suite.cartItemRepo.On("FindByID", ctx, items[1].ID).Return(&items[1], nil)
	suite.inventoryRepo.On("FindByProductID", items[1].ProductID).Return(inventoryB, nil)
	suite.cartItemRepo.On("DeleteWithUnreserve", ctx, items[1].ID, inventoryB.ID).Return(nil)
	suite.cartRepo.On("FindByID", ctx, cart.ID).Return(cart, nil)

	suite.cartRepo.On("Update", ctx, cart).Return(nil)

	// Execute
	err := suite.service.ClearCart(ctx, userID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, models.CartStatusAbandoned, cart.Status)
	suite.cartRepo.AssertExpectations(t)
	suite.cartItemRepo.AssertExpectations(t)
}

func TestBulkAddItems_InvalidUserID(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()
	items := dto.BulkUpdateRequest{}

	// Execute
	_, err := suite.service.BulkAddItems(ctx, 0, items)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidUserID, err)
}

func TestBulkAddItems_ValidationError(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()
	userID := uint(1)
	// Create a request that will fail validation (e.g., Quantity=0)
	items := dto.BulkUpdateRequest{
		UserID: userID,
		Items: []struct {
			ProductID uint "json:\"product_id\" validate:\"required\""
			Quantity  int  "json:\"quantity\" validate:\"required,gt=0\""
		}{
			{ProductID: 1, Quantity: 0},
		},
	}

	// Execute
	_, err := suite.service.BulkAddItems(ctx, userID, items)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")
}

func TestBulkAddItems_PartialFailure(t *testing.T) {
	// Setup
	suite := new(CartServiceTestSuite)
	suite.SetupTest()
	ctx := context.Background()
	userID := uint(1)
	items := dto.BulkUpdateRequest{
		UserID: userID,
		Items: []struct {
			ProductID uint "json:\"product_id\" validate:\"required\""
			Quantity  int  "json:\"quantity\" validate:\"required,gt=0\""
		}{
			{ProductID: 1, Quantity: 1},
			{ProductID: 2, Quantity: 1},
		},
	}

	// Mock AddItemToCart to fail on the second item
	// We can't easily mock the internal AddItemToCart call without refactoring.
	// A better approach for this specific test would be an integration test.
	// Here, we'll simulate the error by having a dependency fail.
	cart := &models.Cart{Model: gorm.Model{ID: 1}, UserID: userID}
	product1 := &models.Product{ID: "1", MerchantID: "merch-1"}
	inventory1 := &models.Inventory{Model: gorm.Model{ID: 10}, ProductID: "1", StockQuantity: 10}

	suite.cartRepo.On("FindActiveCart", ctx, userID).Return(cart, nil)
	// First item success
	suite.productRepo.On("FindByID", "1").Return(product1, nil)
	suite.inventoryRepo.On("FindByProductID", "1").Return(inventory1, nil)
	suite.cartItemRepo.On("FindByProductIDAndCartID", ctx, "1", cart.ID).Return(nil, gorm.ErrRecordNotFound)
	suite.cartItemRepo.On("Create", ctx, mock.Anything).Return(nil)
	suite.inventoryRepo.On("UpdateInventoryQuantity", ctx, inventory1.ID, -1).Return(nil)

	// Second item failure
	suite.productRepo.On("FindByID", "2").Return(nil, errors.New("product not found"))

	// Execute
	_, err := suite.service.BulkAddItems(ctx, userID, items)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to add item 2")
}
*/
```
</xaiArtifact>

---
### utils/blacklist.go
- Size: 0.97 KB
- Lines: 32
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="ee01162e-e457-486e-9154-7b93911a151e" artifact_version_id="97b3aac6-0641-4a07-8f96-bf998d293f2d" title="utils/blacklist.go" contentType="text/go">
```go
package utils

import (
    "context"
    "errors"
    "log"
    "time"
    //"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// Add adds a token to the Redis blacklist with an expiration
func Add(token string) error {
    if RedisClient == nil {
        log.Println("RedisClient is nil, cannot add token to blacklist")
        return errors.New("redis client not initialized")
    }
    // Set token in Redis with a 24-hour expiration
    return RedisClient.Set(ctx, "blacklist:"+token, "true", 24*time.Hour).Err()
}

// IsBlacklisted checks if a token is in the Redis blacklist
func IsBlacklisted(token string) bool {
    if RedisClient == nil {
        log.Println("RedisClient is nil, skipping blacklist check")
        return false // Fallback to allow operation if Redis is unavailable
    }
    if RedisClient == nil { return false }
    _, err := RedisClient.Get(ctx, "blacklist:"+token).Result() // Line 22
    return err == nil // Token exists in Redis if no error
}
```
</xaiArtifact>

---
### utils/redis.go
- Size: 1.41 KB
- Lines: 56
- Last Modified: 2025-09-19 21:21:08

<xaiArtifact artifact_id="6ce65764-7db5-4b45-9591-279a12f342c0" artifact_version_id="c0006d9d-9d0f-429d-8db7-fec5fe51dc39" title="utils/redis.go" contentType="text/go">
```go
package utils

import (
	"api-customer-merchant/internal/config"
	"context"
	"crypto/tls"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis(conf *config.Config) {
    
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     conf.RedisAddr,
        Password: conf.RedisPass,
        DB:       conf.RedisDB,
        TLSConfig: &tls.Config{
            MinVersion: tls.VersionTLS12, 
            InsecureSkipVerify: true, // Use with caution, only for testing
        },
    })

    ctx := context.Background()
    if err := RedisClient.Ping(ctx).Err(); err != nil {
       log.Printf("Failed to connect to Redis: %v, continuing without caching", err)
        RedisClient = nil // Fallback to avoid crashes
    } else {
        log.Println("Connected to Redis successfully")
    }
}

// Helper to get cached value or fetch and cache
func GetOrSetCache(ctx context.Context, key string, ttl time.Duration, fetch func() (any, error)) (any, error) {
    val, err := RedisClient.Get(ctx, key).Result()
    if err == nil {
        return val, nil // Deserialize if needed (e.g., JSON)
    }
    if err != redis.Nil {
        return nil, err
    }

    data, err := fetch()
    if err != nil {
        return nil, err
    }

    // Serialize if complex (e.g., JSON marshal)
    if err := RedisClient.Set(ctx, key, data, ttl).Err(); err != nil {
        return nil, err
    }
    return data, nil
}
```
</xaiArtifact>

---

---
## 📊 Summary
- Total files: 58
- Total size: 216.54 KB
- File types: .go
