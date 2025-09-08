package models

import (
	"time"
	"gorm.io/datatypes" // For JSONB if needed
)

// MerchantApplication holds the information required for a merchant onboarding application.
// This model is fully managed (migrated and queried) by the merchant API (Go/Gin).
type MerchantApplication struct {
	ID                         string         `gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4();index"`
	StoreName                  string         `gorm:"column:store_name;not null;index"`
	RegistrantName             string         `gorm:"column:registrant_name;not null"`
	PersonalEmail              string         `gorm:"column:personal_email;not null;unique;index"`
	WorkEmail                  string         `gorm:"column:work_email;not null;unique;index"`
	PersonalAddress            datatypes.JSON `gorm:"column:personal_address;type:jsonb;not null"` // JSONB for structured address
	WorkAddress                datatypes.JSON `gorm:"column:work_address;type:jsonb;not null"`
	BusinessRegistrationNumber string         `gorm:"column:business_registration_number;not null;unique"`
	Certificate                string         `gorm:"column:certificate"` // URL or base64 for business certificate
	Logo                       string         `gorm:"column:logo"`        // URL for store logo
	CategoryID                 string         `gorm:"column:category_id;type:uuid;index"` // FK to category.id
	Description                string         `gorm:"column:description"`
	Status                     string         `gorm:"column:status;default:pending;index"` // e.g., pending, more_info, approved, rejected
	RejectionReason            string         `gorm:"column:rejection_reason"`
	SubmittedAt                time.Time      `gorm:"column:submitted_at;default:current_timestamp"`
	ReviewedAt                 *time.Time     `gorm:"column:reviewed_at"` // Nullable
	ReviewedBy                 string         `gorm:"column:reviewed_by;type:uuid"` // FK to admin.id
	CreatedAt                  time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt                  time.Time      `gorm:"column:updated_at;autoUpdateTime"`
}

func (MerchantApplication) TableName() string {
	return "merchant_application" // Singular for consistency
}

// Merchant holds the active merchant account details after approval.
// This model is ONLY QUERIED by the merchant API (Go/Gin); schema migrations are handled by the admin API (Express).
// Do NOT include this in AutoMigrate() in db.go to avoid migration attempts from Go.
type Merchant struct {
	ID               string    `gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4();index"`
	ApplicationID    string    `gorm:"column:application_id;type:uuid;not null;unique;index"` // FK to merchant_application.id
	UserID           string    `gorm:"column:user_id;type:uuid;not null;unique;index"`       // FK to users.id (merchant user account)
	StoreName        string    `gorm:"column:store_name;not null;index"`
	Status           string    `gorm:"column:status;default:active;index"` // e.g., active, suspended
	CommissionTier   string    `gorm:"column:commission_tier;default:standard"`
	CommissionRate   float64   `gorm:"column:commission_rate;default:5.00"`
	AccountBalance   float64   `gorm:"column:account_balance;default:0.00"` // Current earnings minus fees
	TotalSales       float64   `gorm:"column:total_sales;default:0.00"`     // Lifetime sales
	TotalPayouts     float64   `gorm:"column:total_payouts;default:0.00"`   // Lifetime payouts
	StripeAccountID  string    `gorm:"column:stripe_account_id"`             // For connected account/payments
	PayoutSchedule   string    `gorm:"column:payout_schedule;default:weekly"` // e.g., weekly, monthly
	LastPayoutDate   *time.Time `gorm:"column:last_payout_date"`              // Nullable
	Logo             string    `gorm:"column:logo"`                          // Store logo URL
	Banner           string    `gorm:"column:banner"`                        // Store banner URL
	Policies         string    `gorm:"column:policies;type:jsonb"`           // JSONB for shipping/return policies
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time `gorm:"column:updated_at;autoUpdateTime"`
	 //Associations (not stored, but for queries):
	 Products      []Product     `gorm:"foreignKey:MerchantID"` // One-to-many with products
	 Orders        []Order       `gorm:"foreignKey:MerchantID"` // Orders fulfilled by this merchant
	 Payouts       []Payout      `gorm:"foreignKey:MerchantID"` // Payout history
	// Catalog can be queried via products association
}

func (Merchant) TableName() string {
	return "merchant" // Singular for consistency
}