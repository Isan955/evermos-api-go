package repository

import (
	"evermos-api/config"
	"evermos-api/internal/entity"
)

type StoreRepository interface {
	Create(store *entity.Store) error
	FindByUserID(userID uint) (*entity.Store, error)
	FindByID(id uint) (*entity.Store, error)
}

type storeRepository struct{}

func NewStoreRepository() StoreRepository {
	return &storeRepository{}
}

func (r *storeRepository) Create(store *entity.Store) error {
	return config.DB.Create(store).Error
}

func (r *storeRepository) FindByUserID(userID uint) (*entity.Store, error) {
	var store entity.Store
	err := config.DB.Where("user_id = ?", userID).First(&store).Error
	return &store, err
}

func (r *storeRepository) FindByID(id uint) (*entity.Store, error) {
	var store entity.Store
	err := config.DB.First(&store, id).Error
	return &store, err
}
