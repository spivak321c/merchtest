package order

import (
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"
	"api-customer-merchant/internal/domain/notifications"
	"api-customer-merchant/internal/domain/pricing"
	"context"
	"fmt"

	"errors"

	"gorm.io/gorm"
)

type OrderService struct {
    orderRepo          *repositories.OrderRepository
    cartRepo           *repositories.CartRepository
    cartItemRepo       *repositories.CartItemRepository
    orderItemRepo      *repositories.OrderItemRepository
    inventoryRepo      *repositories.InventoryRepository
    //paymentRepo        *repositories.PaymentRepository
    pricingService     *pricing.PricingService
    notificationService *notifications.NotificationService
}

// NewOrderService creates a new OrderService with dependencies
func NewOrderService(
    orderRepo *repositories.OrderRepository,
    cartRepo *repositories.CartRepository,
    cartItemRepo *repositories.CartItemRepository,
    orderItemRepo *repositories.OrderItemRepository,
    inventoryRepo *repositories.InventoryRepository,
    //paymentRepo *repositories.PaymentRepository,
    pricingService *pricing.PricingService,
    notificationService *notifications.NotificationService,
) *OrderService {
    return &OrderService{
        orderRepo:          orderRepo,
        cartRepo:           cartRepo,
        cartItemRepo:       cartItemRepo,
        orderItemRepo:      orderItemRepo,
        inventoryRepo:      inventoryRepo,
        //paymentRepo:        paymentRepo,
        pricingService:     pricingService,        
        notificationService: notificationService,  
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
	// Assume pricingService injected in NewOrderService
		
    // Calculate total amount
    var totalAmount float64
    for _, item := range cartItems {
        totalAmount += float64(item.Quantity) * item.Product.Price
    }

order := &models.Order{
		UserID:      userID,
		TotalAmount: totalAmount,
		Status:      models.OrderStatusPending,
	}
	uniqueMerchants := make(map[uint]bool)
	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}
 // Assume pricingService injected in NewOrderService
	shipping, err := s.pricingService.CalculateShipping(cart, "user_address") // Fetch from user
	if err != nil {
		return nil, err
	}
	tax, err := s.pricingService.CalculateTax(cart, "user_country")
	if err != nil {
		return nil, err
	}
	discount, err := s.pricingService.ApplyPromotion(cart, "promo_code") // From request
	if err != nil {
		return nil, err
	}
	order.TotalAmount = totalAmount + shipping + tax - discount
	order.ShippingCost = shipping
	order.TaxAmount = tax

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

	 go func() {
        ctx := context.Background()
        // Assume user_email is fetched from user data; replace "user_email" with actual email
        if err := s.notificationService.SendOrderConfirmation(ctx, order.ID, "user_email"); err != nil {
            // Log, don't fail
            fmt.Printf("Failed to send confirmation for order %d: %v\n", order.ID, err)
        }
        // For each unique merchant in items
        for merchantID := range uniqueMerchants {
            if err := s.notificationService.SendMerchantNewOrder(merchantID, order.ID); err != nil {
                fmt.Printf("Failed to notify merchant %d for order %d: %v\n", merchantID, order.ID, err)
            }
        }
    }()

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


func (s *OrderService) UpdateOrderItemStatus(orderItemID uint, status models.FulfillmentStatus, merchantID string) error {
    item, err := s.orderItemRepo.FindByID(orderItemID)
    if err != nil || item.MerchantID != merchantID {
        return errors.New("invalid item or permission")
    }
    item.FulfillmentStatus = status
    return s.orderItemRepo.Update(item)
}

func (s *OrderService) ProcessRefund(orderID uint, amount float64) error {
    // payment, err := s.paymentRepo.FindByOrderID(orderID)
    // if err != nil || payment.Status != models.PaymentStatusCompleted {
    //     return errors.New("cannot refund")
    // }
    // Placeholder for Stripe refund
    return nil
}