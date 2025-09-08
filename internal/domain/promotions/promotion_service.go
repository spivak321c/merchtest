package promotions

import (
    "errors"
    "time"
    "api-customer-merchant/internal/db/models"
    "api-customer-merchant/internal/db/repositories"
)

type PromotionService struct {
    repo *repositories.PromotionRepository
}

func NewPromotionService(repo *repositories.PromotionRepository) *PromotionService {
    return &PromotionService{repo: repo}
}

func (s *PromotionService) CreatePromotion(promo *models.Promotion) error {
    if promo.Code == "" {
        return errors.New("code required")
    }
    return s.repo.Create(promo)
}

func (s *PromotionService) ValidateCode(code string) (*models.Promotion, error) {
    promo, err := s.repo.FindByCode(code)
    if err != nil || time.Now().Before(promo.ValidFrom) || time.Now().After(promo.ValidTo) {
        return nil, errors.New("invalid promo")
    }
    return promo, nil
}