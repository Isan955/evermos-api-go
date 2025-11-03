package entity

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
	User   User   `json:"-" gorm:"foreignKey:UserID"`
}
