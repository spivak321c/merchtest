package cart

import (
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"
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
	logger      *zap.Logger
	validator   *validator.Validate
}

func NewCartService(cartRepo *repositories.CartRepository, cartItemRepo *repositories.CartItemRepository, productRepo *repositories.ProductRepository, inventoryRepo *repositories.InventoryRepository, logger *zap.Logger) *CartService {
	return &CartService{
		cartRepo:      cartRepo,
		cartItemRepo:  cartItemRepo,
		productRepo:   productRepo,
		inventoryRepo: inventoryRepo,
		logger:      logger,
		validator:   validator.New(),
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
func (s *CartService) AddItemToCart(ctx context.Context, userID, productID, quantity uint) (*models.Cart, error) {
	if userID == 0 {
		return nil, ErrInvalidUserID
	}
	if productID == 0 {
		return nil, errors.New("invalid product ID")
	}
	if quantity == 0 {
		return nil, ErrInvalidQuantity
	}

	cart, err := s.GetActiveCart(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Fixed: Assume productRepo.FindByID takes (ctx, uint)
	product, err := s.productRepo.FindByID(ctx, productID) // Changed: uint, added ctx
	if err != nil {
		return nil, ErrProductNotFound
	}

	// Fixed: Assume FindByProductID takes (ctx, uint)
	inventory, err := s.inventoryRepo.FindByProductID(ctx, productID) // Changed: uint, ctx
	if err != nil {
		return nil, ErrInventoryNotFound
	}
	// Fixed: Assume field is Quantity
	if inventory.Quantity < int(quantity) { // Changed to Quantity
		return nil, ErrInsufficientStock
	}

	// Fixed: Tx returns (*Cart, error)
	err = s.cartRepo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		cartItem, err := s.cartItemRepo.FindByProductIDAndCartID(ctx, productID, cart.ID)
		newQty := quantity
		if err == nil {
			newQty += uint(cartItem.Quantity)
			if inventory.Quantity < int(newQty) {
				return ErrInsufficientStock
			}
			// Fixed: inventory.ID as uint
			if err := s.cartItemRepo.UpdateQuantityWithReservation(ctx, cartItem.ID, int(newQty), inventory.ID); err != nil {
				return err
			}
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			cartItem = &models.CartItem{
				CartID:    cart.ID,
				ProductID: productID,
				Quantity:  int(quantity),
				MerchantID: product.MerchantID,
			}
			if err := s.cartItemRepo.Create(ctx, cartItem); err != nil {
				return err
			}
			// Fixed: Assume method takes (ctx, uint, int)
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
func (s *CartService) UpdateCartItemQuantity(cartItemID uint, quantity int) (*models.Cart, error) {
	if cartItemID == 0 {
		return nil, errors.New("invalid cart item ID")
	}
	if quantity <= 0 {
		return nil, errors.New("quantity must be positive")
	}

	// Get cart item
	cartItem, err := s.cartItemRepo.FindByID(cartItemID)
	if err != nil {
		return nil, errors.New("cart item not found")
	}

	// Check stock availability
	inventory, err := s.inventoryRepo.FindByProductID(cartItem.ProductID)
	if err != nil {
		return nil, errors.New("inventory not found")
	}
	if inventory.StockQuantity < quantity {
		return nil, errors.New("insufficient stock")
	}

	// Update quantity
	if err := s.cartItemRepo.UpdateQuantity(cartItemID, quantity); err != nil {
		return nil, err
	}

	return s.cartRepo.FindByID(cartItem.CartID)
}

// RemoveCartItem removes an item from the cart
func (s *CartService) RemoveCartItem(cartItemID uint) (*models.Cart, error) {
	if cartItemID == 0 {
		return nil, errors.New("invalid cart item ID")
	}

	// Get cart item to find cart ID
	cartItem, err := s.cartItemRepo.FindByID(cartItemID)
	if err != nil {
		return nil, errors.New("cart item not found")
	}

	// Delete cart item
	if err := s.cartItemRepo.Delete(cartItemID); err != nil {
		return nil, err
	}

	return s.cartRepo.FindByID(cartItem.CartID)
}

func (s *CartService) GetCartItemByID(cartItemID uint) (*models.CartItem, error) {
	if cartItemID == 0 {
		return nil, errors.New("invalid cart item ID")
	}
	return s.cartItemRepo.FindByID(cartItemID)
}
