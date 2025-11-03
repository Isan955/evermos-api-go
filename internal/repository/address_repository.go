package repository

import (
	"evermos-api/config"
	"evermos-api/internal/entity"
)

type AddressRepository interface {
	Create(address *entity.Address) error
	FindAllByUserID(userID uint) ([]entity.Address, error)
	FindByID(id, userID uint) (*entity.Address, error)
	Update(address *entity.Address) error
	Delete(id, userID uint) error
}

type addressRepository struct{}

func NewAddressRepository() AddressRepository {
	return &addressRepository{}
}

func (r *addressRepository) Create(address *entity.Address) error {
	return config.DB.Create(address).Error
}

func (r *addressRepository) FindAllByUserID(userID uint) ([]entity.Address, error) {
	var addresses []entity.Address
	err := config.DB.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}

func (r *addressRepository) FindByID(id, userID uint) (*entity.Address, error) {
	var address entity.Address
	err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&address).Error
	return &address, err
}

func (r *addressRepository) Update(address *entity.Address) error {
	return config.DB.Save(address).Error
}

func (r *addressRepository) Delete(id, userID uint) error {
	return config.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&entity.Address{}).Error
}
