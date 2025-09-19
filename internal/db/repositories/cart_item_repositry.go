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

func (r *CartItemRepository) FindByProductIDAndCartID(ctx context.Context, productID, cartID uint) (*models.CartItem, error) {
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