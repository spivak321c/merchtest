package repositories

import (
    "api-customer-merchant/internal/db"
    "api-customer-merchant/internal/db/models"
    "gorm.io/gorm"
)

type ReturnRepository struct {
    db *gorm.DB
}

func NewReturnRepository() *ReturnRepository {
    return &ReturnRepository{db: db.DB}
}

func (r *ReturnRepository) Create(req *models.ReturnRequest) error {
    return r.db.Create(req).Error
}

func (r *ReturnRepository) FindByID(id uint) (*models.ReturnRequest, error) {
    var req models.ReturnRequest
    err := r.db.First(&req, id).Error
    return &req, err
}

func (r *ReturnRepository) Update(req *models.ReturnRequest) error {
    return r.db.Save(req).Error
}