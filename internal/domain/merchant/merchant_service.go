package merchant

import (
    "context"

    "api-customer-merchant/internal/db/models"
    "api-customer-merchant/internal/db/repositories"
)

type MerchantService struct {
    appRepo *repositories.MerchantApplicationRepository
    repo    *repositories.MerchantRepository
}

func NewMerchantService(appRepo *repositories.MerchantApplicationRepository, repo *repositories.MerchantRepository) *MerchantService {
    return &MerchantService{appRepo: appRepo, repo: repo}
}

// SubmitApplication allows a prospective merchant to submit an application.
func (s *MerchantService) SubmitApplication(ctx context.Context, app *models.MerchantApplication) (*models.MerchantApplication, error) {
    if err := s.appRepo.Create(ctx, app); err != nil {
        return nil, err
    }
    return app, nil
}

// GetApplication returns an application by ID (for applicant to check their own status).
func (s *MerchantService) GetApplication(ctx context.Context, id string) (*models.MerchantApplication, error) {
    return s.appRepo.GetByID(ctx, id)
}

// GetMerchantByUserID returns an active merchant account for the authenticated user.
func (s *MerchantService) GetMerchantByUserID(ctx context.Context, uid string) (*models.Merchant, error) {
    return s.repo.GetByUserID(ctx, uid)
}

// GetMerchantByID returns an active merchant by ID.
func (s *MerchantService) GetMerchantByID(ctx context.Context, id string) (*models.Merchant, error) {
    return s.repo.GetByID(ctx, id)
}