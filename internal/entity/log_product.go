package entity

import "gorm.io/gorm"

type LogProduct struct {
	gorm.Model
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Qty           int     `json:"qty"`
}
