package usecase

import (
	"errors"
	"evermos-api/internal/entity"
	"evermos-api/internal/repository"
)

type CategoryUsecase interface {
	Create(name string) error
	GetAll(name string, page, limit int) ([]entity.Category, int64, error)
	Update(id uint, name string) error
	Delete(id uint) error
}

type categoryUsecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{repo}
}

func (u *categoryUsecase) Create(name string) error {
	return u.repo.Create(&entity.Category{Name: name})
}

func (u *categoryUsecase) GetAll(name string, page, limit int) ([]entity.Category, int64, error) {
	return u.repo.FindAll(name, page, limit)
}

func (u *categoryUsecase) Update(id uint, name string) error {
	category, err := u.repo.FindByID(id)
	if err != nil {
		return errors.New("kategori tidak ditemukan")
	}
	category.Name = name
	return u.repo.Update(category)
}

func (u *categoryUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
