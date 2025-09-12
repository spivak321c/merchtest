package returns

import (
    "errors"
    "api-customer-merchant/internal/db/models"
    "api-customer-merchant/internal/db/repositories"
)

type ReturnService struct {
    repo          *repositories.ReturnRepository
    orderItemRepo *repositories.OrderItemRepository
}

func NewReturnService(repo *repositories.ReturnRepository, orderItemRepo *repositories.OrderItemRepository) *ReturnService {
    return &ReturnService{repo: repo, orderItemRepo: orderItemRepo}
}

func (s *ReturnService) CreateReturn(orderItemID uint, reason string, userID uint) error {
    item, err := s.orderItemRepo.FindByID(orderItemID)
    if err != nil || item.Order.UserID != userID {
        return errors.New("invalid item")
    }
    req := &models.ReturnRequest{OrderItemID: orderItemID, Reason: reason, Status: "Pending"}
    return s.repo.Create(req)
}

func (s *ReturnService) ApproveReturn(returnID uint, merchantID string) error {
    req, err := s.repo.FindByID(returnID)
    if err != nil {
        return err
    }
    item, _ := s.orderItemRepo.FindByID(req.OrderItemID)
    if item.MerchantID != merchantID {
        return errors.New("permission denied")
    }
    req.Status = "Approved"
    // Trigger refund, restock
    return s.repo.Update(req)
}