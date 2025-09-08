package payment

/*
import (
	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/db/repositories"
	"errors"
	"os"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
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
func (s *PaymentService) ProcessPayment(orderID uint, amount float64) (*models.Payment, error) {
    if orderID == 0 {
        return nil, errors.New("invalid order ID")
    }
    if amount <= 0 {
        return nil, errors.New("amount must be positive")
    }

    // Verify order exists
    order, err := s.orderRepo.FindByID(orderID)
    if err != nil {
        return nil, errors.New("order not found")
    }

    // Verify order amount matches
    if order.TotalAmount != amount {
        return nil, errors.New("payment amount does not match order total")
    }

    // Stripe setup (key from env)
    stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

    // Create Payment Intent with commission (application fee)
    pi, err := paymentintent.New(&stripe.PaymentIntentParams{  
        Amount:   stripe.Int64(int64(amount * 100)),
        Currency: stripe.String(string(stripe.CurrencyUSD)),
        // Add ApplicationFeeAmount for platform commission
        ApplicationFeeAmount: stripe.Int64(int64(amount * 100 * 0.1)), // 10% fee
        TransferData: &stripe.PaymentIntentTransferDataParams{
            Destination: stripe.String("merchant_stripe_id"), // From merchant model
        },
    })
    if err != nil {
        return nil, err
    }

    payment := &models.Payment{
        OrderID: orderID,
        Amount:  amount,
        Status:  models.PaymentStatusPending,
        //ExternalID: pi.ID,
    }
    if err := s.paymentRepo.Create(payment); err != nil {
        return nil, err
    }

    // Assume confirmation; in real, handle webhook
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

*/