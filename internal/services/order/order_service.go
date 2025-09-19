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
