package usecase

import (
	"errors"
	"evermos-api/internal/entity"
	"evermos-api/internal/repository"
)

type ProductUsecase interface {
	GetAll() ([]entity.Product, error)
	GetByID(id uint) (*entity.Product, error)
	Create(product *entity.Product) error
	Update(id uint, product *entity.Product) error
	Delete(id uint) error
}

type productUsecase struct {
	productRepo repository.ProductRepository
}

func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase {
	return &productUsecase{productRepo}
}

func (u *productUsecase) GetAll() ([]entity.Product, error) {
	return u.productRepo.GetAll()
}

func (u *productUsecase) GetByID(id uint) (*entity.Product, error) {
	return u.productRepo.GetByID(id)
}

func (u *productUsecase) Create(product *entity.Product) error {
	return u.productRepo.Create(product)
}

func (u *productUsecase) Update(id uint, product *entity.Product) error {
	existing, err := u.productRepo.GetByID(id)
	if err != nil {
		return errors.New("product not found")
	}

	existing.Name = product.Name
	existing.Description = product.Description
	existing.Price = product.Price
	existing.Stock = product.Stock
	existing.CategoryID = product.CategoryID

	return u.productRepo.Update(existing)
}

func (u *productUsecase) Delete(id uint) error {
	return u.productRepo.Delete(id)
}
