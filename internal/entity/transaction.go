package entity

import (
	"gorm.io/gorm"
)



type Transaction struct {
	gorm.Model
	UserID           uint              `json:"user_id"`
	Status           string            `json:"status" gorm:"default:'pending'"`
	TotalAmount      float64           `json:"total_amount"`
	TransactionItems []TransactionItem `json:"items" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
