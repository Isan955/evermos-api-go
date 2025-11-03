package entity


import "gorm.io/gorm"

type TransactionItem struct {
	gorm.Model
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id"`
	Qty           int     `json:"qty"`
	Price         float64 `json:"price"`
}