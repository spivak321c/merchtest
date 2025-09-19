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