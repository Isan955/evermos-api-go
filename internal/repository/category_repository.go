package repository

import (
	"evermos-api/config"
	"evermos-api/internal/entity"
)

type CategoryRepository interface {
	Create(category *entity.Category) error
	FindAll(name string, page, limit int) ([]entity.Category, int64, error)
	FindByID(id uint) (*entity.Category, error)
	Update(category *entity.Category) error
	Delete(id uint) error
}

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) Create(category *entity.Category) error {
	return config.DB.Create(category).Error
}

func (r *categoryRepository) FindAll(name string, page, limit int) ([]entity.Category, int64, error) {
	var categories []entity.Category
	var total int64
	db := config.DB.Model(&entity.Category{})
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	db.Count(&total)
	err := db.Offset((page-1)*limit).Limit(limit).Find(&categories).Error
	return categories, total, err
}

func (r *categoryRepository) FindByID(id uint) (*entity.Category, error) {
	var category entity.Category
	err := config.DB.First(&category, id).Error
	return &category, err
}

func (r *categoryRepository) Update(category *entity.Category) error {
	return config.DB.Save(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return config.DB.Delete(&entity.Category{}, id).Error
}
