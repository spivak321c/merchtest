package models

import "gorm.io/gorm"

type ReturnRequest struct {
    gorm.Model
    OrderItemID uint   `gorm:"not null"`
    Reason      string
    Status      string `gorm:"default:'Pending'"`
}