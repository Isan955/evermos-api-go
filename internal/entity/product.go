package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	StoreID     uint    `json:"store_id"`
	CategoryID  uint    `json:"category_id"`
	ImageURL    string  `json:"image_url"`
}
