package repositories

import (
    "context"

    "api-customer-merchant/internal/db"
    "api-customer-merchant/internal/db/models"
)

// MerchantApplicationRepository handles CRUD for merchant applications
// Note: Admin service (in Express) will be responsible for updating status/approval.
type MerchantApplicationRepository struct{}

func NewMerchantApplicationRepository() *MerchantApplicationRepository { return &MerchantApplicationRepository{} }

func (r *MerchantApplicationRepository) Create(ctx context.Context, m *models.MerchantApplication) error {
    return db.DB.WithContext(ctx).Create(m).Error
}

func (r *MerchantApplicationRepository) GetByID(ctx context.Context, id string) (*models.MerchantApplication, error) {
    var m models.MerchantApplication
    if err := db.DB.WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &m, nil
}

func (r *MerchantApplicationRepository) GetByUserEmail(ctx context.Context, email string) (*models.MerchantApplication, error) {
    var m models.MerchantApplication
    if err := db.DB.WithContext(ctx).Where("personal_email = ? OR work_email = ?", email, email).First(&m).Error; err != nil {
        return nil, err
    }
    return &m, nil
}

// MerchantRepository handles active merchants
type MerchantRepository struct{}

func NewMerchantRepository() *MerchantRepository { return &MerchantRepository{} }

func (r *MerchantRepository) GetByID(ctx context.Context, id string) (*models.Merchant, error) {
    var m models.Merchant
    if err := db.DB.WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &m, nil
}

func (r *MerchantRepository) GetByUserID(ctx context.Context, uid string) (*models.Merchant, error) {
    var m models.Merchant
    if err := db.DB.WithContext(ctx).Where("user_id = ?", uid).First(&m).Error; err != nil {
        return nil, err
    }
    return &m, nil
}