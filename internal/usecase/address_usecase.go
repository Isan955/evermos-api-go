package usecase

import (
	"errors"
	"evermos-api/internal/entity"
	"evermos-api/internal/repository"
)

type AddressUsecase interface {
	Create(userID uint, req entity.Address) error
	GetAll(userID uint) ([]entity.Address, error)
	Update(userID, id uint, req entity.Address) error
	Delete(userID, id uint) error
}

type addressUsecase struct {
	addressRepo repository.AddressRepository
}

func NewAddressUsecase(repo repository.AddressRepository) AddressUsecase {
	return &addressUsecase{repo}
}

func (u *addressUsecase) Create(userID uint, req entity.Address) error {
	req.UserID = userID
	return u.addressRepo.Create(&req)
}

func (u *addressUsecase) GetAll(userID uint) ([]entity.Address, error) {
	return u.addressRepo.FindAllByUserID(userID)
}

func (u *addressUsecase) Update(userID, id uint, req entity.Address) error {
	address, err := u.addressRepo.FindByID(id, userID)
	if err != nil {
		return errors.New("alamat tidak ditemukan atau bukan milik anda")
	}
	address.Receiver = req.Receiver
	address.Phone = req.Phone
	address.Province = req.Province
	address.City = req.City
	address.District = req.District
	address.PostalCode = req.PostalCode
	address.Detail = req.Detail
	address.IsPrimary = req.IsPrimary
	return u.addressRepo.Update(address)
}

func (u *addressUsecase) Delete(userID, id uint) error {
	return u.addressRepo.Delete(id, userID)
}
