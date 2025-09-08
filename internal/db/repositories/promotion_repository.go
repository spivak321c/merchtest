package repositories

import (
    "api-customer-merchant/internal/db"
    "api-customer-merchant/internal/db/models"
    "gorm.io/gorm"
)

type PromotionRepository struct {
    db *gorm.DB
}

func NewPromotionRepository() *PromotionRepository {
    return &PromotionRepository{db: db.DB}
}

func (r *PromotionRepository) Create(promo *models.Promotion) error {
    return r.db.Create(promo).Error
}

func (r *PromotionRepository) FindByCode(code string) (*models.Promotion, error) {
    var promo models.Promotion
    err := r.db.Where("code = ?", code).First(&promo).Error
    return &promo, err
}