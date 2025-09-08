package models

import (
    "time"
    "gorm.io/gorm"
)

type Promotion struct {
    gorm.Model
    Code       string    `gorm:"unique"`
    Discount   float64
    ValidFrom  time.Time
    ValidTo    time.Time
    MerchantID uint
}