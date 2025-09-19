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
func (s *CartService) AddItemToCart(ctx context.Context, userID,  quantity uint,productID string) (*models.Cart, error) {
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
	product, err := s.productRepo.FindByID(productID) // Changed: uint, added ctx
	if err != nil {
		return nil, ErrProductNotFound
	}

	// Fixed: Assume FindByProductID takes (ctx, uint)
	inventory, err := s.inventoryRepo.FindByProductID(productID) // Changed: uint, ctx
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
func (s *CartService) UpdateCartItemQuantity(ctx context.Context, cartItemID uint, quantity int) (*models.Cart, error) {
	if cartItemID == 0 {
		return nil, errors.New("invalid cart item ID")
	}
	if quantity <= 0 {
		return nil, ErrInvalidQuantity
	}

	cartItem, err := s.cartItemRepo.FindByID(ctx, cartItemID)
	if err != nil {
		return nil, ErrCartItemNotFound
	}

	// Fixed: ctx, uint
	inventory, err := s.inventoryRepo.FindByProductID(ctx, cartItem.ProductID)
	if err != nil {
		return nil, ErrInventoryNotFound
	}
	if inventory.Quantity < quantity {
		return nil, ErrInsufficientStock
	}

	// Fixed: inventory.ID uint
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
		return nil, ErrCartItemNotFound
	}

	// Fixed: ctx, uint
	inventory, err := s.inventoryRepo.FindByProductID(ctx, cartItem.ProductID)
	if err != nil {
		return nil, ErrInventoryNotFound
	}

	// Fixed: inventory.ID uint
	if err := s.cartItemRepo.DeleteWithUnreserve(ctx, cartItemID, inventory.ID); err != nil {
		return nil, err
	}

	return s.cartRepo.FindByID(ctx, cartItem.CartID)
}

func (s *CartService) GetCartItemByID(ctx context.Context, cartItemID uint) (*models.CartItem, error) { // Added ctx
	if cartItemID == 0 {
		return nil, errors.New("invalid cart item ID")
	}
	return s.cartItemRepo.FindByID(ctx, cartItemID)
}

// ClearCart, BulkAddItems ... (add ctx to calls; stub Bulk if not used)
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
func (s *CartService) BulkAddItems(ctx context.Context, userID uint, items []dto.BulkUpdateRequest) (*models.Cart, error) { // Fixed DTO name
	// Validation loop...
	for _, item := range items {
		if err := s.validator.Struct(&item); err != nil {
			return nil, err
		}
		// Add each (loop AddItemToCart)
		_, err := s.AddItemToCart(ctx, userID, item.ProductID, uint(item.Quantity))
		if err != nil {
			return nil, err
		}
	}
	return s.GetCart(ctx, userID)
}