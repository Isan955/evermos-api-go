package entity

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID     uint   `json:"user_id"`
	Receiver   string `json:"receiver"`
	Phone      string `json:"phone"`
	Province   string `json:"province"`
	City       string `json:"city"`
	District   string `json:"district"`
	PostalCode string `json:"postal_code"`
	Detail     string `json:"detail"`
	IsPrimary  bool   `json:"is_primary" gorm:"default:false"`
}
