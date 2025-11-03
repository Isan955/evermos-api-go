package repository

import (
	"evermos-api/config"
	"evermos-api/internal/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll() ([]entity.Product, error)
	GetByID(id uint) (*entity.Product, error)
	Create(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		db: config.DB,
	}
}

func (r *productRepository) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) GetByID(id uint) (*entity.Product, error) {
	var product entity.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *productRepository) Create(product *entity.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(product *entity.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Product{}, id).Error
}
